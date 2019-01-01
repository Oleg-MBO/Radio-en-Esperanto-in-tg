package repository

import (
	"fmt"
	"sync"

	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/google/uuid"
)

type mockupChannelRepository struct {
	cache []*radiobot.Channel
	m     sync.Mutex
}

// NewMockupChannelRepository is mockup repository for channels
func NewMockupChannelRepository() radiobot.ChannelRepository {
	chList := make([]*radiobot.Channel, 0)
	return &mockupChannelRepository{cache: chList}
}

// RegisterChannel is used to register channel
func (ch *mockupChannelRepository) RegisterChannel(radioCh *radiobot.Channel) (uuid.UUID, error) {
	ch.m.Lock()
	defer ch.m.Unlock()
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}
	radioCh.ID = id

	ch.cache = append(ch.cache, radioCh)
	return id, err
}

// FindChannelByID is used to find channel by id
func (ch *mockupChannelRepository) FindChannelByID(id uuid.UUID) (*radiobot.Channel, error) {
	ch.m.Lock()
	defer ch.m.Unlock()
	for _, radioCh := range ch.cache {
		if radioCh.ID == id {
			return radioCh, nil
		}
	}
	return nil, fmt.Errorf("Not Found")
}

// FindChannelByName is used to find channel by name
func (ch *mockupChannelRepository) FindChannelByName(name string) (*radiobot.Channel, error) {
	ch.m.Lock()
	defer ch.m.Unlock()
	for _, radioCh := range ch.cache {
		if radioCh.Name == name {
			return radioCh, nil
		}
	}
	return nil, fmt.Errorf("Not Found")
}

// GetChannels is used get channels
func (ch *mockupChannelRepository) GetChannels(count, offset int) ([]*radiobot.Channel, error) {
	ch.m.Lock()
	defer ch.m.Unlock()

	radioChannels := make([]*radiobot.Channel, 0, count)
	for i := offset; i < offset+count && i < len(ch.cache); i++ {
		radioChannels = append(radioChannels, ch.cache[i])
	}
	return radioChannels, nil
}
