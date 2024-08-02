package converter

import (
	"strconv"

	"github.com/pkg/errors"
)

func StringToInt(from string) (int, error) {
	_v, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(_v), nil
}
