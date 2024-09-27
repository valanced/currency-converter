package converter

import "context"

//go:generate mockgen -source=./dependencies.go -destination=./mocks_test.go -package=converter_test

type Rater interface {
	FetchRate(_ context.Context, from, to string) (float64, error)
}
