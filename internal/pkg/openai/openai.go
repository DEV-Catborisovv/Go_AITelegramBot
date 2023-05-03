// OpenAI Client with Golang
// Created by Catborisovv (c) 2020-2024

package openai

import (
	"context"
	"telegramBot/configs"

	openai "github.com/sashabaranov/go-openai"
)

// Функция генерации ответа от ChatGPT3
func GenerateResponse(request string) (error, string) {
	client := openai.NewClient(configs.OPEN_AI_TOKEN)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: request,
				},
			},
		},
	)

	if err != nil {
		return err, ""
	}

	return nil, resp.Choices[0].Message.Content
}
