const Redis = require('ioredis');
var redis = new Redis({
  sentinels: [{ host: 'localhost', port: 26379 }],
  name: 'master-group-name'
});
redis.set('foo', 'bar').then( _ => redis.get('foo').then(value => console.log(value)));
