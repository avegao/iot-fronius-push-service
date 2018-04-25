package froniusCurrentDataInverter

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
)

type CurrentDataInverter struct {
	Body struct {
		// DayEnergy Energy generated on current day
		DayEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"DAY_ENERGY"`
		// Pac AC power
		Pac struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"PAC"`
		// TotalEnergy Energy generated overall
		TotalEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"TOTAL_ENERGY"`
		// YearEnergy Energy generated in current year
		YearEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"YEAR_ENERGY"`
	} `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}

func (currentData CurrentDataInverter) ToGrpcRequest() *pb.CurrenDataInverterRequest {
	dayEnergyArray := make([]int32, 0)

	for _, element := range currentData.Body.DayEnergy.Values {
		dayEnergyArray = append(dayEnergyArray, int32(element))
	}

	pacArray := make([]int32, 0)

	for _, element := range currentData.Body.Pac.Values {
		pacArray = append(pacArray, int32(element))
	}

	totalEnergyArray := make([]int32, 0)

	for _, element := range currentData.Body.TotalEnergy.Values {
		totalEnergyArray = append(totalEnergyArray, int32(element))
	}

	yearEnergyArray := make([]int32, 0)

	for _, element := range currentData.Body.YearEnergy.Values {
		yearEnergyArray = append(yearEnergyArray, int32(element))
	}

	timestamp := currentData.Head.Timestamp.Unix()

	return &pb.CurrenDataInverterRequest{
		DayEnergy:   dayEnergyArray,
		Pac:         pacArray,
		TotalEnergy: totalEnergyArray,
		YearEnergy:  yearEnergyArray,
		Timestamp:   timestamp,
	}
}
