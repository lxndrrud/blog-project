package com.lxndrrud.blog_backend.services;

import com.lxndrrud.blog_backend.dto.Post;

import java.util.List;
import java.util.Optional;

public interface IPostService {
    List<Post> getAll();
    Optional<Post> getById(long id);
}
