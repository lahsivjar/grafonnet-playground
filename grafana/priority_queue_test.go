package grafana

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPriorityQueue(t *testing.T) {
	assert := assert.New(t)

	pq := NewPriorityQueue()
	assert.NotNil(pq)
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	pq := NewPriorityQueue()
	assert.NotNil(pq)

	pq.Push(&Item{Key: "1", ProcessAt: time.Now().Add(-2 * time.Second)})

	assert.Equal(1, pq.Size())
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	pq := NewPriorityQueue()
	assert.NotNil(pq)

	pq.Push(&Item{Key: "1", ProcessAt: time.Now().Add(-2 * time.Second)})
	pq.Push(&Item{Key: "2", ProcessAt: time.Now().Add(-1 * time.Second)})

	assert.Equal(2, pq.Size())
	assert.Equal("1", pq.Pop().Key)
	assert.Equal(1, pq.Size())
	assert.Equal("2", pq.Pop().Key)
	assert.Equal(0, pq.Size())
}

func TestPopConditionally(t *testing.T) {
	assert := assert.New(t)

	pq := NewPriorityQueue()
	assert.NotNil(pq)

	pq.Push(&Item{Key: "1", ProcessAt: time.Now().Add(-2 * time.Second)})
	pq.Push(&Item{Key: "2", ProcessAt: time.Now().Add(-1 * time.Second)})
	pq.Push(&Item{Key: "3", ProcessAt: time.Now().Add(-1 * time.Minute)})

	ifBeforeNow := func(i *Item) bool {
		return i.ProcessAt.Before(time.Now())
	}

	assert.Equal(3, pq.Size())
	assert.Equal("3", pq.PopConditionally(ifBeforeNow).Key)
	assert.Equal("1", pq.PopConditionally(ifBeforeNow).Key)
	assert.Equal("2", pq.PopConditionally(ifBeforeNow).Key)
	assert.Equal(0, pq.Size())
}

func TestPopConditionally_Late(t *testing.T) {
	assert := assert.New(t)

	pq := NewPriorityQueue()
	assert.NotNil(pq)

	pq.Push(&Item{Key: "1", ProcessAt: time.Now().Add(2 * time.Second)})

	ifBeforeNow := func(i *Item) bool {
		return i.ProcessAt.Before(time.Now())
	}

	assert.Equal(1, pq.Size())
	assert.Nil(pq.PopConditionally(ifBeforeNow))
	assert.Equal(1, pq.Size())

	time.Sleep(2 * time.Second)
	assert.Equal("1", pq.PopConditionally(ifBeforeNow).Key)
	assert.Equal(0, pq.Size())
}
