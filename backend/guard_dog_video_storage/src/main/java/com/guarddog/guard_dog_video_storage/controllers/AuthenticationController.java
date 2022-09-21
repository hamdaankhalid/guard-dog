package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.servlet.support.ServletUriComponentsBuilder;

import java.net.URI;
import java.security.Principal;
import java.util.List;

@RestController
public class AuthenticationController {

    @Autowired
    private UserService userService;

    @GetMapping("/users")
    public ResponseEntity<List<ServiceUser>> getUsers() {
        return ResponseEntity.ok().body(userService.getUsers());
    }

    @GetMapping("/identity")
    public ResponseEntity<ServiceUser> getIdentity(Principal principal) {
        String requester = principal.getName();
        ServiceUser user = userService.getUser(requester);

        return ResponseEntity.ok().body(user);
    }

    @PostMapping("/signup")
    public ResponseEntity<ServiceUser> saveUser(@RequestBody ServiceUser user) {
        if(userService.getUser(user.getEmail()) != null) {
            return ResponseEntity.badRequest().body(null);
        }

        URI uri = URI.create(ServletUriComponentsBuilder.fromCurrentContextPath().path("/signup").toUriString());
        user.setRole("USER");
        return ResponseEntity.created(uri).body(userService.save(user));
    }
}
