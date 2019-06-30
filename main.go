package main

import (
	// 公共引入

	"github.com/micro/go-micro/util/log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"


	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"

	"github.com/gomsa/nats-api/hander"
	natsPB "github.com/gomsa/nats-api/proto/nats"
	templatePB "github.com/gomsa/nats-api/proto/template"

)

func main() {
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
	}
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()

	natsPB.RegisterNatsHandler(srv.Server(), &hander.Nats{})
	templatePB.RegisterTemplatesHandler(srv.Server(), &hander.Template{})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	if err := client.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
