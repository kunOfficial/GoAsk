package cache

import (
	"GoAsk/dao"
	"GoAsk/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	questionViewPrefix = "view_question:"
)

func getQuestionViewKey(qid int) string {
	return questionViewPrefix + strconv.Itoa(qid)
}

// AddQuestionView 浏览数自增
func (client *CacheClient) AddQuestionView(qid int) error {
	err := client.Incr(getQuestionViewKey(qid)).Err()
	if err == redis.Nil { // 如果键不存在
		// NOTE: 暂时设置的无过期时间
		return client.Set(getQuestionViewKey(qid), "1", 0).Err()
	}
	return err
}

// GetQuestionView 获取浏览数
func (client *CacheClient) GetQuestionView(qid int) (view int, err error) {
	viewStr, err := client.Get(getQuestionViewKey(qid)).Result()
	if err != nil {
		return 0, err
	}
	viewNum, err := strconv.Atoi(viewStr)
	if err != nil {
		return 0, err
	}
	return viewNum, nil
}

// syncViews 将 redis 中的 view 同步到 mysql 中
func syncViews() {
	// Redis是单线程的，所以此时会阻塞其他查询缓存请求，为了不阻塞太久，所以设置 500ms 的执行时间
	//fmt.Println("syncViews")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	cacheClient := NewCacheClient(ctx)
	dbClient := dao.NewDbClient(ctx)
	//fmt.Println(cacheClient.Keys(questionViewPrefix + "*").String())
	iter := cacheClient.Scan(0, questionViewPrefix+"*", 1000).Iterator()
	for iter.Next() {
		key := iter.Val()
		fmt.Println("keys:", key)
		qid, err := strconv.ParseUint(key[len(questionViewPrefix):], 10, 64)
		if err != nil {
			utils.Logger.Info(err)
		}
		viewNumStr, err := cacheClient.Get(iter.Val()).Result()
		if err != nil {
			utils.Logger.Info(err)
		}
		viewNum, err := strconv.ParseUint(viewNumStr, 10, 64)
		if err != nil {
			utils.Logger.Info(err)
		}
		found, err := dbClient.AddQuestionView(uint(qid), uint(viewNum))
		if err != nil {
			utils.Logger.Info(err)
		}
		if !found { // 如果没找到该问题，说明问题已经被删除了，所以需要删掉该键
			cacheClient.Del(key)
		}
		// 用完后删除该key, 其实这里要保证线程安全的话要当做事务来做，但是由于这里 redis 和 mysql 操作混杂，所以不太好处理
		cacheClient.Del(key)
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
