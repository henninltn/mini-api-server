package handler

import (
	"../db"
	"github.com/gin-gonic/gin"
	"log"
)

func WordRoutes(r *gin.Engine) {
	r.GET("/glossary", getAllWords)

	r.GET("/glossary/:title", getWord)

	r.POST("/glossary", postWord)

	r.POST("/glossary/:title", updateWord)
}

func getAllWords(c *gin.Context) {
	wm := db.WordMapper{}
	res, err := wm.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, res)
}

func getWord(c *gin.Context) {
	wm := db.WordMapper{}
	res, err := wm.Find(c.Param("title"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, res)
}

func postWord(c *gin.Context) {
	wm := db.WordMapper{}
	var word db.Word
	if c.BindJSON(&word) == nil {
		wm.Word = word
		err := wm.Insert()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func updateWord(c *gin.Context) {
	wm := db.WordMapper{}
	res, err := wm.Find(c.Param("title"))
	if err != nil {
		log.Fatal(err)
	}
	wm.Word = res
	var word db.Word
	if c.BindJSON(&word) == nil {
		res.Title = word.Title
		res.Contents = word.Contents
		err = wm.Update(&res)
		if err != nil {
			log.Fatal(err)
		}
	}
}
