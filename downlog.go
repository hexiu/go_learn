/********************************************************
 Describe :
 Author : jaxiuhe
 Email : axiu@jaxiu.cn
 Createtime : 2018/10/29 7:00 PM
 CopyRight : Copyright 2018 Tencent Inc.
 Permit : GPL
********************************************************/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
	"time"
	"strconv"
)

var date string
var param map[string]string = map[string]string{
	"date": date,
}

const USERID = "186"
const USERRTX = "caifanglei"

func ReadyParam(days int, date, platId, moduleMark, class string, iplist []string) (param map[string]string) {
	var params []map[string]interface{}
	for _, ip := range iplist {
		fmt.Println(ip)
		param := map[string]interface{}{
			"date": date,
			"days": days,
			"query": []map[string]interface{}{
				map[string]interface{}{
					"dimens": map[string]interface{}{
						"platId": platId,
					},
					"_ip_":        ip,
					"module_mark": moduleMark,
					"target":      class,
					"type":        "server",
				},
			},
		}
		params = append(params, param)
	}
	jsonP, _ := json.Marshal(params)
	paramA := map[string]string{
		"data":     string(jsonP),
		"user_id":  USERID,
		"user_rtx": USERRTX,
	}
	paramAP, _ := json.Marshal(paramA)
	fmt.Println(string(paramAP))
	return paramA
}

func ReadIP() (iplist []string) {
	absPath, _ := filepath.Abs(".")
	file := filepath.Join(absPath, "src", "learn", "iplist")
	fmt.Println("file:", file)
	f, _ := os.Open(file)
	fb := bufio.NewReader(f)
	for {
		line, _, err := fb.ReadLine()
		iplist = append(iplist, string(line))
		if err == io.EOF {
			break
		}
	}
	return
}

func main() {
	//var iplist string
	var timestep time.Duration = 1000
	if len(os.Args) > 1 {
		interge,err := strconv.Atoi(os.Args[1])
		if err!=nil {
			timestep = time.Duration(1000)
		}
		timestep = time.Duration(interge)
	}
	fmt.Println(ReadIP())
	Iplist := ReadIP()
	step := 20
	for i := 0; i < len(Iplist); {
		var iplist []string
		if i + step < len(Iplist) {
			iplist = Iplist[i : i+step]
			i += step
		} else {
			iplist = Iplist[i:len(Iplist)]
			i += step
		}
		param := ReadyParam(1, "2018-10-29", "200004", "206_4_down_platform", "404_rate", iplist)
		var Apiurl string = "http://10.231.144.92/cgi-bin/netman2.0/core/getStructedFeatureData.cgi?data=" + param["data"] + "&user_id=" + param["user_id"] + "&user_rtx=" + param["user_rtx"]
		fmt.Println(param, Apiurl)
		go GetData(Apiurl)
		time.Sleep(timestep * time.Microsecond)
	}
}

func GetData(ApiUrl string) {
	clt := http.DefaultClient
	httpReq, _ := http.NewRequest("GET", ApiUrl, nil)
	resp, err := clt.Do(httpReq)
	fmt.Println("error:", err)
	fmt.Println(ApiUrl)
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
