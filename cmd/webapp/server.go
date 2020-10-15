package main

import (
	product "github.com/ckalagara/go-grpc-server/cmd/webapp/service/product/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.WithFields(
			log.Fields{
				"code":    "SERVICE_INIT_ERR",
				"message": "failed to listen on " + port,
			}).Panic(err.Error())
	}

	log.Infof("Started listening on port %v", port)
	server := grpc.NewServer()
	productServer := product.ServiceImpl{}
	product.RegisterProductServiceServer(server, &productServer)
	err = server.Serve(listener)

	if err != nil {
		log.WithFields(
			log.Fields{
				"code":    "SERVICE_INIT_ERR",
				"message": "failed to serve over grpc",
			}).Panic(err.Error())
	}

}
