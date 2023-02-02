package com.lxndrrud.blog_backend.dto;

import lombok.Data;

import javax.persistence.*;
import java.io.Serializable;

@Data
@Entity(name="users")
public class User implements Serializable {
    @Id
    private long id;
    @Column
    private String login;
    @Column
    private String password;
    @Column(name = "id_role")
    private long roleID;
    @ManyToOne
    @JoinColumn(name = "id_role")
    private Role role;
}
