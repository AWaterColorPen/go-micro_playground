package druid

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/ahmetb/go-linq"
    "github.com/shunfei/godruid"
    log "github.com/sirupsen/logrus"
    "time"
)

type SingleQuery interface {
    BuildQuery() (*godruid.Query, error)
    ParseResult() (interface{}, error)
}

func logHandler(v interface{}) {
    b, _ := json.Marshal(v)
    log.Info(string(b))
}

func query(client *godruid.Client, single SingleQuery) error {
    query, err := single.BuildQuery()
    if err != nil {
        log.Error(err)
        return err
    }

    logHandler(query)

    if err := client.Query(*query); err != nil {
        log.Error(err)
        return err
    }

    logHandler(query)

    return wrapper.ParseResult()
}

func parseTime(t time.Time) string {
    return t.Format("2006-01-02T15:04:05-0700")
}

func buildIntervals(t []time.Time) ([]string, error) {
    if len(t) != 2 {
        return nil, errors.New("invalid time intervals")
    }

    interval := fmt.Sprintf("%v/%v", parseTime(t[0]), parseTime(t[1]))
    return []string{interval}, nil
}

func buildDimensions(dimensions []string) []godruid.DimSpec {
    ds := make([]godruid.DimSpec, 0)
    linq.From(dimensions).Select(func(v interface{}) interface{} {
        return v
    }).ToSlice(&ds)
    return ds
}

func buildTopN(dataSource string) *godruid.Query {
    return nil
}

func buildGroupBy(dataSource string, intervals []time.Time, dimensions []string) (*godruid.Query, error) {
    // _intervals, err := buildIntervals(intervals)
    // if err != nil {
    //     return nil, err
    // }
    //
    // _dimensions := buildDimensions(dimensions)

    return nil, nil
}

func queryTopN() {
    query := &godruid.QueryTopN{
        DataSource:   "stream-report-request",
        Intervals:    []string{"2020-03-10T00:00:00+00:00/2020-03-17T00:00:00+00:00"},
        Granularity:  godruid.GranAll,
        Dimension:    "adGroupId",
        Aggregations: []godruid.Aggregation{godruid.AggCount("count")},
        Filter:       godruid.FilterNot(godruid.FilterSelector("city", nil)),
        Metric:       godruid.TopNMetricNumeric("count"),
        Threshold:    10000,
    }

    client := godruid.Client{
        Url:   "http://9.138.128.91:8082",
    }

    if err := client.Query(query); err != nil {
        log.Print(err)
    }

    b, _ := json.Marshal(query.QueryResult)
    log.Print(string(b))
}

func queryGroupBy() {
    query := &godruid.QueryGroupBy{
        DataSource:   "stream-report-request",
        Intervals:    []string{"2020-03-10T00:00:00+00:00/2020-03-17T00:00:00+00:00"},
        Granularity:  godruid.GranAll,
        Dimensions:   []godruid.DimSpec{"adGroupId"},
        Aggregations: []godruid.Aggregation{godruid.AggCount("count")},
        Filter:        godruid.FilterNot(godruid.FilterSelector("city", nil)),
    }

    client := godruid.Client{
        Url:   "http://9.138.128.91:8082",
    }

    if err := client.Query(query); err != nil {
        log.Print(err)
    }

    b, _ := json.Marshal(query.QueryResult)
    log.Print(string(b))
}
