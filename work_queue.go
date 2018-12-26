package main
import "sync"
import "fmt"
type Worker interface {
	work interface{}
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

var wg sync.WaitGroup

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO: initialize struct; start nWorkers workers as goroutines
	q.Results = make(chan interface{} , maxJobs)
	q.Jobs = make(chan Worker , maxJobs)
	
	for i := 0 ; i < int(nWorkers) ; i++ {
	wg.Add(1)
	
	go q.worker()	
	}
	
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	select {
	case job :=  <- queue.Jobs : 

		queue.Results <- job.Run()
		wg.Done()
	default : 
	close(queue.Jobs)
	}
	
	return 
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: close .Jobs and remove all remaining jobs from the channel.
	
	select {
	case  jb:= <- queue.Jobs :
		fmt.Println("emptying queue --> elements to remove : " , jb)
	default : 
	close(queue.Jobs)
	}

}
func main(){}