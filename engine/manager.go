package engine

import (
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type Manager struct {
	sourceMap   map[string]Source
	cron		*cron.Cron
}

func (m *Manager) Register(sources ...Source) error {
	for _, source := range sources {
		m.sourceMap[source.Name()] = source
		if _, err := m.cron.AddFunc(source.Job()); err != nil {
			return err
		}
	}

	for _, source := range sources {
		if _, err := m.cron.AddFunc(source.Job()); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) Deregister() {
	m.cron.Stop()
	for k, v := range m.sourceMap {
		if err := v.Close(); err != nil {
			log.Error(err)
		}

		delete(m.sourceMap, k)
	}
}

func (m *Manager) Get(name string) Source {
	return m.sourceMap[name]
}
