package devices

//
//import (
//	"fmt"
//	"github.com/0xcafed00d/joystick"
//	evdev "github.com/gvalkov/golang-evdev"
//	"os"
//	"path/filepath"
//	"syscall"
//	"unsafe"
//)
//
//func getJoysticks() {
//	joystickPath := "/dev/input/js0" // Change this to the appropriate device path
//
//	fd, err := syscall.Open(joystickPath, syscall.O_RDONLY, 0)
//	if err != nil {
//		fmt.Println("Error opening joystick:", err)
//		return
//	}
//	defer syscall.Close(fd)
//
//	// Read joystick events
//	var event syscall.JoystickEvent
//	for {
//		_, _, errno := syscall.Syscall(syscall.SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(&event)), uintptr(unsafe.Sizeof(event)))
//		if errno != 0 {
//			fmt.Println("Error reading joystick event:", errno)
//			return
//		}
//
//		fmt.Printf("Time: %d, Value: %d, Type: %d, Number: %d\n", event.Time, event.Value, event.Type, event.Number)
//	}
//}
//
//func detectAxis() {
//	jsid := joystick.
//	js, err := joystick.Open(jsid)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("Joystick Name: %s", js.Name())
//	fmt.Printf("   Axis Count: %d", js.AxisCount())
//	fmt.Printf(" Button Count: %d", js.ButtonCount())
//
//	state, err := js.Read()
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("Axis Data: %v", state.AxisData)
//	js.Close()
//}
