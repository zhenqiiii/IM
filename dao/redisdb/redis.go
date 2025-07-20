package redisdb

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 上下文
var ctx = context.Background()

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.pwd"),
		DB:       1, //使用默认DB
	})
}

// 根据用户email获取验证码：用于比对验证码
func Get(email string) (code string, err error) {
	// 不存在该code时,报Nil error
	code, err = RDB.Get(ctx, email).Result()
	if err != nil {
		log.Println("查找失败：" + err.Error())
		return "", err
	}

	return code, nil

}

// 设置验证码
func Set(email string, code string) error {
	// 若用户在过期时间内发送了多次验证码请求
	// Set会自动覆盖旧值，redis中的验证码保持最新的那个
	// 当然需要在前端设置按钮的下一次点击间隔时间
	// code过期时间30分钟
	err := RDB.Set(ctx, email, code, time.Minute*30).Err()
	if err != nil {
		log.Println("验证码键值对创建失败：" + err.Error())
		return err
	}

	return nil

}
