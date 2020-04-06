package translate

import (
	tk "github.com/kyai/google-translate-tk"
	"github/wbellmelodyw/gin-wechat/cache"
	"github/wbellmelodyw/gin-wechat/logger"
)

const TkkKey = "ttk:key"

func GetToken(text string) string {
	tkk, ok := cache.GetToken(TkkKey)
	if !ok {
		//取不出来就请求google
		logger.Module("get-token").Sugar().Info("set", TkkKey)
		tkk, _ = tk.GetTKK()
		if err := cache.SetToken(TkkKey, tkk); err != nil {
			logger.Module("get-token").Sugar().Error("set err", TkkKey)
		}
	}
	logger.Module("get-token").Sugar().Info("get", TkkKey)
	tk := tk.GetTK(text, tkk)
	return tk
}
