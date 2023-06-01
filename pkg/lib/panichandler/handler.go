package panichandler

import (
	"fmt"
	"go.uber.org/zap"
	"runtime/debug"
)

type zapHandler struct {
	logger *zap.Logger
}

func (h zapHandler) Handle() {
	if r := recover(); r != nil {
		fmt.Println("got panic: ", r)
		fmt.Println("stacktrace from panic: \n", debug.Stack())
		h.logger.Panic("got panic", zap.ByteString("stackTrace", debug.Stack()))
	}
}

//
func ZapHandler(logger *zap.Logger) zapHandler {
	return zapHandler{
		logger: logger,
	}
}
