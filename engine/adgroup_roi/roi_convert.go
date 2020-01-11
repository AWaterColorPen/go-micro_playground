package adgroup_roi

import (
	"github.com/ahmetb/go-linq"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

var (
	roiFieldNames []string
)

func (d *insightDetail) toAdgroupRoi() *adgroupRoi {
	item := &adgroupRoi{

	}

	return item
}

func (d *insightDetail) parseRoi() map[int32]*roi {
	m := make(map[string][]interface{})

	v := reflect.ValueOf(d)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		v := reflect.ValueOf(i)
		name := t.Field(i).Name
		for _, n := range roiFieldNames {
			if strings.Contains(name, n) {
				m[n] = append(m[n], v.Field(i).Interface())
			}
		}
	}

	l := linq.From(m).DistinctBy(func(v interface{}) interface{} {
		return len(v.(linq.KeyValue).Value.([]interface{}))
	}).Count()

	if l != 1 {
		log.Error("parse roi array length no match. ", m)
	}

	rm := make(map[int32]*roi)
	for i := 0; i < 31; i++ {
		rm[1] = &roi{
		}
	}

	return rm
}

func merge(x *adgroupRoi, y *adgroupRoi) *adgroupRoi {
	item := &adgroupRoi{

	}


	return item
}

func init()  {
	t := reflect.TypeOf(&roi{})
	for i := 0; i < t.NumField(); i++ {
		roiFieldNames = append(roiFieldNames, t.Field(i).Name)
	}
}