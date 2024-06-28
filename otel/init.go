package otel

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init(ctx context.Context, logger *zap.Logger, serviceName, grpcEndpoint, httpEndpoint string) func() {
	shutdownTracer := InitTracer(ctx, logger, serviceName, grpcEndpoint)
	shutdownLog := InitLog(ctx, logger, serviceName, httpEndpoint)
	shutdownMeter := InitMeter(ctx, logger, serviceName, grpcEndpoint)
	return func() {
		shutdownTracer()
		shutdownLog()
		shutdownMeter()
	}
}

func InitFromViper(ctx context.Context, logger *zap.Logger, viper *viper.Viper) func() {
	return Init(
		ctx,
		logger,
		viper.GetString("otel.service"),
		viper.GetString("otel.endpoints.grpc"),
		viper.GetString("otel.endpoints.http"),
	)
}
