package main

import (
	"net/http"
	"os"
	"strings"
	
	"image_gallery/category"
	"image_gallery/database"
	"image_gallery/home"
	"image_gallery/image"
	cLog "image_gallery/logger"
	"image_gallery/router"
	
	"github.com/gorilla/handlers"
)

func main() {
	logger := cLog.GetLogger()

	logger.Info("Server started on port 8080")
	apiRouter := router.Router{
		Logger: logger,
	}

	// Home handler
	apiRouter.AddHandler(&home.Handler{
		Logger: logger,
	})

	// Category handler
	apiRouter.AddHandler(&category.Handler{
		Logger: logger,
	})

	// Images handler
	apiRouter.AddHandler(&image.Handler{
		Logger: logger,
	})
	

	err := database.Connect()

	if err != nil {
		logger.Fatalf("could not connect to db: %v", err)
	}

	muxRouter := apiRouter.Configure()

	err = http.ListenAndServe(
		":8080",
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedOrigins(strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")),
			handlers.AllowedHeaders([]string{"Content-Type"}),
			handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"}),
		)(muxRouter),
	)

	logger.Fatal(err)
}
