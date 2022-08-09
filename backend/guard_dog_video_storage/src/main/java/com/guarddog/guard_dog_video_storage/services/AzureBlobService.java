package com.guarddog.guard_dog_video_storage.services;

import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobStorageException;

public class AzureBlobService {
    private BlobServiceClient blobServiceClient = new BlobServiceClientBuilder()
            .endpoint("https://your-storage-account-url.storage.windows.net")
            .sasToken("token")
            .buildClient();

    private BlobContainerClient blobContainerClient;



    public AzureBlobService() {
        try {
            this.blobContainerClient = blobServiceClient.createBlobContainer("my-container-name");
        } catch (BlobStorageException ex) {
            // The container may already exist, so don't throw an error
            if (!ex.getErrorCode().equals(BlobErrorCode.CONTAINER_ALREADY_EXISTS)) {
                throw ex;
            }
        }
    }


    public void transferToCloud(String filepath) {
        this.blobContainerClient.getBlobClient(filepath).upload(filepath);
    }

}


