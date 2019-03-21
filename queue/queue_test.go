package queue_test

import (
	"testing"

	"github.com/zboyco/goutlis/queue"
)

var lq *queue.LinkQueue

func init() {
	lq = &queue.LinkQueue{}
}

func Test_Enqueue(t *testing.T) {
	t.Log("测试入队")
	for i := 0; i < 100; i++ {
		lq.Enqueue(i)
	}
}

func Test_Dequeue(t *testing.T) {
	t.Log("测试出队")
	for {
		item, err := lq.Dequeue()
		if err != nil {
			// log.Println(err)
			// time.Sleep(100 * time.Millisecond)
			break
		}
		t.Log(item.(int))
	}
}

func Benchmark_LinkQueueEnqueue(b *testing.B) {
	q := &queue.LinkQueue{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func Benchmark_LinkQueueDequeue(b *testing.B) {
	q := &queue.LinkQueue{}
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := q.Dequeue()
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func Benchmark_CircularQueueEnqueue(b *testing.B) {
	q, _ := queue.InitCircularQueue(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func Benchmark_CircularQueueDequeue(b *testing.B) {
	q, _ := queue.InitCircularQueue(b.N)
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
