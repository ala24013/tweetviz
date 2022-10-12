package tweetviz

import (
	"github.com/goccy/go-json"
)

type message struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// CreateTweetlistMessage generates a message to provide the current
// tweetlist to the client
func CreateTweetlistMessage(msg []byte) *message {
	return &message{
		"tweetlist",
		string(msg),
	}
}

// CreateDoneLoadingMessage generates a message to tell the client
// that the new stream has been established with twitter
func CreateDoneLoadingMessage() *message {
	return &message{
		"doneLoading",
		"",
	}
}

// serialize returns the jsonified version of the message
func (m *message) serialize() ([]byte, error) {
	return json.Marshal(m)
}
