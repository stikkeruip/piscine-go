package piscine

import "github.com/01-edu/z01"

type DoorState bool

const (
	OPEN  DoorState = true
	CLOSE DoorState = false
)

type Door struct {
	state DoorState
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor *Door) bool {
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}

func OpenDoor(door *Door) {
	door.state = OPEN
}
