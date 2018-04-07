package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
)

type CurrentPowerflow struct {
	Body struct {
		Inverters map[int]inverter `json:"Inverters"`
		Site      site             `json:"Site"`
		SmartLoads struct {
			Ohmpilots map[int]ohmpilot `json:"Ohmpilots,omitempty"`
		} `json:"Smartloads,omitempty"`
		Version string `json:"Version"`
	} `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}

func (powerflow CurrentPowerflow) ToGrpcRequest() (request pb.Powerflow) {
	request.Site = powerflow.Body.Site.ToGrpcRequest()

	inverters := powerflow.Body.Inverters
	request.Inverter = make([]*pb.InverterPowerflow, 0)

	for _, inverter := range inverters {
		request.Inverter = append(request.Inverter, inverter.ToGrpcRequest())
	}

	ohmpilots := powerflow.Body.SmartLoads.Ohmpilots
	request.Ohmpilot = make([]*pb.OhmpilotPowerflow, 0)

	for _, ohmpilot := range ohmpilots {
		request.Ohmpilot = append(request.Ohmpilot, ohmpilot.ToGrpcRequest())
	}

	return
}
