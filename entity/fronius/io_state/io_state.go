package froniusIoState

import (
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
	"strings"
	"strconv"
)

type IoState map[string]IoStatePin

type IoStatePin struct {
	Function  string `json:"function"`
	Type      string `json:"type"`
	Direction string `json:"direction"`
	Set       bool   `json:"set"`
}

func (state IoState) ToGrpcRequest() (pb.CurrentIoState) {
	pins := make([]*pb.IoStatePin, 0)

	for pinNumber, pin := range state {
		grpcPin := &pb.IoStatePin{
			PinNumber: pinNumberStringToInt(pinNumber),
			Function:  pin.Function,
			Type:      pin.Type,
			Direction: pin.Direction,
			Set:       pin.Set,
		}

		pins = append(pins, grpcPin)
	}

	return pb.CurrentIoState{
		Pins: pins,
	}
}

func pinNumberStringToInt(pinNumberString string) (int32) {
	splited := strings.Split(pinNumberString, " ")
	numberInt64, _ := strconv.ParseInt(splited[1], 10, 0)

	return int32(numberInt64)
}
