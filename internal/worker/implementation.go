package worker

import "fmt"

type WorkerImplementation struct {
}

func (wi WorkerImplementation) DoWork() {
	fmt.Println("doing some stuff...")
}
