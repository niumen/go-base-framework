package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"strconv"
	"strings"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPaht := path.Join(logPath, logFileName+".log")
	writerDebug, err := rotatelogs.New(
		fmt.Sprint(path.Join(logPath, logFileName)+"-debugInfo.log"+".%Y%m%d"),
		rotatelogs.WithLinkName(path.Join(logPath, logFileName)+"-debugInfo.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                                             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),                                 // 日志切割时间间隔
	)

	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	//根据等级定义不同的输出
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writerDebug, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writerDebug,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]  %time% - %msg% - %file% \n",
	})

	log.SetReportCaller(true)
	log.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]  %time% - %msg% - %file% \n",
	})

	//添加自定义hook
	log.AddHook(lfHook)
}

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]  %time% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {

	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)
	if entry.Caller != nil {
		output = strings.Replace(output, "%file%", strings.Split(entry.Caller.File, "github.com/")[len(strings.Split(entry.Caller.File, "github.com/"))-1]+":"+strconv.Itoa(entry.Caller.Line), 1)
	}

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)
	for k, v := range entry.Data {
		if s, ok := v.(string); ok {
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}



func Logger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) gin.HandlerFunc {
	logClient := log.New()
	baseLogPaht := path.Join(logPath, logFileName+".log")
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPaht), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                                             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime),                                 // 日志切割时间间隔
	)


	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}


	writeMap := lfshook.WriterMap{
		log.InfoLevel:  writer,
		log.FatalLevel: writer,
		log.DebugLevel: writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.PanicLevel: writer,
	}

	lfHook := lfshook.NewHook(writeMap, &Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]  %time% - %msg%\n",
	})
	logClient.AddHook(lfHook)
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}

}