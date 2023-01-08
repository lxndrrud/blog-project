package mocks

import (
	"database/sql"
	"time"

	"github.com/lxndrrud/blog-project/models"
)

type SuccessPostRepoMock struct {
}

func (prm *SuccessPostRepoMock) GetPostById(idPost int64) (models.Post, error) {
	post := models.Post{
		Id:          idPost,
		IdAuthor:    1,
		AuthorLogin: "test",
		Title:       "test post",
		Annotation:  "test annotation",
		Picture:     "test picture",
		Text:        "test text",
		Views:       1,
		Approved:    true,
		Rejected:    false,
		TimeStart:   sql.NullTime{},
		TimeEnd:     sql.NullTime{},
		CreatedAt:   time.Now(),
	}
	return post, nil
}

func (prm *SuccessPostRepoMock) GetApprovedPosts() ([]models.Post, error) {
	return []models.Post{
		{
			Id:          1,
			IdAuthor:    1,
			AuthorLogin: "test",
			Title:       "test post",
			Annotation:  "test annotation",
			Picture:     "test picture",
			Text:        "test text",
			Views:       1,
			Approved:    true,
			Rejected:    false,
			TimeStart:   sql.NullTime{},
			TimeEnd:     sql.NullTime{},
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (prm *SuccessPostRepoMock) GetApprovedPost(idPost int64) (models.Post, error) {
	post := models.Post{
		Id:          idPost,
		IdAuthor:    1,
		AuthorLogin: "test",
		Title:       "test post",
		Annotation:  "test annotation",
		Picture:     "test picture",
		Text:        "test text",
		Views:       1,
		Approved:    true,
		Rejected:    false,
		TimeStart:   sql.NullTime{},
		TimeEnd:     sql.NullTime{},
		CreatedAt:   time.Now(),
	}
	return post, nil
}

func (prm *SuccessPostRepoMock) GetPostsNeedToApprove() ([]models.Post, error) {
	return []models.Post{
		{
			Id:          2,
			IdAuthor:    1,
			AuthorLogin: "test",
			Title:       "test post",
			Annotation:  "test annotation",
			Picture:     "test picture",
			Text:        "test text",
			Views:       1,
			Approved:    false,
			Rejected:    false,
			TimeStart:   sql.NullTime{},
			TimeEnd:     sql.NullTime{},
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (prm *SuccessPostRepoMock) GetPostNeedToApprove(idPost int64) (models.Post, error) {
	post := models.Post{
		Id:          idPost,
		IdAuthor:    1,
		AuthorLogin: "test",
		Title:       "test post",
		Annotation:  "test annotation",
		Picture:     "test picture",
		Text:        "test text",
		Views:       1,
		Approved:    false,
		Rejected:    false,
		TimeStart:   sql.NullTime{},
		TimeEnd:     sql.NullTime{},
		CreatedAt:   time.Now(),
	}
	return post, nil
}

func (prm *SuccessPostRepoMock) GetUserPosts(idUser int64) ([]models.Post, error) {
	return []models.Post{
		{
			Id:          2,
			IdAuthor:    idUser,
			AuthorLogin: "test",
			Title:       "test post",
			Annotation:  "test annotation",
			Picture:     "test picture",
			Text:        "test text",
			Views:       1,
			Approved:    false,
			Rejected:    false,
			TimeStart:   sql.NullTime{},
			TimeEnd:     sql.NullTime{},
			CreatedAt:   time.Now(),
		},
	}, nil
}

func (prm *SuccessPostRepoMock) InsertPost(post models.Post) error {
	return nil
}

func (prm *SuccessPostRepoMock) AddViews(idPost int64, viewsQuantity int64) error {
	return nil
}

func (prm *SuccessPostRepoMock) ApprovePost(idPost int64) error {
	return nil
}

func (prm *SuccessPostRepoMock) RejectPost(idPost int64) error {
	return nil
}

func (prm *SuccessPostRepoMock) DeletePost(idPost int64) error {
	return nil
}
