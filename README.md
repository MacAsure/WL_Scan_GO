# Weblogic_Scan_GO

----
weblogic漏洞检测工具!!!


# 功能特点
***
**支持漏洞检测**
***
POC种类：
```
check_cve_2014_4210
check_cve_2016_0638
check_cve_2016_3510
check_cve_2017_3248
check_cve_2017_3506
check_cve_2017_10271
check_cve_2018_2628
check_cve_2018_2893
check_cve_2018_2894
check_cve_2018_3191
check_cve_2018_3252
check_cve_2019_2618
check_cve_2019_2725_2729
check_cve_2019_2890
check_cve_2020_2555
check_cve_2020_2883
check_cve_2020_14882
```

# 使用方法

## 环境准备
***
```
git clone https://github.com/iceberg-N/thinkphp5.x_Scan.git
```

## 使用方法
***
-h 提供命令帮助文档
![ThinkPHP5.x_Scan-1.png](./images/ThinkPHP5.x_Scan-1.png)

**单个目标**   
只有域名，默认为http
```
python3 ThinkPHP5_X_Scan.py -u http://example.com
```

![ThinkPHP5.x_Scan-2.png](./images/ThinkPHP5.x_Scan-2.png)

**漏洞利用**

```
python3 ThinkPHP5_X_Scan.py -u http://example.com -c whoami
```

![ThinkPHP5.x_Scan-3.png](./images/ThinkPHP5.x_Scan-3.png)

![./images/ThinkPHP5.x_Scan-4.png](./images/ThinkPHP5.x_Scan-4.png)

**上传webshell**   
支持一句话木马，webshell名为"iceberg.php"，密码为"iceberg"
```
python3 ThinkPHP5_X_Scan.py -w http://example.com
```

![ThinkPHP5.x_Scan-5.png](./images/ThinkPHP5.x_Scan-5.png)

![ThinkPHP5.x_Scan-6.png](./images/ThinkPHP5.x_Scan-6.png)

**批量检测**   
支持txt格式
```
python3 ThinkPHP5_X_Scan.py -f ./ip.txt -o ./result.txt
```

![ThinkPHP5.x_Scan-1.png](./images/ThinkPHP5.x_Scan-7.png)

![ThinkPHP5.x_Scan-8.png](./images/ThinkPHP5.x_Scan-8.png)


# 更多
***
thinkphp5.x的各种版本的漏洞原理及代码分析，请搜索微信公众号"MrHatSec"。   
更多红队技巧，请搜索公众号"信安泥石流"。
