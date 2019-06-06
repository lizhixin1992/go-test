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

//返回一个集合的全部成员，该集合是所有给定集合的交集。
//不存在的 key 被视为空集。
//当给定集合当中有一个空集时，结果也为空集(根据集合运算定律)。
func SInter(key1, key2 string) (value interface{}) {
	value, err := redisClient.SInter(key1, key2).Result()
	if err != nil {
		log.Fatal("redis SInter is err, err : ", err)
	}
	return value
}

//这个命令类似于 SINTER key [key …] 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集。
//如果 destination 集合已经存在，则将其覆盖。
//destination 可以是 key 本身。
func SInterStore(destination, key1, key2 string) {
	_, err := redisClient.SInterStore(destination, key1, key2).Result()
	if err != nil {
		log.Fatal("redis SInterStore is err, err : ", err)
	}
}

//返回一个集合的全部成员，该集合是所有给定集合的并集。
//不存在的 key 被视为空集。
func SUnion(key1, key2 string) (value interface{}) {
	value, err := redisClient.SUnion(key1, key2).Result()
	if err != nil {
		log.Fatal("redis SUnion is err, err : ", err)
	}
	return value
}

//这个命令类似于 SUNION key [key …] 命令，但它将结果保存到 destination 集合，而不是简单地返回结果集。
//如果 destination 已经存在，则将其覆盖。
//destination 可以是 key 本身。
func SUnionStore(destination, key1, key2 string) {
	_, err := redisClient.SUnionStore(destination, key1, key2).Result()
	if err != nil {
		log.Fatal("redis SUnionStore is err, err : ", err)
	}
}

//返回一个集合的全部成员，该集合是所有给定集合之间的差集。
//不存在的 key 被视为空集
//例如：test1:1,2,3,4,5,6,7	test2:4,5,6,7,8,9,0		sdiff test1 test2	返回	1,2,3		sdiff test2 test1 返回8,9,0
func SDiff(key1, key2 string) (value interface{}) {
	value, err := redisClient.SDiff(key1, key2).Result()
	if err != nil {
		log.Fatal("redis SDiff is err, err : ", err)
	}
	return value
}

//这个命令的作用和 SDIFF key [key …] 类似，但它将结果保存到 destination 集合，而不是简单地返回结果集。
//如果 destination 集合已经存在，则将其覆盖。
//destination 可以是 key 本身
func SDiffStore(destination, key1, key2 string) {
	_, err := redisClient.SDiffStore(destination, key1, key2).Result()
	if err != nil {
		log.Fatal("redis SDiffStore is err, err : ", err)
	}
}

//将一个或多个 member 元素及其 score 值加入到有序集 key 当中。
//如果某个 member 已经是有序集的成员，那么更新这个 member 的 score 值，并通过重新插入这个 member 元素，来保证该 member 在正确的位置上。
//如果 key 不存在，则创建一个空的有序集并执行 ZADD 操作。
//当 key 存在但不是有序集类型时，返回一个错误。
func ZAdd(key string, member redis.Z) {
	_, err := redisClient.ZAdd(key, member).Result()
	if err != nil {
		log.Fatal("redis ZAdd is err, err : ", err)
	}
}

//返回有序集 key 中，成员 member 的 score 值。
//如果 member 元素不是有序集 key 的成员，或 key 不存在，返回 nil
func ZScore(key, member string) (value interface{}) {
	value, err := redisClient.ZScore(key, member).Result()
	if err != nil {
		log.Fatal("redis ZScore is err, err : ", err)
	}
	return value
}

//为有序集 key 的成员 member 的 score 值加上增量 increment 。
//可以通过传递一个负数值 increment ，让 score 减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
//当 key 不存在，或 member 不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
//当 key 不是有序集类型时，返回一个错误。
//score 值可以是整数值或双精度浮点数
func ZIncrBy(key string, increment float64, member string) {
	_, err := redisClient.ZIncrBy(key, increment, member).Result()
	if err != nil {
		log.Fatal("redis ZIncrBy is err, err : ", err)
	}
}

//返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
func ZCount(key, min, max string) (value int64) {
	value, err := redisClient.ZCount(key, min, max).Result()
	if err != nil {
		log.Fatal("redis ZCount is err, err : ", err)
	}
	return value
}

