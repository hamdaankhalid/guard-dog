package com.guarddog.guard_dog_video_storage.controllers;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HealthController {

    @GetMapping("/health")
    public ResponseEntity<String> health() {
        System.out.println("Server is up and healthy :)");
        return new ResponseEntity<>("Hello World! I am healthy :)", HttpStatus.OK);
    }
}
