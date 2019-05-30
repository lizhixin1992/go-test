package commons

import (
	"github.com/go-redis/redis"
	"github.com/lizhixin1992/test/conf"
	"github.com/pelletier/go-toml"
	"log"
	"time"
)

var redisClient = New()

func New() *redis.Client {
	tree := conf.GlobalConf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     tree.Get("Addr").(string),
		Password: "",
		DB:       int(tree.Get("DB").(int64)),
	})
}

func SetString(key, value string, time time.Duration) {
	err := redisClient.Set(key, value, time).Err()
	if err != nil {
		log.Fatal("redis SetString is err, err : ", err)
		//panic(err)
	}
}

//只在键 key 不存在的情况下， 将键 key 的值设置为 value
//若键 key 已经存在， 则 SETNX 命令不做任何动作
func SetNXString(key, value string, time time.Duration) (flag bool) {
	flag, err := redisClient.SetNX(key, value, time).Result()
	if err != nil {
		flag = false
		log.Fatal("redis SetNXString is err, err : ", err)
		//panic(err)
	}
	return flag
}

//和SetNX相反，只在键已经存在时， 才对键进行设置操作
func SetXXString(key, value string, time time.Duration) (flag bool) {
	flag, err := redisClient.SetXX(key, value, time).Result()
	if err != nil {
		flag = false
		log.Fatal("redis SetXXString is err, err : ", err)
		//panic(err)
	}
	return flag
}

func GetString(key string) (value string) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		log.Fatal("redis GetString is err, err : ", err)
		//panic(err)
	}
	return value
}

//将键 key 的值设为 value ， 并返回键 key 在被设置之前的旧值
//如果键 key 没有旧值， 也即是说， 键 key 在被设置之前并不存在， 那么命令返回 nil,并且按照传入的key/value保存到redis中
func GetSetString(kev, value string) (oldValue string) {
	oldValue, err := redisClient.GetSet(kev, value).Result()
	if err != nil {
		log.Fatal("redis GetString is err, err : ", err)
	}
	return oldValue
}

//如果键 key 已经存在并且它的值是一个字符串， APPEND 命令将把 value 追加到键 key 现有值的末尾
//如果 key 不存在， APPEND 就简单地将键 key 的值设为 value ， 就像执行 SET key value 一样
func AppendString(key, value string) {
	_, err := redisClient.Append(key, value).Result()
	if err != nil {
		log.Fatal("redis AppendString is err, err : ", err)
	}
}

//将哈希表 hash 中域 field 的值设置为 value
func HSet(key, field string, value interface{}) {
	_, err := redisClient.HSet(key, field, value).Result()
	if err != nil {
		log.Fatal("redis HSet is err, err : ", err)
	}
}

//当且仅当域 field 尚未存在于哈希表的情况下， 将它的值设置为 value
//如果给定域已经存在于哈希表当中， 那么命令将放弃执行设置操作
//如果哈希表 hash 不存在， 那么一个新的哈希表将被创建并执行 HSETNX 命令
func HSetNX(key, field string, value interface{}) {
	_, err := redisClient.HSetNX(key, field, value).Result()
	if err != nil {
		log.Fatal("redis HSetNX is err, err : ", err)
	}
}

//返回哈希表中给定域的值
func HGet(key, field string) (value string) {
	value, err := redisClient.HGet(key, field).Result()
	if err != nil {
		log.Fatal("redis HGet is err, err : ", err)
	}
	return value
}

//检查给定域 field 是否存在于哈希表 hash 当中
func HExists(key, field string) (flag bool) {
	flag, err := redisClient.HExists(key, field).Result()
	if err != nil {
		log.Fatal("redis HExists is err, err : ", err)
	}
	return flag
}

//删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
func HDel(key string, fields ...string) {
	for _, value := range fields {
		_, err := redisClient.HDel(key, value).Result()
		if err != nil {
			log.Fatal("redis HDel is err, err : ", err)
		}
	}
}

//返回哈希表 key 中域的数量
func HLen(key string) (i int) {
	size, err := redisClient.HLen(key).Result()
	if err != nil {
		log.Fatal("redis HLen is err, err : ", err)
	}
	return int(size)
}

