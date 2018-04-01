package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
)

type ohmpilot struct {
	// PowerAcTotal Current power consumption in Watt
	PowerAcTotal float64 `json:"P_AC_Total"`

	State fronius.OhmpilotState `json:"State"`

	// Temperature Temperature of storage / tank in degree Celsius
	Temperature float64 `json:"Temperature"`
}

func (ohmpilot ohmpilot) ToGrpcRequest() *pb.OhmpilotPowerflow {
	return &pb.OhmpilotPowerflow{
		PowerAcTotal: ohmpilot.PowerAcTotal,
		State:        ohmpilot.State.String(),
		Temperature:  ohmpilot.Temperature,
	}
}
