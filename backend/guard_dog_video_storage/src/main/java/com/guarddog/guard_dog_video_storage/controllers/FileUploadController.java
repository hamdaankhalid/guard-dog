package com.guarddog.guard_dog_video_storage.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.guarddog.guard_dog_video_storage.dto.MiniFileUploadRequest;
import com.guarddog.guard_dog_video_storage.dto.VideoMetadataDto;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.jobs.InferenceRequest;
import com.guarddog.guard_dog_video_storage.services.CloudStoreService;
import com.guarddog.guard_dog_video_storage.services.MetadataService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import lombok.Getter;
import lombok.Setter;
import org.jobrunr.scheduling.BackgroundJobRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.security.Principal;


@RestController
public class FileUploadController {
    @Autowired
    private CloudStoreService cloudStoreService;

    @Autowired
    private MetadataService metadataService;

    @Autowired
    private UserService userService;

    @PostMapping(path = "/miniupload", consumes = { MediaType.MULTIPART_FORM_DATA_VALUE })
    public ResponseEntity uploadFile(
            @ModelAttribute MiniFileUploadRequest request,
            Principal principal
    ) throws IOException {
        VideoMetadataDto videoMetadata = new ObjectMapper().readValue(request.getMetadata(), VideoMetadataDto.class);
        String url = cloudStoreService.uploadBlob(request.getFile(), videoMetadata.getName());
        System.out.println("part: " + videoMetadata.getPart() + ", " + request.getFile().getName());
        VideoMetadata savedMetadata = metadataService.upload(videoMetadata, url);

        // TODO: Emit to kafka the URL, USER_ID,
        BackgroundJobRequest.enqueue(
            new InferenceRequest(
                    savedMetadata.getId(),
                    userService.getUser(principal.getName()).getId()
            )
        );
        return ResponseEntity.ok().build();
    }
}
