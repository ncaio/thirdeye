package controllers

import (
    "github.com/revel/revel"
    "bitbucket.org/tebeka/selenium"
    "io/ioutil"
    "h12.me/socks"
    "net/http"
)

type App struct {
	*revel.Controller
}
//
//
//
func (c App) Index() revel.Result {
    iptor := torexitip()
	return c.Render(iptor)
}
//
//
//
func (c App) Thirdeye(target string) revel.Result {
    click(target)
    return c.Render (target)

}
//
//
//
func click(url string) string {
	caps := selenium.Capabilities{
	"browserName": "firefox",
	}
	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()
	wd.Get(url)
    	img, _ := wd.Screenshot()
    	ioutil.WriteFile("/go/src/thirdeye/public/img/screen.png", img, 0x755)
//
// 
//
    	filename := "screen.png"
	return filename
}
//
//
//
 func torexitip() string {
     dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
     tr := &http.Transport{Dial: dialSocksProxy}
     httpClient := &http.Client{Transport: tr}
     req, _ := http.NewRequest("GET", "http://myexternalip.com/raw", nil)
     res, err := httpClient.Do(req)
     if err != nil {
     }
     defer res.Body.Close()
     contents, err := ioutil.ReadAll(res.Body)
     if err != nil {
     }
     return string(contents)
 }

