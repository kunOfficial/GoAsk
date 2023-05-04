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
	answerLikesPrefix    = "likes_answer:"
	addAnswerLikesPrefix = "likes_answer_add:"
)

func getAddAnswerLikesKey(qid int) string {
	return addAnswerLikesPrefix + strconv.Itoa(qid)
}

func getAnswerLikesKey(qid int) string {
	return answerLikesPrefix + strconv.Itoa(qid)
}

// AddLike 往 "likes_answer_add:" 和 "likes_answer:" 里添加 Like
func (client *CacheClient) AddLike(qid uint, uid uint) (liked bool, err error) {
	// 开启一个事务
	result, err := client.SAdd(getAnswerLikesKey(int(qid)), uid).Result()
	if err != nil { // 如果出问题
		return false, err
	}
	if result == 0 { // 如果返回值为0，说明插入失败，即已经点过赞
		return true, nil
	}
	err = client.SAdd(getAddAnswerLikesKey(int(qid)), uid).Err()
	if err != nil {
		return false, err
	}
	return false, nil

}

// LoadLikes 批量往 "likes_answer:" 里添加 Likes
func (client *CacheClient) LoadLikes(qid uint, uids []uint) error {
	err := client.SAdd(getAnswerLikesKey(int(qid)), uids).Err()
	if err == redis.Nil {
		return nil
	}
	return err
}

func (client *CacheClient) GetLikes(aid uint) (FoundKey bool, likes int64, err error) {
	likes, err = client.SCard(getAnswerLikesKey(int(aid))).Result()
	if err != nil {
		if err == redis.Nil {
			return false, 0, err
		} else {
			return false, 0, nil
		}
	}
	return true, likes, nil
}

func (client *CacheClient) IsLiked(qid uint, uid uint) (FoundKey bool, isLiked bool, err error) {
	isLiked, err = client.SIsMember(getAnswerLikesKey(int(qid)), uid).Result()
	if err != nil {
		if err == redis.Nil { // 如果键不存在
			return false, false, nil
		} else {
			return false, false, err
		}
	}
	return true, isLiked, nil
}

// RemoveLike 从 cache 中删除点赞记录
func (client *CacheClient) RemoveLike(qid uint, uid uint) (inAddSet bool, err error) {
	// 开启事务
	pipe := client.TxPipeline()
	_ = pipe.SRem(getAnswerLikesKey(int(qid)), uid)
	rem2 := pipe.SRem(getAddAnswerLikesKey(int(qid)), uid)
	_, txErr := pipe.Exec()
	if txErr != nil {
		return false, txErr
	}
	if rem2.Err() == redis.Nil {
		return false, nil
	}
	return true, nil
}

// syncLikes 将 redis 中的 view 同步到 mysql 中
func syncLikes() {
	//fmt.Println("syncLikes")
	// Redis是单线程的，所以此时会阻塞其他查询缓存请求，为了不阻塞太久，所以设置 500ms 的执行时间
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()
	cacheClient := NewCacheClient(ctx)
	dbClient := dao.NewDbClient(ctx)
	iter := cacheClient.Scan(0, addAnswerLikesPrefix+"*", 1000).Iterator()
	for iter.Next() { // 对每个回答进行遍历
		key := iter.Val()
		fmt.Println("keys:", key)
		aid, err := strconv.ParseUint(key, 10, 64)
		if err != nil {
			utils.Logger.Info(err)
		}
		// 获取该回答的点赞数据
		strUids, err := cacheClient.SMembers(iter.Val()).Result()
		if err != nil {
			utils.Logger.Info(err)
		}
		uids := make([]uint, len(strUids))
		for i, strUid := range strUids {
			uid, err := strconv.ParseUint(strUid, 10, 64)
			if err != nil {
				utils.Logger.Info(err)
			}
			uids[i] = uint(uid)
		}
		// 批量添加到数据库中
		err = dbClient.AddAnswerLikes(uint(aid), uids)
		if err != nil {
			utils.Logger.Info(err)
		}
		err = cacheClient.Del(iter.Val()).Err()
		if err != nil {
			utils.Logger.Info(err)
		}
	}
}
