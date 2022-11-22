package handlers

import "net/http"

func UploadModel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// handle large file download, insert data into postgres

	// return success
}
