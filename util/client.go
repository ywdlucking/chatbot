package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

const url = "http://bi.shie.com.cn/superset/sql_json/"

var INIT = 100000
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetGroupOwner(certNo string) string {
	ok := true
	no := false
	sql := "SELECT org_name FROM t_group_order  WHERE  grp_order_no =(SELECT  grp_order_no from t_group_order_detail WHERE  certi_code ='" + certNo + "' and check_status  = '1'  and  order_status='0')"
	data := make(map[string]interface{})
	INIT = INIT + 1
	data["client_id"] = randStr(10)
	data["ctas_method"] = "TABLE"
	data["database_id"] = 18
	data["expand_data"] = ok
	data["json"] = ok
	data["queryLimit"] = 100
	data["runAsync"] = no
	data["schema"] = "hzihc_gw_prod2023"
	data["select_as_cta"] = no
	data["sql"] = sql
	data["sql_editor_id"] = "u2p1H6ibM"
	data["tab"] = "1"
	data["tmp_table_name"] = ""
	byte, _ := json.Marshal(data)
	println(string(byte))
	req, err := http.NewRequest("POST", url, strings.NewReader(string(byte)))
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "session=.eJwljzFuAzEMBP-i2oUoUhTpzxwokoINH2LgLmkS5O85IOUUM9j9Kds68nyU-7L9zFvZnlHuhQJNAh0nUkuNHthBeBDH0IpjLgAZDaoRVGwaEhDItoha40puvrDPWgVkakMFxVBeA9rEamgADNZ19lgxmCuTZfPuSldetNyKn8faPt-v_Lj2eANuRF4zYYiwGWiX2aeLZoQLxZg6_PL2t9uel_P9uOjrzOP_Ekj5_QOCRkHO.Y75zBQ.MY8NYTUOt21kriIQZ3yp2KD6pBk")
	req.Header.Set("Host", "bi.shie.com.cn")
	req.Header.Set("Origin", "http://bi.shie.com.cn")
	req.Header.Set("Referer", "http://bi.shie.com.cn/superset/sqllab/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("X-CSRFToken", "ImMyMTYyNDRjMGVlMTc4ODZhYTE5NThiNWJjODllZGRjODRkN2I5N2Mi.Y74qbw.IS5PS3Vnw-Ik2d7fKkQjUwH5fSY")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("请求错误：", err)
		return "查询好像失败了，请稍后询问"
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("数据错误：", err)
		return "查询好像失败了，请稍后询问"
	}
	var biResult BiResult
	err = json.Unmarshal(respByte, &biResult)
	if err != nil {
		fmt.Println("数据错误：", err)
		return "查询好像失败了，请稍后询问"
	}
	d, _ := json.Marshal(biResult.Data)
	println(biResult.Data)
	return string(d)
}

type BiResult struct {
	Data   []interface{} `json:"data"`
	Status string        `json:"status"`
}
