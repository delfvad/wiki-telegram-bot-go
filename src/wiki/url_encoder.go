package wiki

import (
	"net/url"
)

// Converts request for using as a part of url
func urlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
