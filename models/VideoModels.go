package models

import (
	"github.com/KXX747/vodieserver/middleware"
)

/**
视频数据
 */


 //添加视频
func(u *VideoInfo) CreateVideoModles() error {

	uuid :=middleware.CreateUuid()
	createTime :=middleware.DateFormatTimeToStringTime()

	u.VideoId = uuid.String()
	u.DisplayCtime = createTime

	_, error:=engineConn.InsertOne(u)
	if error!=nil {
		return error
	}

	return nil
}

//查询视频
func(u *ReuqestVideoParmas) FindVideoModles() (VideoInfo,error) {

	var videoInfo VideoInfo
	_,err:=engineConn.Where("video_id=?",u.VideoId).Get(&videoInfo)
	if err!=nil {
		return videoInfo,err
	}

	err =videoInfo.CheckValidator()
	if err!=nil {
		return videoInfo,err
	}


	return videoInfo,nil
}



//
//添加视频评价
func(u *VideoComment) CreateVideoCommentModles() error {

	uuid :=middleware.CreateUuid()
	createTime :=middleware.DateFormatTimeToStringTime()

	u.ConmmentId= uuid.String()
	u.Commenttime = createTime

	_, error:=engineConn.InsertOne(u)
	if error!=nil {
		return error
	}

	return nil
}






//查询视频评价
func(u *ReuqestVideoCommentParmas) FindVideoCommentModles() (VideoComment,error) {

	var mVideoComment VideoComment
	_,err:=engineConn.Where("video_id=? or user_name=? or conmment_id=?",u.VideoId,u.UserName,u.ConmmentId).Get(&mVideoComment)
	if err!=nil {
		return mVideoComment,err
	}

	return mVideoComment,nil
}