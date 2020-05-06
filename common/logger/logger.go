package logger

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/olivere/elastic/v7"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type RotateLogOption struct {
	Dir			string `json:"dir"`
}

type ElasticLogOption struct {
	Url			string `json:"url"`
	UserName	string `json:"username"`
	Password	string `json:"password"`
	IndexName	string `json:"index"`
}

type Metadata struct {
	Ip			string `json:"ip"`
}

var (
	RotateLog RotateLogOption
	ElasticLog ElasticLogOption
	Meta Metadata
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.SetOutput(ioutil.Discard)
}

func rotateLog(dir string, loglevel log.Level) *rotatelogs.RotateLogs {
	logfile := fmt.Sprintf("%v/%v", dir, loglevel)
	rotateLog, err := rotatelogs.New(
		logfile + ".%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(logfile),
		rotatelogs.WithMaxAge(time.Hour * 24),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		log.Error(err)
	}

	return rotateLog
}

func log4local(dir string)  {
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.TraceLevel: rotateLog(dir, log.TraceLevel),
			log.DebugLevel: rotateLog(dir, log.DebugLevel),
			log.InfoLevel:  rotateLog(dir, log.InfoLevel),
			log.WarnLevel:  rotateLog(dir, log.WarnLevel),
			log.ErrorLevel: rotateLog(dir, log.ErrorLevel),
			log.FatalLevel: rotateLog(dir, log.FatalLevel),
		},
		&log.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		},
	))
}

func Log4local() func() {
	return func() {
		if RotateLog.Dir == "" {
			log.Error("log 4 local failed. empty rotate log dir")
			return
		}

		log4local(RotateLog.Dir)
	}
}

func log4elasticSearch() {
	client, err := elastic.NewClient(
		elastic.SetURL(ElasticLog.Url),
		elastic.SetBasicAuth(ElasticLog.UserName, ElasticLog.Password),
		elastic.SetSniff(false))

	if err != nil {
		log.Error(err)
		return
	}

	hook, err := elogrus.NewAsyncElasticHook(client, Meta.Ip, log.InfoLevel, ElasticLog.IndexName)
	if err != nil {
		log.Error(err)
		return
	}

	log.AddHook(hook)
}

func Log4elasticSearch() func() {
	return func() {
		if Meta.Ip == "" ||
			ElasticLog.Url == "" ||
			ElasticLog.UserName == "" ||
			ElasticLog.Password == "" ||
			ElasticLog.IndexName == "" {
			log.Error("log 4 elastic search failed. invalid meta option or elastic log option", Meta, ElasticLog)
			return
		}

		log4elasticSearch()
	}
}
