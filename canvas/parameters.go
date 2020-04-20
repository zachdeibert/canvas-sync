package canvas

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type parameterType struct {
	predicate func(reflect.Type) (bool, error)
	serialize func(string, interface{}) (string, error)
}

// RegisterParameterType registers a new type of parameter and how to serialize it
func (c *Canvas) RegisterParameterType(predicate func(reflect.Type) (bool, error), serialize func(string, interface{}) (string, error)) error {
	if predicate == nil || serialize == nil {
		return errors.New("Invalid argument")
	}
	c.parameterTypes = append(c.parameterTypes, parameterType{
		predicate: predicate,
		serialize: serialize,
	})
	return nil
}

// RegisterParameterType2 registers a new type of parameter and how to serialize it
func (c *Canvas) RegisterParameterType2(valType reflect.Type, serialize func(string, interface{}) (string, error)) error {
	return c.RegisterParameterType(func(t reflect.Type) (bool, error) {
		return t == valType, nil
	}, serialize)
}

// RegisterParameterType3 registers a new type of parameter and how to serialize it
func (c *Canvas) RegisterParameterType3(predicate func(reflect.Type) (bool, error), serialize func(interface{}) (string, error)) error {
	return c.RegisterParameterType(predicate, func(name string, val interface{}) (string, error) {
		valStr, err := serialize(val)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s=%s", name, valStr), nil
	})
}

// RegisterParameterType4 registers a new type of parameter and how to serialize it
func (c *Canvas) RegisterParameterType4(valType reflect.Type, serialize func(interface{}) (string, error)) error {
	return c.RegisterParameterType(func(t reflect.Type) (bool, error) {
		return t == valType, nil
	}, func(name string, val interface{}) (string, error) {
		valStr, err := serialize(val)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s=%s", name, valStr), nil
	})
}

func (c *Canvas) serializeParameter(name string, val interface{}) (string, error) {
	valType := reflect.TypeOf(val)
	for _, t := range c.parameterTypes {
		if match, err := t.predicate(valType); err != nil {
			return "", err
		} else if match {
			return t.serialize(name, val)
		}
	}
	return "", fmt.Errorf("Unknown type %T", val)
}

func (c *Canvas) registerDefaultParameterTypes() error {
	c.RegisterParameterType4(reflect.TypeOf(""), func(val interface{}) (string, error) {
		return url.QueryEscape(val.(string)), nil
	})
	c.RegisterParameterType(func(t reflect.Type) (bool, error) {
		return t.Kind() == reflect.Slice, nil
	}, func(name string, val interface{}) (string, error) {
		pName := fmt.Sprintf("%s[]", url.QueryEscape(name))
		v := reflect.ValueOf(val)
		strs := make([]string, v.Len())
		for i := range strs {
			var err error
			if strs[i], err = c.serializeParameter(pName, v.Index(i).Interface()); err != nil {
				return "", err
			}
		}
		return strings.Join(strs, "&"), nil
	})
	c.RegisterParameterType4(reflect.TypeOf(time.Time{}), func(val interface{}) (string, error) {
		return val.(time.Time).Format("2006-01-02"), nil
	})
	c.RegisterParameterType3(func(t reflect.Type) (bool, error) {
		return t.Kind() == reflect.String, nil
	}, func(val interface{}) (string, error) {
		return url.QueryEscape(reflect.ValueOf(val).String()), nil
	})
	return nil
}
