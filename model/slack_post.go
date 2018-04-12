package model

import "net/url"

type SlackPost struct {
	Token       string `json:"token"`
	TeamId      string `json:"team_id"`
	ChannelId   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	ThreadTs    string `json:"thread_ts"`
	Timestamp   string `json:"timestamp"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Text        string `json:"text"`
	TriggerWord string `json:"trigger_word"`
}

func (s *SlackPost) CopyWithUnescaping() (*SlackPost, error) {
	token, err := url.QueryUnescape(s.Token)
	if err != nil {
		return nil, err
	}
	teamId, err := url.QueryUnescape(s.TeamId)
	if err != nil {
		return nil, err
	}
	channelId, err := url.QueryUnescape(s.ChannelId)
	if err != nil {
		return nil, err
	}
	channelName, err := url.QueryUnescape(s.ChannelName)
	if err != nil {
		return nil, err
	}
	threadTs, err := url.QueryUnescape(s.ThreadTs)
	if err != nil {
		return nil, err
	}
	timestamp, err := url.QueryUnescape(s.Timestamp)
	if err != nil {
		return nil, err
	}
	userId, err := url.QueryUnescape(s.UserId)
	if err != nil {
		return nil, err
	}
	userName, err := url.QueryUnescape(s.UserName)
	if err != nil {
		return nil, err
	}
	text, err := url.QueryUnescape(s.Text)
	if err != nil {
		return nil, err
	}
	trigger, err := url.QueryUnescape(s.TriggerWord)
	if err != nil {
		return nil, err
	}
	return &SlackPost{
		Token:       token,
		TeamId:      teamId,
		ChannelId:   channelId,
		ChannelName: channelName,
		ThreadTs:    threadTs,
		Timestamp:   timestamp,
		UserId:      userId,
		UserName:    userName,
		Text:        text,
		TriggerWord: trigger,
	}, nil
}
