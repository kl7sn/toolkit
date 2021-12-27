package main

/*
http代理测试


* [Mocking a HTTP access with http.Transport in Golang - oinume journal](http://oinume.hatenablog.com/entry/mocking-http-access-in-golang)
* [Go http访问使用代理](http://www.cnblogs.com/damir/archive/2012/05/06/2486663.html)
* [GO HTTP client客户端使用 - 海运的博客](https://www.haiyun.me/archives/1051.html)
* [Making Tor HTTP Requests with Go | DevDungeon](http://www.devdungeon.com/content/making-tor-http-requests-go)
* [go - golang: How to do a https request with proxy - Stack Overflow](https://stackoverflow.com/questions/42662369/golang-how-to-do-a-https-request-with-proxy)
* [go - Set UserAgent in http request - Stack Overflow](https://stackoverflow.com/questions/13263492/set-useragent-in-http-request)

*/

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func First() {
	/*
		1. 普通请求
	*/

	webUrl := "http://ip.gs/"
	resp, err := http.Get(webUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	// if resp.StatusCode == http.StatusOK {
	// 	fmt.Println(resp.StatusCode)
	// }

	time.Sleep(time.Second * 3)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Second(webUrl, proxyUrl string) {
	/*
		1. 代理请求
		2. 跳过https不安全验证
	*/
	// webUrl := "http://ip.gs/"
	// proxyUrl := "http://115.215.71.12:808"

	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5, //超时时间
	}

	resp, err := client.Get(webUrl)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func Third(webUrl, proxyUrl string) {
	/*
		1. 代理请求
		2. 跳过https不安全验证
		3. 自定义请求头 User-Agent

	*/
	// webUrl := "http://ip.gs/"
	// proxyUrl := "http://171.215.227.125:9000"

	request, _ := http.NewRequest("GET", webUrl, nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	request.AddCookie(&http.Cookie{
		Name:  "shimo_dev_sid",
		Value: "s%3A7bdbc5987bd247d7a243b36f13a59b8e.9EGW4YHzt1SSg9GnRVdpsDfbsIE6yXfFx8BTsg4l32I",
	})
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// webUrl := "https://httpbin.org/ws/?transport=polling&sid=653501010916544512"
	// proxyUrl := "https://172.16.21.219:443"

	// Third(webUrl, proxyUrl)
	// Third(webUrl, proxyUrl)

	get()

}

func get() {

	url := "https://ws.shimodev.com/ws/?transport=polling&sid=654964889832169472"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	client.Timeout = time.Second * 10
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Cookie", "shimo_dev_sid=s%3A7bdbc5987bd247d7a243b36f13a59b8e.9EGW4YHzt1SSg9GnRVdpsDfbsIE6yXfFx8BTsg4l32I")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}