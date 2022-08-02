package com.guarddog.guard_dog_video_storage.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;

@RestController
public class FileUploadController {

    @PostMapping("/mini-upload")
    public ResponseEntity uploadFile(@RequestParam("file") MultipartFile file) throws IOException {

        // Write file locally for now
        String filePath = "/Users/hamdaankhalid/Desktop/guard-dog/backend/guard_dog_video_storage/src/main/temp";
        File dest = new File(filePath+"/"+file.getName());
        file.transferTo(dest);

        return ResponseEntity.ok().build();
    }
}
