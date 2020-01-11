package engine

type Source interface {
	Name() string
	Job() (string, func ())
	Close() error
}