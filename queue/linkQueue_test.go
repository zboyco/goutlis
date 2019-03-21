package queue_test

import (
	"testing"

	"github.com/zboyco/goutlis/queue"
)

var q *queue.LinkQueue

func init() {
	q = &queue.LinkQueue{}
}

func Test_Enqueue(t *testing.T) {
	t.Log("测试入队")
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
}

func Test_Dequeue(t *testing.T) {
	t.Log("测试出队")
	for {
		item, err := q.Dequeue()
		if err != nil {
			// log.Println(err)
			// time.Sleep(100 * time.Millisecond)
			break
		}
		t.Log(item.(int))
	}
}

func Benchmark_Enqueue(b *testing.B) {
	b.N = 1000000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func Benchmark_Dequeue(b *testing.B) {
	b.N = 1000000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := q.Dequeue()
		if err != nil {
			b.Error(err)
			return
		}
	}
}
