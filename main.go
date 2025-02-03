package main

import (
	"context"

	"golang.org/x/oauth2"
	goauth2 "google.golang.org/api/oauth2/v2"
)

func main() {
	_, _ = goauth2.NewService(context.Background())

	_ = oauth2.StaticTokenSource(oauth2.Token{})
}
