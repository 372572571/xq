package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func MarshalString(v interface{}) string {
	str_by, err := json.Marshal(v)

	if err != nil {
		return ""
	}

	str := string(str_by)
	return str
}

// str_io bytes.NewBufferString("")
func MarshalEX(v interface{}, str_io io.Writer) error {
	encoder := json.NewEncoder(str_io)
	err := encoder.Encode(v)
	return err
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalByte(v interface{}) []byte {
	str_by, err := json.Marshal(v)

	if err != nil {
		str_by = []byte("")
	}

	return str_by
}

func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := unmarshalUseNumber(decoder, v); err != nil {
		return formatError(string(data), err)
	}

	return nil
}

func unmarshalUseNumber(decoder *json.Decoder, v interface{}) error {
	decoder.UseNumber()
	return decoder.Decode(v)
}

func formatError(v string, err error) error {
	return fmt.Errorf("string: `%s`, error: `%w`", v, err)
}
