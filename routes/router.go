package routes

import "github.com/gorilla/mux"

func NewRouter( ) *mux.Router {
     router := mux.NewRouter().StrictSlash(true)

     //遍历 web.go 中定义的所有 webRoutes
     for _,route := range webRoutes{
         // 将每个 web 路由应用到路由器
	 router.Methods(route.Method).
	     Path(route.Pattern).
	     Name(route.Name).
	     Handler(route.HandlerFunc)
     }
     return router
}
