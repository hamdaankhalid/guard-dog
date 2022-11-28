package dal

import (
	"errors"

	"github.com/google/uuid"
)

type MockQueries struct {
	ErrorOnly       bool
	Error           error
	MlNotifications []MlNotification
	Models          []ModelWithoutData
	Model           Model
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
	return nil
}

func (q *MockQueries) RetrieveModel(id uuid.UUID) (Model, error) {
	if q.ErrorOnly {
		return Model{}, q.Error
	}
	return q.Model, nil
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
	return nil
}
