package tasks

import (
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Task struct {
	Text      string `json:"text"`
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Index     int    `json:"index"`
}

type Tasks []Task

// Implement sort interface to sort tasks by index and name
func (tasks Tasks) Len() int {
	return len(tasks)
}

func (tasks Tasks) Less(i, j int) bool {
	if tasks[i].Index == tasks[j].Index {
		return tasks[i].Text < tasks[j].Text
	}

	return tasks[i].Index < tasks[j].Index
}

func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

// Simulated DB
var tasksIndex map[string]Task

func init() {
	tasksIndex = make(map[string]Task)
}

func GetTasks() []Task {
	tasks := make(Tasks, 0, len(tasksIndex))

	for _, task := range tasksIndex {
		tasks = append(tasks, task)
	}

	slices.SortFunc(tasks, func(t1, t2 Task) int {
		if t1.Index == t2.Index {
			return strings.Compare(t1.Text, t2.Text)
		}

		if t1.Index < t2.Index {
			return -1
		} else if t1.Index > t2.Index {
			return 1
		} else {
			return 0
		}

	})

	return tasks

}

func AddTask(task Task) string {
	task.ID = uuid.New().String()
	task.Index = len(tasksIndex)

	tasksIndex[task.ID] = task

	return task.ID
}

func RemoveTask(id string) {
	delete(tasksIndex, id)
}
