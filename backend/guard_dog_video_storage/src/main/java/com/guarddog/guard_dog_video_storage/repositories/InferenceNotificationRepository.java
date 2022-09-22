package com.guarddog.guard_dog_video_storage.repositories;

import com.guarddog.guard_dog_video_storage.entities.InferenceNotification;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Collection;

public interface InferenceNotificationRepository extends JpaRepository<InferenceNotification, Integer> {
    Collection<InferenceNotification> findAllByServiceUserId(int userId);
}
