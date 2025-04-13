package utils
import (
	"time"
	"github.com/google/go-github/v60/github"
)
func FormatDate(t time.Time) string {
	return t.Format("Jan 02, 2006 15:04 MST")
}
func FormatGitHubTime(t *github.Timestamp) string {
	if t == nil {
		return "N/A"
	}
	return t.Format("Jan 02, 2006 15:04 MST")
}
