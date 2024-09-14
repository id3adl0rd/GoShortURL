package values

import (
	"go-short-me/internal/component/gsm/core/domain/seedwork/tools"
	"math/rand"
)

type CreatedUrl struct {
	Value string
}

func NewCreatedUrl() (*CreatedUrl, error) {
	generatedUrl := tools.UrlGenerator(rand.Uint64())

	return &CreatedUrl{Value: generatedUrl}, nil
}
