package main

import (
	"log"
	"net/rpc"
	"rpc/1/Models"
)

func client(arg Models.Message,serviceMethod string)(err error,reply Models.Message){
	client, err:=rpc.DialHTTP("tcp",":8081")
	if err!=nil {
		log.Fatal("dialing ",err)
		return err,Models.Message{}
	}
	defer client.Close()
	args:= Models.Message{K: arg.K,V:arg.V}
	err = client.Call(serviceMethod,args,&reply)
	if err!=nil {
		log.Fatal("error:",err)
		return err,Models.Message{}
	}
	return
}
