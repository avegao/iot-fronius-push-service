package service

import (
	"github.com/avegao/gocondi"
	"google.golang.org/grpc"
	"github.com/avegao/iot-fronius-push-service/resource/grpc"
)

func CreateConnection() (connection *grpc.ClientConn, err error) {
	var grpcOptions []grpc.DialOption

	grpcOptions = append(grpcOptions, grpc.WithInsecure())
	address := gocondi.GetContainer().GetStringParameter("fronius_service_address")
	connection, err = grpc.Dial(address, grpcOptions...)

	if nil != err {
		gocondi.GetContainer().GetLogger().WithError(err).Fatalf("Fail to connect with %s", address)

		return
	}

	gocondi.GetContainer().GetLogger().Debugf("gRPC connection status with %v = %s", address, connection.GetState().String())

	return
}

func CreateClient(connection *grpc.ClientConn) iot_fronius.FroniusClient {
	return iot_fronius.NewFroniusClient(connection)
}
