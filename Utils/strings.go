package Utils

import "fmt"

func ConcatAPIPath(apiPrefix, endPoint string) string {
	return fmt.Sprintf("%s/%s", apiPrefix, endPoint)
}
