package url

import "go-short-me/internal/component/gsm/core/domain/url/values"

type Url struct {
	Url values.Url
}

func NewUrl(url string) (*Url, error) {
	u, err := values.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return &Url{Url: *u}, nil
}

func (u *Url) CreateShortURL(url string) (string, error) {
	nUrl, err := NewUrl(url)
	if err != nil {
		return "", err
	}

	return nUrl.Url.Value, nil
}
