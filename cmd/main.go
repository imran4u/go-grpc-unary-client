package main

import (
	"errors"

	"google.golang.org/grpc"
)



func main() {
 var opts []grpc.DialOption

 opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())

 conn, err := grpc.NewClient("localhost:5050", opts...)
 if err != nil {	
	panic(err)	
 }
 defer conn.Close()

 helloAdapter, err := mygrpc.NewHelloAdapter(conn)
 if err != nil {	
	panic(err)	
 }	

 name := "imran"
 println(helloAdapter.SayHello(name))	
}