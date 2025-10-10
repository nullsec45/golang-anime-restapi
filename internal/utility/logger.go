package utility

import (
	"encoding/json"
	"io"
	"path/filepath"
	"github.com/sirupsen/logrus"	
	"os"
	"strings"
)

func CreateLog(level string, message string, category string, fields ...logrus.Fields) {
	level=strings.ToLower(strings.TrimSpace(level))

	if level == "" {
		level="info"
	}

	if category == "" {
		category="application"
	}

	_= os.MkdirAll("logs", 0o755)
	filePath := "logs/application.log"
	if strings.EqualFold(category, "activity") {
		filePath = "logs/activity.log"
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0o666)

	if err != nil {
		logrus.WithError(err).Error("Failed to open log file, fallback to stderr")
	}else{
		defer f.Close()
	}

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetOutput(io.MultiWriter(f))
	logger.SetFormatter(&logrus.JSONFormatter{})


	entry := logger.WithField("file", filepath.Base(filePath))
	if len(fields) > 0 && fields[0] != nil {
		entry = entry.WithFields(fields[0])	
	}	

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl=logrus.InfoLevel
	}

	entry.Log(lvl, message)
}

func CreateLogWithPayload(level, message, category string, payload any) {
	fields := logrus.Fields{
		"payload":SafePayload(payload, 8<<10),
	}

	CreateLog(level, message, category, fields)
}

func SafePayload(v any, maxBytes int) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "<payload_unserializable>"
	}
	if maxBytes > 0 && len(b) > maxBytes {
		return string(b[:maxBytes]) + "...<truncated>"
	}
	return string(b)
}