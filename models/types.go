package models

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/asaskevich/govalidator"
)

var(
	engineConn *xorm.Engine
	dberr  error
	devSql ="root:@tcp(127.0.0.1:3306)/stream?charset=utf8"
)


func init() {
	engine, err := xorm.NewEngine("mysql", devSql)
	dberr = err
	if dberr!=nil{
		panic(dberr.Error())
	}
	engineConn =engine
	engineConn.ShowSQL(true)
	//设置连接池的空闲数大小
	engineConn.SetMaxIdleConns(100)
	//设置最大打开连接数
	engineConn.SetMaxOpenConns(100)


	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engineConn.SetTableMapper(core.SnakeMapper{})
	//
	engineConn.Sync(new(UserInfo),new(VideoInfo),new(VideoComment))

	govalidator.SetNilPtrAllowedByRequired(true)
	govalidator.SetFieldsRequiredByDefault(true)
}

/**
用户信息定义
 */
type UserInfo struct {
	/**
	Id int 					`json:"id" 					valid:"id,optional"`
	UserName string 		`json:"user_name" 			valid:"user_name,required"`
	Pwd string 				`json:"pwd" 				valid:"pwd,required"`
	CreateUserDate string	`json:"create_user_date" 	valid:"create_user_date,required"`
	IsVip string			`json:"is_vip" 				valid:"is_vip,required"`
	**/
	Id int 					`json:"id" 				`
	UserName string 		`json:"user_name" 			`
	Pwd string 				`json:"pwd" 				`
	CreateUserDate string	`json:"create_user_date" 	`
	IsVip string			`json:"is_vip" 				`
}


//查询用户
type ReuqestUserParmas struct {
	UserName string 		`json:"user_name" 	`
}


//video data model
type VideoInfo struct {
	Id int `json:"id"`
	VideoId string `json:"video_id"` //视频id
	UserName string `json:"user_name"` //视频创建着
	VideoName string 	`json:"video_name"` //视频名称
	DisplayCtime string `json:"display_ctime"` //创建时间
	IsVip string `json:"is_vip"` //vip用户
}


//视频查询用户
type ReuqestVideoParmas struct {
	VideoId string `json:"video_id"`
}

//视频评价
type VideoComment struct {
	Id int `json:"id"`
	ConmmentId string `json:"conmment_id"` //视频评论id
	VideoId string `json:"video_id"` //视频id
	UserName string `json:"user_name"` //用户
	CommentContent string `json:"comment_content"` //评价的内容
	IsLike bool `json:"is_like"` //是否点赞
	Commenttime string `json:"display_ctime"` //评论时间
	IsVisible bool `json:"is_visible"` //评论是否可见  true可见
}


//视频评价
type ReuqestVideoCommentParmas struct {
	ConmmentId string `json:"conmment_id"` //视频评论id
	VideoId string `json:"video_id"` //视频id
	UserName string `json:"user_name"` //用户
}






func(userInfo *UserInfo) CheckValidator() error{
	/**
	if _, err := govalidator.ValidateStruct(userInfo); err != nil {
		return err
	} else {
		return nil
	}

	**/
	return nil

}

func(reuqestUserParmas *ReuqestUserParmas) CheckValidator() error {
/**
	if _, err := govalidator.ValidateStruct(reuqestUserParmas); err != nil {
		return err
	} else {
		return nil
	}
	**/
	return nil
}


func(videoInfo *VideoInfo) CheckValidator() error {
	/**
		if _, err := govalidator.ValidateStruct(reuqestUserParmas); err != nil {
			return err
		} else {
			return nil
		}
		**/
	return nil
}


func(reuqestVideoParmas *ReuqestVideoParmas) CheckValidator() error {
	/**
		if _, err := govalidator.ValidateStruct(reuqestUserParmas); err != nil {
			return err
		} else {
			return nil
		}
		**/
	return nil
}



func(mVideoComment *VideoComment) CheckValidator() error {
	/**
		if _, err := govalidator.ValidateStruct(reuqestUserParmas); err != nil {
			return err
		} else {
			return nil
		}
		**/
	return nil
}


func(mReuqestVideoCommentParmas *ReuqestVideoCommentParmas) CheckValidator() error {
	/**
		if _, err := govalidator.ValidateStruct(reuqestUserParmas); err != nil {
			return err
		} else {
			return nil
		}
		**/
	return nil
}