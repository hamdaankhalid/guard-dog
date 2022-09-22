package com.guarddog.guard_dog_video_storage.services;

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

    public Collection<InferenceNotification> getAll(ServiceUser user) {
        return inferenceNotificationRepository.findAllByServiceUserId(user.getId());
    }

    public boolean save(ServiceUser user, VideoMetadata videoMetadata, String details) {
        try {
            inferenceNotificationRepository.save(new InferenceNotification(-1, user, videoMetadata, details));
            return true;
        } catch (Exception e) {
            System.out.println("Error Saving Inference Notification: " + e.getMessage());
            return false;
        }
    }
}
