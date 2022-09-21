package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.repositories.ModelRegistryRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ModelRegistryService {

    @Autowired
    ModelRegistryRepository modelRepo;

    // upload a model with association to a user
    public void register() {
    }
}
