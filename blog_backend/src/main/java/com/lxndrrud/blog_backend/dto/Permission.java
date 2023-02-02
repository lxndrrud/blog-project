package com.lxndrrud.blog_backend.dto;

import lombok.Data;

import javax.persistence.*;
import java.io.Serializable;
import java.util.List;

@Data
@Entity(name = "permissions")
public class Permission implements Serializable {
    @Id
    private long id;
    @Column
    private String title;
    @Column
    private String code;
    @ManyToMany
    @JoinTable(
        name = "roles_permissions",
        joinColumns = @JoinColumn(name = "id_role"),
        inverseJoinColumns = @JoinColumn(name = "id_permissions")
    )
    private List<Role> roles;
}
