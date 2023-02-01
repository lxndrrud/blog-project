package com.lxndrrud.blog_backend.controllers;


import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/posts")
public class PostsController {
    @GetMapping("/")
    public List<String> getAll() {
        ArrayList<String> resultList = new ArrayList<>();
        resultList.add("kek");
        return resultList;
    }
}
