package com.lxndrrud.blog_backend.services;

import com.lxndrrud.blog_backend.dto.Post;
import com.lxndrrud.blog_backend.repositories.IPostRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class PostService implements IPostService {
    private final IPostRepository postRepository;

    public PostService(
            IPostRepository postRepository
    ) {
        this.postRepository = postRepository;
    }

    public List<Post> getAll() {
        return this.postRepository.getAll();
    }

    @Override
    public Optional<Post> getById(long id) {
        return this.postRepository.getById(id);
    }
}
