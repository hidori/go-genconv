package converter

import "fmt"

func IntToString(from int) (string, error) {
	return fmt.Sprintf("%d", from), nil
}
