// Here we have an imagined job queue system, where there are two heaps: one,
// PriorityQueue, for storing jobs which are ready to be processed and are
// sorted by their priority; the other, DelayQueue, for storing delayed jobs
// and stored by their timestamp.
//
// The PriorityQueue and DelayQueue structs wrap Heap, and implement simple
// wrappers for Insert based on the deadline or priority, and Pop which
// handles the interface conversion.
package binheap

import "time"
import "fmt"

type Job struct {
	Deadline time.Time
	Priority int
	Body     string
}

type PriorityQueue struct{ Heap }

func (q *PriorityQueue) Insert(job *Job) {
	q.Heap.Insert(job.Priority, job)
}

func (q *PriorityQueue) Pop() (*Job) {
	if item := q.Heap.Pop(); item != nil {
		return item.(*Job)
	}

	return nil
}

type DelayQueue struct{ Heap }

func (q *DelayQueue) Insert(job *Job) {
	q.Heap.Insert(int(job.Deadline.Unix()), job)
}

func (q *DelayQueue) Pop() (*Job) {
	if item := q.Heap.Pop(); item != nil {
		return item.(*Job)
	}

	return nil
}

func ExampleHeap_custom() {
	priortyq, delayq := &PriorityQueue{}, &DelayQueue{}

	priortyq.Insert(&Job{Deadline: time.Now(), Priority: 20, Body: "job 1"})
	delayq.Insert(&Job{Deadline: time.Now(), Priority: 10, Body: "job 2"})

	if job := delayq.Pop(); job != nil {
		priortyq.Insert(job)
	}

	job := priortyq.Pop()
	fmt.Println(job.Body)

	job = priortyq.Pop()
	fmt.Println(job.Body)

	// Output: job 2
	// job 1
}
