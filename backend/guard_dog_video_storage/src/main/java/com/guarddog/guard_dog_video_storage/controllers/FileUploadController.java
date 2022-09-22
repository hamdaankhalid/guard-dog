package com.guarddog.guard_dog_video_storage.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.guarddog.guard_dog_video_storage.dto.VideoMetadataDto;
import com.guarddog.guard_dog_video_storage.services.CloudStoreService;
import com.guarddog.guard_dog_video_storage.services.MetadataService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;


@RestController
public class FileUploadController {
    @Autowired
    private CloudStoreService cloudStoreService;

    @Autowired
    private MetadataService metadataService;

    @PostMapping(path = "/miniupload", consumes = { "multipart/form-data" })
    public ResponseEntity uploadFile(
            @RequestParam("base64file") MultipartFile file,
            @RequestParam("metadata") String metadata
    ) throws IOException {
        VideoMetadataDto videoMetadata = new ObjectMapper().readValue(metadata, VideoMetadataDto.class);
        String url = cloudStoreService.uploadBlob(file, videoMetadata.getName());
        metadataService.upload(videoMetadata, url);

        return ResponseEntity.ok().build();
    }
}
