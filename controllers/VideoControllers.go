package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/KXX747/vodieserver/models"
	"encoding/json"
	"io"
	"github.com/KXX747/vodieserver/defs"
)


/**
添加视频
 */
func CreateNewVideo(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {

	b,_:=ioutil.ReadAll(r.Body)
	var videoInfo models.VideoInfo
	json.Unmarshal(b,&videoInfo)

	err :=videoInfo.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	err=videoInfo.CreateVideoModles()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	responseData,_:=json.Marshal(videoInfo)
	defs.SendNormalResponse(w,string(responseData),200)

}

/**
查询视频
*/
func FindVideo(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {

	b,_:=ioutil.ReadAll(r.Body)
	var reuqestVideoParmas models.ReuqestVideoParmas
	json.Unmarshal(b,&reuqestVideoParmas)

	err :=reuqestVideoParmas.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	videoInfo,err:=reuqestVideoParmas.FindVideoModles()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	responseData,_:=json.Marshal(videoInfo)
	defs.SendNormalResponse(w,string(responseData),200)

}

//创建视频评论
func CreateNewVideoComment(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {
	b,_:=ioutil.ReadAll(r.Body)
	var mVideoComment models.VideoComment
	json.Unmarshal(b,&mVideoComment)

	err :=mVideoComment.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}


	err=mVideoComment.CreateVideoCommentModles()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	responseData,_:=json.Marshal(mVideoComment)
	defs.SendNormalResponse(w,string(responseData),200)


}


//查询视频评论
func FindVideoComment(w http.ResponseWriter,r *http.Request ,p httprouter.Params)  {

	b,_:=ioutil.ReadAll(r.Body)
	var mReuqestVideoCommentParmas models.ReuqestVideoCommentParmas
	json.Unmarshal(b,&mReuqestVideoCommentParmas)

	err :=mReuqestVideoCommentParmas.CheckValidator()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}

	mVideoComment,err:=mReuqestVideoCommentParmas.FindVideoCommentModles()
	if err!=nil {
		io.WriteString(w,err.Error())
		return
	}
	responseData,_:=json.Marshal(mVideoComment)
	defs.SendNormalResponse(w,string(responseData),200)

}