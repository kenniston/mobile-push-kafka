package br.com.dynamiclight.push

import org.springframework.boot.context.properties.ConfigurationProperties
import org.springframework.stereotype.Component

@Component
@ConfigurationProperties(prefix = "push")
data class PushConfiguration (
    /**
     * Kafka topic to send push messages
     */
    var topic: String? = null
)

