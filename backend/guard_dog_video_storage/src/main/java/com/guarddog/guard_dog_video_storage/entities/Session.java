package com.guarddog.guard_dog_video_storage.entities;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.util.Date;
import java.util.Set;

@Entity
@NoArgsConstructor
@Getter
@Setter
@Table(uniqueConstraints={
        @UniqueConstraint(columnNames = {"deviceName", "sessionStart"})
})
public class Session {

    public Session(int userId, String deviceName, Date sessionStart, int duration, Unit durationUnit, Set<VideoMetadata> VideoMetadatas) {
        this.userId = userId;
        this.deviceName = deviceName;
        this.sessionStart = sessionStart;
        this.duration = duration;
        this.durationUnit = durationUnit;
        this.videoMetadatas = VideoMetadatas;
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", nullable = false)
    private Long id;

    @Column(name = "user_id", nullable = false)
    private int userId;

    // deviceName + sessionStart should be unique
    @Column(nullable = false)
    private String deviceName;

    @Column(nullable = false)
    private Date sessionStart;

    @Column(nullable = false)
    private int duration;

    @Column(nullable = false)
    private Unit durationUnit;

    @OneToMany(cascade=CascadeType.ALL, mappedBy = "session")
    private Set<VideoMetadata> videoMetadatas;

}
