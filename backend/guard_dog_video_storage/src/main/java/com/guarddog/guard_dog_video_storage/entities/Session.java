package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.fasterxml.jackson.annotation.JsonManagedReference;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.util.Date;
import java.util.Set;

@Entity
@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
@Table(uniqueConstraints={
        @UniqueConstraint(columnNames = {"deviceName", "sessionStart"})
})
public class Session {

    public Session(ServiceUser serviceUser, String deviceName, Date sessionStart, int duration, Unit durationUnit, Set<VideoMetadata> videoMetadatas) {
        this.serviceUser = serviceUser;
        this.deviceName = deviceName;
        this.sessionStart = sessionStart;
        this.duration = duration;
        this.durationUnit = durationUnit;
        this.videoMetadatas = videoMetadatas;
    }

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name = "id", nullable = false)
    private Long id;

    @JsonBackReference
    @ManyToOne
    @JoinColumn(name = "service_user_id")
    private ServiceUser serviceUser;

    @JsonManagedReference
    @OneToMany(cascade=CascadeType.ALL, mappedBy = "parentSession")
    private Set<VideoMetadata> videoMetadatas;

    // deviceName + sessionStart should be unique
    @Column(nullable = false)
    private String deviceName;

    @Column(nullable = false)
    private Date sessionStart;

    @Column(nullable = false)
    private int duration;

    @Column(nullable = false)
    private Unit durationUnit;

}
