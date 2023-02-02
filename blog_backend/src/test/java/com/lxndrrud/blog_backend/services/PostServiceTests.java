package com.lxndrrud.blog_backend.services;

import com.lxndrrud.blog_backend.dto.Post;
import com.lxndrrud.blog_backend.repositories.IPostRepository;
import com.lxndrrud.blog_backend.repositories.JdbcPostRepository;

import org.junit.Test;

import org.mockito.Mockito;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;

public class PostServiceTests {
    @Test
    public void testGetAllPosts() {
        IPostRepository jdbcPostRepo = Mockito.mock(JdbcPostRepository.class);
        ArrayList<Post> arrayList = new ArrayList<>();
        Post post = new Post();
        arrayList.add(post);
        Mockito.when(jdbcPostRepo.getAll()).thenReturn(arrayList);

        PostService postService = new PostService(jdbcPostRepo);

        List<Post> result = postService.getAll();

        assertEquals(result.size(), 1);
        assertEquals(result, arrayList);
    }
    @Test
    public void testGetExistingPostById() {
        IPostRepository jdbcPostRepo = Mockito.mock(JdbcPostRepository.class);
        Post post = new Post();
        post.setId(1);
        Mockito.when(jdbcPostRepo.getById(1)).thenReturn(Optional.of(post));

        PostService postService = new PostService(jdbcPostRepo);

        Optional<Post> result = postService.getById(1);

        assertTrue(result.isPresent());
        assertEquals(result.get(), post);
    }

    @Test
    public void testGetEmptyValueById() {
        IPostRepository jdbcPostRepo = Mockito.mock(JdbcPostRepository.class);
        Mockito.when(jdbcPostRepo.getById(1)).thenReturn(Optional.empty());

        PostService postService = new PostService(jdbcPostRepo);

        Optional<Post> result = postService.getById(1);

        assertFalse(result.isPresent());
    }
}