//返回有序集 key 中，指定区间内的成员。
//其中成员的位置按 score 值递增(从小到大)来排序。
//具有相同 score 值的成员按字典序(lexicographical order )来排列
//下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。 你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推
//超出范围的下标并不会引起错误。 比如说，当 start 的值比有序集的最大下标还要大，或是 start > stop 时， ZRANGE 命令只是简单地返回一个空列表。 另一方面，假如 stop 参数的值比有序集的最大下标还要大，那么 Redis 将 stop 当作最大下标来处理
func ZRange(key string, start, stop int64) (value interface{}) {
	value, err := redisClient.ZRange(key, start, stop).Result()
	if err != nil {
		log.Fatal("redis ZRange is err, err : ", err)
	}
	return value
}

//和ZRange基本一样
//通过使用 WITHSCORES 选项，来让成员和它的 score 值一并返回，返回列表以 value1,score1, ..., valueN,scoreN 的格式表示
func ZRangeWithScores(key string, start, stop int64) (value interface{}) {
	value, err := redisClient.ZRangeWithScores(key, start, stop).Result()
	if err != nil {
		log.Fatal("redis ZRangeWithScores is err, err : ", err)
	}
	return value
}

//返回有序集 key 中，指定区间内的成员。
//其中成员的位置按 score 值递减(从大到小)来排列。 具有相同 score 值的成员按字典序的逆序(reverse lexicographical order)排列
func ZRevRange(key string, start, stop int64) (value interface{}) {
	value, err := redisClient.ZRevRange(key, start, stop).Result()
	if err != nil {
		log.Fatal("redis ZRevRange is err, err : ", err)
	}
	return value
}

//和ZRevRange基本一样
//通过使用 WITHSCORES 选项，来让成员和它的 score 值一并返回，返回列表以 value1,score1, ..., valueN,scoreN 的格式表示
func ZRevRangeWithScores(key string, start, stop int64) (value interface{}) {
	value, err := redisClient.ZRevRangeWithScores(key, start, stop).Result()
	if err != nil {
		log.Fatal("redis ZRevRangeWithScores is err, err : ", err)
	}
	return value
}

//返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
//具有相同 score 值的成员按字典序(lexicographical order)来排列(该属性是有序集提供的，不需要额外的计算)。
//可选的 LIMIT 参数指定返回结果的数量及区间(就像SQL中的 SELECT LIMIT offset, count )，注意当 offset 很大时，定位 offset 的操作可能需要遍历整个有序集，此过程最坏复杂度为 O(N) 时间
//min 和 max 可以是 -inf 和 +inf ，这样一来，你就可以在不知道有序集的最低和最高 score 值的情况
//默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)
func ZRangeByScore(key string, opt redis.ZRangeBy) (value interface{}) {
	value, err := redisClient.ZRangeByScore(key, opt).Result()
	if err != nil {
		log.Fatal("redis ZRangeByScore is err, err : ", err)
	}
	return value
}

//和ZRangeByScore基本一样
//通过使用 WITHSCORES 选项，来让成员和它的 score 值一并返回，返回列表以 value1,score1, ..., valueN,scoreN 的格式表示
func ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) (value interface{}) {
	value, err := redisClient.ZRangeByScoreWithScores(key, opt).Result()
	if err != nil {
		log.Fatal("redis ZRangeByScoreWithScores is err, err : ", err)
	}
	return value
}

//和ZRangeByScore类似
//返回有序集 key 中， score 值介于 max 和 min 之间(默认包括等于 max 或 min )的所有的成员。有序集成员按 score 值递减(从大到小)的次序排列。
//具有相同 score 值的成员按字典序的逆序(reverse lexicographical order )排列
func ZRevRangeByScore(key string, opt redis.ZRangeBy) (value interface{}) {
	value, err := redisClient.ZRevRangeByScore(key, opt).Result()
	if err != nil {
		log.Fatal("redis ZRevRangeByScore is err, err : ", err)
	}
	return value
}

//移除有序集 key 中的一个或多个成员，不存在的成员将被忽略。
//当 key 存在但不是有序集类型时，返回一个错误
func ZRem(key string, members ...interface{}) {
	_, err := redisClient.ZRem(key, members).Result()
	if err != nil {
		log.Fatal("redis ZRem is err, err : ", err)
	}
}

//移除有序集 key 中，指定排名(rank)区间内的所有成员。
//区间分别以下标参数 start 和 stop 指出，包含 start 和 stop 在内。
//下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。 你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推
func ZRemRangeByRank(key string, start, stop int64) {
	_, err := redisClient.ZRemRangeByRank(key, start, stop).Result()
	if err != nil {
		log.Fatal("redis ZRemRangeByRank is err, err : ", err)
	}
}

