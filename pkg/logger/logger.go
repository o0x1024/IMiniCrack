package logger

import (
	"IMiniCrack/pkg/global"
	"fmt"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Custom logger implements wails' logging but emits all logs to frontend
type CustomLogger struct {
	prefix    string
	mu        sync.Mutex
	errorOnly bool
}

// NewCustomLogger creates a new custom logger with the given prefix
func NewCustomLogger(prefix string) *CustomLogger {
	return &CustomLogger{
		prefix: "[" + prefix + "] ",
	}
}

// Info level message
func (c *CustomLogger) Info(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))

}

// Infof - formatted message
func (c *CustomLogger) Infof(message string, args ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, fmt.Sprintf(message, args...)))
}

// InfoFields - message with fields
func (c *CustomLogger) InfoFields(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Debug level message
func (c *CustomLogger) Debug(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Debugf - formatted message
func (c *CustomLogger) Debugf(message string, args ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, fmt.Sprintf(message, args...)))
}

// DebugFields - message with fields
func (c *CustomLogger) DebugFields(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Warn level message
func (c *CustomLogger) Warn(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Warnf - formatted message
func (c *CustomLogger) Warnf(message string, args ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, fmt.Sprintf(message, args...)))
}

// WarnFields - message with fields
func (c *CustomLogger) WarnFields(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Error level message
func (c *CustomLogger) Error(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Errorf - formatted message
func (c *CustomLogger) Errorf(message string, args ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, fmt.Sprintf(message, args...)))
}

// ErrorFields - message with fields
func (c *CustomLogger) ErrorFields(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Fatal level message
func (c *CustomLogger) Fatal(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}

// Fatalf - formatted message
func (c *CustomLogger) Fatalf(message string, args ...interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, fmt.Sprintf(message, args...)))
}

// FatalFields - message with fields
func (c *CustomLogger) FatalFields(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	runtime.EventsEmit(*global.WailsCtx, fmt.Sprintf("%s%s", c.prefix, message))
}
