package main

import (
    // "fmt"
    "strconv"
    "gofr.dev/pkg/gofr"
    "strings"
)

type album struct {
    ID   int    `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price int `json:"price"`
}
func adduser(ctx *gofr.Context) (interface{},error){
    data := ctx.PathParam("data")
    data1:=strings.Split(data,",")
    i:=data1[0]
    t:=data1[1]
    a:=data1[2]
    p:=data1[3]
    i1,err1:=strconv.Atoi(i)
    p1,err2:=strconv.Atoi(p)
    
        _, err := ctx.DB().ExecContext(ctx, "INSERT IGNORE INTO album (id,title,artist,price) VALUES (?,?,?,?)", i1,t,a,p1)
       
        if err1 != nil {
            return nil, err 
        }
        if err2 != nil {
            return nil, err 
        }
        
        return nil,err
}
func deleteAlbumById(ctx *gofr.Context) (interface{}, error){
    id := ctx.PathParam("id")
    i ,err:=strconv.Atoi(id)
    if err != nil {
            return nil, err 
        }
    
        _, err1 := ctx.DB().ExecContext(ctx, "DELETE FROM album WHERE Id=(?)",i)

        return nil, err1
}
func getAlbumById(ctx *gofr.Context) (interface{}, error){
    id := ctx.PathParam("id")
    i,err :=strconv.Atoi(id)
    if err != nil {
            return nil, err 
    }
    var albums[] album
    var ans album
    
        rows, err1 := ctx.DB().QueryContext(ctx, "SELECT * FROM album")
        if err1 != nil {
            return nil, err1
        }

        for rows.Next() {
            var alb album
            if err1 := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
                return nil, err1
            }
            albums=append(albums,alb)
        }
        for _,a := range albums{
            if a.ID==i{
                ans=a
            }
        }
        return ans,nil

}
func getall(ctx *gofr.Context) (interface{}, error){
    var albums []album
    
        rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM album")
        if err != nil {
            return nil, err
        }

        for rows.Next() {
            var alb album
            if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
                return nil, err
            }

            albums = append(albums, alb)
        }
        return albums, nil
}
func main() {
    // initialise gofr object
    app := gofr.New()

    
    app.GET("/album/get/{id}",getAlbumById)

    app.POST("/album/delete/{id}",deleteAlbumById)
    app.POST("/album/{data}",adduser)
    // app.POST("/customer/update/{id}",updateid)
    app.GET("/album",getall)
    
    
    app.Start()
}