//移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
func ZRemRangeByScore(key, min, max string) {
	_, err := redisClient.ZRemRangeByScore(key, min, max).Result()
	if err != nil {
		log.Fatal("redis ZRemRangeByScore is err, err : ", err)
	}
}

//当有序集合的所有成员都具有相同的分值时， 有序集合的元素会根据成员的字典序（lexicographical ordering）来进行排序， 而这个命令则可以返回给定的有序集合键 key 中， 值介于 min 和 max 之间的成员
//如果有序集合里面的成员带有不同的分值， 那么命令返回的结果是未指定的（unspecified）。
//命令会使用 C 语言的 memcmp() 函数， 对集合中的每个成员进行逐个字节的对比（byte-by-byte compare）， 并按照从低到高的顺序， 返回排序后的集合成员。 如果两个字符串有一部分内容是相同的话， 那么命令会认为较长的字符串比较短的字符串要大。
//可选的 LIMIT offset count 参数用于获取指定范围内的匹配元素 （就像 SQL 中的 SELECT LIMIT offset count 语句）。 需要注意的一点是， 如果 offset 参数的值非常大的话， 那么命令在返回结果之前， 需要先遍历至 offset 所指定的位置， 这个操作会为命令加上最多 O(N) 复杂度
//合法的 min 和 max 参数必须包含 ( 或者 [ ， 其中 ( 表示开区间（指定的值不会被包含在范围之内）， 而 [ 则表示闭区间（指定的值会被包含在范围之内）。
//特殊值 + 和 - 在 min 参数以及 max 参数中具有特殊的意义， 其中 + 表示正无限， 而 - 表示负无限。 因此， 向一个所有成员的分值都相同的有序集合发送命令 ZRANGEBYLEX <zset> - + ， 命令将返回有序集合中的所有元素
func ZRangeByLex(key string, opt redis.ZRangeBy) (value interface{}) {
	value, err := redisClient.ZRangeByLex(key, opt).Result()
	if err != nil {
		log.Fatal("redis ZRangeByLex is err, err : ", err)
	}
	return value
}

//将任意数量的元素添加到指定的 HyperLogLog 里面。
//作为这个命令的副作用， HyperLogLog 内部可能会被更新， 以便反映一个不同的唯一元素估计数量（也即是集合的基数）。
//如果 HyperLogLog 估计的近似基数（approximated cardinality）在命令执行之后出现了变化， 那么命令返回 1 ， 否则返回 0 。 如果命令执行时给定的键不存在， 那么程序将先创建一个空的 HyperLogLog 结构， 然后再执行命令。
//调用 PFADD key element [element …] 命令时可以只给定键名而不给定元素：
//如果给定键已经是一个 HyperLogLog ， 那么这种调用不会产生任何效果；
//但如果给定的键不存在， 那么命令会创建一个空的 HyperLogLog ， 并向客户端返回 1
func PFAdd(key string, els ...interface{}) {
	_, err := redisClient.PFAdd(key, els).Result()
	if err != nil {
		log.Fatal("redis PFAdd is err, err : ", err)
	}
}

//当 PFCOUNT key [key …] 命令作用于单个键时， 返回储存在给定键的 HyperLogLog 的近似基数， 如果键不存在， 那么返回 0 。
//当 PFCOUNT key [key …] 命令作用于多个键时， 返回所有给定 HyperLogLog 的并集的近似基数， 这个近似基数是通过将所有给定 HyperLogLog 合并至一个临时 HyperLogLog 来计算得出的。
//通过 HyperLogLog 数据结构， 用户可以使用少量固定大小的内存， 来储存集合中的唯一元素 （每个 HyperLogLog 只需使用 12k 字节内存，以及几个字节的内存来储存键本身）。
//命令返回的可见集合（observed set）基数并不是精确值， 而是一个带有 0.81% 标准错误（standard error）的近似值。
//举个例子， 为了记录一天会执行多少次各不相同的搜索查询， 一个程序可以在每次执行搜索查询时调用一次 PFADD key element [element …] ， 并通过调用 PFCOUNT key [key …] 命令来获取这个记录的近似结果
func PFCount(key string) (value int64) {
	value, err := redisClient.PFCount(key).Result()
	if err != nil {
		log.Fatal("redis PFCount is err, err : ", err)
	}
	return value
}

