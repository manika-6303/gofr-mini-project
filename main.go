package main

import (
    // "fmt"
    "strconv"
    "gofr.dev/pkg/gofr"
)

type Customer struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    // Address string `json:"address"`

}
func post(ctx *gofr.Context) (interface{}, error){
    name := ctx.PathParam("name")

        // Inserting a customer row in database using SQL
        _, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

        return nil, err
}
func deleteCustomerById(ctx *gofr.Context) (interface{}, error){
    id := ctx.PathParam("id")
    i ,err:=strconv.Atoi(id)
    if err != nil {
            return nil, err
        }
        // Inserting a customer row in database using SQL
        _, err1 := ctx.DB().ExecContext(ctx, "DELETE FROM customers WHERE Id=(?)",i)

        return nil, err1
}
func getCustomerBYId(ctx *gofr.Context) (interface{}, error){
    id := ctx.PathParam("id")
    i,err :=strconv.Atoi(id)
    var customers[] Customer
    var ans Customer
        // Getting the customer from the database using SQL
        rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
        if err != nil {
            return nil, err
        }

        for rows.Next() {
            var customer Customer
            if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
                return nil, err
            }
            customers=append(customers,customer)
        }
        for _,a := range customers{
            if a.ID==i{
                ans=a
            }
        }
        return ans,nil

}
func get(ctx *gofr.Context) (interface{}, error){
    var customers []Customer
        // Getting the customer from the database using SQL
        rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
        if err != nil {
            return nil, err
        }

        for rows.Next() {
            var customer Customer
            if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
                return nil, err
            }

            customers = append(customers, customer)
        }

        // return the customer
        return customers, nil
}
func main() {
    // initialise gofr object
    app := gofr.New()

    app.POST("/customer/delete/{id}",deleteCustomerById)
    app.POST("/customer/{name}",post)
    // app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
    //     name := ctx.PathParam("name")

    //     // Inserting a customer row in database using SQL
    //     _, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

    //     return nil, err
    // })
    app.GET("/customer",get)

    app.GET("/customer/{id}",getCustomerBYId)
    
    // app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
    //     var customers []Customer

    //     // Getting the customer from the database using SQL
    //     rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
    //     if err != nil {
    //         return nil, err
    //     }

    //     for rows.Next() {
    //         var customer Customer
    //         if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
    //             return nil, err
    //         }

    //         customers = append(customers, customer)
    //     }

    //     // return the customer
    //     return customers, nil
    // })

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
    app.Start()
}
