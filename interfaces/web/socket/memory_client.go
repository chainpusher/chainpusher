package socket

type MemoryClient struct {
	id    int64
	emits []interface{}
}

func (m *MemoryClient) GetId() int64 {
	return m.id
}

func (m *MemoryClient) Emit(message interface{}) error {
	m.emits = append(m.emits, message)
	return nil
}

func (m *MemoryClient) Read() chan []byte {
	c := make(chan []byte)
	return c
}

func (m *MemoryClient) Close() error {
	return nil
}

func (m *MemoryClient) GetEmits() []interface{} {
	return m.emits
}

func NewMemoryClient(id int64) Client {
	return &MemoryClient{id: id, emits: make([]interface{}, 0)}
}
