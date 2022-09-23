package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.dto.VideoMetadataDto;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.entities.Session;
import com.guarddog.guard_dog_video_storage.entities.Unit;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import com.guarddog.guard_dog_video_storage.repositories.SessionRepository;
import com.guarddog.guard_dog_video_storage.repositories.UserRepository;
import com.guarddog.guard_dog_video_storage.repositories.VideoMetadataRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class MetadataService {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private SessionRepository sessionRepository;

    @Autowired
    private VideoMetadataRepository videoMetadataRepository;

    public VideoMetadata upload(VideoMetadataDto videoMetadataDto, String url) {
        // If the sessionStart + deviceName doesn't already exist create a new session
        String deviceName = videoMetadataDto.getDeviceName();
        Date sessionStart = videoMetadataDto.getSessionStart();
        int durationSeconds = videoMetadataDto.getDurationInSeconds();
        Optional<ServiceUser> user = userRepository.findById(videoMetadataDto.getUserId());

        Session session;
        boolean sessionExists = sessionRepository.existsByDeviceNameAndSessionStart(deviceName, sessionStart);
        if (sessionExists){
            session = sessionRepository.findOneByDeviceNameAndSessionStart(deviceName, sessionStart);
        } else {
            session = sessionRepository.save(new Session(user.get(), deviceName, sessionStart, durationSeconds, Unit.SECONDS, new HashSet<>()));
        }

        // persist videoMetadata for session and associate the above session with it
        VideoMetadata videoMetadata = new VideoMetadata(
                -1,
                session,
                videoMetadataDto.getPart(),
                durationSeconds,
                deviceName,
                url
        );


        VideoMetadata savedMetaData = videoMetadataRepository.save(videoMetadata);
        session.getVideoMetadatas().add(videoMetadata);
        sessionRepository.save(session);
        return savedMetaData;
    }


    public VideoMetadata getById(int id) {
        return videoMetadataRepository.findById(id).get();
    }
}
