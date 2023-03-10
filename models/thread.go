package models

import "time"

type Thread struct{
    Id  int
    Uuid  string
    Topic  string
    UserId  int
    CreatedAt  time.Time
}

func (thread *Thread)CreatedAtDate() string{
     return thread.CreatedAt.Format(" Jan 2, 2006 at 3:04pm")
}

func (thread *Thread)NumReplies()(count int){
     rows,err := Db.Query("select count(*) from posts where thread_id = ?",thread.Id)
     if err != nil{
        return
     }
     for rows.Next(){
         if err = rows.Scan(&count);err != nil{
	    return
	 }
     }
     rows.Close()
     return
}

func (thread *Thread) Posts()(posts []Post,err error){
     rows,err := Db.Query("select id,uuid,body,user_id,thread_id,created_at  from posts where thread_id = ?",thread.Id)
     if err != nil{
        return
     }
     for rows.Next(){
	 post := Post{}
	 if err = rows.Scan(&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt);err != nil{
	    return
	 }
	 posts= append(posts,post)
     }
     rows.Close()
     return
}

func Threads()(threads []Thread,err error){
     rows,err := Db.Query("select id,uuid,topic,user_id,created_at from threads order by created_at desc")
     if err != nil{
         return
     }

     for rows.Next(){
         conv := Thread{}

	 if err = rows.Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt);err != nil{
	    return
	 }
	 threads = append(threads,conv)
     }
     rows.Close()
     return
}

func ThreadByUUID(uuid string)(conv Thread,err error){
     conv = Thread{}
     err = Db.QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid = ?",uuid).
           Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.UserId,&conv.CreatedAt)
     return
}

func (thread *Thread)User()(user User){
     user = User{}
     Db.QueryRow("select id,uuid,name,email,created_at from users where id = ?",thread.UserId).
        Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)

     return
}
