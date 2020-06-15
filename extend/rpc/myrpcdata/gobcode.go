package rpc

import (
	"bytes"
	"encoding/gob"
)

type RPCData struct {
	Name string
	Args []interface{}
}

func encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)
	err := bufEnc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	err := bufDec.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

















