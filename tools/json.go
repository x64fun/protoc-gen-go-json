package tools

import (
	"encoding/json"
	"fmt"
)

type JSONString string

func String(i interface{}) (JSONString, error) {
	switch v := i.(type) {
	case string:
		return JSONString(v), nil
	default:
		return "", fmt.Errorf("json: unsupported %T into JSONString data type", i)
	}
}
func GetString(v JSONString) (string, error) {
	return string(v), nil
}

type JSONNumber interface{}

func Number(i interface{}) (JSONNumber, error) {
	switch v := i.(type) {
	case int, int8, int16, int32, int64:
		return v, nil
	case uint, uint8, uint16, uint32, uint64:
		return v, nil
	case float32, float64:
		return v, nil
	case nil:
		return v, nil
	default:
		return 0, fmt.Errorf("json: unsupported %T into JSONNumber data type", i)
	}
}
func GetInt(v JSONNumber) (int, error) {
	switch tmpV := v.(type) {
	case int:
		return tmpV, nil
	case int8:
		return int(tmpV), nil
	case int16:
		return int(tmpV), nil
	case int32:
		return int(tmpV), nil
	case int64:
		return int(tmpV), nil
	case uint:
		return int(tmpV), nil
	case uint8:
		return int(tmpV), nil
	case uint16:
		return int(tmpV), nil
	case uint32:
		return int(tmpV), nil
	case uint64:
		return int(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetInt8(v JSONNumber) (int8, error) {
	switch tmpV := v.(type) {
	case int:
		return int8(tmpV), nil
	case int8:
		return tmpV, nil
	case int16:
		return int8(tmpV), nil
	case int32:
		return int8(tmpV), nil
	case int64:
		return int8(tmpV), nil
	case uint:
		return int8(tmpV), nil
	case uint8:
		return int8(tmpV), nil
	case uint16:
		return int8(tmpV), nil
	case uint32:
		return int8(tmpV), nil
	case uint64:
		return int8(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetInt16(v JSONNumber) (int16, error) {
	switch tmpV := v.(type) {
	case int:
		return int16(tmpV), nil
	case int8:
		return int16(tmpV), nil
	case int16:
		return tmpV, nil
	case int32:
		return int16(tmpV), nil
	case int64:
		return int16(tmpV), nil
	case uint:
		return int16(tmpV), nil
	case uint8:
		return int16(tmpV), nil
	case uint16:
		return int16(tmpV), nil
	case uint32:
		return int16(tmpV), nil
	case uint64:
		return int16(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetInt32(v JSONNumber) (int32, error) {
	switch tmpV := v.(type) {
	case int:
		return int32(tmpV), nil
	case int8:
		return int32(tmpV), nil
	case int16:
		return int32(tmpV), nil
	case int32:
		return tmpV, nil
	case int64:
		return int32(tmpV), nil
	case uint:
		return int32(tmpV), nil
	case uint8:
		return int32(tmpV), nil
	case uint16:
		return int32(tmpV), nil
	case uint32:
		return int32(tmpV), nil
	case uint64:
		return int32(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetInt64(v JSONNumber) (int64, error) {
	switch tmpV := v.(type) {
	case int:
		return int64(tmpV), nil
	case int8:
		return int64(tmpV), nil
	case int16:
		return int64(tmpV), nil
	case int32:
		return int64(tmpV), nil
	case int64:
		return tmpV, nil
	case uint:
		return int64(tmpV), nil
	case uint8:
		return int64(tmpV), nil
	case uint16:
		return int64(tmpV), nil
	case uint32:
		return int64(tmpV), nil
	case uint64:
		return int64(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetUint(v JSONNumber) (uint, error) {
	switch tmpV := v.(type) {
	case int:
		return uint(tmpV), nil
	case int8:
		return uint(tmpV), nil
	case int16:
		return uint(tmpV), nil
	case int32:
		return uint(tmpV), nil
	case int64:
		return uint(tmpV), nil
	case uint:
		return tmpV, nil
	case uint8:
		return uint(tmpV), nil
	case uint16:
		return uint(tmpV), nil
	case uint32:
		return uint(tmpV), nil
	case uint64:
		return uint(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetUint8(v JSONNumber) (uint8, error) {
	switch tmpV := v.(type) {
	case int:
		return uint8(tmpV), nil
	case int8:
		return uint8(tmpV), nil
	case int16:
		return uint8(tmpV), nil
	case int32:
		return uint8(tmpV), nil
	case int64:
		return uint8(tmpV), nil
	case uint:
		return uint8(tmpV), nil
	case uint8:
		return tmpV, nil
	case uint16:
		return uint8(tmpV), nil
	case uint32:
		return uint8(tmpV), nil
	case uint64:
		return uint8(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetUint16(v JSONNumber) (uint16, error) {
	switch tmpV := v.(type) {
	case int:
		return uint16(tmpV), nil
	case int8:
		return uint16(tmpV), nil
	case int16:
		return uint16(tmpV), nil
	case int32:
		return uint16(tmpV), nil
	case int64:
		return uint16(tmpV), nil
	case uint:
		return uint16(tmpV), nil
	case uint8:
		return uint16(tmpV), nil
	case uint16:
		return tmpV, nil
	case uint32:
		return uint16(tmpV), nil
	case uint64:
		return uint16(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetUint32(v JSONNumber) (uint32, error) {
	switch tmpV := v.(type) {
	case int:
		return uint32(tmpV), nil
	case int8:
		return uint32(tmpV), nil
	case int16:
		return uint32(tmpV), nil
	case int32:
		return uint32(tmpV), nil
	case int64:
		return uint32(tmpV), nil
	case uint:
		return uint32(tmpV), nil
	case uint8:
		return uint32(tmpV), nil
	case uint16:
		return uint32(tmpV), nil
	case uint32:
		return tmpV, nil
	case uint64:
		return uint32(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetUint64(v JSONNumber) (uint64, error) {
	switch tmpV := v.(type) {
	case int:
		return uint64(tmpV), nil
	case int8:
		return uint64(tmpV), nil
	case int16:
		return uint64(tmpV), nil
	case int32:
		return uint64(tmpV), nil
	case int64:
		return uint64(tmpV), nil
	case uint:
		return uint64(tmpV), nil
	case uint8:
		return uint64(tmpV), nil
	case uint16:
		return uint64(tmpV), nil
	case uint32:
		return uint64(tmpV), nil
	case uint64:
		return tmpV, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetFloat32(v JSONNumber) (float32, error) {
	switch tmpV := v.(type) {
	case int:
		return float32(tmpV), nil
	case int8:
		return float32(tmpV), nil
	case int16:
		return float32(tmpV), nil
	case int32:
		return float32(tmpV), nil
	case int64:
		return float32(tmpV), nil
	case uint:
		return float32(tmpV), nil
	case uint8:
		return float32(tmpV), nil
	case uint16:
		return float32(tmpV), nil
	case uint32:
		return float32(tmpV), nil
	case uint64:
		return float32(tmpV), nil
	case float32:
		return tmpV, nil
	case float64:
		return float32(tmpV), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}
func GetFloat64(v JSONNumber) (float64, error) {
	switch tmpV := v.(type) {
	case int:
		return float64(tmpV), nil
	case int8:
		return float64(tmpV), nil
	case int16:
		return float64(tmpV), nil
	case int32:
		return float64(tmpV), nil
	case int64:
		return float64(tmpV), nil
	case uint:
		return float64(tmpV), nil
	case uint8:
		return float64(tmpV), nil
	case uint16:
		return float64(tmpV), nil
	case uint32:
		return float64(tmpV), nil
	case uint64:
		return float64(tmpV), nil
	case float32:
		return float64(tmpV), nil
	case float64:
		return tmpV, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("json: cannot unmarshal %T into Go struct field", v)
	}
}

type JSONBoolean bool

func Boolean(i interface{}) (JSONBoolean, error) {
	switch v := i.(type) {
	case bool:
		return JSONBoolean(v), nil
	default:
		return false, fmt.Errorf("json: unsupported %T into JSONBoolean data type", i)
	}
}
func GetBool(v JSONBoolean) (bool, error) {
	return bool(v), nil
}

type JSONObject map[string]interface{}

func Object(i interface{}) (JSONObject, error) {
	switch v := i.(type) {
	case []byte:
		tmp := make(map[string]interface{})
		if err := json.Unmarshal(v, &tmp); err != nil {
			return nil, err
		}
		return tmp, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("json: unsupported %T into JSONObject data type", i)
	}
}
func GetBytes(v JSONObject) ([]byte, error) {
	return json.Marshal(&v)
}
