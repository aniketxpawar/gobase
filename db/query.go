package db

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (db *Database) ExecuteQuery(query string) (string, error) {
	parts := strings.Fields(query)
	lenParts := len(parts)

	if lenParts < 2{
		return "",fmt.Errorf("Invalid Query")
	}
	command := strings.ToUpper(parts[0])
	key := parts[1]

	switch command {
	case "SET":
		if lenParts < 3{
			return "",fmt.Errorf("SET requires key and value")
		}
		var value interface{}
		err := json.Unmarshal([]byte(strings.Join(parts[2:]," ")), &value)
		if err != nil{
			value = strings.Join(parts[2:]," ")
		}
		err = db.Set(key, value)
		if err != nil{
			return "",err
		}
		return "OK",nil
	case "GET":
		value, exists := db.Get(key)
		if !exists{
			return "nil", nil
		}
		output,_ := json.Marshal(&value)
		return string(output), nil
	case "PUSH":
		if lenParts < 3{
			return "",fmt.Errorf("PUSH requires key and value")
		}
		err := db.Append(key, parts[2])
		if err != nil{
			return "nil", err
		}
		return "OK", nil
	case "JSON.GET":
		if lenParts < 3{
			return "",fmt.Errorf("JSON.GET requires key and JSON key")
		}
		value, err := db.GetJSONKey(key, parts[2])
		if err != nil {
			return "nil", err
		}
		output,_ := json.Marshal(&value)
		return string(output), nil 
	default:
		return "", fmt.Errorf("Unknown Command")
	}
}