package com.guarddog.guard_dog_video_storage.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;

@Service
public class CloudStoreService {
    @Autowired
    AzureBlobService azureBlobService;

    public String uploadBlob(MultipartFile file, String filename) throws IOException {
        return azureBlobService.transferToCloud(filename, file);
    }
}
