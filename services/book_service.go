package services

import (
	"encoding/json"
	"github.com/birjemin/gin-structure/cache"
	"github.com/birjemin/gin-structure/models"
	"github.com/birjemin/gin-structure/repos"
)

type IBookService interface {
	List(m map[string]interface{}) map[string]interface{}
	Save(book models.Book) bool
	Get(id uint) string
	Del(book models.Book) bool
}

type bookService struct {
	repoB repos.IBookRepository
}

func NewBookService() IBookService {
	return &bookService{repoB: repos.NewBookRepository()}
}

func (b bookService) List(m map[string]interface{}) map[string]interface{} {
	total, books := b.repoB.List(m)
	maps := make(map[string]interface{}, 2)
	maps["Total"] = total
	maps["List"] = books
	return maps
}

func (b bookService) Save(book models.Book) bool {
	err := b.repoB.Save(book)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (b bookService) Get(id uint) string {
	res, _ := cache.Get("tian")
	if len(res) != 0 {
		return res
	}
	book, err := b.repoB.Get(id)
	str, _ := json.Marshal(book)
	cache.Set("tian", string(str), 600)
	if err != nil {
		return ""
	} else {
		return string(str)
	}
}

func (b bookService) Del(book models.Book) bool {
	err := b.repoB.Del(book)
	if err != nil {
		return false
	} else {
		return true
	}
}
