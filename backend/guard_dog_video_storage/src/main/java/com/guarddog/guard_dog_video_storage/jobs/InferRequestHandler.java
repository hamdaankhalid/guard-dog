package com.guarddog.guard_dog_video_storage.jobs;

import ai.onnxruntime.OrtEnvironment;
import ai.onnxruntime.OrtSession;
import com.guarddog.guard_dog_video_storage.entities.ModelRegistry;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.services.InferenceNotificationService;
import com.guarddog.guard_dog_video_storage.services.MetadataService;
import com.guarddog.guard_dog_video_storage.services.ModelRegistryService;
import com.guarddog.guard_dog_video_storage.services.UserService;
import org.jobrunr.jobs.annotations.Job;
import org.jobrunr.jobs.lambdas.JobRequestHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class InferRequestHandler implements JobRequestHandler<InferRequest> {
    @Autowired
    private InferenceNotificationService inferenceNotificationService;

    @Autowired
    private MetadataService videoMetadatService;

    @Autowired
    private ModelRegistryService modelRegistryService;

    @Autowired
    private UserService userService;

    @Override
    @Job(name = "InferJob")
    public void run(InferRequest inferRequest) throws Exception {
        ServiceUser user = userService.getUser(inferRequest.getUserId());
        VideoMetadata videoMetadata = videoMetadatService.getById(inferRequest.getVideoMetadataId());

        System.out.println("LOAD MODEL HERE FOR INFERENCE");
        System.out.println("RUN INFERENCE");
        System.out.println("IF INFERENCE PRODUCES SIGNIFICANT RESULT SAVE NOTIFICATION");

        inferenceNotificationService.save(user, videoMetadata, inferRequest.getDetails());

        // Join blazor community HK

        // load model
        // ModelRegistry mr = modelRegistryService.getModel(user, inferRequest.getModelId());
        // byte[] modelAsByte = mr.getModelByteData();
        // OrtEnvironment env = OrtEnvironment.getEnvironment();
        // String tempModelFile = Path();

        // session = env.createSession(tempModelFile,new OrtSession.SessionOptions());

        // run inference

        // based on result write save notification
    }
}
