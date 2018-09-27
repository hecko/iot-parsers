package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func get_temp(s string) float32 {
	decoded, _ := hex.DecodeString(s[:4])
	buf := bytes.NewReader(decoded)
	var pl int16
	binary.Read(buf, binary.BigEndian, &pl)
	return float32(pl) / 16
}

func get_humi(s string) uint8 {
	decoded, _ := hex.DecodeString(s[6:8])
	buf := bytes.NewReader(decoded)
	var pl uint8
	binary.Read(buf, binary.BigEndian, &pl)
	return pl
}

func get_batt(s string) uint16 {
	decoded, _ := hex.DecodeString(s[8:12])
	buf := bytes.NewReader(decoded)
	var pl uint16
	binary.Read(buf, binary.BigEndian, &pl)
	return pl
}

func data_parse(packet string) {
	fmt.Printf("Packet in raw form: %s\n", packet)
	fmt.Printf("Temperature: %f\nHumidity: %d\nBattery: %d\n", get_temp(packet), get_humi(packet), get_batt(packet))
	fmt.Println("===")
}

func main() {
	packets := [...]string{"018100280266", "FEFC00440258"}
	for _, packet := range packets {
		data_parse(packet)
	}
}
