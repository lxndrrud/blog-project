package services

import (
	"database/sql"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/lxndrrud/blog-project/models"
)

// Интерфейсы зависимостей сервиса
type IPostServicePostRepo interface {
	GetPostById(idPost int64) (models.Post, error)
	GetUserPosts(idUser int64) ([]models.Post, error)
	GetApprovedPosts() ([]models.Post, error)
	GetApprovedPost(idPost int64) (models.Post, error)
	GetPostsNeedToApprove() ([]models.Post, error)
	GetPostNeedToApprove(idPost int64) (models.Post, error)
	InsertPost(post models.Post) error
	AddViews(idPost int64, viewsQuantity int64) error
	ApprovePost(idPost int64) error
	RejectPost(idPost int64) error
	DeletePost(idPost int64) error
}

type IPostServiceUserRepo interface {
	GetUserById(idUser int64) (models.User, error)
}

type IPostServicePermissionInfra interface {
	GetPermissionsList(token string) ([]models.Permission, models.IError)
	GetUserIdByToken(token string) (int64, models.IError)
}

type IPostServiceFileProcessor interface {
	SaveFileOnDisk(file multipart.File, header *multipart.FileHeader) (string, models.IError)
	DeleteFile(filename string) models.IError
}

type IPostServicePermissionChecker interface {
	CanCreatePosts(permissions []models.Permission) bool
	CanModeratePosts(permissions []models.Permission) bool
}

type IPostService interface {
	GetUserPosts(token string) ([]models.Post, models.IError)
	GetApprovedPosts() ([]models.Post, models.IError)
	GetApprovedPost(idPost int64) (models.Post, models.IError)
	GetPostsNeedToApprove(token string) ([]models.Post, models.IError)
	GetPostNeedToApprove(token string, idPost int64) (models.Post, models.IError)
	CreatePost(title, text, annotation, token string,
		timeStart *time.Time, timeEnd *time.Time, pictureFile multipart.File,
		pictureHeader *multipart.FileHeader) models.IError
	ApprovePost(token string, idPost int64) models.IError
	RejectPost(token string, idPost int64) models.IError
	DeletePost(token string, idPost int64) models.IError
}

type postService struct {
	postRepo          IPostServicePostRepo
	userRepo          IPostServiceUserRepo
	permissionInfra   IPostServicePermissionInfra
	fileProcessor     IPostServiceFileProcessor
	permissionChecker IPostServicePermissionChecker
}

func NewPostService(
	postRepo IPostServicePostRepo,
	userRepo IPostServiceUserRepo,
	permissionInfra IPostServicePermissionInfra,
	fileProcessor IPostServiceFileProcessor,
	permissionChecker IPostServicePermissionChecker,
) IPostService {
	return &postService{
		postRepo:          postRepo,
		userRepo:          userRepo,
		permissionInfra:   permissionInfra,
		fileProcessor:     fileProcessor,
		permissionChecker: permissionChecker,
	}
}

func (c postService) GetUserPosts(token string) ([]models.Post, models.IError) {
	idUser, errService := c.permissionInfra.GetUserIdByToken(token)
	if errService != nil {
		return []models.Post{}, errService
	}
	posts, err := c.postRepo.GetUserPosts(idUser)
	if err != nil {
		return []models.Post{}, models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка"+err.Error())
	}
	return posts, nil
}

