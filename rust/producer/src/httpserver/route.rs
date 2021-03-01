use actix_web::{get, post, web, HttpRequest, HttpResponse};
use serde::{Deserialize, Serialize};

// Push Message Structure
#[derive(Debug, Serialize, Deserialize)]
pub struct Push {
    id: String,
    message: String
}

// Default Health Check endpoint for the server.
#[get("/health")]
pub async fn health(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().json("{\n  \"message\": \"The producer service is working properly.\"\n}")
}

#[post("/v1/push/send")]
pub async fn send(_message: web::Json<Push>) -> HttpResponse {
    println!("model: {:?}", &_message);
    HttpResponse::Ok().json(_message.0)
}
