package Common

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func AnalyzeUrl(Url string) string {
	if strings.Contains(Url, "://") {
		return Url
	} else {
		seult := RequestGET("http://" + Url)
		if seult != 200 {
			return "https://" + Url
		} else {
			return "http://" + Url
		}

	}
}

// 解析域名
func Lookup(target string) string {
	u, err := url.Parse(target)
	if err != nil {
		return ""
	}
	h := strings.Split(u.Host, ":")
	if net.ParseIP(h[0]) == nil {
		ns, err := net.LookupHost(h[0])
		if err != nil {
			fmt.Printf("[-] %v 域名解析失败", target)
			return target
		}
		if len(h) == 2 {
			return u.Scheme + "://" + ns[0] + ":" + h[1]
		} else {
			return u.Scheme + "://" + ns[0]
		}
	} else {
		return target
	}

}

// 判断是否添加端口
func Findport(target string) string {
	u, err := url.Parse(target)
	if err != nil {
		return ""
	}
	if !strings.Contains(u.Host, ":") {
		return u.Scheme + "://" + u.Host + ":7001"
	} else {
		return u.Scheme + "://" + u.Host
	}
}

// 初始化url
func Initurl(target string) string {
	target_http := AnalyzeUrl(target)
	target_ip := Lookup(target_http)
	target_port := Findport(target_ip)
	return target_port
}

func Init(target string) string {
	target_http := AnalyzeUrl(target)
	target_port := Findport(target_http)
	return target_port
}
