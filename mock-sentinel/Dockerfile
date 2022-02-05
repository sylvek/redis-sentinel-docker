FROM ruby:alpine

COPY sentinel.rb /sentinel.rb
COPY command /command
COPY ping /ping
COPY sentinel-get-master-addr-by-name-redis-cluster-master /sentinel-get-master-addr-by-name-redis-cluster-master
COPY sentinel-sentinels-redis-cluster-master /sentinel-sentinels-redis-cluster-master

EXPOSE 26379
CMD ["ruby", "/sentinel.rb"]