package utils
import (
	"testing"
	"github.com/magiconair/properties/assert"
)
const (
	input = "!test 1 2 3"
)
func TestParserLenInput(t *testing.T) {
	expectedLen := 3
	_, nb := ParseArgs(input)
	if nb != expectedLen {
		t.Errorf("Expected %d, got %d", expectedLen, nb)
	}
}
func TestParserInput(t *testing.T) {
	expectedArgs := []string{"1", "2", "3"}
	args, _ := ParseArgs(input)
	assert.Equal(t, expectedArgs, args)
}
