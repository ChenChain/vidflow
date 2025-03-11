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
	a, b := cli.Chat(context.TODO(), "7478984023528865829", "7478984023528865829", "")
	fmt.Println(a, b)
}

// ![图片](
// https://lf-bot-studio-plugin-resource.coze.cn/obj/bot-studio-platform-plugin-tos/artist/image/352310b8ae214f3b9ce3e749314da96f.png
func Test_Text2Pic(t *testing.T) {
	cli, _ := NewCli()
	a, b := cli.Chat(context.TODO(), "7478991161739804684", "111", "半亩方塘一鉴开，天光云影共徘徊。问渠那得清如许？为有源头活水来。")
	fmt.Println(a, b)
}

// https://lf-bot-studio-plugin-resource.coze.cn/obj/bot-studio-platform-plugin-tos/artist/image/21a0282ad1314d1b9a502f336202c6e2.wav
func Test_GenBgm(t *testing.T) {
	cli, _ := NewCli()
	a, b := cli.Chat(context.TODO(), "7480421964000067622", "111", "半亩方塘一鉴开，天光云影共徘徊。问渠那得清如许？为有源头活水来。")
	fmt.Println(a, b)
}
