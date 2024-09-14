package values

import (
	"go-short-me/internal/component/gsm/core/domain/seedwork/validator"
)

type Url struct {
	Value string
}

func NewUrl(value string) (*Url, error) {
	err := validator.UrlParse(value)
	if err != nil {
		return nil, err
	}

	return &Url{Value: value}, nil
}
