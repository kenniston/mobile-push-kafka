use actix_web::{get, middleware::Logger, App, HttpRequest, HttpResponse, HttpServer};
use clap::{Arg, App as Clap, ArgMatches};
use std::str::FromStr;
use env_logger::Env;
use log::{info, LevelFilter};
use serde_gelf::GelfLevel;
use log4rs::{
    append::{ console::{ConsoleAppender, Target} },
    encode::pattern::PatternEncoder,
    config::{Appender, Config, Root}
};

// use gelf::{Logger as Gelf, UdpBackend};
// use log::LevelFilter as LogLevelFilter;

fn config_args() -> ArgMatches<'static> {
    let name = env!("CARGO_PKG_NAME");
    let version = env!("CARGO_PKG_VERSION");
    let authors = env!("CARGO_PKG_AUTHORS");

    Clap::new(name)
        .version(version)
        .author(authors)
        .arg(Arg::with_name("port")
            .short("p")
            .required(true)
            .default_value("4004")
            .env("RS_PORT")
            .help("Configure the server port."))
        .arg(Arg::with_name("log-level")
            .default_value("info")
            .env("RS_LOG_LEVEL")
            .help("Configure the server log level.")
        )
        .arg(Arg::with_name("graylog-server")
            .default_value("localhost")
            .env("RS_GRAYLOG_SERVER")
            .help("Configure the server address.")
        )
        .arg(Arg::with_name("graylog-port")
            .default_value("8888")
            .env("RS_GRAYLOG_PORT")
            .help("Configure the server port (TCP).")
        )
        .get_matches()
}

// Configure Stdout Logger and Graylog Logger. This function get params from
// ArgMatches (Clap Crate) to configure the output level, graylog server and port.
fn logger_config(args: &ArgMatches) {
    let level_param: &str = args.value_of("log-level").unwrap();
    let graylog_server: &str = args.value_of("graylog-server").unwrap();
    let graylog_port: u64 = args.value_of("graylog-port").unwrap().parse().unwrap();
    let log_level = LevelFilter::from_str(&level_param).unwrap();

    // Build a stdout logger
    let stdout = ConsoleAppender::builder()
        .encoder(Box::new(PatternEncoder::new("{d(%Y-%m-%d %H:%M:%S %Z)} {h({l})} [{t}:{L}] - {m}{n}")))
        .target(Target::Stdout)
        .build();

    let graylog_buffer = log4rs_gelf::BufferAppender::builder()
        .set_level(GelfLevel::from(log_level.to_level().unwrap()))
        .set_hostname(graylog_server)
        .set_port(graylog_port)
        .set_use_tls(false)
        .build()
        .unwrap();

    // Configure the Root Logger
    let config = Config::builder()
        .appender(Appender::builder().build("stdout", Box::new(stdout)))
        .appender(Appender::builder().build("graylog", Box::new(graylog_buffer)))
        .build(Root::builder()
            .appender("graylog")
            .appender("stdout")
            .build(log_level),
        )
        .unwrap();

    let _handle = log4rs::init_config(config);
}

// Default Health Check endpoint for the server.
#[get("/health")]
async fn health(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().body("{\n  \"message\": \"The producer service is working properly.\"\n}")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Configure server arguments
    let matches = config_args();

    // Configure Graylog
    // let backend = UdpBackend::new("localhost:8888");
    // let gelf_logger = Gelf::new(Box::new(backend.unwrap()));
    // let _ = gelf_logger.unwrap().install(LogLevelFilter::Info);
    logger_config(&matches);

    // Configure the log level
    let _ = env_logger::Builder::from_env(Env::default().default_filter_or("info")).try_init();

    // Get server port from program params or env
    let port: u16 = matches.value_of("port").unwrap().parse().unwrap();

    info!("Initializing server on port {}...", port);
    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(health)
    })
    .bind(("0.0.0.0", port))?
    .run()
    .await
}
