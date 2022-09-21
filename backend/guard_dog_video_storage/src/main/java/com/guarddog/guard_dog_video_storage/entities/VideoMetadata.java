package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@Entity
@NoArgsConstructor
@Getter
@Setter
public class VideoMetadata {

    public VideoMetadata(int part, int duration, String filename, String url, Session session) {
        this.parentSession = session;
        this.part = part;
        this.duration = duration;
        this.filename = filename;
        this.url = url;
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", nullable = false)
    private Long id;

    @JsonBackReference
    @ManyToOne
    @JoinColumn(name = "parent_session_id", nullable = false)
    private Session parentSession;

    private int part;

    private int duration;

    private String filename;

    private String url;
}
