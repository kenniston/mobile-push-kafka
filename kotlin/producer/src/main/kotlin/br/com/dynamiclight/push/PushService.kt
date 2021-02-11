package br.com.dynamiclight.push

import org.slf4j.LoggerFactory
import org.springframework.kafka.core.KafkaTemplate
import org.springframework.stereotype.Service

@Service
class PushService(
    val config: PushConfiguration,
    val kafkaTemplate: KafkaTemplate<String, String>
) {
    private val logger = LoggerFactory.getLogger(PushService::class.java)

    fun send(message: Push) {
        logger.info("Message: ${message.toString()} - Topic: ${config.topic}")

        config.topic?.let {
            kafkaTemplate.send(it, message.toString())
        }
    }

}