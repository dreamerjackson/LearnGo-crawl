package rpcsupport

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string,serveice interface{}) error{
	rpc.Register(serveice)

	listten,err:= net.Listen("tcp",host)

	if err!=nil{
		return err
	}

	for{
		conn,err:= listten.Accept()
		if err!=nil{
			log.Printf("accept error : %v",err)
			continue
		}

		go jsonrpc.ServeConn(conn)

	}


}

func NewClient(host string) (*rpc.Client,error){
	conn,err:= net.Dial("tcp",host)

	if err!=nil{
		return nil,err
	}

	return jsonrpc.NewClient(conn),nil

}