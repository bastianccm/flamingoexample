package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type controller struct{}

func (c *controller) action(ctx context.Context, req *web.Request) web.Result {
	n := rand.Intn(500)

	time.Sleep(time.Duration(n) * time.Millisecond)

	return &web.Response{
		Status: 200,
		Body:   strings.NewReader(fmt.Sprintf("Hello World! I slept for %d ms", n)),
	}
}
