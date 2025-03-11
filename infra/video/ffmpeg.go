package video

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

func GenerateVideoBySinglePicAndMusic(picUrl, musicUrl, videoPath string) error {
	ffmpeg := "ffmpeg -loop 1 -i %s -i %s -c:v libx264 -tune stillimage -c:a aac -b:a 192k -pix_fmt yuv420p -shortest %s"
	ffmpeg = fmt.Sprintf(ffmpeg, picUrl, musicUrl, videoPath)
	err := exec.Command("/bin/sh", "-c", ffmpeg).Run()
	return errors.Wrap(err, "ffmpeg failed")
}

func CutNSecondMusic(musicUrl, targetUrl string, n int64) error {
	cmd := "ffmpeg -i %s -t %d -c copy %s"
	cmd = fmt.Sprintf(cmd, musicUrl, n, targetUrl)
	err := exec.Command("/bin/sh", "-c", cmd).Run()
	str := exec.Command("/bin/sh", "-c", cmd).String()
	return errors.Wrap(err, str)
}
