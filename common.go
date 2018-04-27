package umfsdk

import (
	"io/ioutil"
	"os"
)

func ReadFileAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
