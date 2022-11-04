package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IPostsRepo interface {
	GetUserPosts(idUser int64) ([]models.Post, error)
	GetApprovedPosts() ([]models.Post, error)
	InsertPost(post models.Post) error
}

func NewPostsRepo(db *sqlx.DB) IPostsRepo {
	return &postsRepo{
		db: db,
	}
}

type postsRepo struct {
	db *sqlx.DB
}

func (c *postsRepo) GetApprovedPosts() ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE p.approved = TRUE 
			AND (p.time_start IS NULL OR p.time_start <= NOW())
			AND (p.time_end IS NULL OR p.time_end > NOW()) 
		ORDER BY views DESC
		`)

	return posts, err
}

func (c *postsRepo) GetPostsNeedToApprove() ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE p.approved = FALSE AND p.rejected = FALSE
		ORDER BY views DESC
		`)

	return posts, err
}

func (c *postsRepo) GetUserPosts(idUser int64) ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE p.approved = TRUE 
			AND (p.time_start IS NULL OR p.time_start <= NOW())
			AND (p.time_end IS NULL OR p.time_end > NOW()) 
			AND u.id = $1
		ORDER BY views DESC
		`,
		idUser)

	return posts, err
}

func (c postsRepo) InsertPost(post models.Post) error {
	_, err := c.db.Exec(
		`
		INSERT INTO "public"."posts"(title, annotation, text, id_author, time_start, time_end) 
		VALUES
			($1, $2, $3, $4, $5, $6);
		`,
		post.Title,
		post.Annotation,
		post.Text,
		post.IdAuthor,
		post.TimeStart,
		post.TimeEnd)
	return err
}
