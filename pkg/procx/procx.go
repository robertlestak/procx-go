package procx

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
)

var (
	ErrNoData = errors.New("no data")
)

func Procx(args []string) (io.Reader, error) {
	cmd := exec.Command("procx", args...)
	cmd.Env = os.Environ()
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	e := cmd.Run()
	if e != nil {
		var err bytes.Buffer
		err.WriteString(err.String())
		err.WriteString(out.String())
		return nil, errors.New(err.String())
	}
	if out.Len() == 0 {
		return nil, ErrNoData
	}
	return &out, nil
}
