package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()

	// memanggil konstruktor service categori
	//defendensi : categori repo, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate) // proses injek ke dalam service

	// membuat objek categori controller, dengan memanggil konstruktor new categori
	// defendensi : categori service
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
