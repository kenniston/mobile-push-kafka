use actix_web::{get, post, web, HttpRequest, HttpResponse};
use serde::{Deserialize, Serialize};
use crate::httpserver::kafka::KafkaClient;

// Push Message Structure
#[derive(Debug, Serialize, Deserialize)]
pub struct Push {
    id: String,
    message: String
}

static KAFKA_CLIENT: KafkaClient = KafkaClient::default();

pub fn configure<F>(factory: F) where F: Fn() -> (String, String) {
    let (server, topic) = factory();
    KAFKA_CLIENT.configure(server, topic);
}

pub fn configure2(factory: &dyn Fn() -> (String, String)) -> Self {

}

// Default Health Check endpoint for the server.
#[get("/health")]
pub async fn health(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().json("{\n  \"message\": \"The producer service is working properly.\"\n}")
}

#[post("/v1/push/send")]
pub async fn send(_message: web::Json<Push>) -> HttpResponse {
    let json = serde_json::to_string(&&_message.0);
    KAFKA_CLIENT.send_message(json.unwrap());
    HttpResponse::Ok().json(_message.0)
}


