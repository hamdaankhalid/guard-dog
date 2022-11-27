package processingqueue

type MockQueue struct {
	InnerState []string
}

func (q *MockQueue) BeginProcessing() {}

func (q *MockQueue) Enqueue(taskName string, task interface{}) {
	q.InnerState = append(q.InnerState, taskName)
}
