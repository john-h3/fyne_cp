package main

import (
	"container/list"
	"crypto/md5"
	"time"
)

type Record struct {
	Time    time.Time
	Content []byte
}

func (r Record) String() string {
	return string(r.Content)
}

type Records struct {
	capacity int
	elements map[[16]byte]*list.Element
	order    *list.List
}

func NewRecords(capacity int) *Records {
	return &Records{
		capacity: capacity,
		elements: make(map[[16]byte]*list.Element),
		order:    list.New(),
	}
}

func (r *Records) Add(content []byte) bool {
	md5Bytes := md5.Sum(content)
	if e, ok := r.elements[md5Bytes]; ok {
		if e == r.order.Front() {
			return false
		}
		rd := e.Value.(Record)
		rd.Time = time.Now()
		e.Value = rd
		r.order.MoveToFront(e)
	} else {
		if r.order.Len() < 5 {
			r.capacity++
		} else {
			// evict last
			last := r.order.Back()
			rd := r.order.Remove(last).(Record)
			delete(r.elements, md5.Sum(rd.Content))
		}
		r.elements[md5Bytes] = r.order.PushFront(Record{time.Now(), content})
	}
	return true
}

func (r *Records) Len() int {
	return r.order.Len()
}

func (r *Records) Slice() [][]byte {
	slice := make([][]byte, 0)
	e := r.order.Front()
	for e != nil {
		slice = append(slice, e.Value.(Record).Content)
		e = e.Next()
	}
	return slice
}
