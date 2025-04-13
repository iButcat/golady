package utils
import (
	"context"
	"errors"
	"github.com/google/go-github/github"
)
func IsPrivate(ghClient *github.Client, owner, name string) (bool, error) {
	repo, _, err := ghClient.Repositories.Get(context.Background(),
		owner, name)
	if err != nil {
		return false, errors.New("not found, you might need to have a token with the right scope")
	}
	return repo.GetPrivate(), nil
}
