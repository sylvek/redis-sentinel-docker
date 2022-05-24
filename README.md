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

## Example of client configuration :

```
var redis = new Redis({
  sentinels: [{ host: 'localhost', port: 26379 }],
  name: 'master-group-name'
});
redis.set('foo', 'bar');
```

A full example is available on `example`.

```
$> cd example
$> npm install
$> npm start run

// should display
> example@1.0.0 start
> node test.js "run"

bar
```

You could replace `master-group-name` by your own "master group name".
You should rename files :
- "sentinel-get-master-addr-by-name-master-group-name"
- "sentinel-sentinels-master-group-name"

by you own group name. :)