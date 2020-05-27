package grafana

import (
	"container/heap"
	"sync"
	"time"
)

// PriorityQueue defines a thread safe priority queue
type PriorityQueue struct {
	lock  sync.Mutex
	queue queue
}

// Item defines a struct that the PriorityQueue consumes
type Item struct {
	Key        string
	RetryCount int
	ProcessAt  time.Time
}

func (item *Item) priority() int64 {
	return -1 * item.ProcessAt.UnixNano()
}

type queue []*Item

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].priority() > q[j].priority()
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(*Item))
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

// Push pushes a new item to priority queue with a given value and priority
func (pq *PriorityQueue) Push(item *Item) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	heap.Push(&pq.queue, item)
}

// Pop returns the lowest priority item from the queue
func (pq *PriorityQueue) Pop() *Item {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	raw := heap.Pop(&pq.queue)
	return raw.(*Item)
}

// PopConditionally pods an item if the given condition is met
func (pq *PriorityQueue) PopConditionally(f func(*Item) bool) *Item {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	if len(pq.queue) > 0 {
		item := pq.queue[0]

		if f(item) {
			raw := heap.Pop(&pq.queue)
			return raw.(*Item)
		}
	}
	return nil
}

// Size returns the current size of the priority queue
func (pq *PriorityQueue) Size() int {
	return pq.queue.Len()
}
