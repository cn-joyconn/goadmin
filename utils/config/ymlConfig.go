package config

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"sync"
// 	gologs "github.com/cn-joyconn/gologs"
// 	filetool "github.com/cn-joyconn/goutils/filetool"
// 	yaml "gopkg.in/yaml.v2"
// //	gologs "github.com/cn-joyconn/gologs"

// )


// // YamlConfig is a config which represents the yaml configuration.
// type YamlConfig struct {
// 	data map[string]interface{}
// 	sync.RWMutex
// }
// // Parse returns a ConfigContainer with parsed yaml config map.
// func (c *YamlConfig) Parse(filename string) (error) {
// 	if filetool.IsExist(filename) {
// 		configBytes, err := filetool.ReadFileToBytes(filename)
// 		if err != nil {
// 			gologs.GetLogger("").Error(err.Error())
// 			return err
// 		}
// 		err = yaml.Unmarshal(configBytes, &c.data)
// 		if err != nil {
// 			errors.New(fmt.Sprintf("解析%s文件失败", filename))
// 			return err
// 		}
// 		return nil
// 	} else {
// 		return errors.New(fmt.Sprintf("未找到: %s", filename))
// 	}
// }

// // Unmarshaler is similar to Sub
// func (c *YamlConfig) Unmarshaler(prefix string, obj interface{}) error {
// 	sub, err := c.sub(prefix)
// 	if err != nil {
// 		return err
// 	}

// 	bytes, err := yaml.Marshal(sub)
// 	if err != nil {
// 		return err
// 	}
// 	return yaml.Unmarshal(bytes, obj)
// }

// func (c *YamlConfig) Sub(key string) (*YamlConfig, error) {
// 	sub, err := c.sub(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &YamlConfig{
// 		data: sub,
// 	}, nil
// }

// func (c *YamlConfig) sub(key string) (map[string]interface{}, error) {
// 	tmpData := c.data
// 	keys := strings.Split(key, ".")
// 	for idx, k := range keys {
// 		if v, ok := tmpData[k]; ok {
// 			switch v.(type) {
// 			case map[string]interface{}:
// 				{
// 					tmpData = v.(map[string]interface{})
// 					if idx == len(keys)-1 {
// 						return tmpData, nil
// 					}
// 				}
// 			case map[interface{}]interface{}:
// 				{
// 					tmpData = v.(map[interface{}]interface{})
// 					if idx == len(keys)-1 {
// 						return tmpData, nil
// 					}
// 				}
// 			default:
// 				return nil, errors.New(fmt.Sprintf("the key is invalid: %s", key))
// 			}
// 		}
// 	}

// 	return tmpData, nil
// }

// // Bool returns the boolean value for a given key.
// func (c *YamlConfig) Bool(key string) (bool, error) {
// 	v, err := c.getData(key)
// 	if err != nil {
// 		return false, err
// 	}
// 	return ParseBool(v)
// }

// // DefaultBool return the bool value if has no error
// // otherwise return the defaultVal
// func (c *YamlConfig) DefaultBool(key string, defaultVal bool) bool {
// 	v, err := c.Bool(key)
// 	if err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // Int returns the integer value for a given key.
// func (c *YamlConfig) Int(key string) (int, error) {
// 	if v, err := c.getData(key); err != nil {
// 		return 0, err
// 	} else if vv, ok := v.(int); ok {
// 		return vv, nil
// 	} else if vv, ok := v.(int64); ok {
// 		return int(vv), nil
// 	}
// 	return 0, errors.New("not int value")
// }

// // DefaultInt returns the integer value for a given key.
// // if err != nil return defaultVal
// func (c *YamlConfig) DefaultInt(key string, defaultVal int) int {
// 	v, err := c.Int(key)
// 	if err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // Int64 returns the int64 value for a given key.
// func (c *YamlConfig) Int64(key string) (int64, error) {
// 	if v, err := c.getData(key); err != nil {
// 		return 0, err
// 	} else if vv, ok := v.(int64); ok {
// 		return vv, nil
// 	}
// 	return 0, errors.New("not bool value")
// }

// // DefaultInt64 returns the int64 value for a given key.
// // if err != nil return defaultVal
// func (c *YamlConfig) DefaultInt64(key string, defaultVal int64) int64 {
// 	v, err := c.Int64(key)
// 	if err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // Float returns the float value for a given key.
// func (c *YamlConfig) Float(key string) (float64, error) {
// 	if v, err := c.getData(key); err != nil {
// 		return 0.0, err
// 	} else if vv, ok := v.(float64); ok {
// 		return vv, nil
// 	} else if vv, ok := v.(int); ok {
// 		return float64(vv), nil
// 	} else if vv, ok := v.(int64); ok {
// 		return float64(vv), nil
// 	}
// 	return 0.0, errors.New("not float64 value")
// }

