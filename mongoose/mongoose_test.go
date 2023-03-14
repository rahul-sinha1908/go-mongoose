package mongoose

import (
	"fmt"
	"testing"
)

const defaultPort = 27017

func TestURLNoSRV(t *testing.T) {
	dbInfo := DBConnection{
		Host:     "testHost",
		Port:     0,
		Database: "testDatabase",
		User:     "testUser",
		Password: "testPassword",
	}
	InitiateDB(dbInfo)
	expectedURL := fmt.Sprintf("mongodb://%s:%s@%s:%d", dbInfo.User, dbInfo.Password, dbInfo.Host, defaultPort)
	if connectionURL != expectedURL {
		t.Errorf(`
		expected -- %s
		actual ---- %s
		`, expectedURL, connectionURL)
		t.Fail()
	}
}

func TestURLSRV(t *testing.T) {
	dbInfo := DBConnection{
		Host:     "testHost",
		Port:     0,
		Database: "testDatabase",
		User:     "testUser",
		Password: "testPassword",
		SRV:      true,
	}
	InitiateDB(dbInfo)
	expectedURL := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbInfo.User, dbInfo.Password, dbInfo.Host)
	if connectionURL != expectedURL {
		t.Errorf(`
		expected -- %s
		actual ---- %s
		`, expectedURL, connectionURL)
		t.Fail()
	}
}
