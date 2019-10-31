package main

import "fmt"

type Device interface {
	isEnabled() bool
	enable()
	disable()
	getVolume() int
	setVolume(percent int)
	getChannel() int
	setChannel(channel int)
	printStatus()
}

type Radio struct {
	on      bool
	volume  int
	channel int
}

func (r *Radio) isEnabled() bool {
	return r.on
}

func (r *Radio) enable() {
	r.on = true
}

func (r *Radio) disable() {
	r.on = false
}

func (r *Radio) getVolume() int {
	return r.volume
}

func (r *Radio) setVolume(volume int) {
	if volume > 100 {
		r.volume = 100
	} else if volume < 0 {
		r.volume = 0
	} else {
		r.volume = volume
	}
}

func (r *Radio) getChannel() int {
	return r.channel
}

func (r *Radio) setChannel(channel int) {
	r.channel = channel
}

func (r *Radio) printStatus() {
	fmt.Println("------------------------------------")
	fmt.Println("| I'm radio.")
	if r.on {
		fmt.Println("| I'm enabled")
	} else {
		fmt.Println("| I'm disabled")
	}
	fmt.Printf("| Current volume is %d percent\n", r.volume)
	fmt.Printf("| Current channel is %d\n", r.channel)
	fmt.Printf("------------------------------------\n\n")
}

type Tv struct {
	on      bool
	volume  int
	channel int
}

func (t *Tv) isEnabled() bool {
	return t.on
}

func (t *Tv) enable() {
	t.on = true
}

func (t *Tv) disable() {
	t.on = false
}

func (t *Tv) getVolume() int {
	return t.volume
}

func (t *Tv) setVolume(volume int) {
	if volume > 100 {
		t.volume = 100
	} else if volume < 0 {
		t.volume = 0
	} else {
		t.volume = volume
	}
}

func (t *Tv) getChannel() int {
	return t.channel
}

func (t *Tv) setChannel(channel int) {
	t.channel = channel
}

func (t *Tv) printStatus() {
	fmt.Println("------------------------------------")
	fmt.Println("| I'm TV set.")
	if t.on {
		fmt.Println("| I'm enabled")
	} else {
		fmt.Println("| I'm disabled")
	}
	fmt.Printf("| Current volume is %d percent\n", t.volume)
	fmt.Printf("| Current channel is %d\n", t.channel)
	fmt.Printf("------------------------------------\n\n")
}

type Remote interface {
	power()
	volumeDown()
	volumeUp()
	channelDown()
	channelUp()
}

type BasicRemote struct {
	device Device
}

func (b BasicRemote) power() {
	fmt.Println("Remote: power toggle")
	if b.device.isEnabled() {
		b.device.disable()
	} else {
		b.device.enable()
	}
}

func (b BasicRemote) volumeDown() {
	fmt.Println("Remote: volume down")
	b.device.setVolume(b.device.getVolume() - 10)
}

func (b BasicRemote) volumeUp() {
	fmt.Println("Remote: volume up")
	b.device.setVolume(b.device.getVolume() + 10)
}

func (b BasicRemote) channelDown() {
	fmt.Println("Remote: channel down")
	b.device.setChannel(b.device.getChannel() - 1)
}

func (b BasicRemote) channelUp() {
	fmt.Println("Remote: channel up")
	b.device.setChannel(b.device.getChannel() + 1)
}

type AdvancedRemote struct {
	BasicRemote
}

func (ar AdvancedRemote) mute() {
	fmt.Println("Remote: mute")
	ar.device.setVolume(0)
}

func testDevice(device Device) {
	fmt.Println("Tests with basic remote.")
	var basicRemote BasicRemote
	basicRemote.device = device
	basicRemote.power()
	basicRemote.channelUp()
	basicRemote.volumeUp()
	basicRemote.device.printStatus()

	fmt.Println("Tests with advanced remote.")
	var advancedRemote AdvancedRemote
	advancedRemote.device = device
	advancedRemote.power()
	advancedRemote.mute()
	device.printStatus()
}

func main() {
	var tv Device
	tv = &Radio{}
	var radio Device
	radio = &Radio{}
	testDevice(tv)
	testDevice(radio)
}
