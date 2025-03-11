package img

import (
	"bytes"
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"html/template"
	"image"
	"image/jpeg"
	"io"
	"os"
	"sync"
)

func GenJpegByH5(content, h5, filePath string) error {
	return nil
}

func generateImgData(buf *[]byte, h5Page string) error {
	chromedpCtx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	return chromedp.Run(chromedpCtx, fullScreenshotActions(100, buf, h5Page))
}

func saveBufferAsJpeg(buf []byte, filePath string) error {
	// 将字节缓冲区转换为 image.Image 对象
	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return err
	}

	// 创建目标文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将 image 对象编码为 JPEG 格式并保存到文件
	return jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
}

func fullScreenshotActions(quality int, res *[]byte, html string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			defer cancel()
			var wg sync.WaitGroup
			wg.Add(1)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					cancel()
					wg.Done()
				}
			})

			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			if err := page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx); err != nil {
				return err
			}
			wg.Wait()
			return nil
		}),
		chromedp.FullScreenshot(res, quality),
	}
}

func generateH5WithTemplate(content, h5Template string) (string, error) {
	h5Tpl, err := template.New("generateH5WithTemplate").Parse(h5Template)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBuffer(nil)
	var writer io.Writer = buffer

	err = h5Tpl.Execute(writer, content)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
