package main

import (
	"context"
	"fmt"
	"log"

	netxdcustomer "github.com/SWETHA0705/netxd_customer/customer"
	netxdcustomerconstants "github.com/SWETHA0705/netxd_server/netxd_customer_constants"
	"google.golang.org/grpc"
)


func main(){
	conn, err := grpc.Dial(netxdcustomerconstants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := netxdcustomer.NewCustomerServiceClient(conn)
	response, err := client.Transaction(context.Background(), &netxdcustomer.Transaction{
		FromAccount: 12345,
		ToAccount:   6789,
		Amount:      1000,
	})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	fmt.Printf("Response: %s\n",response.Message)
}
