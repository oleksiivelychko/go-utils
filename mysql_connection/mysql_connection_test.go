package mysql_connection

import (
	"testing"
)

func TestMySQLConnectionConnectionStringSuccess(t *testing.T) {
	conn, err := NewMySQLConnection("test", "test", "test")
	if conn == nil {
		t.Error(err.Error())
	}
}

func TestMySQLConnectionConnectionStringFailed(t *testing.T) {
	conn, err := NewMySQLConnection("", "", "")
	if conn != nil {
		t.Error(err.Error())
	}
}

func TestMySQLConnectionConnectionClose(t *testing.T) {
	conn, _ := NewMySQLConnection("test", "test", "test")
	err := conn.Close()
	if err != nil {
		t.Error(err.Error())
	}
}
