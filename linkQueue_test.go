package goutlis_test

import (
	"testing"

	"github.com/zboyco/goutlis"
)

var lq *goutlis.LinkQueue

func init() {
	lq = &goutlis.LinkQueue{}
}

func Test_LinkQueueEnqueue(t *testing.T) {
	t.Log("测试入队")
	for i := 0; i < 100; i++ {
		lq.Enqueue(i)
	}
}

func Test_LinkQueueDequeue(t *testing.T) {
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
	q := &goutlis.LinkQueue{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func Benchmark_LinkQueueDequeue(b *testing.B) {
	q := &goutlis.LinkQueue{}
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
