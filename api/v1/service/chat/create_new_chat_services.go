package servicechat

import (
	"context"
	"fmt"
	dtochat "github.com/RyaWcksn/nann-e/dtos/chat"
	dtoroles "github.com/RyaWcksn/nann-e/dtos/roles"
	dtochild "github.com/RyaWcksn/nann-e/dtos/user_child"
	entitychat "github.com/RyaWcksn/nann-e/entities/chat"
	"github.com/sashabaranov/go-openai"
)

var messages []openai.ChatCompletionMessage

func (c *ChatService) CreateNewChat(ctx context.Context, payload *dtochat.CreateNewChatRequest) (*entitychat.CreateNewChatDetail, error) {
	functionName := "ChatService.CreateNewChat"

	client := openai.NewClient("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")

	parentId := ctx.Value("ctxParentId").(string)
	payload.ParentId = parentId

	childDetails, getChildErr := c.childRepo.GetOneUserChild(ctx, &dtochild.GetOneUserChildRequest{
		ParentId:  parentId,
		ChildName: payload.ChildName,
	})
	if getChildErr != nil {
		c.l.Errorf("[%s : c.childRepo.GetOneUserChild] : %s", functionName, getChildErr)
		return nil, getChildErr
	}

	roleDetails, getRoleErr := c.rolesRepo.GetOneRole(ctx, &dtoroles.GetOneRoleRequest{
		ParentId: parentId,
		RoleName: childDetails.RoleName,
	})
	if getRoleErr != nil {
		c.l.Errorf("[%s : c.childRepo.GetOneUserChild] : %s", functionName, getRoleErr)
		return nil, getRoleErr
	}

	if len(messages) == 0 && len(messages) < 2 {

		messages = []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf(`
I want you to take the role as %s for %s named %s and your name will be %s. 
You have to following this set of rules:
%s

You will only talk something related to %s.
Please provide a simple response and easily can be understand by the %s. 
Introduce yourself to the %s and tell him/her what you will be helping with. Also provide me 2 simple and
related follow up question based on the topic or your previous response and 2 simple random question that can
continue the conversation with the %s. After that give me a random short fact that related to %s and can
easily understand by the %s.

Whenever you are prompted to provide a reply, always provide me a response as the following template:
Response : [ChatGPT as %s response]

Related Follow up question : [2 related question, in numbered format]

Random question : [2 random question, in numbered format]

Random fact : [Short fact that related to %s and can easily understand by the %s]
`,
					roleDetails.RoleDescription,
					roleDetails.ChildDescription,
					childDetails.ChildName,
					childDetails.RoleName,
					roleDetails.Rules,
					roleDetails.Topic,
					roleDetails.ChildDescription,
					roleDetails.ChildDescription,
					roleDetails.ChildDescription,
					roleDetails.ChildDescription,
					roleDetails.ChildDescription,
					childDetails.RoleName,
					roleDetails.ChildDescription,
					roleDetails.ChildDescription),
			},
		}
	}

	newQuestion := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: payload.Question,
	}

	messages = append(messages, newQuestion)

	resp, err := c.askingQuestion(client, messages)
	if err != nil {
		return nil, err
	}

	newAnswer := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	}

	messages = append(messages, newAnswer)

	payload.Answer = resp.Choices[0].Message.Content
	createChatErr := c.chatRepo.CreateNewChat(ctx, payload)
	if createChatErr != nil {
		return nil, createChatErr
	}

	res := entitychat.CreateNewChatDetail{
		ChildName: childDetails.ChildName,
		RoleName:  childDetails.RoleName,
		Question:  payload.Question,
		Answer:    payload.Answer,
	}

	return &res, nil
}

func (c *ChatService) askingQuestion(client *openai.Client, newQuestion []openai.ChatCompletionMessage) (*openai.
	ChatCompletionResponse, error) {
	functionName := "ChatService.askingQuestion"
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: newQuestion,
		},
	)
	if err != nil {
		c.l.Errorf("[%s : client.CreateChatCompletion] : %s", functionName, err)
		return nil, err
	}
	return &resp, err
}
