package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IPostsRepo interface {
	GetPostById(idPost int64) (models.Post, error)
	GetUserPosts(idUser int64) ([]models.Post, error)
	GetPostsNeedToApprove() ([]models.Post, error)
	GetPostNeedToApprove(idPost int64) (models.Post, error)
	GetApprovedPosts() ([]models.Post, error)
	GetApprovedPost(idPost int64) (models.Post, error)
	InsertPost(post models.Post) error
	AddViews(idPost int64, viewsQuantity int64) error
	ApprovePost(idPost int64) error
	RejectPost(idPost int64) error
	DeletePost(idPost int64) error
}

func NewPostsRepo(db *sqlx.DB) IPostsRepo {
	return &postsRepo{
		db: db,
	}
}

type postsRepo struct {
	db *sqlx.DB
}

func (c *postsRepo) GetPostById(idPost int64) (models.Post, error) {
	var post models.Post

	err := c.db.Get(
		&post,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.picture as post_picture, p.created_at as post_created_at,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE p.id = $1
		`,
		idPost)
	return post, err
}

func (c *postsRepo) GetApprovedPosts() ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.picture as post_picture, p.created_at as post_created_at,
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

func (c *postsRepo) GetApprovedPost(idPost int64) (models.Post, error) {
	var post models.Post
	err := c.db.Get(
		&post,
		`
			SELECT
				p.id as post_id, 
				p.title as post_title, p.text as post_text, p.annotation as post_annotation,
				p.picture as post_picture, p.created_at as post_created_at,
				p.views as post_views, 
				p.approved as post_approved, 
				p.rejected as post_rejected, 
				p.time_start as post_time_start, p.time_end as post_time_end,
				p.id_author as post_id_author, u.login as post_author_login
			FROM "public"."posts" p
			INNER JOIN public.users AS u ON u.id = p.id_author
			WHERE p.approved = TRUE AND p.id = $1
				AND (p.time_start IS NULL OR p.time_start <= NOW())
				AND (p.time_end IS NULL OR p.time_end > NOW()) 
		`,
		idPost)
	return post, err
}

func (c *postsRepo) GetPostsNeedToApprove() ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.picture as post_picture, p.created_at as post_created_at,
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

func (c *postsRepo) GetPostNeedToApprove(idPost int64) (models.Post, error) {
	var post models.Post

	err := c.db.Get(
		&post,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.picture as post_picture, p.created_at as post_created_at,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE p.approved = FALSE AND p.rejected = FALSE AND p.id = $1
		ORDER BY views DESC
		`,
		idPost)

	return post, err
}

func (c *postsRepo) GetUserPosts(idUser int64) ([]models.Post, error) {
	var posts []models.Post

	err := c.db.Select(
		&posts,
		`
		SELECT 
			p.id as post_id, 
			p.title as post_title, p.text as post_text, p.annotation as post_annotation,
			p.picture as post_picture, p.created_at as post_created_at,
			p.views as post_views, 
			p.approved as post_approved, 
			p.rejected as post_rejected, 
			p.time_start as post_time_start, p.time_end as post_time_end,
			p.id_author as post_id_author, u.login as post_author_login
		FROM "public"."posts" p
		INNER JOIN public.users AS u ON u.id = p.id_author
		WHERE u.id = $1
		ORDER BY post_created_at DESC
		`,
		idUser)

	return posts, err
}

func (c *postsRepo) InsertPost(post models.Post) error {
	_, err := c.db.Exec(
		`
		INSERT INTO "public"."posts"(title, annotation, text, id_author, time_start, time_end, picture) 
		VALUES
			($1, $2, $3, $4, $5, $6, $7);
		`,
		post.Title,
		post.Annotation,
		post.Text,
		post.IdAuthor,
		post.TimeStart,
		post.TimeEnd,
		post.Picture)
	return err
}

func (c *postsRepo) AddViews(idPost int64, viewsQuantity int64) error {
	_, err := c.db.Exec(
		`
			UPDATE "public"."posts"
			SET views = views + $2
			WHERE id = $1
		`,
		idPost,
		viewsQuantity)
	return err
}

func (c *postsRepo) ApprovePost(idPost int64) error {
	_, err := c.db.Exec(
		`
			UPDATE public.posts
			SET approved = TRUE
			WHERE id = $1 AND rejected = FALSE
		`,
		idPost)
	return err
}

func (c *postsRepo) RejectPost(idPost int64) error {
	_, err := c.db.Exec(
		`
			UPDATE public.posts
			SET rejected = TRUE
			WHERE id = $1 AND approved = FALSE
		`,
		idPost)
	return err
}

func (c *postsRepo) DeletePost(idPost int64) error {
	_, err := c.db.Exec(
		`
			DELETE FROM public.posts
			WHERE id = $1
		`,
		idPost)
	return err
}
