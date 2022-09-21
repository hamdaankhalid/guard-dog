package com.guarddog.guard_dog_video_storage.repositories;

import com.guarddog.guard_dog_video_storage.entities.Session;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Date;

@Repository
public interface SessionRepository extends JpaRepository<Session, Integer> {
    Session findOneByDeviceNameAndSessionStart(String deviceName, Date sessionStart);
    boolean existsByDeviceNameAndSessionStart(String deviceName, Date sessionStart);
}
