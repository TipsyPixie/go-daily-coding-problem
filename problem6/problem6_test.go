package problem6

import (
    "testing"
)

func TestAdd(t *testing.T) {
    head := NewNode(NodePointer(0))
    err := Add(head)
    if err != nil {
        t.FailNow()
    }
    if head.both == NodePointer(0) || head.both.dereferencePointer() == nil {
        t.Fail()
    }
}

func TestGet(t *testing.T) {
    head := NewNode(NodePointer(0))
    for i := 0; i < 100; i++ {
        err := Add(head)
        if err != nil {
            t.FailNow()
        }
    }
    node0, err := Get(head, 0)
    if err != nil || node0 != head {
        t.FailNow()
    }
    node96, err96 := Get(head, 96)
    node97, err97 := Get(head, 97)
    node98, err98 := Get(head, 98)
    if err96 != nil || err97 != nil || err98 != nil || node97.both != node96.getPointer()^node98.getPointer() {
        t.FailNow()
    }
}
