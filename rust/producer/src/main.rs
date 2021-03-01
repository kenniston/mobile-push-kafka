mod httpserver;

fn main() -> std::io::Result<()> {
    httpserver::server::Server::run()
}
