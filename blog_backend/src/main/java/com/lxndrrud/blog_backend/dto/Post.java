package com.lxndrrud.blog_backend.dto;

import lombok.Data;

import java.io.Serializable;
import java.util.Date;

@Data
public class Post implements Serializable {
    private long id;
    private String text;
    private Date createdAt;
}
