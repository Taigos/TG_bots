package storage

const (
	StateMainMenu = iota
	StateWaitingForName
	StateWaitingForNumbers
)

type Database interface {
	SetUserState(userID int64, state int)
	GetUserState(userID int64) int
}

type MemoryStorage struct {
	userStates map[int64]int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		userStates: make(map[int64]int),
	}
}

func (m *MemoryStorage) SetUserState(userID int64, state int) {
	m.userStates[userID] = state
}

func (m *MemoryStorage) GetUserState(userID int64) int {
	return m.userStates[userID]
}
