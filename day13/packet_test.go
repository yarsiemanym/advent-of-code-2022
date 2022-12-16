package day13

import (
	"testing"
)

func Test_Packet_parsePacket_1(t *testing.T) {
	text := "[2]"
	packet := parsePacket(text).(*packet)

	if packet == nil {
		t.Error("Expected a packet but got nil.")
	} else if packet.data == nil {
		t.Error("Expected a packet data but got nil.")
	} else if len(packet.data) != 1 {
		t.Errorf("Expected 1 but got %d.", len(packet.data))
	} else if packet.data[0].(int) != 2 {
		t.Errorf("Expected 2 but got %d.", packet.data[0].(int))
	}
}

func Test_Packet_parsePacket_2(t *testing.T) {
	text := "[2,3]"
	packet := parsePacket(text).(*packet)

	if packet == nil {
		t.Error("Expected a packet but got nil.")
	} else if packet.data == nil {
		t.Error("Expected a packet data but got nil.")
	} else if len(packet.data) != 2 {
		t.Errorf("Expected 2 but got %d.", len(packet.data))
	} else if packet.data[0].(int) != 2 {
		t.Errorf("Expected 2 but got %d.", packet.data[0].(int))
	} else if packet.data[1].(int) != 3 {
		t.Errorf("Expected 3 but got %d.", packet.data[0].(int))
	}
}

func Test_Packet_parsePacket_3(t *testing.T) {
	text := "[2,[3,4]]"
	packet := parsePacket(text).(*packet)

	if packet == nil {
		t.Error("Expected a packet but got nil.")
	} else if packet.data == nil {
		t.Error("Expected a packet data but got nil.")
	} else if len(packet.data) != 2 {
		t.Errorf("Expected 2 but got %d.", len(packet.data))
	} else if packet.data[0].(int) != 2 {
		t.Errorf("Expected 2 but got %d.", packet.data[0].(int))
	} else if len(packet.data[1].([]interface{})) != 2 {
		t.Errorf("Expected 2 but got %d.", len(packet.data[1].([]interface{})))
	} else if packet.data[1].([]interface{})[0].(int) != 3 {
		t.Errorf("Expected 3 but got %d.", packet.data[1].([]interface{})[0].(int))
	} else if packet.data[1].([]interface{})[1].(int) != 4 {
		t.Errorf("Expected 4 but got %d.", packet.data[1].([]interface{})[1].(int))
	}
}

func Test_Packet_parsePacket_4(t *testing.T) {
	text := "[2,[3,4],5]"
	packet := parsePacket(text).(*packet)

	if packet == nil {
		t.Error("Expected a packet but got nil.")
	} else if packet.data == nil {
		t.Error("Expected a packet data but got nil.")
	} else if len(packet.data) != 3 {
		t.Errorf("Expected 3 but got %d.", len(packet.data))
	} else if packet.data[0].(int) != 2 {
		t.Errorf("Expected 2 but got %d.", packet.data[0].(int))
	} else if len(packet.data[1].([]interface{})) != 2 {
		t.Errorf("Expected 2 but got %d.", len(packet.data[1].([]interface{})))
	} else if packet.data[1].([]interface{})[0].(int) != 3 {
		t.Errorf("Expected 3 but got %d.", packet.data[1].([]interface{})[0].(int))
	} else if packet.data[1].([]interface{})[1].(int) != 4 {
		t.Errorf("Expected 4 but got %d.", packet.data[1].([]interface{})[1].(int))
	} else if packet.data[2].(int) != 5 {
		t.Errorf("Expected 5 but got %d.", packet.data[2].(int))
	}
}

func Test_Packet_parsePacket_5(t *testing.T) {
	text := "[]"
	packet := parsePacket(text).(*packet)

	if packet == nil {
		t.Error("Expected a packet but got nil.")
	} else if packet.data == nil {
		t.Error("Expected a packet data but got nil.")
	} else if len(packet.data) != 0 {
		t.Errorf("Expected 0 but got %d.", len(packet.data))
	}
}
