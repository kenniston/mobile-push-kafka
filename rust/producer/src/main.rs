use actix_web::{get, App, HttpRequest, HttpResponse, HttpServer};
use actix_web::middleware::Logger;
use clap::{Arg, App as Clap};
use env_logger::Env;

#[get("/health")]
async fn health(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().body("{\n  \"message\": \"The producer service is working properly.\"\n}")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    let name = env!("CARGO_PKG_NAME");
    let version = env!("CARGO_PKG_VERSION");
    let authors = env!("CARGO_PKG_AUTHORS");

    let matches = Clap::new(name)
        .version(version)
        .author(authors)
        .arg(Arg::with_name("port")
            .short("p")
            .required(true)
            .default_value("4004")
            .env("PORT")
            .help("Configure the Server Port."))
        .get_matches();

    let port: u16 = matches.value_of("port").unwrap().parse().unwrap();

    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(health)
    })
    .bind(("0.0.0.0", port))?
    .run()
    .await
}
