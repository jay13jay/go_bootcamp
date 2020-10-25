package main

import "fmt"

func main() {
	speed := 100
	force := 2.5
	label := 30032

	fmt.Printf("Speed: %v\tForce: %v\n", speed, force)
	// power := speed * int(force)
	power := float64(speed) * force
	fmt.Printf("Speed * Force:\t%v\n", power)

	fmt.Printf("Label:\t%v\n", label)
	// newlabel := float64(label)
	// fmt.Printf("Label:\t%v\n", newlabel)
	fmt.Printf("Label:\t%v", stringlabel)

}
