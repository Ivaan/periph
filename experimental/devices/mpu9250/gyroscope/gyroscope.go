package gyroscope

const (
	GYRO_FS_SEL_250DPS  = 0
	GYRO_FS_SEL_500DPS  = 8
	GYRO_FS_SEL_1000DPS = 0x10
	GYRO_FS_SEL_2000DPS = 0x18

	GYRO_FS_SENS_250DPS  = 250 / 32768.0
	GYRO_FS_SENS_500DPS  = 500 / 32768.0
	GYRO_FS_SENS_1000DPS = 1000 / 32768.0
	GYRO_FS_SENS_2000DPS = 2000 / 32768.0
)

//Sensitivity returns the conversion factor from registry output to degrees per second
func Sensitivity(selector int) float32 {
	switch selector {
	case GYRO_FS_SEL_250DPS:
		return GYRO_FS_SENS_250DPS
	case GYRO_FS_SEL_500DPS:
		return GYRO_FS_SENS_500DPS
	case GYRO_FS_SEL_1000DPS:
		return GYRO_FS_SENS_1000DPS
	case GYRO_FS_SEL_2000DPS:
		return GYRO_FS_SENS_2000DPS
	}
	return GYRO_FS_SENS_250DPS
}
