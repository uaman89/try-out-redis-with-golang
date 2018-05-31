# Try-out-redis-with-golang

This is an example interaction of Redis wi Golang.

Here has been chosen three clients fo Go:
* [go-redis/redis](https://github.com/go-redis/redis)
* [radix](https://github.com/mediocregopher/radix.v2)
* [redigo](https://github.com/gomodule/redigo)

Each of them are supporting novadys and has documentaion with examples at https://godoc.org.
They works well with base operation like SET / GET / DEL.


Other ones redis-client for Go hasn't been maintainer for years, so they has been omitted.

## go-redis/redis
[docs](https://godoc.org/github.com/go-redis/redis)

Has the highest rating on github ang biggest number of [features](https://github.com/go-redis/redis#redis-client-for-golang).
On my opinion it was easier to understand and start work with.

## radix
[docs](https://godoc.org/ghttps://godoc.org/ithub.com/mediocregopher/radix.v2)

Pretty similar to go-redis/redis.
Has lower community.
MIT license.

## redigo
([docs]https://godoc.org/github.com/gomodule/redigo/redis)

Has smallest feature list comparing to previous ones.
Implements A Print-like API with support for all Redis commands.
Has no supporting []redis sentinel](https://redis.io/topics/sentinel)

## Examples
Simple examples of usage:

[go-redis](./go-redis)

[radix](./radix)

[redigo](./redigo)