// // DefaultFloat returns the float64 value for a given key.
// // if err != nil return defaultVal
// func (c *YamlConfig) DefaultFloat(key string, defaultVal float64) float64 {
// 	v, err := c.Float(key)
// 	if err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // String returns the string value for a given key.
// func (c *YamlConfig) String(key string) (string, error) {
// 	if v, err := c.getData(key); err == nil {
// 		if vv, ok := v.(string); ok {
// 			return vv, nil
// 		}
// 	}
// 	return "", nil
// }

// // DefaultString returns the string value for a given key.
// // if err != nil return defaultVal
// func (c *YamlConfig) DefaultString(key string, defaultVal string) string {
// 	v, err := c.String(key)
// 	if v == "" || err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // Strings returns the []string value for a given key.
// func (c *YamlConfig) Strings(key string) ([]string, error) {
// 	v, err := c.String(key)
// 	if v == "" || err != nil {
// 		return nil, err
// 	}
// 	return strings.Split(v, ";"), nil
// }

// // DefaultStrings returns the []string value for a given key.
// // if err != nil return defaultVal
// func (c *YamlConfig) DefaultStrings(key string, defaultVal []string) []string {
// 	v, err := c.Strings(key)
// 	if v == nil || err != nil {
// 		return defaultVal
// 	}
// 	return v
// }

// // GetSection returns map for the given section
// func (c *YamlConfig) GetSection(section string) (map[string]string, error) {

// 	if v, ok := c.data[section]; ok {
// 		return v.(map[string]string), nil
// 	}
// 	return nil, errors.New("not exist section")
// }

// // SaveConfigFile save the config into file
// func (c *YamlConfig) SaveConfigFile(filename string) (err error) {
// 	// Write configuration file by filename.
// 	configBytes,err := yaml.Marshal(c.data)
// 	if err != nil {
// 		gologs.GetLogger("").Error(err.Error())
// 		return
// 	}
// 	_,err =filetool.WriteBytesToFile(filename,configBytes)
// 	if err != nil {
// 		gologs.GetLogger("").Error(err.Error())
// 		return
// 	}	
// 	return err
// }

// // Set writes a new value for key.
// func (c *YamlConfig) Set(key, val string) error {
// 	c.Lock()
// 	defer c.Unlock()
// 	c.data[key] = val
// 	return nil
// }

// // DIY returns the raw value by a given key.
// func (c *YamlConfig) DIY(key string) (v interface{}, err error) {
// 	return c.getData(key)
// }

// func (c *YamlConfig) getData(key string) (interface{}, error) {

// 	if len(key) == 0 {
// 		return nil, errors.New("key is empty")
// 	}
// 	c.RLock()
// 	defer c.RUnlock()

// 	keys := strings.Split(c.key(key), ".")
// 	var tmpData map[interface {}]interface{}
// 	tmpData = c.data
// 	for idx, k := range keys {
// 		if v, ok := tmpData[k]; ok {
// 			switch v.(type) {
// 			case map[string]interface{}:
// 				{
// 					tmpData = v.(map[string]interface{})
// 					if idx == len(keys)-1 {
// 						return tmpData, nil
// 					}
// 				}
// 			case map[interface {}]interface {}:
// 				{
// 					tmpData = v.(map[interface {}]interface{})
// 					if idx == len(keys)-1 {
// 						return tmpData, nil
// 					}
// 				}
// 			default:
// 				{
// 					return v, nil
// 				}

// 			}
// 		}
// 	}
// 	return nil, fmt.Errorf("not exist key %q", key)
// }

// func (c *YamlConfig) key(key string) string {
// 	return key
// }

// func ParseBool(val interface{}) (value bool, err error) {
// 	if val != nil {
// 		switch v := val.(type) {
// 		case bool:
// 			return v, nil
// 		case string:
// 			switch v {
// 			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
// 				return true, nil
// 			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
// 				return false, nil
// 			}
// 		case int8, int32, int64:
// 			strV := fmt.Sprintf("%d", v)
// 			if strV == "1" {
// 				return true, nil
// 			} else if strV == "0" {
// 				return false, nil
// 			}
// 		case float64:
// 			if v == 1.0 {
// 				return true, nil
// 			} else if v == 0.0 {
// 				return false, nil
// 			}
// 		}
// 		return false, fmt.Errorf("parsing %q: invalid syntax", val)
// 	}
// 	return false, fmt.Errorf("parsing <nil>: invalid syntax")
// }