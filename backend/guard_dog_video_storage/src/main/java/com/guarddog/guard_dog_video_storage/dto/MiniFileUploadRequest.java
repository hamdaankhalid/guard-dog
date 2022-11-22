package com.guarddog.guard_dog_video_storage.dto;

import lombok.Getter;
import lombok.Setter;
import org.springframework.web.multipart.MultipartFile;

@Getter
@Setter
public class MiniFileUploadRequest{
    private MultipartFile file;
    private String metadata;
};