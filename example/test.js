const Redis = require('ioredis');
var redis = new Redis({
  sentinels: [{ host: 'localhost', port: 26379 }],
  name: 'redis-cluster-master'
});
redis.set('foo', 'bar');