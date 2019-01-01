package repository

import (
	radiobot "github.com/Oleg-MBO/Radio-en-Esperanto"
	"github.com/globalsign/mgo"
)

type mongoChatRepository struct {
	Collection *mgo.Collection
}

// NewMongoChatRepository mongo repository for chats
func NewMongoChatRepository(collection *mgo.Collection) radiobot.ChatRepository {
	return &mongoChatRepository{Collection: collection}
}

// RegisterChat is used for register chats
func (mchat *mongoChatRepository) RegisterChat(*radiobot.Chat) error {
	panic("implement me")
}

// FindChat is used to find chat by id
func (mchat *mongoChatRepository) FindChat(id int64) (*radiobot.Chat, error) {
	panic("implement me")
}

// SubscribeChat is used to subscribe chat on channel
func (mchat *mongoChatRepository) SubscribeChat(chat *radiobot.Chat, channel *radiobot.Channel) error {
	panic("implement me")
}

// UnsubscribeChat is used to unsubscribe chat on channel
func (mchat *mongoChatRepository) UnsubscribeChat(*radiobot.Chat, *radiobot.Channel) error {
	panic("implement me")
}

// GetAllChats is used to get all chats with count and offset
func (mchat *mongoChatRepository) GetAllChats(count, offset int) ([]*radiobot.Chat, error) {
	panic("implement me")
}

// GetAllChatsSubscribedOn is used to fetch all chats which subscribed on channel
func (mchat *mongoChatRepository) GetAllChatsSubscribedOn(ch *radiobot.Channel, count, offset int) ([]*radiobot.Chat, error) {
	panic("implement me")
}

// GetAllChatsIDSubscribedOn is used to fetch all chats ID which subscribed on channel
func (mchat *mongoChatRepository) GetAllChatsIDSubscribedOn(ch *radiobot.Channel, count, offset int) ([]int64, error) {
	panic("implement me")
}