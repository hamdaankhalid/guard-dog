package com.guarddog.guard_dog_video_storage.entities;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@Entity
@NoArgsConstructor
@Getter
@Setter
public class VideoMetadata {

    public VideoMetadata(int part, int duration, String filename, Session session, String url) {
        this.part = part;
        this.duration = duration;
        this.filename = filename;
        this.session = session;
        this.url = url;
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", nullable = false)
    private Long id;

    private int part;

    private int duration;

    private String filename;

    @ManyToOne(fetch = FetchType.LAZY)
    private Session session;

    private String url;
}
