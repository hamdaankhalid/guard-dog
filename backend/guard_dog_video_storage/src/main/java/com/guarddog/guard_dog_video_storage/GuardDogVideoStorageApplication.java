package com.guarddog.guard_dog_video_storage;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@SpringBootApplication
public class GuardDogVideoStorageApplication {
	public static void main(String[] args) {
		SpringApplication.run(GuardDogVideoStorageApplication.class, args);
	}

}
