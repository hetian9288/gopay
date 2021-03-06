package alipay

import (
	"testing"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
	"github.com/iGoogle-ink/gopay/pkg/xlog"
)

func TestClient_ZhimaCreditScoreGet(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("transaction_id", util.GetRandomString(48))
	bm.Set("product_code", "w1010100100000000001")

	// 芝麻分
	aliRsp, err := client.ZhimaCreditScoreGet(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditScoreGet(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
