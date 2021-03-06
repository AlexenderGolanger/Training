package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	id     int
	name   string
	age    string
	salary int
}

type Storage interface {
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]Employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) insert(e Employee) error {
	s.data[e.id] = e

	return nil
}

func (s *memoryStorage) get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("Employee with such id doesn't exist")
	}

	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e Employee) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (s *dumbStorage) get(id int) (Employee, error) {
	e := Employee{
		id: id,
	}

	return e, nil
}

func (s *dumbStorage) delete(id int) error {
	fmt.Printf("вставка пользователя с id: %d прошла успешно\n", id)
	return nil
}
