package downloads

import (
	"strings"

	"github.com/anvh2/dowmuzic/plugins/prompts"
)

// Song -
func Song() error {
	url, err := prompts.String("url")
	if err != nil {
		return err
	}

	if !strings.Contains(url, "") {

	}
	return nil
}
