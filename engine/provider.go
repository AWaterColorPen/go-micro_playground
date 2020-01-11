package engine

type Provider interface {
	Query(QueryOption) []interface{}
}