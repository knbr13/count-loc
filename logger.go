package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	// LogLevelDebug is for debug messages
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is for informational messages
	LogLevelInfo
	// LogLevelWarn is for warning messages
	LogLevelWarn
	// LogLevelError is for error messages
	LogLevelError
	// LogLevelSilent suppresses all log output
	LogLevelSilent
)

// Logger provides thread-safe logging functionality
type Logger struct {
	level      LogLevel
	logger     *log.Logger
	errorLog   *log.Logger
	mu         sync.Mutex
	errorCount int
	warnCount  int
}

// Global logger instance
var defaultLogger *Logger

func init() {
	defaultLogger = NewLogger(LogLevelInfo, os.Stdout, os.Stderr)
}

// NewLogger creates a new Logger instance
func NewLogger(level LogLevel, out io.Writer, errOut io.Writer) *Logger {
	return &Logger{
		level:    level,
		logger:   log.New(out, "", 0),
		errorLog: log.New(errOut, "", 0),
	}
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// SetOutput sets the output writer
func (l *Logger) SetOutput(out io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.SetOutput(out)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelDebug {
		l.logger.Printf("[DEBUG] "+format, args...)
	}
}

// Info logs an informational message
func (l *Logger) Info(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelInfo {
		l.logger.Printf("[INFO] "+format, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelWarn {
		l.warnCount++
		l.logger.Printf("[WARN] "+format, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelError {
		l.errorCount++
		l.errorLog.Printf("[ERROR] "+format, args...)
	}
}

// GetErrorCount returns the number of errors logged
func (l *Logger) GetErrorCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.errorCount
}

// GetWarnCount returns the number of warnings logged
func (l *Logger) GetWarnCount() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.warnCount
}

// Package-level functions that use the default logger

// SetLogLevel sets the log level for the default logger
func SetLogLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

// LogDebug logs a debug message using the default logger
func LogDebug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// LogInfo logs an informational message using the default logger
func LogInfo(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// LogWarn logs a warning message using the default logger
func LogWarn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// LogError logs an error message using the default logger
func LogError(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// FileError represents an error that occurred while processing a file
type FileError struct {
	FilePath string
	Err      error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("error processing file %s: %v", e.FilePath, e.Err)
}

// NewFileError creates a new FileError
func NewFileError(filePath string, err error) *FileError {
	return &FileError{
		FilePath: filePath,
		Err:      err,
	}
}

// DirectoryError represents an error that occurred while processing a directory
type DirectoryError struct {
	DirPath string
	Err     error
}

func (e *DirectoryError) Error() string {
	return fmt.Sprintf("error processing directory %s: %v", e.DirPath, e.Err)
}

// NewDirectoryError creates a new DirectoryError
func NewDirectoryError(dirPath string, err error) *DirectoryError {
	return &DirectoryError{
		DirPath: dirPath,
		Err:     err,
	}
}

// PermissionError represents a permission denied error
type PermissionError struct {
	Path string
	Err  error
}

func (e *PermissionError) Error() string {
	return fmt.Sprintf("permission denied: %s", e.Path)
}

// NewPermissionError creates a new PermissionError
func NewPermissionError(path string, err error) *PermissionError {
	return &PermissionError{
		Path: path,
		Err:  err,
	}
}

// IsPermissionError checks if an error is a permission error
func IsPermissionError(err error) bool {
	return os.IsPermission(err)
}

// LogFileError logs a file processing error
func LogFileError(filePath string, err error) {
	if IsPermissionError(err) {
		LogWarn("Permission denied: %s", filePath)
	} else {
		LogError("Failed to process file %s: %v", filePath, err)
	}
}

// LogDirectoryError logs a directory processing error
func LogDirectoryError(dirPath string, err error) {
	if IsPermissionError(err) {
		LogWarn("Permission denied for directory: %s", dirPath)
	} else {
		LogError("Failed to process directory %s: %v", dirPath, err)
	}
}
