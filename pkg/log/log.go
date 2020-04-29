/*
Package log  library is an abstract over the zap library.
*/
package log

import (
        "github.com/kesavand/golang/pkg/log/zaplog"
        "github.com/kesavand/golang/pkg/log/api"
)


/*
NewLogger Creates new logger based on the logger type
*/
func NewLogger(loglib string) (api.Logger, error) {

        /*call the logger based on the loglib*/
        return zaplog.NewZapLogger()

}

