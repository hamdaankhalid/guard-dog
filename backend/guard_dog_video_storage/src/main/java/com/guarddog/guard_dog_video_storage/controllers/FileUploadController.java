package com.guarddog.guard_dog_video_storage.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.guarddog.guard_dog_video_storage.dto.VideoMetadata;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;


@RestController
public class FileUploadController {

    @PostMapping(path = "/miniupload", consumes = { "multipart/form-data" })
    public ResponseEntity uploadFile(
            @RequestParam("base64file") MultipartFile file,
            @RequestParam("metadata") String metadata
    ) throws IOException {
        // Write metadata to DB
        VideoMetadata videoMetadata = new ObjectMapper().readValue(metadata, VideoMetadata.class);
        
        // Write file locally for now
        String filePath = "/Users/hamdaankhalid/Desktop/guard-dog/backend/guard_dog_video_storage/src/main/temp";
        File dest = new File(filePath+"/"+videoMetadata.getName());
        file.transferTo(dest);

        return ResponseEntity.ok().build();
    }
}
