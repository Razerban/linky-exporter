// Copyright 2023 Ahmed Abdelkafi
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prom

type LinkyTimeSerie struct {
	LinkyId                            string
	Version                            string
	LinkyDate                          float64
	ContractTypeName                   string
	PriceLabel                         string
	TotalEnergyUsed                    float64
	EnergyUsedIndex1                   float64
	EnergyUsedIndex2                   float64
	EnergyUsedIndex3                   float64
	EnergyUsedIndex4                   float64
	EnergyUsedIndex5                   float64
	EnergyUsedIndex6                   float64
	EnergyUsedIndex7                   float64
	EnergyUsedIndex8                   float64
	EnergyUsedIndex9                   float64
	EnergyUsedIndex10                  float64
	EnergyUsedDistributorIndex1        float64
	EnergyUsedDistributorIndex2        float64
	EnergyUsedDistributorIndex3        float64
	EnergyUsedDistributorIndex4        float64
	TotalEnergyProduced                float64
	TotalReactiveEnergyQ1              float64
	TotalReactiveEnergyQ2              float64
	TotalReactiveEnergyQ3              float64
	TotalReactiveEnergyQ4              float64
	IntensityP1                        float64
	IntensityP2                        float64
	IntensityP3                        float64
	VoltageP1                          float64
	VoltageP2                          float64
	VoltageP3                          float64
	ReferencePower                     float64
	BreakingPower                      float64
	PowerUsed                          float64
	PowerUsedP1                        float64
	PowerUsedP2                        float64
	PowerUsedP3                        float64
	PowerUsedMax                       float64
	PowerUsedMaxP1                     float64
	PowerUsedMaxP2                     float64
	PowerUsedMaxP3                     float64
	PowerUsedMaxLastYear               float64
	PowerUsedMaxLastYearP1             float64
	PowerUsedMaxLastYearP2             float64
	PowerUsedMaxLastYearP3             float64
	PowerProduced                      float64
	PowerProducedMax                   float64
	PowerProducedLastYear              float64
	UsedLoadCurvePoint                 float64
	UsedLoadCurvePointLastYear         float64
	ProducedLoadCurvePoint             float64
	ProducedLoadCurvePointLastYear     float64
	AverageVoltageP1                   float64
	AverageVoltageP2                   float64
	AverageVoltageP3                   float64
	DryContactStatus                   float64
	CutOffDeviceStatus                 float64
	LinkyTerminalShieldStatus          float64
	SurgeStatus                        float64
	ReferencePowerExceededStatus       float64
	ConsumptionStatus                  float64
	EnergyDirectionStatus              float64
	ContractTypePriceStatus            float64
	ContractTypePriceDistributorStatus float64
	ClockStatus                        float64
	TicStatus                          float64
	EuridisLinkStatus                  float64
	CPLStatus                          float64
	CPLSyncStatus                      float64
	TempoContractColorStatus           float64
	TempoContractNextDayColorStatus    float64
	MovingPeakNoticeStatus             float64
	MovingPeakStatus                   float64
	MovingPeakStart1                   float64
	MovingPeakEnd1                     float64
	MovingPeakStart2                   float64
	MovingPeakEnd2                     float64
	MovingPeakStart3                   float64
	MovingPeakEnd3                     float64
	Prm                                string
	Relay1                             float64
	Relay2                             float64
	Relay3                             float64
	Relay4                             float64
	Relay5                             float64
	Relay6                             float64
	Relay7                             float64
	Relay8                             float64
	CurrentPricingNumber               string
	ContractTypeDayNumber              string
	ContractTypeNextDayNumber          string
	ContractTypeNextDayProfile         string
	PeakNextDayProfile                 string
	// Message1 string
	// Message2 string
}
