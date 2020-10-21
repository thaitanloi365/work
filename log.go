package work

import "fmt"

func logError(key string, err error) {
	fmt.Printf("ERROR: %s - %+v\n", key, err)
}
