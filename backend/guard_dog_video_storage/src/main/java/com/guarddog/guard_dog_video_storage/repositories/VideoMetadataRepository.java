package com.guarddog.guard_dog_video_storage.repositories;

import com.guarddog.guard_dog_video_storage.entities.VideoMetadata;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface VideoMetadataRepository extends JpaRepository<VideoMetadata, Integer> {
}
