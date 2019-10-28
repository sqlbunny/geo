package geo

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
)

func scan(value interface{}, g geom) (err error) {
	var data []byte
	switch value.(type) {
	case []byte:
		data = make([]byte, len(value.([]byte)))
		_, err = hex.Decode(data, value.([]byte))
	case string:
		data, err = hex.DecodeString(value.(string))
	case nil:
		return errors.New("EWKB scan: use null package members to process NULLs")
	default:
		return errors.New("EWKB scan: value is neither byte slice nor string")
	}
	if err != nil {
		return err
	}
	return Unmarshal(data, g)
}

func value(g geom) (driver.Value, error) {
	data := Marshal(g)
	return hex.EncodeToString(data), nil
}
