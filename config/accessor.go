package config

import "fmt"

func (c *Config) GetString(keys string) (string, bool) {
	value, found := c.getValue(keys)
	return fmt.Sprintf("%v", value), found
}

func (c *Config) GetInt(keys string) (int, bool) {
	value, found := c.getValue(keys)
	if intValue, ok := value.(int); ok {
		return intValue, found
	}
	return 0, false
}

func (c *Config) GetBool(keys string) (bool, bool) {
	value, found := c.getValue(keys)
	if boolValue, ok := value.(bool); ok {
		return boolValue, found
	}
	return false, false
}

func (c *Config) GetMapInterface(keys string) (map[string]interface{}, bool) {
	value, found := c.getMap(keys, mapInterfaceName)
	if value, ok := value.(map[string]interface{}); ok {
		return value, found
	}
	return nil, false
}

func (c *Config) GetMapString(keys string) (map[string]string, bool) {
	value, found := c.getMap(keys, mapStringName)
	if value, ok := value.(map[string]string); ok {
		return value, found
	}
	return nil, false
}

func (c *Config) GetMapInt(keys string) (map[string]int, bool) {
	value, found := c.getMap(keys, mapIntName)
	if value, ok := value.(map[string]int); ok {
		return value, found
	}
	return nil, false
}
