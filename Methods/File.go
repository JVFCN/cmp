package Methods

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"os"
	"regexp"
	"strings"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func JsonArrToArr(arr *jsonvalue.V) []string {
	re := regexp.MustCompile(`^\[|]$`)
	cleanStr := re.ReplaceAllString(arr.String(), "")

	return strings.Split(cleanStr, ",")
}
