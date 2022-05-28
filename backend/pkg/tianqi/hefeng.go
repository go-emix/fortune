package tianqi

import (
	"errors"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

const key = "5152bf44b4074ab99e59f5f6c5748b53"

const api = "https://devapi.qweather.com/v7/weather/now?"

func GetTemp() (result interface{}, err error) {
	lid := "101011100"
	var get *http.Response
	get, err = http.Get(api + "key=" + key + "&location=" + lid)
	if err != nil {
		return
	}
	defer get.Body.Close()
	all, err := io.ReadAll(get.Body)
	if err != nil {
		return
	}
	code := gjson.GetBytes(all, "code").String()
	if code != "200" {
		err = errors.New("hefeng api return error")
		return
	}
	result = gjson.GetBytes(all, "now.temp").String()
	return
}