//同时将多个 field-value (域-值)对设置到哈希表 key 中
//此命令会覆盖哈希表中已存在的域
//如果 key 不存在，一个空哈希表被创建并执行 HMSET 操作
func HMSet(key string, fields map[string]interface{}) {
	_, err := redisClient.HMSet(key, fields).Result()
	if err != nil {
		log.Fatal("redis HMSet is err, err : ", err)
	}
}

//将一个或多个值 value 插入到列表 key 的表头
//如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表头： 比如说，对空列表 mylist 执行命令 LPUSH mylist a b c ，列表的值将是 c b a ，这等同于原子性地执行 LPUSH mylist a 、 LPUSH mylist b 和 LPUSH mylist c 三个命令。
//如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。
//当 key 存在但不是列表类型时，返回一个错误。
func LPush(key string, values ...interface{}) {
	//for _, value := range values {
	_, err := redisClient.LPush(key, values).Result()
	if err != nil {
		log.Fatal("redis LPush is err, err : ", err)
	}
	//}
}

//将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表。
//和 LPUSH key value [value …] 命令相反，当 key 不存在时， LPUSHX 命令什么也不做
func LPushX(key string, value interface{}) {
	_, err := redisClient.LPushX(key, value).Result()
	if err != nil {
		log.Fatal("redis LPushX is err, err : ", err)
	}
}

//将一个或多个值 value 插入到列表 key 的表尾(最右边)。
//如果有多个 value 值，那么各个 value 值按从左到右的顺序依次插入到表尾：比如对一个空列表 mylist 执行 RPUSH mylist a b c ，得出的结果列表为 a b c ，等同于执行命令 RPUSH mylist a 、 RPUSH mylist b 、 RPUSH mylist c 。
//如果 key 不存在，一个空列表会被创建并执行 RPUSH 操作。
func RPush(key string, values ...interface{}) {
	//for _, value := range values {
	_, err := redisClient.RPush(key, values).Result()
	if err != nil {
		log.Fatal("redis RPush is err, err : ", err)
	}
	//}
}

//将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表。
//和 RPUSH key value [value …] 命令相反，当 key 不存在时， RPUSHX 命令什么也不做
func RPushX(key string, value interface{}) {
	_, err := redisClient.RPushX(key, value).Result()
	if err != nil {
		log.Fatal("redis RPushX is err, err : ", err)
	}
}

//移除并返回列表 key 的头元素
func LPop(key string) (value interface{}) {
	value, err := redisClient.LPop(key).Result()
	if err != nil {
		log.Fatal("redis LPop is err, err : ", err)
	}
	return value
}

//移除并返回列表 key 的尾元素
func RPop(key string) (value interface{}) {
	value, err := redisClient.RPop(key).Result()
	if err != nil {
		log.Fatal("redis RPop is err, err : ", err)
	}
	return value
}

//命令 RPOPLPUSH 在一个原子时间内，执行以下两个动作：
//将列表 source 中的最后一个元素(尾元素)弹出，并返回给客户端。
//将 source 弹出的元素插入到列表 destination ，作为 destination 列表的的头元素。
//举个例子，你有两个列表 source 和 destination ， source 列表有元素 a, b, c ， destination 列表有元素 x, y, z ，执行 RPOPLPUSH source destination 之后， source 列表包含元素 a, b ， destination 列表包含元素 c, x, y, z ，并且元素 c 会被返回给客户端。
//如果 source 不存在，值 nil 被返回，并且不执行其他动作。
//如果 source 和 destination 相同，则列表中的表尾元素被移动到表头，并返回该元素，可以把这种特殊情况视作列表的旋转(rotation)操作
func RPopLPush(key1, key2 string) (value interface{}) {
	value, err := redisClient.RPopLPush(key1, key2).Result()
	if err != nil {
		log.Fatal("redis RPopLPush is err, err : ", err)
	}
	return value
}

//根据参数 count 的值，移除列表中与参数 value 相等的元素。
//count 的值可以是以下几种：
//count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count 。
//count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值。
//count = 0 : 移除表中所有与 value 相等的值
func LRem(key string, count int, value interface{}) {
	_, err := redisClient.LRem(key, int64(count), value).Result()
	if err != nil {
		log.Fatal("redis LRem is err, err : ", err)
	}
}

//返回列表 key 中，下标为 index 的元素。
//下标(index)参数 start 和 stop 都以 0 为底，也就是说，以 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。
//你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func LIndex(key string, index int) (value interface{}) {
	value, err := redisClient.LIndex(key, int64(index)).Result()
	if err != nil {
		log.Fatal("redis LIndex is err, err : ", err)
	}
	return value
}

