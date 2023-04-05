package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// variables de entorno

	//

	//Server
	sv := gin.Default()

	sv.GET("/", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	sv.POST("/movies", CreateMovie())
	//Run

	if err := sv.Run(":8080"); err != nil {
		panic(err)
	}
	//process

	mv := &Movie{
		ID:    1,
		Title: "The Matrix",
	}
	movies = append(movies, *mv)
}

//package handlers

func CreateMovie() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title"`
		Rating float64 `json:"rating"`
	}
	return func(ctx *gin.Context) {

		//request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"message": "Bad Request", "data": nil})
			return
		}
	}
}

//...

// package service

type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Rating float64 `json:"rating"`
}

var movies = []Movie{}
