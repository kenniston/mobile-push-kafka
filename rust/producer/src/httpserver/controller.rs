use actix_web::{web, HttpRequest, HttpResponse};
use serde::{Deserialize, Serialize};
use crate::httpserver::kafka::KafkaClient;

// Push Message Structure
#[derive(Debug, Serialize, Deserialize)]
pub struct Push {
    id: String,
    message: String
}

pub struct Controller {
    kafka_client: KafkaClient
}

impl Controller {
    pub fn default(server: String, topic: String) -> Self {
        Controller{ kafka_client: KafkaClient::default(server, topic) }
    }

    pub async fn health(&self, _req: HttpRequest) -> HttpResponse {
        HttpResponse::Ok().json("{\n  \"message\": \"The producer service is working properly.\"\n}")
    }

    pub async fn send(&self, _message: web::Json<Push>) -> HttpResponse {
        let json = serde_json::to_string(&&_message.0);
        self.kafka_client.send_message(json.unwrap());
        HttpResponse::Ok().json(_message.0)
    }
}




