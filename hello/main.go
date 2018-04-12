package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hidelbreq/cultivate-bot/model"
)

type Response struct {
	Text string `json:"text"`
}

const TextArrow string = "その矢印、どっち向いている？"
const TextExcellent string = "絶好調！"

func ResponseToPost(p *model.SlackPost) *Response {
	if p.TriggerWord == "すみません" {
		return &Response{Text: TextArrow}
	}

	return &Response{Text: TextExcellent}
}

func Handler(p *model.SlackPost) (*Response, error) {
	post, err := p.CopyWithUnescaping()
	fmt.Println(post, err)
	if err != nil {
		return &Response{
			Text: err.Error(),
		}, err
	}

	return ResponseToPost(p), nil
}

func main() {
	lambda.Start(Handler)
}
