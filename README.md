# Redis + Sentinel + Docker = :heart:

[Tired of trying to make Redis, Sentinel and Docker work together?](https://redis.io/docs/manual/sentinel/#sentinel-docker-nat-and-possible-issues)

As Developer, it's frustrating to deal with Redis configuration.
Locally you should have a single Redis instance and in production we have to handle a connection with Sentinel.
Unfortunately, most of Redis clients do not offer a single way to configure Redis and Sentinel.

And, if you work on MacOS (and M1) you hit the jackpot!

In fact, [Sentinel does only one thing](https://redis.io/docs/manual/sentinel/), it ensures to always maintain integrity between clients and a Redis Master/Replicas. The client connects to the sentinel and asks two things :
- [the current current master address](https://redis.io/docs/manual/sentinel/#obtaining-the-address-of-the-current-master)
- the current sentinel configuration
  
If you answer correctly to these two questions, your client will switch and connect to the correct Redis instance.
  
So the main objective was to mock this handshake and return "localhost:6379" as the current Redis Master.

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
