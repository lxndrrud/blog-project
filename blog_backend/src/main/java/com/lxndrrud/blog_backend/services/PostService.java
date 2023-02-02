package com.lxndrrud.blog_backend.services;

import com.lxndrrud.blog_backend.dto.Post;
import com.lxndrrud.blog_backend.repositories.IPostRepository;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

@Service
public class PostService implements IPostService {
    private final IPostRepository postRepository;

    public PostService(
            IPostRepository postRepository
    ) {
        this.postRepository = postRepository;
    }

    public List<Post> getAll() {
        Iterable<Post> posts = this.postRepository.getAll();
        return new ArrayList<Post>((Collection) posts);
    }
}
