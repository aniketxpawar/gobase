package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type Datatype interface{}

type Storage struct {
	data map[string]Datatype
	file *os.File
	mu sync.RWMutex
}

func NewStorage(filename string) (*Storage, error) {
	storage := &Storage{
		data: make(map[string]Datatype),
	}

	if filename != "" {
		file, err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE,0644)
		if err != nil{
			return nil,err
		}
		storage.file = file
		storage.loadFromFile()
	}
	return storage, nil
}

func (s *Storage) loadFromFile() error {
	decoder := json.NewDecoder(s.file)
	return decoder.Decode(&s.data)
}

func (s *Storage) saveToFile() error {
	encoder := json.NewEncoder(s.file)
	return encoder.Encode(s.data)
}

func (s *Storage) Set(key string, value Datatype) error{
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
	if s.file != nil{
		s.saveToFile()
	}
	return nil
}

func (s *Storage) Get(key string) (Datatype, bool){
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, exists := s.data[key]
	return value, exists
	
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	if s.file != nil{
		s.saveToFile()
	}
}