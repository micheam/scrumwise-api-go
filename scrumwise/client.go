package scrumwise

import (
	"fmt"
	"strings"
)

var (
	BaseURL    string = "https://api.scrumwise.com/service/api"
	ApiVersion string = "v1"
)

func Endpoint(method string) string {
	return fmt.Sprintf("%s/%s/%s",
		strings.TrimRight(BaseURL, "/"),
		strings.Trim(ApiVersion, "/"),
		strings.TrimLeft(method, "/"))
}
