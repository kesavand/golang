/*
The zap logger
*/

package zaplog

import (
	"fmt"

	"github.com/kesavand/golang/pkg/log/api"
	"go.uber.org/zap"
)

/*
ZapLogger is the zap suggared logger
*/
type ZapLogger struct {
	zap.SugaredLogger
}

/*
NewZapLogger creates a new instance of the zap logger
*/
func NewZapLogger() (api.Logger, error) {

	var logger *zap.Logger
	var err error

	if logger, err = zap.NewProduction(); err != nil {
		fmt.Println("Unable to Create Logger")
		panic(err)
	}
	return logger.Sugar(), nil
}

//Debug logs all messages
func (log *ZapLogger) Debug(...interface{}) {
}

//Warn logs warnings
func (log *ZapLogger) Warn(...interface{}) {
}

// Err logs only error messages
func (log *ZapLogger) Error(...interface{}) {
}

//Info log level
func (log *ZapLogger) Info(...interface{}) {
}

//Fatal log level
func (log *ZapLogger) Fatal(...interface{}) {
}
