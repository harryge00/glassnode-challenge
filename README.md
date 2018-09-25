# glassnode-challenge

## How to run
```
make
```
It will build 3 images and run 3 containers.


## Performance
* Protobuf can be used in inter-communication for better performance. JSON is used because of its simplicity and compatibility with `Coinmarket` API.
* Ranking can be cached with `Time to live (TTL)` because in general, it does not change as frequently as the price. Cache can be achieved using Redis/Etcd, or in process memory.

## TODO
Some coins in `cryptocompare` cannot find corresponding price in `coinmarketcap`. Such as `CRAIG, XBS, GIVE, CHASH, DANK...` The price of such coins are omitted.  