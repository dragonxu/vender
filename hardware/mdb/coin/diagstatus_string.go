// Code generated by "stringer -type=DiagStatus -trimprefix=Diag"; DO NOT EDIT.

package coin

import "strconv"

const _DiagStatus_name = "PoweringUpPoweringDownOKKeypadShiftedManualActiveNewInventoryInformationInhibitedGeneralErrorGeneralChecksum1GeneralChecksum2GeneralVoltageDiscriminatorErrorDiscriminatorFlightOpenDiscriminatorReturnOpenDiscriminatorJamDiscriminatorBelowStandardDiscriminatorSensorADiscriminatorSensorBDiscriminatorSensorCDiscriminatorTemperatureDiscriminatorOpticsAccepterErrorAccepterJamAccepterAlarmAccepterEmptyAccepterExitBeforeEnterSeparatorErrorSeparatorSortSensorDispenserErrorStorageErrorStorageCassetteRemovedStorageCashboxSensorStorageAmbientLight"

var _DiagStatus_map = map[DiagStatus]string{
	256:  _DiagStatus_name[0:10],
	512:  _DiagStatus_name[10:22],
	768:  _DiagStatus_name[22:24],
	1024: _DiagStatus_name[24:37],
	1296: _DiagStatus_name[37:49],
	1312: _DiagStatus_name[49:72],
	1536: _DiagStatus_name[72:81],
	4096: _DiagStatus_name[81:93],
	4097: _DiagStatus_name[93:109],
	4098: _DiagStatus_name[109:125],
	4099: _DiagStatus_name[125:139],
	4352: _DiagStatus_name[139:157],
	4368: _DiagStatus_name[157:180],
	4369: _DiagStatus_name[180:203],
	4400: _DiagStatus_name[203:219],
	4417: _DiagStatus_name[219:245],
	4432: _DiagStatus_name[245:265],
	4433: _DiagStatus_name[265:285],
	4434: _DiagStatus_name[285:305],
	4435: _DiagStatus_name[305:329],
	4436: _DiagStatus_name[329:348],
	4608: _DiagStatus_name[348:361],
	4656: _DiagStatus_name[361:372],
	4657: _DiagStatus_name[372:385],
	4672: _DiagStatus_name[385:398],
	4688: _DiagStatus_name[398:421],
	4864: _DiagStatus_name[421:435],
	4880: _DiagStatus_name[435:454],
	5120: _DiagStatus_name[454:468],
	5376: _DiagStatus_name[468:480],
	5378: _DiagStatus_name[480:502],
	5379: _DiagStatus_name[502:522],
	5380: _DiagStatus_name[522:541],
}

func (i DiagStatus) String() string {
	if str, ok := _DiagStatus_map[i]; ok {
		return str
	}
	return "DiagStatus(" + strconv.FormatInt(int64(i), 10) + ")"
}
