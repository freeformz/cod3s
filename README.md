[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

# cod3s

Reflect the requested HTTP code back to you.

```bash
$ http http://cod3s.herokuapp.com/201
HTTP/1.1 201 Created
Connection: keep-alive
Content-Length: 0
Content-Type: text/plain; charset=utf-8
Date: Sat, 28 Mar 2015 01:12:19 GMT
Server: Cowboy
Via: 1.1 vegur

$ http http://cod3s.herokuapp.com/505
HTTP/1.1 505 HTTP Version Not Supported
Connection: keep-alive
Content-Length: 0
Content-Type: text/plain; charset=utf-8
Date: Sat, 28 Mar 2015 01:14:45 GMT
Server: Cowboy
Via: 1.1 vegur
```
