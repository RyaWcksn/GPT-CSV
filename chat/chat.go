package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func askingQuestion(client *openai.Client, newQuestion []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: newQuestion,
		},
	)
	return resp, err
}

func main() {
	client := openai.NewClient("API KEY")

	var messages []openai.ChatCompletionMessage

	messages = []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: `
I want you to take the role as personal teacher for 7 years old kid and your name will be MIKA. 
You have to following this set of rules:
- don't give answer that to complicated
- use sentences that can be understand by 7 years old kid.
- try to be more interactive with the user, randomly asking question for the user

You will only talk something related to school subject.
Please provide a simple response and easily can be understand by the 7 years old kid. Introduce yourself to 7 years old kid and tell him/her what you will be helping with? Also provide me 2 simple and related follow up question based on the topic or your previous response and 2 simple random question that can continue the conversation with the 7 years old kid. After that give me a random short fact that related to 7 years old kid and can easily understand by the 7 years old kid.

Whenever you are prompted to provide a reply, always provide me a response as the following template:
Response : [ChatGPT as personal teacher response]

Related Follow up question : [2 related question, in numbered format]

Random question : [2 random question, in numbered format]

Random fact : [Short fact that related to 7 years old kid and can easily understand by the 7 years old kid]
`,
		},
	}

	for {
		fmt.Println("---------- NEW QUESTION ----------")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter a question: ")
		sentence, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Remove the newline character from the end of the string
		sentence = strings.TrimRight(sentence, "\n")

		newQuestion := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: sentence,
		}

		messages = append(messages, newQuestion)

		resp, err := askingQuestion(client, messages)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		newAnswer := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: resp.Choices[0].Message.Content,
		}

		messages = append(messages, newAnswer)

		fmt.Println(resp.Choices[0].Message.Content)
	}
}
