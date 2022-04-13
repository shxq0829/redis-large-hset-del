# Large HSET keys delete
> If you afraid when large hash files has been deleted on Redis. You have to delete per single hash keys step by step.

> For example, you have a registered hash list of 30 million. If you delete it with the “del” command, you will probably have a huge load on your system. Or you may experience a crash.

> By scanning with the “HSCAN” command, I think it is the healthiest method to do this singular operation. 

## Usage
```shell
$ go build
```
```shell
$ ./import.sh HSET_KEY COUNT //import test data to localhost:6379
```
### Run
> **Warning** cluster mode is DEFAULT. Or -clusterMode=false

> **Notice** If too little data is imported,HSCAN COUNT is invalid.It will display all data.

> Deletes up to "-batchSize" every 3 seconds. Default 1000.

#### Examples
```shell
$ ./import test 1000
$ ./hset_del -key=test
$ ./hset_del -key=test -batchSize=2000
$ ./hset_del -key=test -clusterMode=false
$ ./hset_del -key=test -addr=127.0.0.1:6380
$ ./hset_del -key=test -addr=127.0.0.1:6380 -password=123 
```

## Blog Links
- [Redis HSCAN command](https://redis.io/commands/hscan)
- [Lazy Redis is better Redis (Antirez)](http://www.antirez.com/news/93)
- [Deleting Large Hashes in Redis](https://redisgreen.net/blog/deleting-large-hashes/)


