package main

import "fmt"

func main() {
	/*
		const red = 0
		const green = 1
		const blue = 2
	*/

	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red int = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red   = iota + 3 // 0 + 3
			green            // 1 + 3
			blue             // 2 + 3
		)
	*/

	/*
		const (
			red   = iota * 3 // 0 * 3
			green            // 1 * 3
			blue             // 2 * 3
		)
	*/

	// skipping a value
	const (
		red   = iota * 3 // 0 * 3
		green            // 1 * 3
		_
		blue // 3 * 3
	)

	fmt.Printf("Red = %d, Green = %d & Blue = %d\n", red, green, blue)

	fmt.Println("iota applied")
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	// fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
