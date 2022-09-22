package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.dto.ModelRegistryMetadata;
import com.guarddog.guard_dog_video_storage.entities.ModelRegistry;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.services.ModelRegistryService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.security.Principal;
import java.util.Collection;

@RestController
public class ModelRegistryController {

    @Autowired
    UserService userService;

    @Autowired
    ModelRegistryService modelRegistryService;

    @PostMapping(path = "/model", consumes = { "multipart/form-data" })
    public ResponseEntity registerModel(@RequestParam("file") MultipartFile file, Principal principal) {
        ServiceUser user = userService.getUser(principal.getName());
        boolean success = modelRegistryService.register(user, file);
        if (!success) {
            return ResponseEntity.internalServerError().build();
        }
        return ResponseEntity.ok().build();
    }

    @GetMapping(path = "/model")
    public ResponseEntity<Collection<ModelRegistryMetadata>> getModels(Principal principal) {
        ServiceUser user = userService.getUser(principal.getName());
        return new ResponseEntity<>(modelRegistryService.getModels(user), HttpStatus.OK);
    }

    @GetMapping(path = "/model/{id}")
    public ResponseEntity<Resource> getModel(@PathVariable("id") int id, Principal principal) throws IOException {
        ServiceUser user = userService.getUser(principal.getName());
        ModelRegistry mr = modelRegistryService.getModel(user, id);
        if (mr == null) {
            return new ResponseEntity<>(null, HttpStatus.FORBIDDEN);
        }

        Resource file = new ByteArrayResource(mr.getModelByteData());

        if (!file.exists() && !file.isReadable()) {
            throw new RuntimeException("Could not read the file!");
        }

        return ResponseEntity.ok()
                .contentType(MediaType.APPLICATION_OCTET_STREAM)
                .contentLength(file.contentLength())
                .header(HttpHeaders.CONTENT_DISPOSITION,
                        ContentDisposition.attachment()
                                .filename(mr.getModelName())
                                .build().toString())
                .body(file);

        // return ResponseEntity.ok()
           //     .header(
             //           HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + mr.getModelName() + "\""
               // ).body(file);
    }

    @DeleteMapping(path = "/model/{id}")
    public ResponseEntity deleteModel(@PathVariable("id") int id, Principal principal) {
        ServiceUser user = userService.getUser(principal.getName());
        if (modelRegistryService.deleteModel(user, id)) {
            return ResponseEntity.ok().build();
        } else {
            return ResponseEntity.internalServerError().build();
        }
    }
}
