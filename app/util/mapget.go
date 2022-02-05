// Content managed by Project Forge, see [projectforge.md] for details.
package util

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (m ValueMap) GetRequired(k string) (interface{}, error) {
	v, ok := m[k]
	if !ok {
		const msg = "no value [%s] among candidates [%s]"
		return nil, errors.Errorf(msg, k, StringArrayOxfordComma(m.Keys(), "and"))
	}
	return v, nil
}

func (m ValueMap) GetArray(key string, allowEmpty bool, allowNonArray bool) ([]interface{}, error) {
	return m.ParseArray(key, false, allowEmpty, allowNonArray)
}

func (m ValueMap) GetBool(key string, allowEmpty bool) (bool, error) {
	return m.ParseBool(key, false, allowEmpty)
}

func (m ValueMap) GetInt(key string, allowEmpty bool) (int, error) {
	return m.ParseInt(key, false, allowEmpty)
}

func (m ValueMap) GetInt64(key string, allowEmpty bool) (int64, error) {
	ret, err := m.ParseInt(key, false, allowEmpty)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return int64(ret), nil
}

func (m ValueMap) GetFloat(key string, allowEmpty bool) (float64, error) {
	return m.ParseFloat(key, false, allowEmpty)
}

func (m ValueMap) GetMap(key string, allowEmpty bool) (ValueMap, error) {
	return m.ParseMap(key, false, allowEmpty)
}

func (m ValueMap) GetString(key string, allowEmpty bool) (string, error) {
	return m.ParseString(key, false, allowEmpty)
}

func (m ValueMap) GetStringOpt(key string) string {
	ret, _ := m.ParseString(key, true, true)
	return ret
}

func (m ValueMap) GetStringArray(key string, allowEmpty bool) ([]string, error) {
	ret, err := m.ParseArray(key, false, allowEmpty, true)
	if err != nil {
		return nil, err
	}
	strs := make([]string, 0, len(ret))
	for _, s := range ret {
		strs = append(strs, fmt.Sprint(s))
	}
	return strs, nil
}

func (m ValueMap) GetTime(key string, allowEmpty bool) (*time.Time, error) {
	return m.ParseTime(key, false, allowEmpty)
}

func (m ValueMap) GetUUID(key string, allowEmpty bool) (*uuid.UUID, error) {
	return m.ParseUUID(key, false, allowEmpty)
}

func (m ValueMap) GetType(key string, ret interface{}) error {
	result, err := m.GetPath(key, false)
	if err != nil {
		return errors.Wrap(err, "invalid type")
	}
	switch t := result.(type) {
	case []byte:
		err = FromJSON(t, ret)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal to expected type")
		}
		return nil
	default:
		return errors.Errorf("expected binary json data, encountered %T", t)
	}
}
