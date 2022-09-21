package com.guarddog.guard_dog_video_storage.entities;

import com.fasterxml.jackson.annotation.JsonManagedReference;
import com.sun.istack.NotNull;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.util.Set;

@Entity @Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class ServiceUser {
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE)
    @Column(name = "id", nullable = false)
    private int id;

    @NotNull
    private String email;

    @NotNull
    private String password;

    @Column
    private String role;

    @JsonManagedReference
    @OneToMany(cascade=CascadeType.ALL, mappedBy = "serviceUser")
    private Set<Session> sessions;

    @JsonManagedReference
    @OneToMany(cascade=CascadeType.ALL, mappedBy = "serviceUser")
    private Set<ModelRegistry> models;
}
