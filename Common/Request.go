package Common

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ResponseParameterPOST struct {
	body       string
	statuscode int
}

// 异常捕获
func Error(Url string) {
	err := recover()
	if err != nil {
		log.Printf("[-] %v 扫描异常!", Url)
	}
}

// 基本get请求
func RequestGET(Url string) int {
	defer Error(Url)

	client := &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest("GET", Url, nil)
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3861.400 QQBrowser/10.7.4313.400"},
		"Content-Type": {"text/xml"},
	}
	response, err := client.Do(request)
	if err != nil {
		return 0
	}
	return response.StatusCode
}

func CustomizePost(target string, Url string, body string, headers map[string]string) map[string]string {
	//uri, _ := url.Parse("htttp://127.0.0.1:8080")
	defer Error(target)
	client := &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		//Transport: &http.Transport{
		//	// 设置代理
		//	Proxy: http.ProxyURL(uri),
		//},
	}
	//fmt.Println(target + url)
	request, _ := http.NewRequest("POST", target+Url, strings.NewReader(body))
	request.Header = http.Header{
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0"},
		"Content-Type": {"text/xml"},
		//"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
		//"Accept-Language":           {"zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"},
		//"Accept-Encoding":           {"gzip, deflate"},
		//"Upgrade-Insecure-Requests": {"1"},
	}
	if headers != nil {
		for header, text := range headers {
			request.Header.Set(header, text)
		}
	}
	//fmt.Println(request.Header)
	response, err := client.Do(request)
	//fmt.Println(response.StatusCode)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Printf(string(content))
	result := make(map[string]string)
	result["body"] = string(content)
	result["statuscode"] = strconv.Itoa(response.StatusCode)
	return result
}

//func Update(target string, url string, body string, headers map[string]string) {
//	defer Error(target)
//	client := &http.Client{
//		Timeout: time.Second * 5,
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
//	request, _ := http.NewRequest("POST", target+url, strings.NewReader(body))
//	request.Header = http.Header{
//		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:108.0) Gecko/20100101 Firefox/108.0"},
//		"Content-Type": {"multipart/form-data; boundary=----WebKitFormBoundary9UO3YXctYrZi8iQ"},
//		//"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
//		//"Accept-Language":           {"zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"},
//		//"Accept-Encoding":           {"gzip, deflate"},
//		//"Upgrade-Insecure-Requests": {"1"},
//	}
//
//}
