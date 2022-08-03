package com.guarddog.guard_dog_video_storage.controllers;

import com.guarddog.guard_dog_video_storage.dto.VideoMetadata;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;

@RestController
public class FileUploadController {

    @PostMapping("/miniupload")
    public ResponseEntity uploadFile(@RequestPart("file") MultipartFile file, @RequestPart("metadata") VideoMetadata metadata) throws IOException {
        // Write metadata to DB
        System.out.println(metadata);

        // Stream temp file to the right directory
        
        // Write file locally for now
        String filePath = "/Users/hamdaankhalid/Desktop/guard-dog/backend/guard_dog_video_storage/src/main/temp";
        File dest = new File(filePath+"/"+file.getName());
        file.transferTo(dest);

        return ResponseEntity.ok().build();
    }
}
