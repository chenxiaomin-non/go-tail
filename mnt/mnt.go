package mnt

import (
	sync "sync"

	tail "github.com/chenxiaomin-non/go-tail/tail"
)

var (
	once sync.Once
	mng  *ObsManager
)

const (
	DEFAULT_CONFIG_FILEPATH = "./mnt/test.yaml"
	DEFAULT_MSG_CHAN_SIZE   = 5
)

// ObsManager is a manager for Observers object
// and can be use as a single publisher for all observers
type ObsManager struct {
	observers map[string]*tail.Observer
	msg       chan string
	mu        sync.Mutex
}

// GetObsManager return a singleton ObsManager
func GetObsManager() *ObsManager {
	once.Do(func() {
		mng = &ObsManager{
			observers: map[string]*tail.Observer{},
			msg:       make(chan string, DEFAULT_MSG_CHAN_SIZE),
		}
		mng.Init()
	})
	return mng
}

// AddObserver add a new observer to manager
func (m *ObsManager) AddObserver(obs *tail.Observer, name string) {
	m.observers[name] = obs
	obs.Publisher = m.msg
}

func (m *ObsManager) Init() {
	// parse yaml file
	obsMap, err := ParseYaml(DEFAULT_CONFIG_FILEPATH)
	if err != nil {
		panic(err)
	}

	// init observers & start them
	m.observers = obsMap
	for _, obs := range obsMap {
		obs.Publisher = m.msg
		err := obs.Start()
		if err != nil {
			panic(err)
		}
	}
}

func (m *ObsManager) CountObservers() int {
	return len(m.observers)
}

func (m *ObsManager) PopMessage() string {
	m.mu.Lock()
	msg := <-m.msg
	defer m.mu.Unlock()
	return msg
}