//将多个 HyperLogLog 合并（merge）为一个 HyperLogLog ， 合并后的 HyperLogLog 的基数接近于所有输入 HyperLogLog 的可见集合（observed set）的并集。
//合并得出的 HyperLogLog 会被储存在 destkey 键里面， 如果该键并不存在， 那么命令在执行之前， 会先为该键创建一个空的 HyperLogLog
func PFMerge(dest string, keys string) {
	_, err := redisClient.PFMerge(dest, keys).Result()
	if err != nil {
		log.Fatal("redis PFMerge is err, err : ", err)
	}
}

//检查给定 key 是否存在
//若 key 存在，返回 1 ，否则返回 0
func Exists(key string) (value bool) {
	flag, err := redisClient.Exists(key).Result()
	if err != nil {
		log.Fatal("redis Exists is err, err : ", err)
	} else {
		if flag == 1 {
			value = true
		} else {
			value = false
		}
	}
	return value
}

//返回 key 所储存的值的类型
//none (key不存在)
//string (字符串)
//list (列表)
//set (集合)
//zset (有序集)
//hash (哈希表)
//stream （流）
func Type(key string) (value string) {
	value, err := redisClient.Type(key).Result()
	if err != nil {
		log.Fatal("redis Type is err, err : ", err)
	}
	return value
}

//将 key 改名为 newkey 。
//当 key 和 newkey 相同，或者 key 不存在时，返回一个错误。
//当 newkey 已经存在时， RENAME 命令将覆盖旧值
func Rename(key, newkey string) {
	_, err := redisClient.Rename(key, newkey).Result()
	if err != nil {
		log.Fatal("redis Rename is err, err : ", err)
	}
}

//当且仅当 newkey 不存在时，将 key 改名为 newkey 。
//当 key 不存在时，返回一个错误
//修改成功时，返回 true ； 如果 newkey 已经存在，返回 false
func RenameNX(key, newkey string) (value bool) {
	value, err := redisClient.RenameNX(key, newkey).Result()
	if err != nil {
		log.Fatal("redis RenameNX is err, err : ", err)
	}
	return value
}

//将当前数据库的 key 移动到给定的数据库 db 当中。
//如果当前数据库(源数据库)和给定数据库(目标数据库)有相同名字的给定 key ，或者 key 不存在于当前数据库，那么 MOVE 没有任何效果。
//因此，也可以利用这一特性，将 MOVE 当作锁(locking)原语(primitive)
//移动成功返回 true ，失败则返回 false
func Move(key string, db int64) (value bool) {
	value, err := redisClient.Move(key, db).Result()
	if err != nil {
		log.Fatal("redis Move is err, err : ", err)
	}
	return value
}

//删除给定的一个或多个 key 。
//不存在的 key 会被忽略
func Del(key string) {
	_, err := redisClient.Del(key).Result()
	if err != nil {
		log.Fatal("redis Del is err, err : ", err)
	}
}

//从当前数据库中随机返回(不删除)一个 key
//当数据库不为空时，返回一个 key 。 当数据库为空时，返回 nil
func RandomKey() (value string) {
	value, err := redisClient.RandomKey().Result()
	if err != nil {
		log.Fatal("redis RandomKey is err, err : ", err)
	}
	return value
}

//返回当前数据库的 key 的数量
func DBSize() (value int64) {
	value, err := redisClient.DBSize().Result()
	if err != nil {
		log.Fatal("redis DBSize is err, err : ", err)
	}
	return value
}

//查找所有符合给定模式 pattern 的 key ， 比如说：
//KEYS * 匹配数据库中所有 key 。
//KEYS h?llo 匹配 hello ， hallo 和 hxllo 等。
//KEYS h*llo 匹配 hllo 和 heeeeello 等。
//KEYS h[ae]llo 匹配 hello 和 hallo ，但不匹配 hillo 。
//特殊符号用 \ 隔开
//
//KEYS 的速度非常快，但在一个大的数据库中使用它仍然可能造成性能问题，如果你需要从一个数据集中查找特定的 key ，你最好还是用 Redis 的集合结构(set)来代替
func Keys(pattern string) (value []string) {
	value, err := redisClient.Keys(pattern).Result()
	if err != nil {
		log.Fatal("redis Keys is err, err : ", err)
	}
	return value
}
