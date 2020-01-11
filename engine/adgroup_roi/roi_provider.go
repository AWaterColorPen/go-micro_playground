package adgroup_roi

import (
	"fmt"
	"gomicro-playground/engine"
	"github.com/ahmetb/go-linq"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Option struct {
	DbSource	string
}

type RoiProvider struct {
	db 			*gorm.DB
}

func (r *RoiProvider) query(date uint32, advertiserId uint64, campaignId uint64, adgroupId uint64) <- chan interface{} {
	ch := r.fetchChan(date, advertiserId, campaignId, adgroupId)

	out := make(chan interface{})
	go linq.From(ch).Select(func(v interface{}) interface{} {
		u := v.(*insightDetail)
		return u.toAdgroupRoi()
	}).GroupBy(func(v interface{}) interface{} {
		return "1"
	}, func(v interface{}) interface{} {
		return v
	}).Select(func(v interface{}) interface{} {
		g := v.(linq.Group)
		return linq.From(g.Group).Aggregate(func(v1 interface{}, v2 interface{}) interface{} {
			u1 := v1.(*adgroupRoi)
			u2 := v2.(*adgroupRoi)
			return merge(u1, u2)
		})
	}).GroupBy(func(v interface{}) interface{} {
		return "1"
	}, func(v interface{}) interface{} {
		return v
	}).ToChannel(out)

	return out
}

func (r *RoiProvider) fetchChan(date uint32, advertiserId uint64, campaignId uint64, adgroupId uint64) <- chan interface{} {
	where := fmt.Sprintf("trace_time >= %v", date)
	if advertiserId > 0 {
		where = fmt.Sprintf("%v AND advertiser_id = %v", where, advertiserId)
	}
	if campaignId > 0 {
		where = fmt.Sprintf("%v AND campaignId_id = %v", where, campaignId)
	}
	if adgroupId > 0 {
		where = fmt.Sprintf("%v AND adgroup_id = %v", where, adgroupId)
	}

	return r.fetch(where)
}

func (r *RoiProvider) fetch(query ...interface{}) <- chan interface{} {
	out := make(chan interface{})

	db := r.db.Table(tableName)
	for _, v := range query {
		db = db.Where(v)
	}

	rows, err := db.Rows()
	if err != nil {
		log.Print(err)
		return out
	}

	go func() {
		defer close(out)
		defer rows.Close()
		for rows.Next() {
			v := &insightDetail{}
			if err := db.ScanRows(rows, v); err != nil {
				log.Error(err)
				return
			}

			out <- v
		}
	}()

	return out
}

func (r *RoiProvider) Query(query engine.QueryOption) []interface{} {
	date := 20200101
	ch := r.query(uint32(date), query.AdvertiserId, query.CampaignId, query.AdgroupId)

	var res []interface{}
	linq.From(ch).ToSlice(&res)
	return res
}

func NewProvider(option Option) engine.Provider {
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

func main() {
	dbSource := ""
	provider := NewProvider(Option{DbSource:dbSource})
	provider.Query(nil)
}

