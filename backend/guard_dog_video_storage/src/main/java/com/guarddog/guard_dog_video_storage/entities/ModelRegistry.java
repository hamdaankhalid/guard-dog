package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@Entity @Getter @Setter @AllArgsConstructor @NoArgsConstructor
public class ModelRegistry {
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE)
    @Column(name = "id", nullable = false)
    private int id;

    @JsonBackReference
    @ManyToOne
    @JoinColumn(name = "service_user_id")
    private ServiceUser serviceUser;

    @Column
    private String modelName;

    @Lob
    private byte[] modelByteData;
}
