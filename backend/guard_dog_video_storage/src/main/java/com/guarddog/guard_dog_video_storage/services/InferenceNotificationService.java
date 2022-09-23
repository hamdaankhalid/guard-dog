package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.dto.InferenceNotificationDto;
import com.guarddog.guard_dog_video_storage.entities.InferenceNotification;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.repositories.InferenceNotificationRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collection;

@Service
public class InferenceNotificationService {

    @Autowired
    private InferenceNotificationRepository inferenceNotificationRepository;

    @Autowired
    private UserService userService;

    @Autowired
    private MetadataService metadataService;

    public Collection<InferenceNotificationDto> getAll(ServiceUser user) {
        int userId = user.getId();
        return inferenceNotificationRepository.findAllByServiceUserId(userId);
    }

    public boolean save(int userId, int videoMetadataId, String details) {
        try {
            ServiceUser user = userService.getUser(userId);
            VideoMetadata videoMetadata = metadataService.getById(videoMetadataId);
            inferenceNotificationRepository.save(
                new InferenceNotification(
                    -1,
                    user,
                    videoMetadata,
                    details
                )
            );
            return true;
        } catch (Exception e) {
            System.out.println("Error Saving Inference Notification: " + e.getMessage());
            return false;
        }
    }
}
