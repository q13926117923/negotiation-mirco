package core

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"

	"user/model"
	pb "user/services/proto"
)

func BuildUser(item model.User) *pb.UserModel{
	userModel := pb.UserModel{
		Id: uint32(item.ID),
		UserName: item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
func (*UserService)UserLogin(ctx context.Context,req *pb.UserRequest,resp *pb.UserResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?",req.UserName).First(&user).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err){
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false{
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
func (*UserService)UserRegister(ctx context.Context, req *pb.UserRequest,resp *pb.UserResponse)error{


	if req.Password != req.PasswordConfirm{
		err := errors.New("输入密码不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error;err!=nil{
		return err
	}
	if count>0{
		err:=errors.New("用户名已经存在")
		return err
	}

	user:= model.User{
		UserName: req.UserName,
	}
	if err:=user.SetPassword(req.Password);err!=nil{
		return err //设置密码
	}
	if err := model.DB.Create(&user).Error;err!=nil{
		return err //创建用户
	}
	resp.UserDetail = BuildUser(user)
	return nil


	return nil
}
