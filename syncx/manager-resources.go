package syncx

import "sync"

// manager resources
type (
	ManagerResources struct {
		resources interface{}
		rw        sync.RWMutex
		generate  func() interface{}
		equal     func(a, b interface{}) bool
	}
)

func NewManagerResources(generate func() interface{},
	equal func(a, b interface{}) bool) *ManagerResources {

	return &ManagerResources{
		generate: generate,
		equal:    equal,
	}
}

// if the resources equal `m.resources` then broken.
func (m *ManagerResources) BrokenResources(resources interface{}) {
	m.rw.Lock()
	defer m.rw.Lock()
	if m.equal(m.resources, resources) {
		m.resources = nil
	}
}

// get resources
func (m *ManagerResources) Take() interface{} {
	m.rw.RLock()
	resources := m.resources
	m.rw.RUnlock()

	if resources != nil {
		return resources
	}

	m.rw.Lock()
	defer m.rw.Unlock()

	// maybe other take generate new resources.
	if m.resources == nil {
		m.resources = m.generate()
	}

	return m.resources
}
