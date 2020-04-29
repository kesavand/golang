

package api

/*Logger is the interface for logging library*/
type Logger interface {
        Debug(...interface{})
        Warn(...interface{})
        Error(...interface{})
        Info(...interface{})
        Fatal(...interface{})
}
