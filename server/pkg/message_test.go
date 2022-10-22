package tweetviz

import (
	"bytes"
	"testing"
)

func TestCreateTweetlistMessage(t *testing.T) {
	msg := []byte("hello")
	tmsg := CreateTweetlistMessage(msg)
	if tmsg.Code != "tweetlist" {
		t.Error("invalid code on tweetlist message")
	}
	if tmsg.Msg != "hello" {
		t.Error("invalid msg on tweetlist message")
	}
}

func TestCreateDoneLoadingMessage(t *testing.T) {
	msg := CreateDoneLoadingMessage()
	if msg.Code != "doneLoading" {
		t.Error("invalid code on done loading message")
	}
	if msg.Msg != "" {
		t.Error("invalid msg on done loading message")
	}
}

func TestMessageSerialize(t *testing.T) {
	msg := []byte("hello")
	tmsg := CreateTweetlistMessage(msg)
	s, _ := tmsg.serialize()
	compare := bytes.Compare(s, []byte("{\"code\":\"tweetlist\",\"msg\":\"hello\"}"))
	if compare != 0 {
		t.Error("invalid serialization of message")
	}
}
