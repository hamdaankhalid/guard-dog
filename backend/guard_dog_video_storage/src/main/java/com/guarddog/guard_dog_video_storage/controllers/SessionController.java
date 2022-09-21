package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.Session;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.repositories.UserRepository;
import com.guarddog.guard_dog_video_storage.services.MetadataService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.security.Principal;
import java.util.Collection;


@RestController
public class SessionController {
    @Autowired
    private UserService userService;

    @GetMapping("/sessions")
    public ResponseEntity<Collection<Session>> getSessions(Principal principal) {
        String requester = principal.getName();
        ServiceUser user = userService.getUser(requester);
        Collection<Session> sessions = user.getSessions();
        return new ResponseEntity<>(sessions, HttpStatus.OK);
    }
}
