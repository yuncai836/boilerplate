package entry

import (
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func GracefulServe[T any](c *T, cs chan *T, spawner func(c *T) (func(), error)) {
	ss := make(chan os.Signal, 1)
	defer close(ss)
	signal.Notify(ss, os.Interrupt)
	shutdown, err := spawner(c)
	if err != nil {
		otelzap.L().Error("create grpc serve task fail", zap.Error(err))
		return
	}
	for {
		select {
		case c = <-cs:
			shutdown()
			shutdown, err = spawner(c)
			if err != nil {
				otelzap.L().Error("reload grpc server fail", zap.Error(err), zap.Any("config", c))
			}
		case s := <-ss:
			otelzap.L().Info("receive os signal", zap.String("signal", s.String()))
			shutdown()
			return
		}
	}
}
