package com.guarddog.guard_dog_video_storage.jobs;

import com.guarddog.guard_dog_video_storage.services.InferenceNotificationService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.jobrunr.jobs.annotations.Job;
import org.jobrunr.jobs.lambdas.JobRequestHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class InferenceRequestHandler implements JobRequestHandler<InferenceRequest> {

    @Autowired
    private InferenceNotificationService inferenceNotificationService;

    @Override
    @Job(name = "InferenceJob")
    public void run(InferenceRequest inferenceRequest) throws Exception {
        inferenceNotificationService.save(
                inferenceRequest.getUserId(),
                inferenceRequest.getVideoMetadataId(),
                inferenceRequest.getDetails()
        );
    }
}
