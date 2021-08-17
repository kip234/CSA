package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/1/Models"
)

func server(){
	s:=Models.NewKV()
	rpc.Register(s)
	rpc.HandleHTTP()
	l,e:=net.Listen("tcp",":8081")
	if e!=nil {
		log.Fatal("listen error:",e)
	}
	http.Serve(l,nil)
}