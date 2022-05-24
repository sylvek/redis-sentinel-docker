# Redis+Sentinel+Docker=:heart:

[Tired of trying to make Redis, Sentinel and Docker work together?](https://redis.io/topics/sentinel#sentinel-docker-nat-and-possible-issues)

First, I simply linked a containerized Redis with a sentinel.
Sadely, Sentinel broadcasts the Redis internal IP >:-[

So, I tried to use the `host mode` of Docker, but it doesn't work on MacOS.

Since my need is to develop with an single Redis, i wrote a tiny sentinel server that responds naively.

Battle tested with [ioredis](https://ioredis.readthedocs.io/en/stable/README/).

## How to use ?

```
> git clone https://github.com/sylvek/redis-sentinel-docker.git
> cd redis-sentinel-docker
> docker-compose up
```

## With redis-cli

```
▶ redis-cli -h localhost -p 26379
localhost:26379> ping
PONG
localhost:26379> sentinel get-master-addr-by-name redis-cluster-master
1) "127.0.0.1"
2) "6379"
```

## Example of client configuration :

```
var redis = new Redis({
  sentinels: [{ host: 'localhost', port: 26379 }],
  name: 'redis-cluster-master'
});
redis.set('foo', 'bar');
```
