package com.guarddog.guard_dog_video_storage.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.util.Date;

@AllArgsConstructor
@Getter
public class VideoMetadata {
    private String name;
    private int part;
    private String deviceName;
    private int durationInSeconds;
    private Date sessionStartTime;

    public String toString() {
        return "{" + name + ", "+ part + ", " + deviceName + ", " + durationInSeconds + ", " + sessionStartTime + "}";
    };
}
