package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/syslog"
	"runtime/debug"
	"time"

	log "github.com/sirupsen/logrus"
)

func errTest() {
	err := errors.New("db transaction failed")

	err = fmt.Errorf("GetTransaction(): %w; trace: %s", err, debug.Stack())

	fmt.Println(err)
}

func logrus() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func logrusV2() {
	ctxLogger := log.WithFields(log.Fields{
		"component":   "",
		"action":      "",
		"timestamp":   "",
		"description": "",
		"userID":      "",
		"data":        "",
	})

	ctxLogger.Info("http", "")
}

// Logger with fields
type Logger struct {
	log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: log.New(),
	}
}

func (l *Logger) Info(component, action, desc, userID string, data interface{}) {
	bs, _ := json.Marshal(data)
	l.log.WithFields(log.Fields{
		"component": component,
		"action":    action,
		"timestamp": time.Now(),
		"userID":    userID,
		"data":      string(bs),
	}).Info(desc)

	//l.log.Infof("component = %s; action = %s; timestamp = %s; userID = %s; data = %s\n", component, action, time.Now(), userID, string(bs))
}

func logrusV3() {
	data := struct {
		Title       string
		Description string
		DueDate     time.Time
		Completed   bool
	}{
		Title:       "work1",
		Description: "need to do work1",
		DueDate:     time.Now().Add(24 * time.Hour),
		Completed:   false,
	}

	logger := NewLogger()
	logger.Info("usecase", "main", "info logging", "1", data)
}

// LoggerV2 with formatted fields
type LoggerV2 struct {
	log *log.Logger
}

func NewLoggerV2() *LoggerV2 {
	l := &LoggerV2{
		log: log.New(),
	}

	l.log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		DataKey:           "",
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "@timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
		CallerPrettyfier: nil,
		PrettyPrint:      true,
	})
	return l
}

func (l *LoggerV2) Info(component, action, desc, userID string, data interface{}) {
	bs, _ := json.Marshal(data)
	l.log.WithFields(log.Fields{
		"component": component,
		"action":    action,
		"timestamp": time.Now(),
		"userID":    userID,
		"data":      string(bs),
	}).Info(desc)

}

func logrusV4() {
	data := struct {
		Title       string
		Description string
		DueDate     time.Time
		Completed   bool
	}{
		Title:       "work1",
		Description: "need to do work1",
		DueDate:     time.Now().Add(24 * time.Hour),
		Completed:   false,
	}

	logger := NewLogger()
	logger.Info("usecase", "main", "info logging", "1", data)
}

/*
SyslogHook is allows you to perform additional actions for each log entry, such as sending the log entry to a remote service,
filtering out log entries based on some criteria, or formatting the log entry in a specific way.
*/

type SyslogHook struct {
	writer *syslog.Writer
}

func NewSyslogHook() (*SyslogHook, error) {
	writer, err := syslog.New(syslog.LOG_INFO, "my-app-name")
	if err != nil {
		return nil, err
	}
	return &SyslogHook{writer}, nil
}

func (hook *SyslogHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook *SyslogHook) Fire(entry *log.Entry) error {
	msg, err := entry.String()
	if err != nil {
		return err
	}
	switch entry.Level {
	case log.PanicLevel:
		return hook.writer.Crit(msg)
	case log.FatalLevel:
		return hook.writer.Crit(msg)
	case log.ErrorLevel:
		return hook.writer.Err(msg)
	case log.WarnLevel:
		return hook.writer.Warning(msg)
	case log.InfoLevel:
		return hook.writer.Info(msg)
	case log.DebugLevel:
		return hook.writer.Debug(msg)
	default:
		return nil
	}
}

func logrusV5() {
	_log := log.New()

	// Add a new syslog hook to the logger
	syslogHook, err := NewSyslogHook()
	if err != nil {
		log.Fatal(err)
	}

	// Add a new hook to the logger
	_log.AddHook(syslogHook)

	// Log a message
	_log.Info("Hello, world!")
}

func main() {
	//errTest()

	//logrus()

	//logrusV2()

	//logrusV3()

	//logrusV4()

	//logrusV5()

}

/*
	component
	action
	dynamic field
	timestamp
	description
	publicID
	data
	log level
*/

/*
	>> logger <<

	better use web hook

	>> arch pattern <<

	discuss cases:
		- config
		- several repo instances, separate delivery/usecase/repository folders for each of instances | all will be in one

*/
