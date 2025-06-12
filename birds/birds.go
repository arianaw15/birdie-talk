package birds

import (
	"context"
	"fmt"
)

func GetBirds(ctx context.Context, name string) (*Response, error) {

	msg := fmt.Sprintf("Hello %s, welcome to the birds API!", name)
	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}
