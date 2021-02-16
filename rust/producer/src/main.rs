mod server;
use crate::server::Server;

fn main() -> std::io::Result<()> {
    Server::run()
}
