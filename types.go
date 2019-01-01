package radiobot

import (
	"strconv"

	"github.com/google/uuid"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Channel represent radio channel
type Channel struct {
	ID      uuid.UUID `json:"id" bson:"_id"`
	Name    string    `json:"name" bson:"name"`
	Comment string    `json:"comment" bson:"comment"`
}

// Chat telegram info
type Chat struct {
	ID int64 `json:"id" bson:"_id"`

	// See telebot.ChatType and consts.
	Type telebot.ChatType `json:"type" bson:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title" bson:"title"`

	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Username  string `json:"username" bson:"username"`

	SubscribedChannelsID []uuid.UUID
}

// Recipient returns chat ID (see Recipient interface).
func (c *Chat) Recipient() string {
	if c.Type == telebot.ChatChannel {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}
