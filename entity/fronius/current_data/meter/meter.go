package froniusCurrentDataMeter

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
)

type CurrentDataMeter struct {
	Body CurrentDataMeterBody   `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}

type CurrentDataMeterBody map[string]CurrentDataMeterBodyElement

type CurrentDataMeterBodyElement struct {
	CurrentAcPhase1 float64 `json:"Current_AC_Phase_1"`
	CurrentAcPhase2 float64 `json:"Current_AC_Phase_2"`
	CurrentAcPhase3 float64 `json:"Current_AC_Phase_3"`
	CurrentAcSum    float64 `json:"Current_AC_Sum"`
	Details struct {
		Manufacturer string `json:"Manufacturer"`
		Model        string `json:"Model"`
		Serial       string `json:"Serial"`
	} `json:"Details"`
	Enable                            uint8   `json:"Enable"`
	EnergyReactiveVArAcPhase1Consumed uint32  `json:"EnergyReactive_VArAC_Phase_1_Consumed"`
	EnergyReactiveVArAcPhase1Produced uint32  `json:"EnergyReactive_VArAC_Phase_1_Produced"`
	EnergyReactiveVArAcPhase2Consumed uint32  `json:"EnergyReactive_VArAC_Phase_2_Consumed"`
	EnergyReactiveVArAcPhase2Produced uint32  `json:"EnergyReactive_VArAC_Phase_2_Produced"`
	EnergyReactiveVArAcPhase3Consumed uint32  `json:"EnergyReactive_VArAC_Phase_3_Consumed"`
	EnergyReactiveVArAcPhase3Produced uint32  `json:"EnergyReactive_VArAC_Phase_3_Produced"`
	EnergyReactiveVArAcSumConsumed    uint32  `json:"EnergyReactive_VArAC_Sum_Consumed"`
	EnergyReactiveVArAcSumProduced    uint32  `json:"EnergyReactive_VArAC_Sum_Produced"`
	EnergyRealWAcMinusAbsolute        uint32  `json:"EnergyReal_WAC_Minus_Absolute"`
	EnergyRealWAcPhase1Consumed       uint32  `json:"EnergyReal_WAC_Phase_1_Consumed"`
	EnergyRealWAcPhase1Produced       uint32  `json:"EnergyReal_WAC_Phase_1_Produced"`
	EnergyRealWAcPhase2Consumed       uint32  `json:"EnergyReal_WAC_Phase_2_Consumed"`
	EnergyRealWAcPhase2Produced       uint32  `json:"EnergyReal_WAC_Phase_2_Produced"`
	EnergyRealWAcPhase3Consumed       uint32  `json:"EnergyReal_WAC_Phase_3_Consumed"`
	EnergyRealWAcPhase3Produced       uint32  `json:"EnergyReal_WAC_Phase_3_Produced"`
	EnergyRealWAcPlusAbsolute         uint32  `json:"EnergyReal_WAC_Plus_Absolute"`
	EnergyRealWAcSumConsumed          uint32  `json:"EnergyReal_WAC_Sum_Consumed"`
	EnergyRealWAcSumProduced          uint32  `json:"EnergyReal_WAC_Sum_Produced"`
	FrequencyPhaseAverage             float64 `json:"Frequency_Phase_Average"`
	MeterLocationCurrent              uint8   `json:"Meter_Location_Current"`
	PowerApparentSPhase1              float64 `json:"PowerApparent_S_Phase_1"`
	PowerApparentSPhase2              float64 `json:"PowerApparent_S_Phase_2"`
	PowerApparentSPhase3              float64 `json:"PowerApparent_S_Phase_3"`
	PowerApparentSSum                 float64 `json:"PowerApparent_S_Sum"`
	PowerFactorPhase1                 float64 `json:"PowerFactor_Phase_1"`
	PowerFactorPhase2                 float64 `json:"PowerFactor_Phase_2"`
	PowerFactorPhase3                 float64 `json:"PowerFactor_Phase_3"`
	PowerFactorSum                    float64 `json:"PowerFactor_Sum"`
	PowerReactiveQPhase1              float64 `json:"PowerReactive_Q_Phase_1"`
	PowerReactiveQPhase2              float64 `json:"PowerReactive_Q_Phase_2"`
	PowerReactiveQPhase3              float64 `json:"PowerReactive_Q_Phase_3"`
	PowerReactiveQSum                 float64 `json:"PowerReactive_Q_Sum"`
	PowerRealPPhase1                  float64 `json:"PowerReal_P_Phase_1"`
	PowerRealPPhase2                  float64 `json:"PowerReal_P_Phase_2"`
	PowerRealPPhase3                  float64 `json:"PowerReal_P_Phase_3"`
	PowerRealPSum                     float64 `json:"PowerReal_P_Sum"`
	TimeStamp                         uint32  `json:"TimeStamp"`
	Visible                           uint8   `json:"Visible"`
	VoltageAcPhase1                   float64 `json:"Voltage_AC_Phase_1"`
	VoltageAcPhase2                   float64 `json:"Voltage_AC_Phase_2"`
	VoltageAcPhase3                   float64 `json:"Voltage_AC_Phase_3"`
}

func (body CurrentDataMeterBody) ToGrpcRequest() *pb.CurrenDataMeterRequest {
	elements := make([]*pb.CurrentDataMeter, 0)

	for _, element := range body {
		elements = append(elements, element.ToGrpcRequest())
	}

	return &pb.CurrenDataMeterRequest{
		Elements: elements,
	}
}

func (element CurrentDataMeterBodyElement) ToGrpcRequest() *pb.CurrentDataMeter {
	return &pb.CurrentDataMeter{
		CurrentAcPhase1: element.CurrentAcPhase1,
		CurrentAcPhase2: element.CurrentAcPhase2,
		CurrentAcPhase3: element.CurrentAcPhase3,
		CurrentAcSum:    element.CurrentAcSum,
		MeterDetails: &pb.MeterDetails{
			Manufacturer: element.Details.Manufacturer,
			Model:        element.Details.Model,
			Serial:       element.Details.Serial,
		},
		Enable:                            element.Enable != 0,
		EnergyReactiveVArAcPhase1Consumed: element.EnergyReactiveVArAcPhase1Consumed,
		EnergyReactiveVArAcPhase1Produced: element.EnergyReactiveVArAcPhase1Produced,
		EnergyReactiveVArAcPhase2Consumed: element.EnergyReactiveVArAcPhase2Consumed,
		EnergyReactiveVArAcPhase2Produced: element.EnergyReactiveVArAcPhase2Produced,
		EnergyReactiveVArAcPhase3Consumed: element.EnergyReactiveVArAcPhase3Consumed,
		EnergyReactiveVArAcPhase3Produced: element.EnergyReactiveVArAcPhase3Produced,
		EnergyReactiveVArAcSumConsumed:    element.EnergyReactiveVArAcSumConsumed,
		EnergyReactiveVArAcSumProduced:    element.EnergyReactiveVArAcSumProduced,
		EnergyRealWAcMinusAbsolute:        element.EnergyRealWAcMinusAbsolute,
		EnergyRealWAcPhase1Consumed:       element.EnergyRealWAcPhase1Consumed,
		EnergyRealWAcPhase1Produced:       element.EnergyRealWAcPhase1Produced,
		EnergyRealWAcPhase2Consumed:       element.EnergyRealWAcPhase2Consumed,
		EnergyRealWAcPhase2Produced:       element.EnergyRealWAcPhase2Produced,
		EnergyRealWAcPhase3Consumed:       element.EnergyRealWAcPhase3Consumed,
		EnergyRealWAcPhase3Produced:       element.EnergyRealWAcPhase3Produced,
		EnergyRealWAcSumConsumed:          element.EnergyRealWAcSumConsumed,
		EnergyRealWAcSumProduced:          element.EnergyRealWAcSumProduced,
		FrequencyPhaseAverage:             element.FrequencyPhaseAverage,
		MeterLocationCurrent:              uint32(element.MeterLocationCurrent),
		PowerApparentSPhase1:              element.PowerApparentSPhase1,
		PowerApparentSPhase2:              element.PowerApparentSPhase2,
		PowerApparentSPhase3:              element.PowerApparentSPhase3,
		PowerApparentSSum:                 element.PowerApparentSSum,
		PowerFactorPhase1:                 element.PowerFactorPhase1,
		PowerFactorPhase2:                 element.PowerFactorPhase2,
		PowerFactorPhase3:                 element.PowerFactorPhase3,
		PowerFactorSum:                    element.PowerFactorSum,
		PowerReactiveQPhase1:              element.PowerReactiveQPhase1,
		PowerReactiveQPhase2:              element.PowerReactiveQPhase2,
		PowerReactiveQPhase3:              element.PowerReactiveQPhase3,
		PowerReactiveQSum:                 element.PowerReactiveQSum,
		PowerRealPPhase1:                  element.PowerRealPPhase1,
		PowerRealPPhase2:                  element.PowerRealPPhase2,
		PowerRealPPhase3:                  element.PowerRealPPhase3,
		PowerRealPSum:                     element.PowerRealPSum,
		Timestamp:                         element.TimeStamp,
		Visible:                           element.Visible != 0,
		VoltageAcPhase1:                   element.VoltageAcPhase1,
		VoltageAcPhase2:                   element.VoltageAcPhase2,
		VoltageAcPhase3:                   element.VoltageAcPhase3,
	}
}
