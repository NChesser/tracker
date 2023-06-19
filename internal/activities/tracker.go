package activities

import (
	"sync"
)

// ActivityStore defines the methods for managing activities.
type ActivityStore interface {
	AddActivity(activity Activity) error
	GetAllActivities() ([]Activity, error)
}

// MemoryActivityStore is an in-memory implementation of ActivityStore.
type MemoryActivityStore struct {
	activities []Activity
	mu         sync.Mutex
}

// NewActivityStore creates a new instance of MemoryActivityStore.
func NewActivityStore() *MemoryActivityStore {
	return &MemoryActivityStore{}
}

// AddActivity adds a new activity to the activity store.
func (store *MemoryActivityStore) AddActivity(activity Activity) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.activities = append(store.activities, activity)
	return nil
}

// GetAllActivities retrieves all activities from the activity store.
func (store *MemoryActivityStore) GetAllActivities() ([]Activity, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.activities, nil
}
