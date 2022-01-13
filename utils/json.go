package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSON(name string, js interface{}) error {
	f, err := os.Open(name)
	if err != nil {
		return Wrap(err)
	}
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return Wrap(err)
	}

	return Wrap(json.Unmarshal(body, js))
}
