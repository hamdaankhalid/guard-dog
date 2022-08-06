package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.entities.Session;
import com.guarddog.guard_dog_video_storage.services.MetadataService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class SessionController {
    @Autowired
    private MetadataService metadataService;

    @GetMapping("/sessions")
    public ResponseEntity<List<Session>> getSessions() {
        List<Session> sessions = metadataService.getSessions(1234);

        return new ResponseEntity<>(sessions, HttpStatus.OK);
    }
}
