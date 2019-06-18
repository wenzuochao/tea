package tea

import (
	"testing"

	"github.com/alibabacloud-go/tea/utils"
)

func TestConvert(t *testing.T) {
	in := map[string]interface{}{
		"key": "value",
	}
	out := &struct {
		Key string `json:"key"`
	}{}
	err := Convert(in, out)
	utils.AssertNil(t, err)
	utils.AssertEqual(t, "value", out.Key)
}

func TestConvertNonPtr(t *testing.T) {
	in := map[string]interface{}{
		"key": "value",
	}
	out := struct {
		Key string `json:"key"`
	}{}
	err := Convert(in, out)
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "The out parameter must be pointer", err.Error())
}

func TestConvertType(t *testing.T) {
	in := map[string]interface{}{
		"key": "value",
	}
	out := &struct {
		Key int `json:"key"`
	}{}
	err := Convert(in, out)
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "Convert type fails for field: key, expect type: int, current type: string", err.Error())
}

func TestSDKError(t *testing.T) {
	err := NewSDKError(map[string]interface{}{
		"code":    "code",
		"message": "message",
		"data": map[string]interface{}{
			"httpCode":  "404",
			"requestId": "dfadfa32cgfdcasd4313",
			"hostId":    "github.com/alibabacloud/tea",
		},
	})
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "SDKError: {\"hostId\":\"github.com/alibabacloud/tea\",\"httpCode\":\"404\",\"requestId\":\"dfadfa32cgfdcasd4313\"} message ", err.Error())
}

func TestSDKErrorCode404(t *testing.T) {
	err := NewSDKError(map[string]interface{}{
		"code":    404,
		"message": "message",
	})
	utils.AssertNotNil(t, err)
	utils.AssertEqual(t, "SDKError: 404 message ", err.Error())
}