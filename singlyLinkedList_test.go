package goutlis_test

import (
	"strconv"
	"testing"

	"github.com/zboyco/goutlis"
)

var l *goutlis.SinglyList

func init() {
	l = &goutlis.SinglyList{}
}

func TestAppend(t *testing.T) {
	t.Log("测试追加")
	for i := 0; i < 2; i++ {
		err := l.Append("test text" + strconv.Itoa(i))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestLenght(t *testing.T) {
	t.Log("测试长度获取")
	lenght := l.Lenght()
	if lenght != 2 {
		t.Error("长度不正确")
	}
}

func TestInsertAt(t *testing.T) {
	t.Log("测试插入数据")
	err := l.InsertAt(1, "insert text3")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	t.Log("测试获取元素")
	item, err := l.Get(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(item.(string))
}

func TestRemove(t *testing.T) {
	t.Log("测试删除")
	item, err := l.Remove(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("remove data: " + item.(string))
}

func TestClear(t *testing.T) {
	t.Log("测试清空")
	l.Clear()
	if l.Lenght() > 0 {
		t.Error("清空失败")
		return
	}
}
