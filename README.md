# 参考サイト
https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a
https://qiita.com/shiei_kawa/items/eddf48287455380f618f

## GET,POST,PUT,PATCH,DELETE,OPTIONS
http://sekitaka-1214.hatenablog.com/entry/2016/08/11/153816

## GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】
https://qiita.com/gold-kou/items/45a95d61d253184b0f33

## Go言語のGormを実践投入する時に最低限知っておくべきことのまとめ【ORM】
https://qiita.com/ttiger55/items/3606b8dd570637c12387



## Endpoint(API)
### GET
http://127.0.0.1:8080/fetchAllProducts


.envrc
```
# Go Env
export GO111MODULE=on

# MySQL
export DB_NAME=database
export DB_USER=mysql
export DB_PASSWORD=mysql
export TEST_DB_NAME=database_test
export TEST_DB_USER=root
export TEST_DB_PASSWORD=root
```


## ローカル MySQL Volume 削除
```
$ docker volume ls

$ docker volume rm go-gin-mysql-todo-api_go-gin-mysql-todo-api-data
```

### MySQL接続確認
```
# General User
$ mysql -h 0.0.0.0 -umysql -p'mysql'

# Root User
$ mysql -h 0.0.0.0 -uroot -p'mysql'
```




```
$ go mod tidy
```
