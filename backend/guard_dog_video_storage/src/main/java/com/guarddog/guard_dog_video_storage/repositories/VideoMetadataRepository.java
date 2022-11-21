package com.guarddog.guard_dog_video_storage.repositories;

import com.guarddog.guard_dog_video_storage.entities.Session;
import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Date;

@Repository
public interface VideoMetadataRepository extends JpaRepository<VideoMetadata, Integer> {
    boolean existsByParentSessionAndPart(Session session, int part);
}
