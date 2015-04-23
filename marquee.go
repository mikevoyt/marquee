package main

import (
	"fmt"
	"os/exec"
	"time"
)

var screenlen = 16

func main() {

	go blinkled()

	exec.Command("sh", "-c", "echo 0 > /sys/kernel/lcd/clear").Output()

	offset := 0
	message0 := "               Thanks for everything, and good luck Hasty!  Vince, Mike, Raj, Roy, Geoff, Satish, and Gautam -- 4/22/2015   "
	message1 := "               Mar 10, 2010: Hasty co-founds Fig Card; Apr 28, 2011: PayPal acquires Fig Card; Dec 30, 2012: In exactly one year from this date, Hasty will be awarded a patent for \"DONGLE FACILITATED WIRELESS CONSUMER PAYMENTS\"; Apr 29 2013: PCWeb delivers the first Beacon hardware and demos the acceptance tests; Sep 9 2013: PayPal announces the PayPal Beacon; Sep 18 2013: Apple releases iOS 7 and its iBeacon technology; Sep 19 2013: Hasty wishes he had named it the \"Bacon\"; Nov 24 2013: Mike gets the bootloader running on the first Beacon 2 hardware using the Broadcom 4334; Sep 11, 2014: Whizz starts the final manufacturing run of Beacon 2 hardware at their factory in Malaysia; Apr 24 2015: first OneFob boards arrive to PayPal for bringup; Apr 24 2015: Hasty puts PayPal in the rear-view mirror - \"There was a fine man named Hasty, Who brought us cookies quite tasty, He said \"payments a chore, So Im out the door\", And now were left with no pastries\"    "

	for { //loop forever

		writeline(0, offset, message0)
		writeline(1, offset, message1)
		offset++
		time.Sleep(100 * time.Millisecond)
	}
}

func writeline(line int, offset int, message string) {
	start := offset % len(message)
	end := start + screenlen
	if end > len(message)-1 {
		end = len(message) - 1
	}
	msg := fmt.Sprintf("echo '%v' > /sys/kernel/lcd/line%v", message[start:end], line)
	exec.Command("sh", "-c", msg).Output()

}

func blinkled() {
	colors := []string{"white", "red", "green", "blue", "yellow", "purple"}

	i := 0
	for {
		color := colors[i%len(colors)]
		exec.Command("led-ctrl", color).Output()
		time.Sleep(100 * time.Millisecond)
		i++
	}
}
