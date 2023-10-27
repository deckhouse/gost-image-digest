package calculator

import (
	"io"

	"go.cypherpunks.ru/gogost/v5/gost34112012256"
)

func Calculate(reader io.Reader) ([]byte, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	hasher := gost34112012256.New()
	_, err = hasher.Write(data)
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
