package com.lxndrrud.blog_backend.controllers;


import com.lxndrrud.blog_backend.dto.Post;
import com.lxndrrud.blog_backend.services.IPostService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/posts")
public class PostsController {

    private final IPostService postService;

    public PostsController(
        IPostService postService
    ) {
        this.postService = postService;
    }
    @GetMapping("/")
    public List<Post> getAll() {
        return this.postService.getAll();
    }

    @GetMapping("/{postId}/")
    public Optional<Post> getById(@PathVariable("postId") long id) {
        return this.postService.getById(id);
    }
}
