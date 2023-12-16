package main

import "gofr.dev/pkg/gofr"

type Customer struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
func post(ctx *gofr.Context) (interface{}, error){
    name := ctx.PathParam("name")

        // Inserting a customer row in database using SQL
        _, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

        return nil, err
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

    app.POST("/customer/{name}",post)
    app.Start()
}
