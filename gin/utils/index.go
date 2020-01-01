package utils

import "encoding/json"

func JsonFromStr(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), v)
	return err
}

func JsonFromByte(byte []byte, v interface{}) error {
	err := json.Unmarshal(byte, v)
	return err
}