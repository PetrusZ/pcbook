package sample

import (
	"math/rand"
	"time"

	"github.com/PetrusZ/pcbook/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i7-4790",
			"Core i5-4690",
			"Core i3-4170",
		)
	}

	return randomStringFromSet(
		"Ryzen 5 3600",
		"Ryzen 3 2200",
		"Ryzen 3 1800",
	)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GeForce GTX 1070",
			"GeForce GTX 1080",
			"GeForce GTX 1080 Ti",
			"GeForce GTX 980",
		)
	}

	return randomStringFromSet(
		"Radeon RX 550",
		"Radeon RX 560",
		"Radeon RX 5700",
		"Radeon RX Vega",
	)
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("MacBook", "MacBook Pro", "MacBook Air")
	case "Dell":
		return randomStringFromSet("XPS 13", "XPS 15", "Inspiron")
	default:
		return randomStringFromSet("ThinkPad", "ThinkPad X1", "ThinkPad P1")
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	return &pb.Screen_Resolution{
		Width:  uint32(height),
		Height: uint32(width),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
