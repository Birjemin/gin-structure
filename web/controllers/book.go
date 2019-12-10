package controllers

import (
	"github.com/birjemin/gin-structure/models"
	"github.com/birjemin/gin-structure/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"time"

	//"github.com/spf13/cast"
)

type BookController struct {
	Service services.IBookService
}

var book = &BookController{Service: services.NewBookService()}

//GET http://localhost:8081/api/v1/book?page=1&size=10
func Get(g *gin.Context) {
	log.Println("get .... get")
	m := make(map[string]interface{}, 3)
	page := g.Query("page")
	size := g.Query("size")
	if page == "" {
		JsonReturn(g, -1, "page不能为空", "")
		return
	}
	if size == "" {
		JsonReturn(g, -1, "size不能为空", "")
		return
	}
	m["page"] = page
	m["size"] = size
	m["data"] = book.Service.List(m)
	JsonReturn(g, 0, "", m)
	return
}

//GET http://localhost:8081/api/v1/book/1
func GetBy(g *gin.Context) {
	uid := cast.ToUint(g.Param("id"))
	b := book.Service.Get(uid)
	JsonReturn(g, 0, "", b)
	return
}

//POST http://localhost:8081/api/v1/book
func Post(g *gin.Context) {
	b := models.Book{}
	//book.ID = cast.ToUint(r.PostFormValue("id"))
	b.Name = g.PostForm("name")
	b.Count = g.PostForm("count")
	b.Author = g.PostForm("author")
	b.Type = g.PostForm("type")
	b.CreatedAt = cast.ToUint64(time.Now().Unix())
	r :=  book.Service.Save(b)
	JsonReturn(g, 0, "", r)
	return
}

//PUT http://localhost:8081/api/v1/book/1
func PutBy(g *gin.Context) {
	b := models.Book{}
	b.ID = cast.ToUint(g.Param("id"))
	b.Name = g.PostForm("name")
	b.Count = g.PostForm("count")
	b.Author = g.PostForm("author")
	b.Type = g.PostForm("type")
	r := book.Service.Save(b)
	JsonReturn(g, 0, "", r)
	return
}

//DELETE http://localhost:8081/api/v1/book/2
func DeleteBy(g *gin.Context) {
	b := models.Book{}
	b.ID = cast.ToUint(g.Param("id"))
	r := book.Service.Del(b)
	JsonReturn(g, 0, "", r)
	return
}
