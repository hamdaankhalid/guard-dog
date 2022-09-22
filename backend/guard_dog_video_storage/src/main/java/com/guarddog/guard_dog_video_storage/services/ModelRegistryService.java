package com.guarddog.guard_dog_video_storage.services;

import com.guarddog.guard_dog_video_storage.dto.ModelRegistryMetadata;
import com.guarddog.guard_dog_video_storage.entities.ModelRegistry;
import com.guarddog.guard_dog_video_storage.entities.ServiceUser;
import com.guarddog.guard_dog_video_storage.repositories.ModelRegistryRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.Collection;

@Service
public class ModelRegistryService {

    @Autowired
    ModelRegistryRepository modelRepo;

    public boolean register(ServiceUser user, MultipartFile file) {
        try {

            byte[] fileAsBytes = file.getBytes();

            modelRepo.save(new ModelRegistry(-1, user, file.getOriginalFilename(), fileAsBytes));

            return true;
        } catch (IOException e) {
            return false;
        }
    }

    public Collection<ModelRegistryMetadata> getModels(ServiceUser user) {
        return modelRepo.findAllByServiceUserId(user.getId());
    }

    public ModelRegistry getModel(ServiceUser user, int id) {
        ModelRegistry mr = modelRepo.findById(id).get();
        if (mr == null || mr.getServiceUser() != user) {
            return null;
        }
        return mr;
    }

    public boolean deleteModel(ServiceUser user, int id) {
        ModelRegistry mr = modelRepo.findById(id).get();
        if (mr == null || mr.getServiceUser() != user) {
            return false;
        }
        try {
            modelRepo.deleteById(id);
            return true;
        } catch(Exception e) {
            return false;
        }
    }
}
