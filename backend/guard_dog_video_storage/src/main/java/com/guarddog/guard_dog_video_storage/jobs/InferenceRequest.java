package com.guarddog.guard_dog_video_storage.jobs;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.jobrunr.jobs.lambdas.JobRequest;

/**
 * Not a job but used as the type for generics for inference job
 */
@Getter @Setter @AllArgsConstructor @NoArgsConstructor
public class InferenceRequest implements JobRequest {
    private int videoMetadataId;
    private int userId;
    private String details;

    public InferenceRequest(int videoMetadataId, int userId) {
        this.videoMetadataId = videoMetadataId;
        this.userId = userId;
    }

    @Override
    public Class<InferenceRequestHandler> getJobRequestHandler() {
        return InferenceRequestHandler.class;
    }
}
