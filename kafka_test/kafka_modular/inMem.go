package main

import "sync"

type InMem struct {
	mu sync.Mutex
	data map[int]string
}





// create new in memory db
func NewInMem() *InMem{
	return &InMem{
		data : make(map[int]string),
	}
}


func (db *InMem) Save(id int, data string) {
	db.mu.Lock() //lock
	defer db.mu.Unlock() // must unlock
	db.data[id] = data // shared state
}