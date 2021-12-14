package utils

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

const(
    key = "e380724b77e60eab4a2dc380edd36054"
    url = "https://restapi.amap.com/v3/geocode/regeo?output=json&radius=10&extensions=all&key=" + key
)

type RegeoResponse struct {
    Status      string
    RegeoCode   *RegeoCode
    Info        string
}

type RegeoCode struct {
    FormattedAddress    string  `json:"formatted_address"`
}

/*
通过经纬度获取定位
*/
func Regeo(latitude, longitude float64) (string, error) {
    requestUrl := url + "&location=" + fmt.Sprint(longitude, ",", latitude)

    resp, err := http.Get(requestUrl)
    if err != nil {
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var regeoResp RegeoResponse
    err = json.Unmarshal(body, &regeoResp)
    if err != nil {
        return "", err
    }
    if regeoResp.Status == "0" {
        return "", errors.New(regeoResp.Info)
    }

    return regeoResp.RegeoCode.FormattedAddress, nil
}
