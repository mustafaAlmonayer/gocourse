package intermediate

import "go.uber.org/zap"

var ZapLogger *zap.Logger

func zapLogging() {
	ZapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer ZapLogger.Sync()

	ZapLogger.Info("This is an info message.")
}
