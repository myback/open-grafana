package grpcplugin

import (
	"fmt"
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	glog "github.com/myback/open-grafana/pkg/infra/log"
)

type logWrapper struct {
	Logger glog.Logger

	name        string
	impliedArgs []interface{}
}

func formatArgs(args ...interface{}) []interface{} {
	if len(args) == 0 || len(args)%2 != 0 {
		return args
	}

	var res []any

	for n := 0; n < len(args); n += 2 {
		key := args[n]

		if stringKey, ok := key.(string); ok && stringKey == "timestamp" {
			continue
		}

		res = append(res, key)
		res = append(res, args[n+1])
	}

	return res
}

// Log Emit a message and key/value pairs at a provided log level
func (lw logWrapper) Log(level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.Trace:
		lw.Trace(msg, args...)
	case hclog.Debug:
		lw.Debug(msg, args...)
	case hclog.Warn:
		lw.Warn(msg, args...)
	case hclog.Error:
		lw.Error(msg, args...)
	default:
		lw.Info(msg, args...)
	}
}

// Trace Emit a message and key/value pairs at the TRACE level
func (lw logWrapper) Trace(msg string, args ...interface{}) {
	lw.Logger.Debug(msg, formatArgs(args...)...)
}

// Debug Emit a message and key/value pairs at the DEBUG level
func (lw logWrapper) Debug(msg string, args ...interface{}) {
	lw.Logger.Debug(msg, formatArgs(args...)...)
}

// Info Emit a message and key/value pairs at the INFO level
func (lw logWrapper) Info(msg string, args ...interface{}) {
	lw.Logger.Info(msg, formatArgs(args...)...)
}

// Warn Emit a message and key/value pairs at the WARN level
func (lw logWrapper) Warn(msg string, args ...interface{}) {
	lw.Logger.Warn(msg, formatArgs(args...)...)
}

// Emit a message and key/value pairs at the ERROR level
func (lw logWrapper) Error(msg string, args ...interface{}) {
	lw.Logger.Error(msg, formatArgs(args...)...)
}

// IsTrace Indicate if TRACE logs would be emitted.
func (lw logWrapper) IsTrace() bool { return true }

// IsDebug Indicate if DEBUG logs would be emitted.
func (lw logWrapper) IsDebug() bool { return true }

// IsInfo Indicate if INFO logs would be emitted.
func (lw logWrapper) IsInfo() bool { return true }

// IsWarn Indicate if WARN logs would be emitted.
func (lw logWrapper) IsWarn() bool { return true }

// IsError Indicate if ERROR logs would be emitted.
func (lw logWrapper) IsError() bool { return true }

// ImpliedArgs returns With key/value pairs
func (lw logWrapper) ImpliedArgs() []interface{} {
	return lw.impliedArgs
}

// With Creates a sublogger that will always have the given key/value pairs
func (lw logWrapper) With(args ...interface{}) hclog.Logger {
	return logWrapper{
		Logger:      lw.Logger.New(args...),
		name:        lw.name,
		impliedArgs: args,
	}
}

// Name Returns the Name of the logger
func (lw logWrapper) Name() string {
	return lw.name
}

// Named Create a logger that will prepend the name string on the front of all messages.
// If the logger already has a name, the new value will be appended to the current
// name.
func (lw logWrapper) Named(name string) hclog.Logger {
	if name == "stdio" {
		// discard logs from stdio hashicorp/go-plugin gRPC service since
		// it's not enabled/in use per default.
		// discard debug log of "waiting for stdio data".
		// discard warn log of "received EOF, stopping recv loop".
		return hclog.NewNullLogger()
	}

	if lw.name != "" {
		name = fmt.Sprintf("%s.%s", lw.name, name)
	}

	return logWrapper{
		Logger:      lw.Logger.New(),
		name:        name,
		impliedArgs: lw.impliedArgs,
	}
}

// ResetNamed Create a logger that will prepend the name string on the front of all messages.
// This sets the name of the logger to the value directly, unlike Named which honor
// the current name as well.
func (lw logWrapper) ResetNamed(name string) hclog.Logger {
	return logWrapper{
		Logger:      lw.Logger.New(),
		name:        name,
		impliedArgs: lw.impliedArgs,
	}
}

// SetLevel No-op. The wrapped logger implementation cannot update the level on the fly.
func (lw logWrapper) SetLevel(hclog.Level) {}

// GetLevel get logging level
func (lw logWrapper) GetLevel() hclog.Level {
	return hclog.Debug
}

// StandardLogger Return a value that conforms to the stdlib log.Logger interface
func (lw logWrapper) StandardLogger(*hclog.StandardLoggerOptions) *log.Logger {
	return nil
}

// StandardWriter Return a value that conforms to io.Writer, which can be passed into log.SetOutput()
func (lw logWrapper) StandardWriter(*hclog.StandardLoggerOptions) io.Writer {
	return io.Discard
}
