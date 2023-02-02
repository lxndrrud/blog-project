package com.lxndrrud.blog_backend.repositories;

import com.lxndrrud.blog_backend.dto.Post;
import org.springframework.data.repository.CrudRepository;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.metadata.GenericTableMetaDataProvider;
import org.springframework.stereotype.Repository;

import java.sql.ResultSet;
import java.sql.SQLException;

@Repository
public class JdbcPostRepository implements IPostRepository {
    private final JdbcTemplate jdbcTemplate;

    public JdbcPostRepository(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }
    @Override
    public Iterable<Post> getAll() {
        return jdbcTemplate.query(
                "select id, text, created_at from public.posts",
                this::mapRowToPost);
    }

    private Post mapRowToPost(ResultSet row, int rowNum) throws SQLException {
        Post post = new Post();
        post.setId(row.getLong("id"));
        post.setText(row.getString("text"));
        post.setCreatedAt(row.getDate("created_at"));
        return post;
    }

}
