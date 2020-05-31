package test

import (
	"encoding/base64"
	"fmt"
	"github.com/songouku/gwx/util"
	"testing"
)

func TestAes(t *testing.T) {
	encodingKey := "qh6iHD9Ct3b6sL1ixrQcmEaSzEA4eFrLZNt1V9wOK8Z"
	aesKey, err := base64.StdEncoding.DecodeString(encodingKey + "=")
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	content := "IktqWSySM5R1D99mty2Mju07XUf7syWSPK+Ev5mLyqRVoTlWIQPF1PHWk0ZKrNlspiWdaaMf9s3tsAXJ0eDxwVuqcRbmirOE9QdAC3uetJ3VBT676pru7IqC6l+NH9DH+P14/CQJCvUawyTBasDF4h/zcoTOS9XFfCXYaqVGy0kpcSODL/vNSQXo70AtswYB6nW/LMpXg7yLqDxoB46rdDBuQaBxZRMyP2Sz7h3LfXam3KnHhqBX1g0snWt439bi9A1mgxycXYk5U70g6RkVCCz9n1NL9YsSy0owq8s09hli/iZQC1dFDkMCSJqp4b3qMtCZRt/l3QQiRa70Aub4ra4gf9UR3T8yGPuYYt7W0wL5amjitzV3hiYapXYr8CBNJD0aH9R14opWDgsXF8jKA9CvquR5AoD0QxqsNDMyQn/c91gGOihoy4GcSv30DbYlijGVbp/EU8vvjg1RuLOSnPekqBVu68YSj5g5mzNNNsRuCyUx+Cg1lSBfd+oiValu"
	res, err := util.AesDecrypt([]byte(content), aesKey)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("aes result is %s\n", res)
}
