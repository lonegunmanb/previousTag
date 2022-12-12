package pkg

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PreviousTag(t *testing.T) {
	cases := []struct {
		tags       []string
		name       string
		currentTag string
		expected   string
	}{
		{
			tags: []string{
				"0.0.1",
				"0.0.2",
				"v1.0.0"},
			name:       "different_major_version",
			currentTag: "v1.0.0",
			expected:   "0.0.2",
		},
		{
			tags: []string{
				"v0.0.1",
				"v0.0.2",
				"v1.0.0"},
			name:       "different_major_version_with_v",
			currentTag: "v1.0.0",
			expected:   "v0.0.2",
		},
		{
			tags: []string{
				"0.0.1",
				"0.0.2",
				"1.0.0"},
			name:       "different_major_version_without_v",
			currentTag: "1.0.0",
			expected:   "0.0.2",
		},
		{
			tags: []string{
				"v0.0.1",
				"v0.0.2",
				"1.0.0"},
			name:       "different_major_version_reverse",
			currentTag: "1.0.0",
			expected:   "v0.0.2",
		},
		{
			tags: []string{
				"0.0.1",
				"0.0.2",
				"0.0.3"},
			name:       "same_major_version",
			currentTag: "0.0.2",
			expected:   "0.0.1",
		},
		{
			tags:       []string{},
			name:       "no_tag",
			currentTag: "0.0.0",
			expected:   "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stub := gostub.Stub(&getTags, func(owner string, repo string) ([]string, error) {
				return c.tags, nil
			})
			defer stub.Reset()

			tag, err := PreviousTag("dummyOwner", "dummyRepo", c.currentTag)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, tag)
		})
	}
}
