package com.guarddog.guard_dog_video_storage.services;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobStorageException;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.multipart.MultipartFile;

import java.io.*;

public class AzureBlobService {
    @Value("${AZURE_CONTAINER_SAS_TOKEN:NA}")
    private String sasToken;

    private BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
            .endpoint("https://guarddogvideostore.blob.core.windows.net/")
            .sasToken(sasToken)
            .buildClient();

    private BlobContainerClient blobContainerClient;

    public AzureBlobService() {
        try {
            this.blobContainerClient = blobServiceClient.createBlobContainer("miniuploads");
        } catch (BlobStorageException ex) {
            // The container may already exist, so don't throw an error
            if (!ex.getErrorCode().equals(BlobErrorCode.CONTAINER_ALREADY_EXISTS)) {
                throw ex;
            }
        }
    }


    public String transferToCloud(String filename, MultipartFile file) throws IOException {
        InputStream fileStream = file.getInputStream();
        BlobClient blobClient = this.blobContainerClient.getBlobClient(filename);
        blobClient.upload(fileStream, file.getSize());
        return blobClient.getBlobUrl();
    }
}


