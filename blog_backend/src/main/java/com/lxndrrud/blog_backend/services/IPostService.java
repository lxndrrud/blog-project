package com.lxndrrud.blog_backend.services;

import com.lxndrrud.blog_backend.dto.Post;

import java.util.List;

public interface IPostService {
    List<Post> getAll();
}
