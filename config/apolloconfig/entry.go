package apolloconfig

import (
	"strconv"
	"strings"
)

type Entry[T any] interface {
	Get() T
}

// To convert config value from string type to another type
type convertFunction[T any] func(string) (T, error)

type entryImpl[T any] struct {
	key          string
	defaultValue T
	convertFn    convertFunction[T]
}

func NewEntry[T any](key string, defaultValue interface{}, convertFn convertFunction[T]) Entry[T] {
	return &entryImpl[T]{
		key:          key,
		defaultValue: defaultValue,
		convertFn:    convertFn,
	}
}

func (e *entryImpl[T]) Get() T {
	// If Apollo config is not enabled, just return the default value
	if !enabled {
		return e.defaultValue
	}

	logger := getLogger().WithFields("key", e.key)

	// If client is not initialized, return the default value
	client := GetClient()
	if client == nil {
		logger.Debugf("apollo client is nil")
		return e.defaultValue
	}

	// Get the string value and convert it to type T
	s := client.GetValue(e.key)

	if e.convertFn == nil {
		logger.Debugf("convertFn is nil")
		return e.defaultValue
	}
	v, err := e.convertFn(s)
	if err != nil {
		logger.Debugf("conversion error: %v", err)
		// If cannot convert, return the default value
		return e.defaultValue
	}
	return v
}

// ----- Convert functions -----

func ToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, parseIntBase, parseIntBitSize)
}

func ToUint32Slice(s string) ([]uint32, error) {
	sArr := strings.Split(s, comma)
	result := make([]uint32, len(sArr))
	for i := range sArr {
		v, err := ToInt(sArr[i])
		if err != nil {
			return nil, err
		}
		result[i] = uint32(v)
	}
	return result, nil
}
