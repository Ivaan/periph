package mpu6050

import (
	"fmt"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/experimental/devices/mpu6050/reg"
)

// DebugF the debug function type.
type DebugF func(string, ...interface{})

// I2CTransport Encapsulates the I2C transport parameters.
type I2CTransport struct {
	device i2c.Dev
	debug  DebugF
}

// NewI2CTransport Creates the I2C transport using the provided device name.
func NewI2CTransport(busName string) (*I2CTransport, error) {
	bus, err := i2creg.Open(busName)
	if err != nil {
		return nil, wrapf("can't open SPI %v", err)
	}
	//TODO: find out where the good place to put this is
	//or determine we don't need it
	//defer bus.Close()

	dev := i2c.Dev{Bus: bus, Addr: reg.DEFAULT_RESET_VALUE}

	return &I2CTransport{device: dev, debug: noop}, nil
}

// EnableDebug Sets the debugging function
func (i *I2CTransport) EnableDebug(f DebugF) {
	i.debug = f
}

func (i *I2CTransport) writeByte(address byte, value byte) error {
	i.debug("write register %x value %x", address, value)
	var (
		write = []byte{address, value}
		read  = make([]byte, 0)
	)

	return i.device.Tx(write, read)
}

func (i *I2CTransport) writeMagReg(address byte, value byte) error {
	return i.writeByte(address, value)
}

func (i *I2CTransport) writeMaskedReg(address byte, mask byte, value byte) error {
	i.debug("write masked %x, mask %x, value %x", address, mask, value)
	maskedValue := mask & value
	i.debug("masked value %x", maskedValue)
	regVal, err := i.readByte(address)
	if err != nil {
		return err
	}
	i.debug("current register %x", regVal)
	regVal = (regVal &^ maskedValue) | maskedValue
	i.debug("new value %x", regVal)
	return i.writeByte(address, regVal)
}

func (i *I2CTransport) readMaskedReg(address byte, mask byte) (byte, error) {
	i.debug("read masked %x, mask %x", address, mask)
	reg, err := i.readByte(address)
	if err != nil {
		return 0, err
	}
	i.debug("masked value %x", reg)
	return reg & mask, nil
}

func (i *I2CTransport) readByte(address byte) (byte, error) {
	i.debug("read register %x", address)
	var (
		write = []byte{address}
		read  = make([]byte, 1)
	)
	if err := i.device.Tx(write, read); err != nil {
		return 0, err
	}

	i.debug("register content %x:", read[0])
	return read[0], nil
}

func (i *I2CTransport) readUint16(address ...byte) (uint16, error) {
	if len(address) != 2 {
		return 0, fmt.Errorf("Only 2 bytes per read")
	}
	h, err := i.readByte(address[0])
	if err != nil {
		return 0, err
	}
	l, err := i.readByte(address[1])
	if err != nil {
		return 0, err
	}
	return uint16(h)<<8 | uint16(l), nil
}

func noop(string, ...interface{}) {}
