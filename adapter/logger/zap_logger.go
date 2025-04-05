package logger

import "go.uber.org/zap"

type ZapLogger struct {
	sugar *zap.SugaredLogger
}

func NewZapLogger() (*ZapLogger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{sugar: l.Sugar()}, nil
}

func (z *ZapLogger) Info(msg string, keysAndValues ...any) {
	z.sugar.Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Error(msg string, keysAndValues ...any) {
	z.sugar.Errorw(msg, keysAndValues...)
}

func (z *ZapLogger) Fatal(msg string, keysAndValues ...any) {
	z.sugar.Fatalw(msg, keysAndValues...)
}
