package dal

import "github.com/google/uuid"

type MockQueries struct {
	ErrorOnly bool
	Error     error
}

func (q *MockQueries) UploadMlNotification(mlNotification *MlNotification) error {
	if q.ErrorOnly {
		return q.Error
	}
	return nil
}

func (q *MockQueries) RetrieveAllMlNotifications(userId int) ([]MlNotification, error) {
	if q.ErrorOnly {
		return nil, q.Error
	}
	id1 := uuid.New()
	modelId1 := uuid.New()
	id2 := uuid.New()
	modelId2 := uuid.New()

	notification1 := MlNotification{
		Id:         id1,
		DeviceName: "dev1",
		SessionId:  1,
		Part:       1,
		ModelId:    modelId1,
		UserId:     1}

	notification2 := MlNotification{
		Id:         id2,
		DeviceName: "dev1",
		SessionId:  2,
		Part:       1,
		ModelId:    modelId2,
		UserId:     1}

	return []MlNotification{notification1, notification2}, nil
}

func (q *MockQueries) RetrieveMlNotification(notificationId uuid.UUID) (MlNotification, error) {
	if q.ErrorOnly {
		return MlNotification{}, q.Error
	}
	id1 := uuid.New()
	modelId1 := uuid.New()

	notification1 := MlNotification{
		Id:         id1,
		DeviceName: "dev1",
		SessionId:  1,
		Part:       1,
		ModelId:    modelId1,
		UserId:     1}

	return notification1, nil
}

func (q *MockQueries) UploadModel(model *Model) error {
	if q.ErrorOnly {
		return q.Error
	}
	return nil
}

func (q *MockQueries) RetrieveModel(id uuid.UUID) (Model, error) {
	if q.ErrorOnly {
		return Model{}, q.Error
	}
	id1 := uuid.New()

	notification1 := Model{
		Id:        id1,
		Filename:  "f1.onnx",
		UserId:    1,
		ModelFile: []byte{}}

	return notification1, nil
}

func (q *MockQueries) RetrieveAllModels(userId int) ([]ModelWithoutData, error) {
	if q.ErrorOnly {
		return nil, q.Error
	}
	id1 := uuid.New()
	id2 := uuid.New()

	notification1 := ModelWithoutData{
		Id:       id1,
		Filename: "f1.onnx",
		UserId:   1,
	}

	notification2 := ModelWithoutData{
		Id:       id2,
		Filename: "f2.onnx",
		UserId:   1,
	}

	return []ModelWithoutData{notification1, notification2}, nil
}

func (q *MockQueries) DeleteModel(modelId uuid.UUID) error {
	if q.ErrorOnly {
		return q.Error
	}
	return nil
}
