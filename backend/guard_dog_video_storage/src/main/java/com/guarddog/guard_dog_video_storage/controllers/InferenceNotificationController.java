package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.dto.InferenceNotificationDto;
import com.guarddog.guard_dog_video_storage.entities.InferenceNotification;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.services.InferenceNotificationService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.security.Principal;
import java.util.Collection;

@RestController
public class InferenceNotificationController {

    @Autowired
    private InferenceNotificationService inferenceNotificationService;
    @Autowired
    private UserService userService;

    @GetMapping(path = "/inferences")
    public ResponseEntity<Collection<InferenceNotificationDto>> getInferenceNotifications(Principal principal) {
        ServiceUser user = userService.getUser(principal.getName());
        Collection<InferenceNotificationDto> notifications = inferenceNotificationService.getAll(user);
        return new ResponseEntity<>(notifications, HttpStatus.OK);
    }
}
