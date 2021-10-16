package logs

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"time"
	"todoList/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 2016-09-27 09:38:21.541541811 +0200 CEST
// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700]
// "GET /apache_pb.gif HTTP/1.0" 200 2326
// "http://www.example.com/start.html"
// "Mozilla/4.08 [en] (Win98; I ;Nav)"

var timeFormat = "02/Jan/2006:15:04:05 -0800"

// Logger is the logrus logger handler
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.Config.LogFilePath
	logFileName := config.Config.LogFileName

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	fmt.Println(fileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Print(err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
