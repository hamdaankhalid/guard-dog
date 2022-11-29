package dal

import (
	"errors"

	"github.com/google/uuid"
)

type MockQueries struct {
	// All attributes are exported so we can use them for testing
	ErrorOnly       bool
	Error           error
	MlNotifications []MlNotification
	Models          []ModelWithoutData
	ModelsWithData  []Model
}

func (q *MockQueries) UploadMlNotification(mlNotification *MlNotification) error {
	if q.ErrorOnly {
		return q.Error
	}
	q.MlNotifications = append(q.MlNotifications, *mlNotification)
	return nil
}

func (q *MockQueries) RetrieveAllMlNotifications(userId int) ([]MlNotification, error) {
	if q.ErrorOnly {
		return nil, q.Error
	}

	return q.MlNotifications, nil
}

func (q *MockQueries) RetrieveMlNotification(notificationId uuid.UUID) (MlNotification, error) {
	if q.ErrorOnly {
		return MlNotification{}, q.Error
	}
	for _, v := range q.MlNotifications {
		if v.Id == notificationId {
			return v, nil
		}
	}

	return MlNotification{}, errors.New("Noitification not found")
}

func (q *MockQueries) UploadModel(model *Model) error {
	if q.ErrorOnly {
		return q.Error
	}

	q.ModelsWithData = append(q.ModelsWithData, *model)

	return nil
}

func (q *MockQueries) RetrieveModel(id uuid.UUID) (Model, error) {
	if q.ErrorOnly {
		return Model{}, q.Error
	}

	for _, v := range q.ModelsWithData {
		if v.Id == id {
			return v, nil
		}
	}

	for _, v := range q.Models {
		if v.Id == id {
			return Model{ModelFile: []byte{}, Id: v.Id, Filename: v.Filename, UserId: v.UserId}, nil
		}
	}
	return Model{}, errors.New("No Model Found")
}

func (q *MockQueries) RetrieveAllModels(userId int) ([]ModelWithoutData, error) {
	if q.ErrorOnly {
		return nil, q.Error
	}

	return q.Models, nil
}

func (q *MockQueries) DeleteModel(modelId uuid.UUID) error {
	if q.ErrorOnly {
		return q.Error
	}

	for index, v := range q.ModelsWithData {
		if v.Id == modelId {
			q.ModelsWithData = append(q.ModelsWithData[:index], q.ModelsWithData[index+1:]...)
		}
	}

	for index, v := range q.Models {
		if v.Id == modelId {
			q.ModelsWithData = append(q.ModelsWithData[:index], q.ModelsWithData[index+1:]...)
		}
	}

	return nil
}
