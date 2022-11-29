package handlers

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/middlewares"
	"github.com/hamdaankhalid/mlengine/tasks"
)

func (router *Router) UploadModel(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("model")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	router.processingQueue.Enqueue(tasks.UploadModelTaskName, &tasks.UploadModelReq{File: &file, Handler: handler, UserId: user.Id})

	w.WriteHeader(http.StatusCreated)
}
