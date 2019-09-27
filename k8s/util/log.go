package util

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/olivere/elastic/v7"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"io/ioutil"
	"time"
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

var (
	rotateLog RotateLogOption
	elasticLog ElasticLogOption
)

func GetRotateLogOption() *RotateLogOption {
	return &rotateLog
}

func GetElasticLogOption() *ElasticLogOption {
	return &elasticLog
}

func _rotateLog(loglevel log.Level) *rotatelogs.RotateLogs {
	logfile := fmt.Sprintf("%v/log-%v", rotateLog.Dir, loglevel)
	rotateLog, err := rotatelogs.New(
		logfile + ".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logfile),
		rotatelogs.WithMaxAge(time.Hour * 24),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		log.Error(err)
	}

	return rotateLog
}

func _log4localFilesystem()  {
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.TraceLevel: _rotateLog(log.TraceLevel),
			log.InfoLevel:  _rotateLog(log.InfoLevel),
			log.WarnLevel:  _rotateLog(log.WarnLevel),
			log.ErrorLevel: _rotateLog(log.ErrorLevel),
			log.FatalLevel: _rotateLog(log.FatalLevel),
		},
		&log.TextFormatter{FullTimestamp: true},
	))
}

func _log4elasticSearch() {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticLog.Url),
		elastic.SetBasicAuth(elasticLog.UserName, elasticLog.Password),
		elastic.SetSniff(false))

	if err != nil {
		log.Error(err)
		return
	}

	hook, err := elogrus.NewAsyncElasticHook(client, currentIpAddress(), log.TraceLevel, elasticLog.IndexName)
	if err != nil {
		log.Error(err)
		return
	}

	log.AddHook(hook)
}

func Initlog() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	log.SetOutput(ioutil.Discard)
	_log4elasticSearch()
	_log4localFilesystem()
}
