package com.lxndrrud.blog_backend.repositories;


import com.lxndrrud.blog_backend.dto.Post;

import java.util.List;
import java.util.Optional;

public interface IPostRepository {
    List<Post> getAll();
    Optional<Post> getById(long id);
}
