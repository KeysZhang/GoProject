# 实验三：简单的web服务程序——登陆注册

</br>

##一. 框架

####本次实验使用了beego框架，beego 是一个快速开发 Go 应用的 HTTP 框架，他可以用来快速开发 API、Web 及后端服务等各种应用，是一个 RESTful 的框架，主要设计灵感来源于 tornado、sinatra 和 flask 这三个框架，但是结合了 Go 本身的一些特性（interface、struct 嵌入等）而设计的一个框架。这个框架符合mvc结构，并且封装很多实用的工具包，因此我选了这个框架

</br>

## 二. curl测试

#### 1. 登陆界面--get()

####命令：curl -v http://localhost:8080

```
详细信息: GET http://localhost:8080/ with 0-byte payload
详细信息: received 1143-byte response of content type text/html; charset=utf-8


StatusCode        : 200
StatusDescription : OK
Content           : ﻿<!doctype html>
                    <html lang="zh">
                    <head>
                    <meta charset="UTF-8">
                    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
                    <meta name="viewport" content="width=device-width, initial-scale=1....
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 1143
                    Content-Type: text/html; charset=utf-8
                    Date: Sat, 04 Nov 2017 03:10:14 GMT
                    Server: beegoServer:1.9.0

                    ﻿<!doctype html>
                    <html lang="zh">
                    <head>
                    <meta chars...
Forms             : {}
Headers           : {[Content-Length, 1143], [Content-Type, text/html; charset=utf-8], [Date, Sat, 04 Nov 2017 03:10:14 GMT], [Server, beegoServer:1.9.0
                    ]}
Images            : {}
InputFields       : {@{innerHTML=; innerText=; outerHTML=<input name="username" type="text" placeholder="Username">; outerText=; tagName=INPUT; name=use
                    rname; type=text; placeholder=Username}, @{innerHTML=; innerText=; outerHTML=<input name="password" type="password" placeholder="Pas
                    sword">; outerText=; tagName=INPUT; name=password; type=password; placeholder=Password}}
Links             : {@{innerHTML=Register; innerText=Register; outerHTML=<a href="/regist">Register</a>; outerText=Register; tagName=A; href=/regist}}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 1143
```

####2.注册界面--get() 

####命令：curl -v http://localhost:8080/regist

```
详细信息: GET http://localhost:8080/regist with 0-byte payload
详细信息: received 1136-byte response of content type text/html; charset=utf-8


StatusCode        : 200
StatusDescription : OK
Content           : <!doctype html>
                    <html lang="zh">
                    <head>
                    <meta charset="UTF-8">
                    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
                    <meta name="viewport" content="width=device-width, initial-scale=1.0...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 1136
                    Content-Type: text/html; charset=utf-8
                    Date: Sat, 04 Nov 2017 03:26:00 GMT
                    Server: beegoServer:1.9.0

                    <!doctype html>
                    <html lang="zh">
                    <head>
                    <meta charse...
Forms             : {}
Headers           : {[Content-Length, 1136], [Content-Type, text/html; charset=utf-8], [Date, Sat, 04 Nov 2017 03:26:00 GMT], [Server, beegoServer:1.9.0
                    ]}
Images            : {}
InputFields       : {@{innerHTML=; innerText=; outerHTML=<input name="username" type="text" placeholder="Username">; outerText=; tagName=INPUT; name=use
                    rname; type=text; placeholder=Username}, @{innerHTML=; innerText=; outerHTML=<input name="password" type="password" placeholder="Pas
                    sword">; outerText=; tagName=INPUT; name=password; type=password; placeholder=Password}}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 1136
```

</br>

## 三. ab测试

####1. 登陆界面--get()

####命令：ab -n 1000 -c 100 http://localhost:8080/

```
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        beegoServer:1.9.0
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        1143 bytes

Concurrency Level:      100
Time taken for tests:   1.214 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      1288000 bytes
HTML transferred:       1143000 bytes
Requests per second:    823.83 [#/sec] (mean)
Time per request:       121.385 [ms] (mean)
Time per request:       1.214 [ms] (mean, across all concurrent requests)
Transfer rate:          1036.22 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.6      1       3
Processing:     7  114  43.6    119     314
Waiting:        4   87  50.5     81     314
Total:          9  115  43.6    121     315

Percentage of the requests served within a certain time (ms)
  50%    121
  66%    127
  75%    133
  80%    137
  90%    160
  95%    191
  98%    241
  99%    261
 100%    315 (longest request)
```

####2.注册界面--get()

#### 命令：ab -n 1000 -c 100 http://localhost:8080/regist

```
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        beegoServer:1.9.0
Server Hostname:        localhost
Server Port:            8080

Document Path:          /regist
Document Length:        1136 bytes

Concurrency Level:      100
Time taken for tests:   1.106 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      1281000 bytes
HTML transferred:       1136000 bytes
Requests per second:    904.34 [#/sec] (mean)
Time per request:       110.578 [ms] (mean)
Time per request:       1.106 [ms] (mean, across all concurrent requests)
Transfer rate:          1131.30 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.5      1       5
Processing:     3  104  27.3    101     152
Waiting:        2   55  33.7     53     143
Total:          4  105  27.4    102     153

Percentage of the requests served within a certain time (ms)
  50%    102
  66%    114
  75%    125
  80%    130
  90%    141
  95%    147
  98%    149
  99%    151
 100%    153 (longest request)
```

#### 3.参数解释

##### ab  -n 1000 -c100 url

#####-n：在测试会话中所执行的请求个数

##### -c：一次产生的请求个数（并发数）

##### Document Path：测试的页面

##### Document Length：页面大小

##### Concurrency Level：测试的并发数

#####Time taken for tests：整个测试持续的时间

##### Complete requests：完成的请求数量

##### Failed requests：失败的请求数量

#####Total transferred:：整个过程中的网络传输量

##### HTML transferred：整个过程中的HTML内容传输量

##### Requests per second：相当于LR中的每秒事务数，这个参数最重要之一

##### Time per request：相当于LR中的平均事务响应时间，最重要的指标之二

