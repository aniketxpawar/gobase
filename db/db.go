package db

import (
	"errors"

	"github.com/aniketxpawar/gobase/storage"
)

type Database struct {
	storage *storage.Storage
}

func NewDatabase(storageFile string) (*Database, error){
	store, err := storage.NewStorage(storageFile)
	if err != nil{
		return nil, err
	}
	return &Database{storage: store}, nil
}

func (db *Database) Set(key string, value interface{}) error {
	return db.storage.Set(key,value)
}

func (db *Database) Get(key string) (interface{}, bool) {
	return db.storage.Get(key)
}

func (db *Database) Delete(key string) {
	db.storage.Delete(key)
}

func (db *Database) Append(key string, value interface{}) error {
	existing,exists := db.storage.Get(key)
	if !exists {
		return errors.New("Key does not exist")
	}
	array, ok := existing.([]interface{})
	if !ok {
		return errors.New("Value is not a Array")
	}
	array = append(array, value)
	return db.storage.Set(key, array)
}

func (db *Database) GetJSONKey(key string, jsonKey string) (interface{}, error) {
	existing, exists := db.storage.Get(key)
	if !exists {
		return nil,errors.New("Key does not exist")
	}
	jsonObj, ok := existing.(map[string]interface{})
	if !ok {
		return nil, errors.New("Value is not a JSON")
	}
	value, exists := jsonObj[jsonKey]
	if !exists {
		return nil, errors.New("Given Key does not exist")
	}
	return value,nil
}