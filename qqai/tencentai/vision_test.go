// vision_test.go
// http://www.jeepxie.net/article/772180.html
package tencentai

import (
	"fmt"
	"os"
	"testing"
)

func TestVisionImgFileToText(t *testing.T) {

	client := NewAiClient(os.Getenv("QQ_AI_APPID"), os.Getenv("QQ_AI_APPKEY"))

	err, vtt := client.VisionImgFileToText("./fm_big.jpg")
	if err != nil {
		t.Fatalf("Expect success but returned: %v", err)
	} else {
		if vtt.Ret != 0 {
			t.Fatalf("Expect ret is 0 but%v msg:%s:", vtt.Ret, vtt.Msg)
		} else {
			fmt.Println("text:", vtt.Data.Text)
		}
	}
}

/*

export QQ_AI_APPID=... QQ_AI_APPKEY=... TZ='Asia/Shanghai'

$ go test
text: 海面上一艘白色的船.
PASS
ok  	tencentai	2.554s

*/
