package health

import "log/slog"

type Service interface {
	HealthCheck() (string, *slog.Logger)
}

type svc struct {
	env    string
	logger *slog.Logger
}

func NewService(env string, logger *slog.Logger) *svc {
	return &svc{
		env:    env,
		logger: logger,
	}
}

func (s *svc) HealthCheck() (string, *slog.Logger) {
	return s.env, s.logger
}
