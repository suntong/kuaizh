// vision_test.go
// http://www.jeepxie.net/article/772180.html
package tencentai

import (
	"fmt"
	"testing"
)

func TestVisionImgFileToText(t *testing.T) {

	client := NewAiClient("your-appid", "your-appkey")

	err, vtt := client.VisionImgFileToText("./fm_big.jpg")
	if err != nil {
		t.Fatalf("Expect success but return:", err)
	} else {
		if vtt.Ret != 0 {
			t.Fatalf("Expect ret is 0 but%v msg:%s", vtt.Ret, vtt.Msg)
		} else {
			fmt.Println("text:", vtt.Data.Text)
		}
	}
}
