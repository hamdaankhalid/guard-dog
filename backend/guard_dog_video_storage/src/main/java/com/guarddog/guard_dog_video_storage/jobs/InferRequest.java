package com.guarddog.guard_dog_video_storage.jobs;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.jobrunr.jobs.lambdas.JobRequest;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class InferRequest implements JobRequest {
    private int videoMetadataId;
    private int userId;
    private int modelId;
    private String details;

    public InferRequest(int videoMetadataId, int userId, int modelId) {
        this.videoMetadataId = videoMetadataId;
        this.userId = userId;
        this.modelId = modelId;
    }

    @Override
    public Class<InferRequestHandler> getJobRequestHandler() {
        return InferRequestHandler.class;
    }
}
