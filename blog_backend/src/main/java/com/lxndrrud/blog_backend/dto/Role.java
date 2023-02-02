package com.lxndrrud.blog_backend.dto;

import lombok.Data;

import javax.persistence.*;
import java.io.Serializable;
import java.util.List;

@Data
@Entity
public class Role implements Serializable {
    @Id
    private long id;
    @Column
    private String title;
    @ManyToMany()
    @JoinTable(
        name = "roles_permissions",
        joinColumns = @JoinColumn(name = "id_permission"),
        inverseJoinColumns = @JoinColumn(name = "id_role")
    )
    private List<Permission> permissions;
}
