# shortrul

[zero教程](https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl.md)

- db

```
create database gozero DEFAULT CHARACTER SET utf8mb4 ;;
use gozero;
CREATE TABLE `shorturl`
(
  `shorten` varchar(255) NOT NULL COMMENT 'shorten key',
  `url` varchar(255) NOT NULL COMMENT 'original url',
  PRIMARY KEY(`shorten`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```


- shorten api 调用

```
curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 29 Aug 2020 10:49:49 GMT
Content-Length: 21

{"shorten":"f35b2a"}
```



- expand api 调用

```
curl -i "http://localhost:8888/expand?shorten=f35b2a"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 29 Aug 2020 10:51:53 GMT
Content-Length: 34

{"url":"http://www.xiaoheiban.cn"}
```



- etcd

```
> etcdctl --endpoints=$ENDPOINTS  get --from-key ""
transform.rpc/4741857488835633933
127.0.0.1:8080
```