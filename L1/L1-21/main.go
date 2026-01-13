package main

import "fmt"

// целевой интерфейс, который ожидает клиент
type USBDevice interface {
	ConnectUSB() string
}

// существующая структура с несовместимым интерфейсом
type LightningPhone struct {
	Name string
}

func (p *LightningPhone) ConnectLightning() string {
	return p.Name + " подключен через Lightning"
}

// адаптер: реализует USBDevice, делегирует вызовы к LightningPhone
type LightningToUSBAdapter struct {
	phone *LightningPhone
}

func (a *LightningToUSBAdapter) ConnectUSB() string {
	return a.phone.ConnectLightning() + " (через USB-адаптер)"
}

// клиентский код работает только с USBDevice
func connectToComputer(device USBDevice) {
	fmt.Println(device.ConnectUSB())
}

func main() {
	iphone := &LightningPhone{Name: "iPhone"}
	adapter := &LightningToUSBAdapter{phone: iphone}
	connectToComputer(adapter) // iPhone подключен через Lightning (через USB-адаптер)
}
