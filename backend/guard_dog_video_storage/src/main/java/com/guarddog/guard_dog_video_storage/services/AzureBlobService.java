package com.guarddog.guard_dog_video_storage.services;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobErrorCode;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.specialized.BlockBlobClient;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.*;
@Service
public class AzureBlobService {
    private String sasToken = System.getenv("AZURE_CONTAINER_SAS_TOKEN");

    private BlobServiceClient blobServiceClient;

    private BlobContainerClient blobContainerClient;

    public AzureBlobService() {
        this.blobServiceClient = new BlobServiceClientBuilder()
                .endpoint("https://guarddogvideostore.blob.core.windows.net/")
                .sasToken(sasToken)
                .buildClient();
        try {
            this.blobContainerClient = blobServiceClient.createBlobContainer("miniuploads");
        } catch (BlobStorageException ex) {
            // The container may already exist, so don't throw an error
            if (ex.getErrorCode().equals(BlobErrorCode.CONTAINER_ALREADY_EXISTS)) {
                this.blobContainerClient = blobServiceClient.getBlobContainerClient("miniuploads");
            } else {
                throw ex;
            }
        }
    }


    public String transferToCloud(String filename, MultipartFile file) throws IOException {
        InputStream fileStream = new BufferedInputStream(file.getInputStream());
        BlockBlobClient blobClient = this.blobContainerClient.getBlobClient(filename).getBlockBlobClient();
        blobClient.upload(fileStream, file.getSize(), true);
        return blobClient.getBlobUrl();
    }
}


