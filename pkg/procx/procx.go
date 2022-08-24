package procx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	ErrNoData = errors.New("no data")
)

func findExec(name string) (string, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}
	if path == "" {
		path = "/usr/local/bin/" + name
	}
	return path, nil
}

func Procx(args []string) (io.Reader, error) {
	path, e := findExec("procx")
	if e != nil {
		return nil, e
	}
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	e = cmd.Run()
	if e != nil {
		var err bytes.Buffer
		err.WriteString(err.String())
		err.WriteString(out.String())
		return nil, errors.New(err.String())
	}
	if len(err.Bytes()) > 0 {
		fmt.Fprint(os.Stderr, err.String())
	}
	if out.Len() == 0 {
		return nil, ErrNoData
	}
	return &out, nil
}
