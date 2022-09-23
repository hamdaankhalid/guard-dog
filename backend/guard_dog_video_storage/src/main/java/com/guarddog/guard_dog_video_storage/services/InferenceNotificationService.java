package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.dto.InferenceNotificationDto;
import com.guarddog.guard_dog_video_storage.dto.ModelRegistryMetadata;
import com.guarddog.guard_dog_video_storage.entities.InferenceNotification;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.jobs.InferRequest;
import com.guarddog.guard_dog_video_storage.repositories.InferenceNotificationRepository;
import org.jobrunr.scheduling.BackgroundJobRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;

@Service
public class InferenceNotificationService {

    @Autowired
    private InferenceNotificationRepository inferenceNotificationRepository;

    @Autowired
    private ModelRegistryService modelRegistryService;

    @Autowired
    private UserService userService;

    @Autowired
    private MetadataService metadataService;

    public Collection<InferenceNotificationDto> getAll(ServiceUser user) {
        int userId = user.getId();
        return inferenceNotificationRepository.findAllByServiceUserId(userId);
    }

    public boolean infer(int userId, int videoMetadataId, String details) {
        try {
            ServiceUser user = userService.getUser(userId);
            Collection<ModelRegistryMetadata> models = modelRegistryService.getModels(user);
            // for each model enqueue a job with each modelId, userId, videoMetadataId, details

            for (ModelRegistryMetadata model: models) {
                BackgroundJobRequest.enqueue(
                        new InferRequest(videoMetadataId, userId, model.getId(), details)
                );
            }

            return true;
        } catch (Exception e) {
            System.out.println("Error Enqueuing Infer Jobs: " + e.getMessage());
            return false;
        }
    }

    public void save(ServiceUser user, VideoMetadata videoMetadata, String details) {
        inferenceNotificationRepository.save(
                new InferenceNotification(-1, user, videoMetadata, details)
        );
    }
}
