package com.guarddog.guard_dog_video_storage.jobs;

import com.guarddog.guard_dog_video_storage.services.InferenceNotificationService;
import org.jobrunr.jobs.annotations.Job;
import org.jobrunr.jobs.lambdas.JobRequestHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class InferRequestHandler implements JobRequestHandler<InferRequest> {
    @Autowired
    InferenceNotificationService inferenceNotificationService;

    @Override
    @Job(name = "InferJob")
    public void run(InferRequest inferRequest) throws Exception {
        // load model

        // run inference

        // based on result write save notification
    }
}
