package inMemory

import (
	"fmt"
	"sync"
)

type InMem struct { // in-memory db struct
	mu sync.Mutex
	data map[int]string
}





// create new in memory db
func NewInMem() *InMem{
	return &InMem{
		data : make(map[int]string),
	}
}

// save data to in-memory db
func (db *InMem) Save(id int, data string) {
	db.mu.Lock() //lock
	defer db.mu.Unlock() // must unlock
	db.data[id] = data // shared state
}
// print all data in in-memory db
func (db *InMem) PrintAll() {
	for id, data := range db.data {
		fmt.Printf("ID: %d, Data: %s\n", id, data)
	}
}