package com.guarddog.guard_dog_video_storage.services;

import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;

@Service
public class CloudStoreService {
    private static String FILE_PATH = "/Users/hamdaankhalid/Desktop/guard-dog/backend/guard_dog_video_storage/src/main/temp";

    public void uploadBlob(MultipartFile file, String filename) throws IOException {
        // Write file locally for now
        File dest = new File(FILE_PATH+"/"+filename);
        file.transferTo(dest);
    }
}
