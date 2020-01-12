# go-echo-api
Simple API example built in go echo

## Init

### Docker Up

```
docker-compose up -d
```

### Create DB table

```
$ docker-compose exec db bash

mysql -u root -p

CREATE TABLE `users` ( `id` int NOT NULL AUTO_INCREMENT, `name` varchar(100) NOT NULL, PRIMARY KEY (`id`) );

```
### Start echo server

```
$ docker-compose exec app bash

realize start --run
```

### Browse to
```
http://localhost:1323
```
