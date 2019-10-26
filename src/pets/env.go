package pets

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var ErrInvalid = errors.New("Error: invalid file contents")

// setEnv sets ENV_VARS from file at fpath
func setEnv() error {
	fpath := ".env"
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()
	m, err := readEnv(f)
	if err != nil {
		return err
	}
	for k, v := range m {
		os.Setenv(k, v)
	}
	return nil
}

// readEnv reads pairs of keys and values into a map and returns it
func readEnv(r io.Reader) (map[string]string, error) {
	m := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "=")
		if len(pair) != 2 {
			return m, ErrInvalid
		}
		m[pair[0]] = pair[1]
	}
	return m, nil
}
