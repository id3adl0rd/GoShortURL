package values

import "net/url"

type Url struct {
	Value string
}

func NewUrl(value string) (*Url, error) {
	_, err := url.Parse(value)
	if err != nil {
		return nil, err
	}
	return &Url{Value: value}, nil
}
