package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ChainUnary(l *slog.Logger, interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		LoggingUnaryInterceptor(l),
		// можно добавить другие interceptors
	)
}

func LoggingUnaryInterceptor(
	logger *slog.Logger,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start)

		if err != nil {
			s, _ := status.FromError(err)

			logger.Error(
				"grpc request failed",
				"method", info.FullMethod,
				"code", s.Code().String(),
				"duration", duration,
			)

			return nil, err
		}

		logger.Info(
			"grpc request",
			"method", info.FullMethod,
			"code", "OK",
			"duration", duration,
		)

		return resp, nil
	}
}

func LoggingUnary(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		start := time.Now()

		resp, err := handler(ctx, req)

		log.Info(
			"grpc request",
			"method", info.FullMethod,
			"duration", time.Since(start),
			"error", err,
		)

		return resp, err
	}
}

func ValidationUnary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		// пример простой валидации
		if req == nil {
			return nil, status.Error(3, "empty request")
		}

		return handler(ctx, req)
	}
}
