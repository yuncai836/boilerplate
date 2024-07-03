package entry

import (
	"context"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"github.com/yuncai836/boilerplates/otel"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func GracefulServe[T any](c *T, cs chan *T, spawner func(ctx context.Context, c *T) (func(), error)) {
	ss := make(chan os.Signal, 1)
	defer close(ss)
	signal.Notify(ss, os.Interrupt)
	var ctx = context.Background()
	ctx, span := otel.Tracer.Start(ctx, "graceful server start")
	shutdown, err := spawner(ctx, c)
	if err != nil {
		otelzap.L().ErrorContext(ctx, "spawn serve task fail", zap.Error(err))
		span.End()
		return
	}
	span.End()
	for {
		select {
		case c = <-cs:
			var ctx = context.Background()
			ctx, span := otel.Tracer.Start(ctx, "graceful server reload")
			shutdown()
			shutdown, err = spawner(ctx, c)
			if err != nil {
				otelzap.L().ErrorContext(ctx, "reload server fail", zap.Error(err), zap.Any("config", c))
				span.End()
				continue
			}
			otelzap.L().InfoContext(ctx, "reload server success")
			span.End()
		case s := <-ss:
			otelzap.L().Info("receive os signal", zap.String("signal", s.String()))
			shutdown()
			return
		}
	}
}
