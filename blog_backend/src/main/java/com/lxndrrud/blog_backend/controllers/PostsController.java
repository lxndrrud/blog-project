package com.lxndrrud.blog_backend.controllers;


import com.lxndrrud.blog_backend.dto.Post;
import com.lxndrrud.blog_backend.services.IPostService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

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
}
