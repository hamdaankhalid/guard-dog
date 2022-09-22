package com.guarddog.guard_dog_video_storage.repositories;

import com.guarddog.guard_dog_video_storage.dto.ModelRegistryMetadata;
import com.guarddog.guard_dog_video_storage.entities.ModelRegistry;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Collection;

@Repository
public interface ModelRegistryRepository extends JpaRepository<ModelRegistry, Integer> {
    Collection<ModelRegistryMetadata> findAllByServiceUserId(int userId);
}
