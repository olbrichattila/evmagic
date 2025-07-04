package helpers

import "encoding/json"

func ToStruct[T any](data []byte) (T, error) {
	var m T
	err := json.Unmarshal(data, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
