package hander

import (
	"context"

	pb "github.com/gomsa/nats-api/proto/template"
	"github.com/gomsa/nats/client"
	templatePB "github.com/gomsa/nats/proto/template"
	"github.com/gomsa/tools/uitl"
)

// Template 角色结构
type Template struct {
}

// List 角色列表
func (srv *Template) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	query := &templatePB.ListQuery{}
	err = uitl.Data2Data(req, query)
	if err != nil {
		return err
	}
	templateRes, err := client.Template.List(ctx, query)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(templateRes, res)
	if err != nil {
		return err
	}
	return err
}

// Get 获取角色
func (srv *Template) Get(ctx context.Context, req *pb.Template, res *pb.Response) (err error) {
	template := &templatePB.Template{}
	err = uitl.Data2Data(req, template)
	if err != nil {
		return err
	}
	templateRes, err := client.Template.Get(ctx, template)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(templateRes, res)
	if err != nil {
		return err
	}
	return err
}

// Create 创建角色
func (srv *Template) Create(ctx context.Context, req *pb.Template, res *pb.Response) (err error) {
	template := &templatePB.Template{}
	err = uitl.Data2Data(req, template)
	if err != nil {
		return err
	}
	templateRes, err := client.Template.Create(ctx, template)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(templateRes, res)
	if err != nil {
		return err
	}
	return err
}

// Update 更新角色
func (srv *Template) Update(ctx context.Context, req *pb.Template, res *pb.Response) (err error) {
	template := &templatePB.Template{}
	err = uitl.Data2Data(req, template)
	if err != nil {
		return err
	}
	templateRes, err := client.Template.Update(ctx, template)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(templateRes, res)
	if err != nil {
		return err
	}
	return err
}

// Delete 删除角色
func (srv *Template) Delete(ctx context.Context, req *pb.Template, res *pb.Response) (err error) {
	template := &templatePB.Template{
		Id: req.Id,
	}
	if err != nil {
		return err
	}
	templateRes, err := client.Template.Delete(ctx, template)
	if err != nil {
		return err
	}
	err = uitl.Data2Data(templateRes, res)
	if err != nil {
		return err
	}
	return err
}
