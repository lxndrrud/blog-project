package com.lxndrrud.blog_backend.repositories;

import com.lxndrrud.blog_backend.dto.Post;
import org.springframework.stereotype.Repository;

import javax.persistence.EntityManagerFactory;
import javax.persistence.TypedQuery;
import java.util.List;
import java.util.Optional;

@Repository
public class JpaPostRepository implements IPostRepository {
    private final EntityManagerFactory entityManagerFactory;

    public JpaPostRepository(
        EntityManagerFactory entityManagerFactory
    ) {
        this.entityManagerFactory = entityManagerFactory;
    }

    @Override
    public List<Post> getAll() {
        TypedQuery<Post> query = entityManagerFactory.createEntityManager().createQuery(
        "SELECT p FROM posts p",
            Post.class);
        return query.getResultList();
    }

    @Override
    public Optional<Post> getById(long id) {
        TypedQuery<Post> query = entityManagerFactory.createEntityManager().createQuery(
        "SELECT p FROM posts p WHERE p.id = :id",
            Post.class);
        query.setParameter("id", id);
        Post post = query.getSingleResult();
        return post == null ? Optional.empty() : Optional.of(post);
    }
}
