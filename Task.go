package main

type Task struct {
	Name      string
	Is_finish bool
}

func NewTask(name string) *Task {
	return &Task{
		Name:      name,
		Is_finish: false,
	}
}
func (task Task) GetName() string {
	return task.Name
}

func (task Task) IsDone() bool {
	return task.Is_finish
}
func (task *Task) Done() {
	task.Is_finish = true
}
func (task *Task) ReDo() {
	task.Is_finish = false
}
