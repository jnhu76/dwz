package services

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	uuid := GetUUID()
	if len(uuid) != 6 {
		t.Errorf("the length of uuid is %d", len(uuid))
	}

	uuid2 := GetUUID()
	if uuid == uuid2 {
		t.Errorf("Generator is wrong.")
	}
}
