package modbus

type pdu struct {
	unitId       uint8
	functionCode uint8
	payload      []byte
}

type Error string

// Error implements the error interface.
func (me Error) Error() (s string) {
	s = string(me)
	return
}

const (
	// coils
	fcReadCoils          uint8 = 0x01
	fcWriteSingleCoil    uint8 = 0x05
	fcWriteMultipleCoils uint8 = 0x0f

	// discrete inputs
	fcReadDiscreteInputs uint8 = 0x02

	// 16-bit input/holding registers
	fcReadHoldingRegisters       uint8 = 0x03
	fcReadInputRegisters         uint8 = 0x04
	fcWriteSingleRegister        uint8 = 0x06
	fcWriteMultipleRegisters     uint8 = 0x10
	fcMaskWriteRegister          uint8 = 0x16
	fcReadWriteMultipleRegisters uint8 = 0x17
	fcReadFifoQueue              uint8 = 0x18

	// file access
	// fcReadFileRecord  uint8 = 0x14
	// fcWriteFileRecord uint8 = 0x15

	// exception codes
	exIllegalFunction         uint8 = 0x01
	exIllegalDataAddress      uint8 = 0x02
	exIllegalDataValue        uint8 = 0x03
	exServerDeviceFailure     uint8 = 0x04
	exAcknowledge             uint8 = 0x05
	exServerDeviceBusy        uint8 = 0x06
	exMemoryParityError       uint8 = 0x08
	exGWPathUnavailable       uint8 = 0x0a
	exGWTargetFailedToRespond uint8 = 0x0b
)
