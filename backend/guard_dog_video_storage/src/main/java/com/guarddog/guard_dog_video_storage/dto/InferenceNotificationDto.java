package com.guarddog.guard_dog_video_storage.dto;

import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;

public interface InferenceNotificationDto {
    int getId();
    String getDetails();
    VideoMetadata getVideoMetadata();
    int getServiceUserId();
}
