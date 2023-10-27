package calculator

import (
	"fmt"
	"io"
	"os"

	"go.cypherpunks.ru/gogost/v5/gost34112012256"
)

func Calculate(reader io.Reader) ([]byte, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	hasher := gost34112012256.New()
	_, err = hasher.Write(data)
	if err != nil {
		return nil, err
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	return hasher.Sum(nil), nil
}
