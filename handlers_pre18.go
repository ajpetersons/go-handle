// +build !go1.8

package handle

type loggingResponseWriter interface {
	commonLoggingResponseWriter
}
