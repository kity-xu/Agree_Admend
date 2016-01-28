package crossframe

import (
	"cn.agree/utils"
	"net/http"
)

func CrossFrameProxyServer() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("Htmls"))))
	err := http.ListenAndServe(":8880", nil)
	if err != nil {
		utils.Error("fail to start Proxy,error is %s", err.Error())
		return
	}
}
