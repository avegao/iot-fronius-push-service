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
	requestInverters := make([]*pb.InverterPowerflow, 0)

	for _, inverter := range inverters {
		requestInverters = append(requestInverters, inverter.ToGrpcRequest())
	}

	ohmpilots := powerflow.Body.SmartLoads.Ohmpilots
	requestOhmpilots := make([]*pb.OhmpilotPowerflow, 0)

	for _, ohmpilot := range ohmpilots {
		requestOhmpilots = append(requestOhmpilots, ohmpilot.ToGrpcRequest())
	}

	return
}
