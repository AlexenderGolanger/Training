package main

import (
	"fmt"
	"github.com/AlexenderGolanger/Training/storage/storage"
)

func main() {
	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))

	spawnEmployees(ds)
	fmt.Println(ds)
}

func spawnEmployees(s Storage) {
	for i := 1; i <= 10; i++ {
		s.insert(Employee{id: i})
	}
}
