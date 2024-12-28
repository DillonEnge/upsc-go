package upscgo

import (
	"os/exec"
	"strings"
)

const (
	STATUS_ONLINE                 = "OL"
	BATTERY_CHARGE                = "battery.charge"
	BATTERY_CHARGE_LOW            = "battery.charge.low"
	BATTERY_CHARGE_WARNING        = "battery.charge.warning"
	BATTERY_MFR_DATE              = "battery.mfr.date"
	BATTERY_RUNTIME               = "battery.runtime"
	BATTERY_RUNTIME_LOW           = "battery.runtime.low"
	BATTERY_TYPE                  = "battery.type"
	BATTERY_VOLTAGE               = "battery.voltage"
	BATTERY_VOLTAGE_NOMINAL       = "battery.voltage.nominal"
	DEVICE_MFR                    = "device.mfr"
	DEVICE_MODEL                  = "device.model"
	DEVICE_SERIAL                 = "device.serial"
	DEVICE_TYPE                   = "device.type"
	DRIVER_NAME                   = "driver.name"
	DRIVER_PARAMETER_POLLFREQ     = "driver.parameter.pollfreq"
	DRIVER_PARAMETER_POLLINTERVAL = "driver.parameter.pollinterval"
	DRIVER_PARAMETER_PORT         = "driver.parameter.port"
	DRIVER_PARAMETER_SYNCHRONOUS  = "driver.parameter.synchronous"
	DRIVER_VERSION                = "driver.version"
	DRIVER_VERSION_DATA           = "driver.version.data"
	DRIVER_VERSION_INTERNAL       = "driver.version.internal"
	DRIVER_VERSION_USB            = "driver.version.usb"
	INPUT_VOLTAGE                 = "input.voltage"
	INPUT_VOLTAGE_NOMINAL         = "input.voltage.nominal"
	OUTPUT_VOLTAGE                = "output.voltage"
	UPS_BEEPER_STATUS             = "ups.beeper.status"
	UPS_DELAY_SHUTDOWN            = "ups.delay.shutdown"
	UPS_DELAY_START               = "ups.delay.start"
	UPS_LOAD                      = "ups.load"
	UPS_MFR                       = "ups.mfr"
	UPS_MODEL                     = "ups.model"
	UPS_PRODUCTID                 = "ups.productid"
	UPS_REALPOWER_NOMINAL         = "ups.realpower.nominal"
	UPS_SERIAL                    = "ups.serial"
	UPS_STATUS                    = "ups.status"
	UPS_TEST_RESULT               = "ups.test.result"
	UPS_TIMER_SHUTDOWN            = "ups.timer.shutdown"
	UPS_TIMER_START               = "ups.timer.start"
	UPS_VENDORID                  = "ups.vendorid"
)

type Uspc struct {
	name string
}

func NewUspc(name string) *Uspc {
	return &Uspc{name: name}
}

func (u *Uspc) Query(v string) (string, error) {
	cmd := exec.Command("upsc", u.name, v)
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
