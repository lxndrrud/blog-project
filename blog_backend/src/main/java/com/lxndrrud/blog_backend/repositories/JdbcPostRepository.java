package com.lxndrrud.blog_backend.repositories;

import com.lxndrrud.blog_backend.dto.Post;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.Optional;

public class JdbcPostRepository implements IPostRepository {
    private final JdbcTemplate jdbcTemplate;

    public JdbcPostRepository(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }
    @Override
    public List<Post> getAll() {
        return jdbcTemplate.query(
            "select id, text, created_at from public.posts",
            this::mapRowToPost);
    }

    @Override
    public Optional<Post> getById(long id) {
        List<Post> posts = jdbcTemplate.query(
            "select id, text, created_at from public.posts where id = ?",
            this::mapRowToPost,
            id);
        return posts.size() == 0
                ? Optional.empty()
                : Optional.of(posts.get(0));
    }

    private Post mapRowToPost(ResultSet row, int rowNum) throws SQLException {
        Post post = new Post();
        post.setId(row.getLong("id"));
        post.setText(row.getString("text"));
        post.setCreatedAt(row.getDate("created_at"));
        return post;
    }

}
