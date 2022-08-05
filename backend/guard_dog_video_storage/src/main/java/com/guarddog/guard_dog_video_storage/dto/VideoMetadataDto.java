package com.guarddog.guard_dog_video_storage.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.Date;

@AllArgsConstructor
@NoArgsConstructor
@Setter
@Getter
public class VideoMetadataDto {
    private String name;
    private int part;
    private String deviceName;
    private int durationInSeconds;
    private Date sessionStart;

    public String toString() {
        return "{" + name + ", "+ part + ", " + deviceName + ", " + durationInSeconds + ", " + sessionStart + "}";
    };
}
