
1. mysqlの準備

``` sh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=mysql -d -p 3306:3306 mysql
docker exec -it mysql /bin/sh

# docker内で下実施
mysql -uroot -p
create database blog;

```

