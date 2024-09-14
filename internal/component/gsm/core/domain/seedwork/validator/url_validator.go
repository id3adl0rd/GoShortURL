package validator

import (
	"errors"
	"net/url"
)

func UrlParse(value string) error {
	if value == "" {
		return errors.New("url is empty")
	}

	_, err := url.Parse(value)
	if err != nil {
		return err
	}

	return nil
}
