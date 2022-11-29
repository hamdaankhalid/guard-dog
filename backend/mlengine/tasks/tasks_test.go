package tasks_test

import (
	"fmt"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/tasks"
)

// mock multipart file struct
type mockFile struct {
}

func (m *mockFile) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (m *mockFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}
func (m *mockFile) Seek(offset int64, whence int) (int64, error) {
	var i int64 = 2
	return i, nil
}
func (m *mockFile) Close() error {
	return nil
}

func TestUploadModelTask(t *testing.T) {
	fmt.Println("FUCK U")
	uploadModelReq := tasks.UploadModelReq{
		UserId:  1,
		File:    &mockFile{},
		Handler: &multipart.FileHeader{Filename: "fakemodel.onnx", Size: 0, Header: textproto.MIMEHeader{}},
	}

	mq := dal.MockQueries{ErrorOnly: false}

	// invoke
	err := tasks.UploadModelTask(&uploadModelReq, &mq)

	// assert
	if err != nil {
		t.Fail()
	}
}
