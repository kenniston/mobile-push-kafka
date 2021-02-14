package br.com.dynamiclight;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.core.env.Environment;
import org.springframework.kafka.annotation.EnableKafka;

@EnableKafka
@SpringBootApplication
public class ProducerApplication implements ApplicationRunner {
	private final Logger logger = LoggerFactory.getLogger(ProducerApplication.class);
	private final Environment env;

	public ProducerApplication(Environment env) {
		this.env = env;
	}

	@Override
	public void run(ApplicationArguments args) throws Exception {
		String serverPort = env.getProperty("server.port");
		logger.info("Starting Server on port " + serverPort);
	}

	public static void main(String[] args) {
		SpringApplication.run(ProducerApplication.class, args);
	}

}
