package br.com.dynamiclight.push;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class PushService {
    private final Logger logger = LoggerFactory.getLogger(PushService.class);
    private final PushConfiguration config;
    private final KafkaTemplate<String, String> kafkaTemplate;

    public PushService(PushConfiguration config, KafkaTemplate<String, String> kafkaTemplate) {
        this.config = config;
        this.kafkaTemplate = kafkaTemplate;
    }

    public void send(Push message) {
        logger.info("Message: " + message.toString() + " - Topic: " + config.getTopic());
        if (config.getTopic() != null) {
            kafkaTemplate.send(config.getTopic(), message.toString());
        }
    }

}
