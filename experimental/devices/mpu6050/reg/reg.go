package reg

const (
	MPU6050_DEFAULT_ADDRESS     = 0xD1
	MPU6050_ALT_DEFAULT_ADDRESS = 0xD2

	MPU6050_SELF_TEST_X_GYRO = 0x00
	MPU6050_SELF_TEST_Y_GYRO = 0x01
	MPU6050_SELF_TEST_Z_GYRO = 0x02

	MPU6050_SELF_TEST_X_ACCEL = 0x0D
	MPU6050_SELF_TEST_Y_ACCEL = 0x0E
	MPU6050_SELF_TEST_Z_ACCEL = 0x0F

	MPU6050_XG_OFFSET_H   = 0x13
	MPU6050_XG_OFFSET_L   = 0x14
	MPU6050_YG_OFFSET_H   = 0x15
	MPU6050_YG_OFFSET_L   = 0x16
	MPU6050_ZG_OFFSET_H   = 0x17
	MPU6050_ZG_OFFSET_L   = 0x18
	MPU6050_SMPLRT_DIV    = 0x19
	MPU6050_CONFIG        = 0x1A
	MPU6050_GYRO_CONFIG   = 0x1B
	MPU6050_ACCEL_CONFIG  = 0x1C
	MPU6050_ACCEL_CONFIG2 = 0x1D
	MPU6050_LP_ACCEL_ODR  = 0x1E
	MPU6050_WOM_THR       = 0x1F

	MPU6050_FIFO_EN        = 0x23
	MPU6050_I2C_MST_CTRL   = 0x24
	MPU6050_I2C_SLV0_ADDR  = 0x25
	MPU6050_I2C_SLV0_REG   = 0x26
	MPU6050_I2C_SLV0_CTRL  = 0x27
	MPU6050_I2C_SLV1_ADDR  = 0x28
	MPU6050_I2C_SLV1_REG   = 0x29
	MPU6050_I2C_SLV1_CTRL  = 0x2A
	MPU6050_I2C_SLV2_ADDR  = 0x2B
	MPU6050_I2C_SLV2_REG   = 0x2C
	MPU6050_I2C_SLV2_CTRL  = 0x2D
	MPU6050_I2C_SLV3_ADDR  = 0x2E
	MPU6050_I2C_SLV3_REG   = 0x2F
	MPU6050_I2C_SLV3_CTRL  = 0x30
	MPU6050_I2C_SLV4_ADDR  = 0x31
	MPU6050_I2C_SLV4_REG   = 0x32
	MPU6050_I2C_SLV4_DO    = 0x33
	MPU6050_I2C_SLV4_CTRL  = 0x34
	MPU6050_I2C_SLV4_DI    = 0x35
	MPU6050_I2C_MST_STATUS = 0x36
	MPU6050_INT_PIN_CFG    = 0x37
	MPU6050_INT_ENABLE     = 0x38

	MPU6050_INT_STATUS       = 0x3A
	MPU6050_ACCEL_XOUT_H     = 0x3B
	MPU6050_ACCEL_XOUT_L     = 0x3C
	MPU6050_ACCEL_YOUT_H     = 0x3D
	MPU6050_ACCEL_YOUT_L     = 0x3E
	MPU6050_ACCEL_ZOUT_H     = 0x3F
	MPU6050_ACCEL_ZOUT_L     = 0x40
	MPU6050_TEMP_OUT_H       = 0x41
	MPU6050_TEMP_OUT_L       = 0x42
	MPU6050_GYRO_XOUT_H      = 0x43
	MPU6050_GYRO_XOUT_L      = 0x44
	MPU6050_GYRO_YOUT_H      = 0x45
	MPU6050_GYRO_YOUT_L      = 0x46
	MPU6050_GYRO_ZOUT_H      = 0x47
	MPU6050_GYRO_ZOUT_L      = 0x48
	MPU6050_EXT_SENS_DATA_00 = 0x49
	MPU6050_EXT_SENS_DATA_01 = 0x4A
	MPU6050_EXT_SENS_DATA_02 = 0x4B
	MPU6050_EXT_SENS_DATA_03 = 0x4C
	MPU6050_EXT_SENS_DATA_04 = 0x4D
	MPU6050_EXT_SENS_DATA_05 = 0x4E
	MPU6050_EXT_SENS_DATA_06 = 0x4F
	MPU6050_EXT_SENS_DATA_07 = 0x50
	MPU6050_EXT_SENS_DATA_08 = 0x51
	MPU6050_EXT_SENS_DATA_09 = 0x52
	MPU6050_EXT_SENS_DATA_10 = 0x53
	MPU6050_EXT_SENS_DATA_11 = 0x54
	MPU6050_EXT_SENS_DATA_12 = 0x55
	MPU6050_EXT_SENS_DATA_13 = 0x56
	MPU6050_EXT_SENS_DATA_14 = 0x57
	MPU6050_EXT_SENS_DATA_15 = 0x58
	MPU6050_EXT_SENS_DATA_16 = 0x59
	MPU6050_EXT_SENS_DATA_17 = 0x5A
	MPU6050_EXT_SENS_DATA_18 = 0x5B
	MPU6050_EXT_SENS_DATA_19 = 0x5C
	MPU6050_EXT_SENS_DATA_20 = 0x5D
	MPU6050_EXT_SENS_DATA_21 = 0x5E
	MPU6050_EXT_SENS_DATA_22 = 0x5F
	MPU6050_EXT_SENS_DATA_23 = 0x60

	MPU6050_I2C_SLV0_DO        = 0x63
	MPU6050_I2C_SLV1_DO        = 0x64
	MPU6050_I2C_SLV2_DO        = 0x65
	MPU6050_I2C_SLV3_DO        = 0x66
	MPU6050_I2C_MST_DELAY_CTRL = 0x67
	MPU6050_SIGNAL_PATH_RESET  = 0x68
	MPU6050_MOT_DETECT_CTRL    = 0x69
	MPU6050_USER_CTRL          = 0x6A
	MPU6050_PWR_MGMT_1         = 0x6B
	MPU6050_PWR_MGMT_2         = 0x6C

	MPU6050_FIFO_COUNTH = 0x72
	MPU6050_FIFO_COUNTL = 0x73
	MPU6050_FIFO_R_W    = 0x74
	MPU6050_WHO_AM_I    = 0x75
	MPU6050_XA_OFFSET_H = 0x77
	MPU6050_XA_OFFSET_L = 0x78

	MPU6050_YA_OFFSET_H = 0x7A
	MPU6050_YA_OFFSET_L = 0x7B

	MPU6050_ZA_OFFSET_H = 0x7D
	MPU6050_ZA_OFFSET_L = 0x7E

	//reset values
	WHOAMI_RESET_VAL            = 0x71
	POWER_MANAGMENT_1_RESET_VAL = 0x01
	DEFAULT_RESET_VALUE         = 0x00

	WHOAMI_DEFAULT_VAL = 0x68

	//CONFIG register masks
	MPU6050_FIFO_MODE_MASK    = 0x40
	MPU6050_EXT_SYNC_SET_MASK = 0x38
	MPU6050_DLPF_CFG_MASK     = 0x07

	//GYRO_CONFIG register masks
	MPU6050_XGYRO_CTEN_MASK  = 0x80
	MPU6050_YGYRO_CTEN_MASK  = 0x40
	MPU6050_ZGYRO_CTEN_MASK  = 0x20
	MPU6050_GYRO_FS_SEL_MASK = 0x18
	MPU6050_FCHOICE_B_MASK   = 0x03

	MPU6050_GYRO_FULL_SCALE_250DPS  = 0
	MPU6050_GYRO_FULL_SCALE_500DPS  = 1
	MPU6050_GYRO_FULL_SCALE_1000DPS = 2
	MPU6050_GYRO_FULL_SCALE_2000DPS = 3

	//ACCEL_CONFIG register masks
	MPU6050_AX_ST_EN_MASK     = 0x80
	MPU6050_AY_ST_EN_MASK     = 0x40
	MPU6050_AZ_ST_EN_MASK     = 0x20
	MPU6050_ACCEL_FS_SEL_MASK = 0x18

	MPU6050_FULL_SCALE_2G  = 0
	MPU6050_FULL_SCALE_4G  = 1
	MPU6050_FULL_SCALE_8G  = 2
	MPU6050_FULL_SCALE_16G = 3

	//ACCEL_CONFIG_2 register masks
	MPU6050_ACCEL_FCHOICE_B_MASK = 0xC0
	MPU6050_A_DLPF_CFG_MASK      = 0x03

	//LP_ACCEL_ODR register masks
	MPU6050_LPOSC_CLKSEL_MASK = 0x0F

	//FIFO_EN register masks
	MPU6050_TEMP_FIFO_EN_MASK = 0x80
	MPU6050_GYRO_XOUT_MASK    = 0x40
	MPU6050_GYRO_YOUT_MASK    = 0x20
	MPU6050_GYRO_ZOUT_MASK    = 0x10
	MPU6050_ACCEL_MASK        = 0x08
	MPU6050_SLV2_MASK         = 0x04
	MPU6050_SLV1_MASK         = 0x02
	MPU6050_SLV0_MASK         = 0x01

	//I2C_MST_CTRL register masks
	MPU6050_MULT_MST_EN_MASK   = 0x80
	MPU6050_WAIT_FOR_ES_MASK   = 0x40
	MPU6050_SLV_3_FIFO_EN_MASK = 0x20
	MPU6050_I2C_MST_P_NSR_MASK = 0x10
	MPU6050_I2C_MST_CLK_MASK   = 0x0F

	//I2C_SLV0_ADDR register masks
	MPU6050_I2C_SLV0_RNW_MASK = 0x80
	MPU6050_I2C_ID_0_MASK     = 0x7F

	//I2C_SLV0_CTRL register masks
	MPU6050_I2C_SLV0_EN_MASK      = 0x80
	MPU6050_I2C_SLV0_BYTE_SW_MASK = 0x40
	MPU6050_I2C_SLV0_REG_DIS_MASK = 0x20
	MPU6050_I2C_SLV0_GRP_MASK     = 0x10
	MPU6050_I2C_SLV0_LENG_MASK    = 0x0F

	//I2C_SLV1_ADDR register masks
	MPU6050_I2C_SLV1_RNW_MASK = 0x80
	MPU6050_I2C_ID_1_MASK     = 0x7F

	//I2C_SLV1_CTRL register masks
	MPU6050_I2C_SLV1_EN_MASK      = 0x80
	MPU6050_I2C_SLV1_BYTE_SW_MASK = 0x40
	MPU6050_I2C_SLV1_REG_DIS_MASK = 0x20
	MPU6050_I2C_SLV1_GRP_MASK     = 0x10
	MPU6050_I2C_SLV1_LENG_MASK    = 0x0F

	//I2C_SLV2_ADDR register masks
	MPU6050_I2C_SLV2_RNW_MASK = 0x80
	MPU6050_I2C_ID_2_MASK     = 0x7F

	//I2C_SLV2_CTRL register masks
	MPU6050_I2C_SLV2_EN_MASK      = 0x80
	MPU6050_I2C_SLV2_BYTE_SW_MASK = 0x40
	MPU6050_I2C_SLV2_REG_DIS_MASK = 0x20
	MPU6050_I2C_SLV2_GRP_MASK     = 0x10
	MPU6050_I2C_SLV2_LENG_MASK    = 0x0F

	//I2C_SLV3_ADDR register masks
	MPU6050_I2C_SLV3_RNW_MASK = 0x80
	MPU6050_I2C_ID_3_MASK     = 0x7F

	//I2C_SLV3_CTRL register masks
	MPU6050_I2C_SLV3_EN_MASK      = 0x80
	MPU6050_I2C_SLV3_BYTE_SW_MASK = 0x40
	MPU6050_I2C_SLV3_REG_DIS_MASK = 0x20
	MPU6050_I2C_SLV3_GRP_MASK     = 0x10
	MPU6050_I2C_SLV3_LENG_MASK    = 0x0F

	//I2C_SLV4_ADDR register masks
	MPU6050_I2C_SLV4_RNW_MASK = 0x80
	MPU6050_I2C_ID_4_MASK     = 0x7F

	//I2C_SLV4_CTRL register masks
	MPU6050_I2C_SLV4_EN_MASK      = 0x80
	MPU6050_SLV4_DONE_INT_EN_MASK = 0x40
	MPU6050_I2C_SLV4_REG_DIS_MASK = 0x20
	MPU6050_I2C_MST_DLY_MASK      = 0x1F

	//I2C_MST_STATUS register masks
	MPU6050_PASS_THROUGH_MASK  = 0x80
	MPU6050_I2C_SLV4_DONE_MASK = 0x40
	MPU6050_I2C_LOST_ARB_MASK  = 0x20
	MPU6050_I2C_SLV4_NACK_MASK = 0x10
	MPU6050_I2C_SLV3_NACK_MASK = 0x08
	MPU6050_I2C_SLV2_NACK_MASK = 0x04
	MPU6050_I2C_SLV1_NACK_MASK = 0x02
	MPU6050_I2C_SLV0_NACK_MASK = 0x01

	//INT_PIN_CFG register masks
	MPU6050_ACTL_MASK              = 0x80
	MPU6050_OPEN_MASK              = 0x40
	MPU6050_LATCH_INT_EN_MASK      = 0x20
	MPU6050_INT_ANYRD_2CLEAR_MASK  = 0x10
	MPU6050_ACTL_FSYNC_MASK        = 0x08
	MPU6050_FSYNC_INT_MODE_EN_MASK = 0x04
	MPU6050_BYPASS_EN_MASK         = 0x02

	//INT_ENABLE register masks
	MPU6050_WOM_EN_MASK        = 0x40
	MPU6050_FIFO_OFLOW_EN_MASK = 0x10
	MPU6050_FSYNC_INT_EN_MASK  = 0x08
	MPU6050_RAW_RDY_EN_MASK    = 0x01

	//INT_STATUS register masks
	MPU6050_WOM_INT_MASK          = 0x40
	MPU6050_FIFO_OFLOW_INT_MASK   = 0x10
	MPU6050_FSYNC_INT_MASK        = 0x08
	MPU6050_RAW_DATA_RDY_INT_MASK = 0x01

	//I2C_MST_DELAY_CTRL register masks
	MPU6050_DELAY_ES_SHADOW_MASK = 0x80
	MPU6050_I2C_SLV4_DLY_EN_MASK = 0x10
	MPU6050_I2C_SLV3_DLY_EN_MASK = 0x08
	MPU6050_I2C_SLV2_DLY_EN_MASK = 0x04
	MPU6050_I2C_SLV1_DLY_EN_MASK = 0x02
	MPU6050_I2C_SLV0_DLY_EN_MASK = 0x01

	//SIGNAL_PATH_RESET register masks
	MPU6050_GYRO_RST_MASK  = 0x04
	MPU6050_ACCEL_RST_MASK = 0x02
	MPU6050_TEMP_RST_MASK  = 0x01

	//MOT_DETECT_CTRL register masks
	MPU6050_ACCEL_INTEL_EN_MASK   = 0x80
	MPU6050_ACCEL_INTEL_MODE_MASK = 0x40

	//USER_CTRL register masks
	MPU6050_FIFO_EN_MASK      = 0x40
	MPU6050_I2C_MST_EN_MASK   = 0x20
	MPU6050_I2C_IF_DIS_MASK   = 0x10
	MPU6050_FIFO_RST_MASK     = 0x04
	MPU6050_I2C_MST_RST_MASK  = 0x02
	MPU6050_SIG_COND_RST_MASK = 0x01

	//PWR_MGMT_1 register masks
	MPU6050_H_RESET_MASK            = 0x80
	MPU6050_SLEEP_MASK              = 0x40
	MPU6050_CYCLE_MASK              = 0x20
	MPU6050_GYRO_STANDBY_CYCLE_MASK = 0x10
	MPU6050_PD_PTAT_MASK            = 0x08
	MPU6050_CLKSEL_MASK             = 0x07

	//PWR_MGMT_2 register masks
	MPU6050_DISABLE_XA_MASK = 0x20
	MPU6050_DISABLE_YA_MASK = 0x10
	MPU6050_DISABLE_ZA_MASK = 0x08
	MPU6050_DISABLE_XG_MASK = 0x04
	MPU6050_DISABLE_YG_MASK = 0x02
	MPU6050_DISABLE_ZG_MASK = 0x01

	MPU6050_DISABLE_XYZA_MASK = 0x38
	MPU6050_DISABLE_XYZG_MASK = 0x07

	//Magnetometer register maps
	MPU6050_MAG_ADDRESS = 0x0C

	MPU6050_MAG_WIA    = 0x00
	MPU6050_MAG_INFO   = 0x01
	MPU6050_MAG_ST1    = 0x02
	MPU6050_MAG_XOUT_L = 0x03
	MPU6050_MAG_XOUT_H = 0x04
	MPU6050_MAG_YOUT_L = 0x05
	MPU6050_MAG_YOUT_H = 0x06
	MPU6050_MAG_ZOUT_L = 0x07
	MPU6050_MAG_ZOUT_H = 0x08
	MPU6050_MAG_ST2    = 0x09
	MPU6050_MAG_CNTL   = 0x0A
	MPU6050_MAG_RSV    = 0x0B //reserved mystery meat
	MPU6050_MAG_ASTC   = 0x0C
	MPU6050_MAG_TS1    = 0x0D
	MPU6050_MAG_TS2    = 0x0E
	MPU6050_MAG_I2CDIS = 0x0F
	MPU6050_MAG_ASAX   = 0x10
	MPU6050_MAG_ASAY   = 0x11
	MPU6050_MAG_ASAZ   = 0x12

	//Magnetometer register masks
	MPU6050_WIA_MASK = 0x48
)
