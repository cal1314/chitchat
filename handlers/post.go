package handlers

import (
   "fmt"
   "chitchat/models"
   "net/http"
   "github.com/nicksnyder/go-i18n/v2/i18n"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(writer http.ResponseWriter,request *http.Request){
     sess,err := session(writer,request)
     if err != nil{
        http.Redirect(writer,request,"/login",302)
     }else{
        err = request.ParseForm()
	if err != nil{
	   fmt.Println("cannot parse form")
	}
	user,err := sess.User()
	if err != nil{
	   fmt.Println("cannot get user from session")
	}
	body := request.PostFormValue("body")
	uuid := request.PostFormValue("uuid")
	thread,err := models.ThreadByUUID(uuid)
	fmt.Println(err)
	if err != nil{
	   msg := localizer.MustLocalize(&i18n.LocalizeConfig{
	          MessageID : "thread_not_found",
	   })
	   errorMessage(writer,request,msg)
	}
	if _,err := user.CreatePost(thread,body);err != nil{
	  fmt.Println("cannot create post")
	}
	url := fmt.Sprint("/thread/read?id=",uuid)
	http.Redirect(writer,request,url,302)
     }
}
