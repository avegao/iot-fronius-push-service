package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
	"encoding/json"
)

type inverter struct {
	BatteryMode fronius.BatteryMode `json:"battery_mode"`

	// DeviceType Device type of inverter
	DeviceType uint16 `json:"DT"`

	// EnergyDay Energy in Wh this day, null if no inverter is connected
	EnergyDay float64`json:"E_Day"`

	// EnergyDay Energy in Wh ever since, null if no inverter is connected
	EnergyTotal float64 `json:"E_Total"`

	// EnergyDay Energy in Wh this year, null if no inverter is connected
	EnergyYear float64 `json:"E_Year"`

	// CurrentPower current power in Watt, null if not running
	CurrentPower float64 `json:"P"`

	// Soc Current state of charge in % ( 0 - 100% )
	Soc uint8 `json:"SOC"`
}

func (i *inverter) UnmarshalJSON(data []byte) error {
	type defaultInverter inverter

	parsingInverter := &defaultInverter{
		BatteryMode: fronius.BatteryModeDisabled,
	}

	err := json.Unmarshal(data, parsingInverter)
	if err != nil {
		return err
	}

	*i = inverter(*parsingInverter)

	return nil
}

func (i inverter) ToGrpcRequest() *pb.InverterPowerflow {
	return &pb.InverterPowerflow{
		BatteryMode:  i.BatteryMode.String(),
		DeviceType:   uint32(i.DeviceType),
		EnergyDay:    i.EnergyDay,
		EnergyYear:   i.EnergyYear,
		EnergyTotal:  i.EnergyTotal,
		CurrentPower: i.CurrentPower,
		Soc:          uint32(i.Soc),
	}
}
