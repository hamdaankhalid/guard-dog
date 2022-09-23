package com.guarddog.guard_dog_video_storage.jobs;

import com.guarddog.guard_dog_video_storage.dto.VideoMetadataDto;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.jobrunr.jobs.lambdas.JobRequest;

/**
 * Not a job but used as the type for generics for inferncing job
 */
@Getter @Setter @AllArgsConstructor @NoArgsConstructor
public class InferenceRequest implements JobRequest {
    private int videoMetadataId;
    private int userId;
    private String details;

    @Override
    public Class<InferenceRequestHandler> getJobRequestHandler() {
        return InferenceRequestHandler.class;
    }
}
