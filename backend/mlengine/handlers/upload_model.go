package handlers

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func UploadModel(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// get user details from the JWT token
	model := dal.Model{ModelFile: file, Id: -1, UserId: user.Id}
	err = dal.UploadModel(&model)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
