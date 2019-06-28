package hander

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/util/log"
	uuid "github.com/satori/go.uuid"

	pb "github.com/gomsa/nats-api/proto/nats"
	"github.com/gomsa/nats-api/providers/redis"
	"github.com/gomsa/nats/client"
	natsPB "github.com/gomsa/nats/proto/nats"
)

// Nats 信息结构
type Nats struct {
}

// MobileVerify 手机验证码
func (srv *Nats) MobileVerify(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	uuid, vrify, err := srv.Verify()
	if err != nil {
		log.Log(err)
		return err
	}
	request := &natsPB.Request{
		Addressee: req.Mobile,
		Event:     `register_verify`,
		Type:      `sms`,
		QueryParams: map[string]string{
			`code`: vrify,
		},
	}
	roleRes, err := client.Nats.ProcessCommonRequest(ctx, request)
		log.Log(err)
	if err != nil {
		return err
	}
	if !roleRes.Valid {
		log.Log(roleRes, err)
		return errors.New("发送验证码失败")
	}
	res.Valid = true
	res.Uuid = uuid
	return err
}

// Verify 生成验证码并保存到 redis
func (srv *Nats) Verify() (key string,vrify string, err error) {
	key = uuid.NewV4().String()
	vrify = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	redis := redis.NewClient()
	// 过期时间默认30分钟
	err = redis.SetNX(key, vrify, 30*time.Minute).Err()
	return key, vrify, err
}
