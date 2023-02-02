package com.lxndrrud.blog_backend.repositories;


import com.lxndrrud.blog_backend.dto.Post;

public interface IPostRepository {
    Iterable<Post> getAll();
}
