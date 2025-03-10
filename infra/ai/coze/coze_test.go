package coze

import (
	"context"
	"fmt"
	"testing"
)

func Test_GenOAuth(t *testing.T) {
	token, ex, err := genOAuth(context.TODO())
	fmt.Println(token, ex, err)
}

func Test_DoChat(t *testing.T) {
	cli, _ := NewCli()
	a, b := cli.Chat(context.TODO(), "7478984023528865829", "7478984023528865829")
	fmt.Println(a, b)
}
