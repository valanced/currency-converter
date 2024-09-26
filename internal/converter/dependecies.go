package converter

import "context"

type Rater interface {
	FetchRate(_ context.Context, from, to string) (float64, error)
}
