package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lxndrrud/blog-project/models"
)

type IFileProcessor interface {
	SaveFileOnDisk(file multipart.File, header *multipart.FileHeader) (string, models.IError)
	DeleteFile(filename string) models.IError
}

func NewFileProcessor() IFileProcessor {
	return new(fileProcessor)
}

type fileProcessor struct {
}

func (c *fileProcessor) SaveFileOnDisk(file multipart.File, header *multipart.FileHeader) (string, models.IError) {
	// Считать содержимое файла
	contents := make([]byte, 500000)
	_, err := file.Read(contents)
	if err != nil {
		fmt.Println(err)
		return "", models.NewError(http.StatusInternalServerError, "Ошибка во время сохранения файла!")
	}
	// Сгенерировать хэш
	h := sha1.New()
	_, err = h.Write([]byte(header.Filename))
	if err != nil {
		fmt.Println(err)
		return "", models.NewError(http.StatusInternalServerError, "Ошибка во время сохранения файла!")
	}
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	filename := sha1_hash + time.Now().Format("2006-01-02_15:04:05") + "." + strings.Split(header.Filename, ".")[1]
	// Создать файл и получить указатель на него
	filePointer, err := os.Create(filepath.Clean("./storage/" + filename))
	if err != nil {
		fmt.Println(err)
		return "", models.NewError(http.StatusInternalServerError, "Ошибка во время сохранения файла!")
	}
	// Записать в файл содержимое
	_, err = filePointer.Write(contents)
	if err != nil {
		fmt.Println(err)
		return "", models.NewError(http.StatusInternalServerError, "Ошибка во время сохранения файла!")
	}
	filePointer.Close()
	return filename, nil
}

func (c *fileProcessor) DeleteFile(filename string) models.IError {
	err := os.Remove(filepath.Clean("./storage" + filename))
	if err != nil {
		return models.NewError(http.StatusInternalServerError, "Не удалось удалить постер!")
	}
	return nil
}
