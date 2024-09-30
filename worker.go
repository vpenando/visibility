package visibility

import (
	"github.com/vpenando/visibility/internal/worker"
)

type Worker interface {
	DoWork()
}

func NewWorker() Worker {
	w := worker.WorkerImplementation{}
	return w
}
