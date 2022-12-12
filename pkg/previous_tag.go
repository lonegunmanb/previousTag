package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"golang.org/x/mod/semver"
	"strings"

	"github.com/google/go-github/v42/github"
)

func PreviousTag(owner, repo, currentTag string) (string, error) {
	if !valid(currentTag) {
		return "", errors.New(fmt.Sprintf("invalid current tag: %s", currentTag))
	}

	tags, err := getTags(owner, repo)
	if err != nil {
		return "", err
	}

	t := linq.From(tags).Where(func(i interface{}) bool {
		return valid(i.(string))
	}).Where(func(i interface{}) bool {
		return semver.Compare(wrap(currentTag), wrap(i.(string))) > 0
	}).Sort(func(i, j interface{}) bool {
		return semver.Compare(wrap(i.(string)), wrap(j.(string))) < 0
	}).Last()

	if t == nil {
		return "", nil
	}

	return t.(string), nil
}

var getTags = func(owner string, repo string) ([]string, error) {
	client := github.NewClient(nil)
	tags, _, err := client.Repositories.ListTags(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}
	if tags == nil {
		return nil, fmt.Errorf("cannot find tags")
	}
	var t []string
	linq.From(tags).Select(func(i interface{}) interface{} {
		return i.(*github.RepositoryTag).GetName()
	}).ToSlice(&t)
	return t, nil
}

func valid(t string) bool {
	t = wrap(t)
	semValid := semver.IsValid(t)
	return semValid && !strings.Contains(t, "rc")
}

func wrap(t string) string {
	if !strings.HasPrefix(t, "v") {
		t = fmt.Sprintf("v%s", t)
	}
	return t
}
