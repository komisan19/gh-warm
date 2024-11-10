package warm

import (
	"bytes"
	"fmt"

	"github.com/cli/go-gh/v2"
)

const (
	gfi = "--good-first-issues=>=10"
)

type ghCLI struct{}

func NewWarm() *ghCLI {
	return &ghCLI{}
}

func (g *ghCLI) SearchRepo(language string) (*bytes.Buffer, *bytes.Buffer, error) {
	lang := fmt.Sprintf("--language=%v", language)
	sout, eout, err := gh.Exec("search", "repos", lang, gfi)

	if err != nil {
		return nil, &eout, err
	}
	return &sout, nil, nil
}
