package goutlis_test

import (
	"testing"

	"github.com/zboyco/goutlis"
)

var cq *goutlis.CircularQueue

func init() {
	cq, _ = goutlis.InitCircularQueue(100)
}

func Test_CircularQueueEnqueue(t *testing.T) {
	t.Log("测试入队")
	for i := 0; i < 100; i++ {
		cq.Enqueue(i)
	}
}

func Test_CircularQueueDequeue(t *testing.T) {
	t.Log("测试出队")
	for {
		item, err := cq.Dequeue()
		if err != nil {
			// log.Println(err)
			// time.Sleep(100 * time.Millisecond)
			break
		}
		t.Log(item.(int))
	}
}

func Benchmark_CircularQueueEnqueue(b *testing.B) {
	q, _ := goutlis.InitCircularQueue(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func Benchmark_CircularQueueDequeue(b *testing.B) {
	q, _ := goutlis.InitCircularQueue(b.N)
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := q.Dequeue()
		if err != nil {
			b.Error(err, i)
			return
		}
	}
}
