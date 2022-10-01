package com.guarddog.guard_dog_video_storage.jobs;

import ai.onnxruntime.OrtEnvironment;
import ai.onnxruntime.OrtException;
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
        ModelRegistry mr = modelRegistryService.getModel(user, inferRequest.getModelId());

        // load model
        OrtSession session = createSession(mr);

        // TODO: run inference from session

        // save inference
        inferenceNotificationService.save(user, videoMetadata, "RESULT FROM INFERENCE");
    }

    private OrtSession createSession(ModelRegistry mr) throws OrtException {
        byte[] modelAsByte = mr.getModelByteData();
        OrtEnvironment env = OrtEnvironment.getEnvironment();
        OrtSession session = env.createSession(modelAsByte);
        return session;
    }
}
