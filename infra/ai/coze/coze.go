package coze

import (
	"context"
	"fmt"
	"github.com/coze-dev/coze-go"
	"github.com/pkg/errors"
	"io"
	"time"
)

type Cli struct {
	Token  string
	Expire int64
}

func NewCli() (*Cli, error) {
	t, ex, err := genOAuth(context.TODO())
	if err != nil {
		return nil, err
	}
	return &Cli{
		Token:  t,
		Expire: ex,
	}, nil
}

func (c *Cli) selfCheck(ctx context.Context) error {
	if time.Now().Unix() > c.Expire {
		t, ex, err := genOAuth(ctx)
		if err != nil {
			return err
		}
		c.Token = t
		c.Expire = ex
	}
	return nil
}

// pom :7478984023528865829
func (c *Cli) Chat(ctx context.Context, botId, userId, msg string) (string, error) {
	if err := c.selfCheck(ctx); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	token := c.Token
	botID := botId
	userID := userId

	authCli := coze.NewTokenAuth(token)

	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	req := &coze.CreateChatsReq{
		BotID:  botID,
		UserID: userID,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText(msg, nil),
		},
	}

	resp, err := cozeCli.Chat.Stream(ctx, req)
	if err != nil {
		return "", errors.Wrapf(err, "Error starting chats")
	}
	res := ""
	defer resp.Close()
	for {
		event, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", errors.Wrapf(err, "Error receiving event")
		}
		if event.Event == coze.ChatEventConversationMessageDelta {
			res += event.Message.Content
		} else if event.Event == coze.ChatEventConversationChatCompleted {
			fmt.Printf("Token usage:%d\n", event.Chat.Usage.TokenCount)
		}
	}
	return res, err
}