//BLPOP 是列表的阻塞式(blocking)弹出原语。
//它是 LPOP key 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BLPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
//当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素
//超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely)
func BLPop(timeout time.Duration, key string) (value interface{}) {
	value, err := redisClient.BLPop(timeout, key).Result()
	if err != nil {
		log.Fatal("redis BLPop is err, err : ", err)
	}
	return value
}

//BRPOP 是列表的阻塞式(blocking)弹出原语。
//它是 RPOP key 命令的阻塞版本，当给定列表内没有任何元素可供弹出的时候，连接将被 BRPOP 命令阻塞，直到等待超时或发现可弹出元素为止。
//当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的尾部元素
//超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely)
func BRPop(timeout time.Duration, key string) (value interface{}) {
	value, err := redisClient.BRPop(timeout, key).Result()
	if err != nil {
		log.Fatal("redis BRPop is err, err : ", err)
	}
	return value
}

//BRPOPLPUSH 是 RPOPLPUSH source destination 的阻塞版本，当给定列表 source 不为空时， BRPOPLPUSH 的表现和 RPOPLPUSH source destination 一样。
//当列表 source 为空时， BRPOPLPUSH 命令将阻塞连接，直到等待超时，或有另一个客户端对 source 执行 LPUSH key value [value …] 或 RPUSH key value [value …] 命令为止。
//超时参数 timeout 接受一个以秒为单位的数字作为值。超时参数设为 0 表示阻塞时间可以无限期延长(block indefinitely)
func BRPopLPush(key1, key2 string, timeout time.Duration) (value interface{}) {
	value, err := redisClient.BRPopLPush(key1, key2, timeout).Result()
	if err != nil {
		log.Fatal("redis BRPopLPush is err, err : ", err)
	}
	return value
}

//将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略。
//假如 key 不存在，则创建一个只包含 member 元素作成员的集合
func SAdd(key string, values ...interface{}) {
	//for _, value := range values {
	_, err := redisClient.SAdd(key, values).Result()
	if err != nil {
		log.Fatal("redis SAdd is err, err : ", err)
	}
	//}
}

//移除并返回集合中的一个随机元素
func SPop(key string) (value interface{}) {
	value, err := redisClient.SPop(key).Result()
	if err != nil {
		log.Fatal("redis SPop is err, err : ", err)
	}
	return value
}

//返回集合中的一个随机元素
//仅仅返回随机元素，而不对集合进行任何改动
func SRandMember(key string) (value interface{}) {
	value, err := redisClient.SRandMember(key).Result()
	if err != nil {
		log.Fatal("redis SRandMember is err, err : ", err)
	}
	return value
}

//如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合。
//如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值
//仅仅返回随机元素，而不对集合进行任何改动
func SRandMemberN(key string, count int) (value interface{}) {
	value, err := redisClient.SRandMemberN(key, int64(count)).Result()
	if err != nil {
		log.Fatal("redis SRandMemberN is err, err : ", err)
	}
	return value
}

//移除集合 key 中的一个或多个 member 元素，不存在的 member 元素会被忽略
func SRem(key string, members ...interface{}) {
	_, err := redisClient.SRem(key, members).Result()
	if err != nil {
		log.Fatal("redis SRem is err, err : ", err)
	}
}

//SMOVE 是原子性操作。
//如果 source 集合不存在或不包含指定的 member 元素，则 SMOVE 命令不执行任何操作，仅返回 0 。否则， member 元素从 source 集合中被移除，并添加到 destination 集合中去。
//当 destination 集合已经包含 member 元素时， SMOVE 命令只是简单地将 source 集合中的 member 元素删除。
//当 source 或 destination 不是集合类型时，返回一个错误
func SMove(key1, key2 string, member interface{}) {
	_, err := redisClient.SMove(key1, key2, member).Result()
	if err != nil {
		log.Fatal("redis SMove is err, err : ", err)
	}
}

//返回集合 key 中的所有成员。
//不存在的 key 被视为空集合。
func SMembers(key string) (value interface{}) {
	value, err := redisClient.SMembers(key).Result()
	if err != nil {
		log.Fatal("redis SMembers is err, err : ", err)
	}
	return value
}
