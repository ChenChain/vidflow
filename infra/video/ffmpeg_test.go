package video

import "testing"

func Test_GenerateVideo(t *testing.T) {
	p := "../../static/testimg1.jpeg"
	a := "../../static/testaudio1.wav"
	err := CutNSecondMusic(a, "../../static/c2.wav", 6)
	if err != nil {
		panic(err)
	}
	err = GenerateVideoBySinglePicAndMusic(p, "../../static/c1.wav", "../../static/v1.mp4")
	if err != nil {
		panic(err)
	}
}
