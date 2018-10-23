package shared

import (
	"errors"
	"reflect"
	"strings"
)

//http://woodair.net/sbs/article/barebones42_socket_data.htm

type BaseStation struct {
	MessageType      string
	TransmissionType string
	SessionID        string
	AircraftID       string
	HexIdent         string
	FlightID         string

	DateMessageGenerated string
	TimeMessageGenerated string
	DateMessageLogged    string
	TimeMessageLogged    string

	Callsign     string
	Altitude     string
	GroundSpeed  string
	Track        string
	Latitude     string
	Longitude    string
	VerticalRate string
	Squawk       string

	Alert      bool
	Emergency  bool
	SPI        bool
	IsOnGround bool
}

func ParseBaseStation(msg string) (*BaseStation, error) {
	format := new(BaseStation)
	if !strings.HasPrefix(msg, "MSG") {
		return format, errors.New("invalid messasge, no MSG prefix")
	}
	if !strings.Contains(msg, ",") {
		return format, errors.New("invalid messasge, not csv format")
	}

	split := strings.Split(msg, ",")

	if len(split) != 22 {
		return format, errors.New("invalid messasge, not enough data fields")
	}

	//Nice to see that golang has reflection, works nicely as well
	ref := reflect.Indirect(reflect.ValueOf(format))
	for i := 0; i < ref.NumField(); i++ {
		if len(split[i]) <= 0 {
			continue
		}
		if ref.Field(i).Type().Kind() == reflect.String {
			ref.Field(i).SetString(split[i])
		} else if ref.Field(i).Type().Kind() == reflect.Bool {
			//TODO parse the value
			ref.Field(i).SetBool(false)
		}
	}

	return format, nil
}
