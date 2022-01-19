package utils

import (
	"encoding/json"
	"io"
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

func WriteJSON(w io.Writer, js interface{}) error {
	data, err := json.Marshal(js)
	if err != nil {
		return Wrap(err)
	}

	if _, err := w.Write(data); err != nil {
		return Wrap(err)
	}

	return nil
}
