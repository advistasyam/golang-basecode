package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	logrus_cloudwatchlogs "github.com/kdar/logrus-cloudwatchlogs"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

func CreateAppLog(cloudwatchSession *session.Session, pd interface{}) *logrus.Entry {
	l := logrus.New()
	logrusCloudwatchHook(cloudwatchSession, l, "golang-basecode-log")

	logData := logrus.Fields{}
	mapstructure.Decode(pd, &logData)

	return l.WithFields(logData)
}

func MiddlewareLogger(cloudwatchSession *session.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			var body map[string]interface{}

			if c.Request().Body != nil {
				bodyBytes, _ := ioutil.ReadAll(c.Request().Body)
				if len(bodyBytes) > 0 {
					errParJson := json.Unmarshal(bodyBytes, &body)

					if errParJson != nil {
						return errors.New("Invalid JSON")
					}
				}

				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			go func() {
				defer Recover("MiddlewareLogger")

				sourceLogData := make(map[string]interface{})
				sourceLogData["at"] = time.Now().Format("02/01/2006 15:04:05 -0700")
				sourceLogData["interface"] = os.Getenv("INTERFACE")
				sourceLogData["method"] = req.Method
				sourceLogData["host"] = req.Host
				sourceLogData["uri"] = req.URL.String()
				sourceLogData["ip"] = req.RemoteAddr
				sourceLogData["real_ip"] = c.RealIP()
				sourceLogData["body"] = body
				sourceLogData["latency"] = strconv.FormatInt(int64(stop.Sub(start)), 10)
				sourceLogData["latency_human"] = stop.Sub(start).String()
				sourceLogData["user_agent"] = req.UserAgent()
				sourceLogData["referer"] = req.Referer()
				sourceLogData["status"] = res.Status
				sourceLogData["path"] = req.URL.Path
				if err != nil {
					sourceLogData["error"] = err.Error()
				}

				l := logrus.New()
				logrusCloudwatchHook(cloudwatchSession, l, "golang-basecode-access")

				logData := logrus.Fields{}
				mapstructure.Decode(sourceLogData, &logData)

				l.WithFields(logData).Info("incoming request")
			}()

			return
		}
	}
}

func logrusCloudwatchHook(cloudwatchSession *session.Session, l *logrus.Logger, stream string) {
	if cloudwatchSession == nil {
		return
	}

	group := os.Getenv("CLOUDWATCH_LOG_GROUP")
	l.Out = ioutil.Discard
	l.Formatter = logrus_cloudwatchlogs.NewProdFormatter()
	hook, err := logrus_cloudwatchlogs.NewHook(group, stream, cloudwatchSession)

	if err != nil {
		return
	}

	l.Hooks.Add(hook)
}
