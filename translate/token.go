package translate

import (
	tk "github.com/kyai/google-translate-tk"
	"github/wbellmelodyw/gin-wechat/cache"
	"log"
)

const TkkKey = "ttk:key"

func GetToken(text string) string {
	tkk, ok := cache.GetToken(TkkKey)
	if !ok {
		//取不出来就请求google
		log.Println("set")
		tkk, _ = tk.GetTKK()
		err := cache.SetToken(TkkKey, tkk)
		log.Print(err)
	}
	tk := tk.GetTK(text, tkk)
	return tk
}
