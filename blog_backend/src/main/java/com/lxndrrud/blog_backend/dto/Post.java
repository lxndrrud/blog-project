package com.lxndrrud.blog_backend.dto;

import lombok.Data;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import java.io.Serializable;
import java.util.Date;

@Data
@Entity
public class Post implements Serializable {
    @Id
    private long id;
    @Column
    private String text;
    @Column(name = "created_at")
    private Date createdAt;
}
