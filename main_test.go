package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomsa/nats-api/hander"
	natsPB "github.com/gomsa/nats-api/proto/nats"
)

func TestMobileVerify(t *testing.T) {

	h := hander.Nats{}
	req := &natsPB.Request{
		Mobile: `13954386521`,
	}
	res := &natsPB.Response{}
	err := h.MobileVerify(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
