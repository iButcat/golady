package utils
import (
	"strconv"
	"strings"
)
func RemoveSingleQuote(ele string) string {
	return strings.ReplaceAll(ele, "'", "")
}
func IntToString(i int) string {
	return strconv.Itoa(i)
}
func IntPtrToString(i *int) string {
	if i == nil {
		return "0"
	}
	return strconv.Itoa(*i)
}
