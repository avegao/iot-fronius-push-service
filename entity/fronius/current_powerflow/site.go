package froniusCurrentPowerflow

import (
	"github.com/avegao/iot-fronius-push-service/entity/fronius"
	pb "github.com/avegao/iot-fronius-push-service/resource/grpc"
)

type site struct {
	Id int

	Mode fronius.SiteMode `json:"Mode"`

	// BatteryStandby True when battery is in standby
	BatteryStandby bool `json:"BatteryStandby"`

	// BackupMode Field is available if configured (false) or active (true)
	// if not available, mandatory config is not set.
	BackupMode bool `json:"BackupMode"`

	// PowerFromGrid This value is null if no meter is enabled (+ from grid, - to grid)
	PowerFromGrid float64 `json:"P_Grid"`

	// PowerLoad This value is null if no meter is enabled (+ generator, - consumer)
	PowerLoad float64 `json:"P_Load"`

	// PowerAkku This value is null if no battery is active (+ charge, - discharge)
	PowerAkku float64 `json:"P_Akku"`

	// PowerFromPV This value is null if inverter is not running (+ production (default))
	PowerFromPv float64 `json:"P_PV"`

	// RelativeSelfConsumption Current relative self consumption in %, null if no smart meter is connected
	RelativeSelfConsumption uint8 `json:"rel_SelfConsumption"`

	// RelativeAutonomy Current relative autonomy in %, null if no smart meter is connected
	RelativeAutonomy uint8 `json:"rel_Autonomy"`

	MeterLocation fronius.MeterLocation `json:"Meter_Location"`

	// EnergyDay Energy [Wh] this day, null if no inverter is connected
	EnergyDay float64 `json:"E_Day"`

	// EnergyYear Energy [Wh] this year, null if no inverter is connected
	EnergyYear float64 `json:"E_Year"`

	// EnergyTotal Energy [Wh] ever since, null if no inverter is connected
	EnergyTotal float64 `json:"E_Total"`
}

func (s site) ToGrpcRequest() *pb.SitePowerflow {
	return &pb.SitePowerflow{
		Mode:                    s.Mode.String(),
		BatteryStandby:          s.BatteryStandby,
		BackupMode:              s.BackupMode,
		PowerFromGrid:           s.PowerFromGrid,
		PowerLoad:               s.PowerLoad,
		PowerAkku:               s.PowerAkku,
		PowerFromPv:             s.PowerFromPv,
		RelativeSelfConsumption: uint32(s.RelativeSelfConsumption),
		RelativeAutonomy:        uint32(s.RelativeAutonomy),
		MeterLocation:           s.MeterLocation.String(),
		EnergyDay:               s.EnergyDay,
		EnergyYear:              s.EnergyYear,
		EnergyTotal:             s.EnergyTotal,
	}
}
