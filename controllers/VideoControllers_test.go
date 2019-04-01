package controllers

import (
	"testing"
	"github.com/KXX747/vodieserver/models"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/KXX747/vodieserver/middleware"
	"encoding/json"
)


var uuid ="7e4c3c38-4b52-4d33-9fff-d661b3468b1c"


func TestVideoWorkFlow(t *testing.T) {

	t.Run("NewVideo",TestCreateNewVideo)
	t.Run("FindVideo",TestFindVideo)

}

//添加视频
func TestCreateNewVideo(t *testing.T) {
	createNewVideo("kuangxinxi")
}

//查询视频
func TestFindVideo(t *testing.T) {
	findVideo(uuid,"kuangxinxi")
}

//添加视频评价
func TestCreateNewVideoComment(t *testing.T) {
	createNewVideoComment("kuangxinxin","7e4c3c38-4b52-4d33-9fff-d661b3468b1c")
}

//查询视频评价
func TestFindVideoComment(t *testing.T) {
	findVideoComment("kuangxinxin","7e4c3c38-4b52-4d33-9fff-d661b3468b1c","b98e92c3-408d-4629-816c-5af36c076e3a")
}


//
func createNewVideoComment(username string,videoid string)  {

	mVideoComment:=models.VideoComment{}
	mVideoComment.Id = 0
	mVideoComment.ConmmentId = middleware.CreateUuid().String()
	mVideoComment.VideoId =videoid
	mVideoComment.UserName = username
	mVideoComment.CommentContent ="昕昕宝宝是哥狗屎，都是"
	mVideoComment.Commenttime = middleware.DateFormatTimeToStringTime()
	mVideoComment.IsVisible=true
	mVideoComment.IsLike = true
	b,_:=json.Marshal(mVideoComment)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/video/CreateNewVideoComment", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("createNewVideo ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION,mVideoComment.UserName)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))
}

//
func findVideoComment(username ,videoId ,ConmmentId string)  {

	mReuqestVideoCommentParmas:=models.ReuqestVideoCommentParmas{}
	mReuqestVideoCommentParmas.VideoId =videoId
	mReuqestVideoCommentParmas.UserName =username
	mReuqestVideoCommentParmas.ConmmentId =ConmmentId
	b ,_:=json.Marshal(mReuqestVideoCommentParmas)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/video/FindVideoComment", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION,mReuqestVideoCommentParmas.UserName)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))

}




//
func createNewVideo(userName string)  {

	mVideoInfo:=models.VideoInfo{}
	mVideoInfo.Id = 0
	mVideoInfo.VideoId = middleware.CreateUuid().String()
	uuid = mVideoInfo.VideoId
	mVideoInfo.UserName =userName
	mVideoInfo.VideoName = "上海地铁开通地铁开通"
	mVideoInfo.DisplayCtime = middleware.DateFormatTimeToStringTime()
	mVideoInfo.IsVip="0"
	b,_:=json.Marshal(mVideoInfo)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/video/createNewVideo", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("createNewVideo ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION,mVideoInfo.UserName)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))
}

//
func findVideo(videoId string,userName string)  {

	videoId ="7e4c3c38-4b52-4d33-9fff-d661b3468b1c"

	mReuqestVideoParmas:=models.ReuqestVideoParmas{}
	mReuqestVideoParmas.VideoId =videoId
	b ,_:=json.Marshal(mReuqestVideoParmas)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/video/findVideoByVideoId", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println("TestFindUser ",err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HEADER_FIELD_SESSION,userName)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body " ,err.Error())
	}
	fmt.Println(string(body))

}