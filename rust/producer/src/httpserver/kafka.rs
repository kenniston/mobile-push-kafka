use rdkafka::config::ClientConfig;
use rdkafka::producer::{FutureProducer};

pub struct KafkaClient {
    producer: &'static FutureProducer,
    topic: String
}

impl KafkaClient {
    pub fn default(server: String, topic: String) -> Self {
        let p: &FutureProducer = &ClientConfig::new()
            .set("bootstrap.servers", server)
            .set("message.timeout.ms", "5000")
            .create()
            .expect("Producer creation error");

        KafkaClient{ producer: p, topic }
    }

    pub fn send_message(&self, message: String) {
        println!("model: {:?}", message);
    }
}