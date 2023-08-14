package cache

import "time"

//说明#
//『友情链接』这个数据不会经常修改，没有必要每次请求都读取数据库。为了方便『友情链接列表』接口添加缓存，这节课我们来开发 cache 包。
//
//1. 缓存机制设计
//Store interface
//
//目前系统中使用到 Redis，cache 包不应该只依赖于 Redis ，后续如使用其他缓存方案，如 Memcached ，可以很方便切换。所以这里我们利用 Go 的 interface 功能，将数据存储抽象化。
//
//RedisStore
//
//RedisStore 是 cahce 包的 Store interface Redis 实现。目前我们的 Redis 里存放多种业务逻辑的数据，如数字验证码、短信验证码等，缓存的信息应该和这些业务数据使用不同的数据库。

type Store interface {
	Set(key string, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()

	IsAlive() error

	//	Increment 当参数只有1个时，为key，增加1
	//	当参数有2个时，第一个参数为key，第二个参数为要增加的值 int64 类型

	Increment(paramaters ...interface{})

	//	Decrement 当参数只有1个时，为key 减去1
	//当参数有2两个时，第一个为key，第二个为要减去的值int64类型

	Decrement(parameters ...interface{})
}
