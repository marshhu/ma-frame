package utils

import "encoding/json"

func ByteToObj(data []byte, obj interface{}) error {
	err := json.Unmarshal(data, obj)
	return err
}

func JsonToObj(str string, obj interface{}) error {
	data := []byte(str)
	err := json.Unmarshal(data, obj)
	return err
}

func ObjToJson(obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if data == nil {
		data = []byte("")
	}
	return string(data), err
}
