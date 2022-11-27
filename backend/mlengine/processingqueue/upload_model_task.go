package processingqueue

import (
	"bytes"
	"io"
	"log"

	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/dal"
)

func uploadModelTask(uploadModelReq *UploadModelReq) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)

		return
	}
	bytes := bytes.NewBuffer(nil)
	n, err := io.Copy(bytes, *uploadModelReq.File)
	if err != nil || n == 0 {
		log.Println("Read Error", err)
		return
	}

	model := dal.Model{ModelFile: bytes.Bytes(), Id: uuid, Filename: uploadModelReq.Handler.Filename, UserId: uploadModelReq.UserId}
	err = dal.UploadModel(&model)
	if err != nil {
		log.Println("Error uploading model file: ", model.Filename)
	}
}
