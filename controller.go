package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/flamingo"

	"flamingo.me/flamingo/v3/framework/opencensus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"

	"go.opencensus.io/trace"

	"flamingo.me/flamingo/v3/framework/web"
)

var measure = stats.Int64("mymeasure", "my measure", stats.UnitMilliseconds)

func init() {
	view.SetReportingPeriod(100 * time.Millisecond)
	opencensus.View(
		"mymetric",
		measure,
		view.Distribution(1, 10, 100, 250, 500),
	)
}

type controller struct {
	logger flamingo.Logger
}

func (c *controller) Inject(logger flamingo.Logger) {
	c.logger = logger.WithField(flamingo.LogKeyModule, "example")
}

func (c *controller) action(ctx context.Context, req *web.Request) web.Result {
	n := rand.Intn(500)

	stats.Record(ctx, measure.M(int64(n)))

	c.logger.WithContext(ctx).WithField("duration", n).Info("sleeping")

	time.Sleep(20 * time.Millisecond)

	ctx, span := trace.StartSpan(ctx, "my/span")
	defer span.End()

	req2, _ := http.NewRequest("GET", "https://google.com", nil)
	req2 = req2.WithContext(ctx)
	r, _ := http.DefaultClient.Do(req2)
	defer r.Body.Close()

	time.Sleep(time.Duration(n) * time.Millisecond)

	return &web.Response{
		Status: 200,
		Body:   strings.NewReader(fmt.Sprintf("Hello World! I slept for %d ms", n)),
	}
}
