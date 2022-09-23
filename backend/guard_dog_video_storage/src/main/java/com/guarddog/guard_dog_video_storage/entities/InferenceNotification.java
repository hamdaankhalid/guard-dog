package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class InferenceNotification {
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE)
    @Column(name = "id", nullable = false)
    private int id;

    @ManyToOne
    @JoinColumn(name = "service_user_id")
    private ServiceUser serviceUser;

    @ManyToOne
    @JoinColumn(name = "video_metadata_id")
    private VideoMetadata videoMetadata;

    @Column
    private String details;
}
