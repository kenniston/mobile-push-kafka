use rdkafka::config::ClientConfig;
use rdkafka::message::OwnedHeaders;
use rdkafka::producer::{FutureProducer, FutureRecord};
use std::ptr::null;

pub struct KafkaClient {
    producer: Option<&'static FutureProducer>,
    topic: Option<String>
}

impl KafkaClient {
    pub fn default() -> Self {
        KafkaClient{ producer: None, topic: None }
    }

    pub fn configure(&self, server: String, topic: String) {
        let p: &FutureProducer = &ClientConfig::new()
            .set("bootstrap.servers", server)
            .set("message.timeout.ms", "5000")
            .create()
            .expect("Producer creation error");

        //self.producer = Some(p);
    }


    pub fn send_message(&self, message: String) {
        println!("model: {:?}", message);
    }
}