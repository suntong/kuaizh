//vision.go
package tencentai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	VISION_IMGTOTEXT_URI = "https://api.ai.qq.com/fcgi-bin/vision/vision_imgtotext"
)

func (this *AiClient) VisionImgFileToText(img string) (error, VisionImgToTextResp) {
	var vtt VisionImgToTextResp

	file, err := os.Open(img)
	if err != nil {
		return err, vtt
	}

	defer file.Close()
	return this.VisionImgReaderToText(file)

}

type VisionImgToTextRespData struct {
	Text string `json:"text, omitempty"`
}

type VisionImgToTextResp struct {
	Ret  int                     `json:"ret, omitempty"`
	Msg  string                  `json:"msg, omitempty"`
	Data VisionImgToTextRespData `json:"data, omitempty"`
}

func (this *AiClient) VisionImgReaderToText(img io.Reader) (error, VisionImgToTextResp) {
	var vtt VisionImgToTextResp
	buffer := &bytes.Buffer{}
	_, err := io.Copy(buffer, img)
	if err != nil {
		return err, vtt
	}

	data := make(map[string]string)
	data["app_id"] = this.Appid
	data["time_stamp"] = fmt.Sprintf("%d", time.Now().Unix())
	data["nonce_str"] = GetRandomString(30)
	data["session_id"] = "100"
	data["image"] = string(base64.StdEncoding.EncodeToString(buffer.Bytes()))

	reqbody := this.GetBody(data)
	//fmt.Println("body is:", reqbody, "image len:", len(data["image"]), " time:", data["time_stamp"], ",now:", time.Now())
	resp, err := this.Post(VISION_IMGTOTEXT_URI, reqbody)
	if err != nil {
		return err, vtt
	}

	json.NewDecoder(resp.Body).Decode(&vtt)
	return nil, vtt
}
