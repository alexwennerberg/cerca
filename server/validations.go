package server

import (
	"cerca/util"
	"fmt"
	"regexp"
)

const maxContentLength = 10000
const maxThreadTitle = 100
// username must consist of letters, digits, or the
// symbols _ and -; unicode OK
const usernameRegex = `^[\pL\pN_-]{1,16}$`

func validateContent(c string) error {
	if len(c) > maxContentLength {
		return fmt.Errorf("post content is too long, max length %d", maxContentLength)
	}
	return nil
}

func validateTitle(t string) error {
	if len(t) > maxThreadTitle {
		return fmt.Errorf("title is too long, max length %d", maxThreadTitle)
	}
	return nil
}

func validateUsername(u string) error {
	matched, err := regexp.MatchString(usernameRegex, u)
	util.Check(err, "invalid regex") // should not happen unless invalid regex
	if !matched {
		return fmt.Errorf("invalid username. must consist of 16 or fewer lowercase letters, _ or -")
	}
	return nil
}

