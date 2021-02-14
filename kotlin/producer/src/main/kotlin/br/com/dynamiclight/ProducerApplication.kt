package br.com.dynamiclight

import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.ApplicationArguments
import org.springframework.boot.ApplicationRunner
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.core.env.Environment
import org.springframework.kafka.annotation.EnableKafka

@EnableKafka
@SpringBootApplication
class ProducerApplication: ApplicationRunner {
	private val logger = LoggerFactory.getLogger(ProducerApplication::class.java)

	@Autowired
	private val env: Environment? = null

	override fun run(args: ApplicationArguments) {
		val serverPort = env?.getProperty("server.port")
		logger.info("Starting Server on port $serverPort")
	}
}

fun main(args: Array<String>) {
	runApplication<ProducerApplication>(*args) {

	}
}
