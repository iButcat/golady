package utils
import "strings"
func ParseArgs(m string) ([]string, int) {
	args := strings.Split(m[1:], " ")
	args = args[1:]
	return args, len(args)
}
