package url

import (
	"context"
	"go-short-me/internal/component/gsm/core/domain/seedwork/interfaces"
	"go-short-me/internal/component/gsm/core/domain/url/values"
)

type Url struct {
	url        values.Url
	createdURL values.CreatedUrl
}

func (u *Url) Url() values.Url {
	return u.url
}

func (u *Url) CreatedURL() values.CreatedUrl {
	return u.createdURL
}

func NewUrl(url string) (*Url, error) {
	u, err := values.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return &Url{url: *u}, nil
}

func (u *Url) CreateShortURL(db interfaces.Database) (string, error) {
	nUrl, err := values.NewCreatedUrl()
	if err != nil {
		return "", err
	}

	u.createdURL.Value = nUrl.Value

	ctx := context.Background()
	err = db.InsertURL(ctx, u)
	if err != nil {
		return "", err
	}

	return u.createdURL.Value, nil
}