func (c postService) GetApprovedPosts() ([]models.Post, models.IError) {
	posts, err := c.postRepo.GetApprovedPosts()
	if err != nil {
		return posts, models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return posts, nil
}

func (c postService) GetApprovedPost(idPost int64) (models.Post, models.IError) {
	// Получить пост
	post, err := c.postRepo.GetApprovedPost(idPost)
	if err == sql.ErrNoRows {
		return models.Post{}, models.NewError(http.StatusNotFound, "Пост не найден!")
	}
	if err != nil {
		return models.Post{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	// Обновить количество просмотров
	err = c.postRepo.AddViews(post.Id, 1)
	if err != nil {
		return models.Post{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	post.Views += 1
	return post, nil
}

func (c postService) GetPostsNeedToApprove(token string) ([]models.Post, models.IError) {
	permissions, errService := c.permissionInfra.GetPermissionsList(token)
	if errService != nil {
		return []models.Post{}, errService
	}
	if !c.permissionChecker.CanModeratePosts(permissions) {
		return []models.Post{}, models.NewError(http.StatusForbidden,
			"Вам запрещено модерировать посты!")
	}
	posts, err := c.postRepo.GetPostsNeedToApprove()
	if err != nil {
		return []models.Post{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return posts, nil
}

func (c postService) GetPostNeedToApprove(token string, idPost int64) (models.Post, models.IError) {
	permissions, errService := c.permissionInfra.GetPermissionsList(token)
	if errService != nil {
		return models.Post{}, errService
	}
	if !c.permissionChecker.CanModeratePosts(permissions) {
		return models.Post{}, models.NewError(http.StatusForbidden,
			"Вам запрещено модерировать посты!")
	}
	post, err := c.postRepo.GetPostNeedToApprove(idPost)
	if err == sql.ErrNoRows {
		return models.Post{}, models.NewError(http.StatusNotFound, "Непроверенный пост с данным идентификатором не найден!")
	}
	if err != nil {
		return models.Post{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return post, nil
}

func (c postService) CreatePost(title, text, annotation, token string,
	timeStart *time.Time, timeEnd *time.Time, pictureFile multipart.File,
	pictureHeader *multipart.FileHeader) models.IError {
	idUser, errService := c.permissionInfra.GetUserIdByToken(token)
	if errService != nil {
		return errService
	}
	// Получить разрешения пользователя по информации с токена
	permissions, errService := c.permissionInfra.GetPermissionsList(token)
	if errService != nil {
		return errService
	}
	// Проверить разрешение на создание постов
	if !c.permissionChecker.CanCreatePosts(permissions) {
		return models.NewError(http.StatusForbidden, "Вам запрещено создавать посты!")
	}
	// Сохранить файл на диск и получить имя файла для записи в базу
	filename, errService := c.fileProcessor.SaveFileOnDisk(pictureFile, pictureHeader)
	if errService != nil {
		return errService
	}

	dbTimeStart := sql.NullTime{}
	if timeStart != nil {
		dbTimeStart = sql.NullTime{
			Valid: true,
			Time:  *timeStart,
		}
	}
	dbTimeEnd := sql.NullTime{}
	if timeEnd != nil {
		dbTimeEnd = sql.NullTime{
			Valid: true,
			Time:  *timeEnd,
		}
	}
	post := models.Post{
		Title:      title,
		Annotation: annotation,
		Text:       text,
		TimeStart:  dbTimeStart,
		TimeEnd:    dbTimeEnd,
		IdAuthor:   idUser,
		Picture:    "/" + filename,
	}
	err := c.postRepo.InsertPost(post)
	if err != nil {
		return models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return nil
}

func (c postService) ApprovePost(token string, idPost int64) models.IError {
	permissions, errService := c.permissionInfra.GetPermissionsList(token)
	if errService != nil {
		return errService
	}
	if !c.permissionChecker.CanModeratePosts(permissions) {
		return models.NewError(http.StatusForbidden, "Вам запрещено модерировать посты!")
	}
	err := c.postRepo.ApprovePost(idPost)
	if err != nil {
		return models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return nil
}

func (c postService) RejectPost(token string, idPost int64) models.IError {
	permissions, errService := c.permissionInfra.GetPermissionsList(token)
	if errService != nil {
		return errService
	}
	if !c.permissionChecker.CanModeratePosts(permissions) {
		return models.NewError(http.StatusForbidden, "Вам запрещено модерировать посты!")
	}
	err := c.postRepo.RejectPost(idPost)
	if err != nil {
		return models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return nil
}

func (c postService) DeletePost(token string, idPost int64) models.IError {
	idUser, errService := c.permissionInfra.GetUserIdByToken(token)
	if errService != nil {
		return errService
	}
	post, err := c.postRepo.GetPostById(idPost)
	if err != nil {
		return models.NewError(http.StatusNotFound, "Пост не найден!")
	}
	if post.IdAuthor != idUser {
		return models.NewError(http.StatusForbidden, "Вы не являетесь автором поста!")
	}
	errService = c.fileProcessor.DeleteFile(post.Picture)
	if errService != nil {
		return errService
	}
	err = c.postRepo.DeletePost(idPost)
	if err != nil {
		return models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return nil
}
