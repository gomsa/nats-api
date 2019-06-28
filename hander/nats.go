package hander

import (
	"context"
	"errors"

	pb "github.com/gomsa/nats-api/proto/nats"
	"github.com/gomsa/nats/client"
	natsPB "github.com/gomsa/nats/proto/nats"
)

// Nats 信息结构
type Nats struct {
}

// MobileVerify 手机验证码
func (srv *Nats) MobileVerify(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
		
	request := &natsPB.Request{
		Addressee: req.,
		Event:     `register_verify`,
		Type:      `sms`,
		QueryParams: map[string]string{
			`code`: `123456`,
		},
	}
	roleRes, err := client.Nats.ProcessCommonRequest(ctx, request)
	if err != nil {
		return err
	}
	if !roleRes.Valid {
		return errors.New("发送验证码失败")
	}
	return err
}
