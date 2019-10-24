package pets

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var ErrInvalid = errors.New("Error: invalid file contents")

// setEnv sets ENV_VARS from the passed in reader
func setEnv(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), "=")
		if len(lineSlice) != 2 {
			return ErrInvalid
		}
		k := lineSlice[0]
		v := lineSlice[1]
		os.Setenv(k, v)
	}
	return nil
}
