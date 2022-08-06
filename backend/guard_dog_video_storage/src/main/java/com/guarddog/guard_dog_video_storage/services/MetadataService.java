package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.dto.VideoMetadataDto;
import com.guarddog.guard_dog_video_storage.entities.Session;
import com.guarddog.guard_dog_video_storage.entities.Unit;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.repositories.SessionRepository;
import com.guarddog.guard_dog_video_storage.repositories.VideoMetadataRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class MetadataService {
    @Autowired
    private SessionRepository sessionRepository;

    @Autowired
    private VideoMetadataRepository videoMetadataRepository;

    public void upload(VideoMetadataDto videoMetadataDto) {
        // If the sessionStart + deviceName doesn't already exist create a new session
        String deviceName = videoMetadataDto.getDeviceName();
        Date sessionStart = videoMetadataDto.getSessionStart();
        int durationSeconds = videoMetadataDto.getDurationInSeconds();

        Session session;
        boolean sessionExists = sessionRepository.existsByDeviceNameAndSessionStart(deviceName, sessionStart);
        if (sessionExists){
            session = sessionRepository.findOneByDeviceNameAndSessionStart(deviceName, sessionStart);
        } else {
            session = sessionRepository.save(new Session(1234, deviceName, sessionStart, durationSeconds, Unit.SECONDS, new HashSet<>()));
        }

        // persist videoMetadata for session and associate the above session with it
        VideoMetadata videoMetadata = new VideoMetadata(
                videoMetadataDto.getPart(),
                durationSeconds,
                deviceName,
                session
        );
        videoMetadataRepository.save(videoMetadata);
        session.getVideoMetadatas().add(videoMetadata);
        sessionRepository.save(session);
    }


    public List<Session> getSessions(int userId) {
        return sessionRepository.findByUserId(userId);
    }
}
