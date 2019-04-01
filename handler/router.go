package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/KXX747/vodieserver/controllers"
	"net/http"
	"github.com/KXX747/vodieserver/middleware"
	"github.com/KXX747/vodieserver/defs"
)


type MiddleHandler struct {
	r *httprouter.Router
	
}

func NewMiddleHandler(r *httprouter.Router) http.Handler {

	m:=MiddleHandler{}
	m.r = r
	return m

}

//利用接口拦截参数
func(m MiddleHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	//检查身份的合法性
	validateApi:=middleware.ValidDateApi(r)
	if !validateApi {
		defs.SendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}
	//
	m.r.ServeHTTP(w,r)
}

//路由
func RegisterHandlers() *httprouter.Router  {
	router :=httprouter.New()
	router = addUserRouter(router)
	//
	router = addVideoRouter(router)

	return router
}


//
func addUserRouter(router *httprouter.Router) *httprouter.Router  {
	//用户信息相关
	router.POST("/user/create",controllers.CreateUser)
	router.POST("/user/find",controllers.FindUser)
	router.POST("/user/update",controllers.UpdateUser)
	router.POST("/user/delete",controllers.DeleteUser)
	return router
}


//视频相关
func addVideoRouter(router *httprouter.Router) *httprouter.Router  {

	//添加查询视频
	router.POST("/video/createNewVideo",controllers.CreateNewVideo)
	router.POST("/video/findVideoByVideoId",controllers.FindVideo)

	//视频评价添加查询
	router.POST("/video/CreateNewVideoComment",controllers.CreateNewVideoComment)
	router.POST("/video/FindVideoComment",controllers.FindVideoComment)


	return router
}

