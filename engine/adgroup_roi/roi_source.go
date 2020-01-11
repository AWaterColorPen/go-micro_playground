package adgroup_roi

import (
	"gomicro-playground/engine"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


type RoiSource struct {
	name 		string

}

func (r *RoiSource) Name() string {
	return r.name
}

func (r *RoiSource) Job() (string, func ()) {
	spec := "0/10 * * * *"
	job := func() {

	}
	return spec, job
}

func (r *RoiSource) Close() {

}

func NewSource(option Option) engine.Provider {
	db, err := gorm.Open("mysql", option.DbSource)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(80)
	db.DB().SetMaxOpenConns(80)
	db.DB().SetConnMaxLifetime(59 * time.Second)

	roi := &RoiProvider{}
	roi.db = db
	return roi
}

