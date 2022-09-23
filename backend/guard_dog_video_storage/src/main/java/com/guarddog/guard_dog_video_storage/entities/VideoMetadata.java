package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@Entity
@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
public class VideoMetadata {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", nullable = false)
    private int id;
    @JsonBackReference
    @ManyToOne
    @JoinColumn(name = "parent_session_id", nullable = false)
    private Session parentSession;

    private int part;

    private int duration;

    private String filename;

    private String url;
}
