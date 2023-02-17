package fyneappsettings

import (
	"context"
	"testing"
	"time"
)

func TestFyneAppSettings_Bridge(t *testing.T) {
	type TestSettings struct {
		Field1 string
		Field2 int
		Field3 bool
	}
	Test := []struct {
		name   string
		result TestSettings
		Start  map[string]interface{}
	}{
		{
			"Test1",
			TestSettings{"ĄĄĄŚŚŚÐÐÐÐÐ  ÐÐÐÐÐ", 0, true},
			map[string]interface{}{"Field1": "", "Field2": 0, "Field3": false},
		},
		{
			"Test2",
			TestSettings{"test1", 1, true},
			map[string]interface{}{"Field1": "cdcdcdc", "Field2": 22, "Field3": false},
		},
	}
	for _, test := range Test {
		fa := MockPreferences{mockMap: test.Start}
		new := TestSettings{}
		testWachSettings := InitBridge(&new, fa)
		testWachSettings.Read()

		if new.Field1 != test.Start["Field1"].(string) {
			t.Errorf("inccorect field Field1 value")

		}
		if new.Field2 != test.Start["Field2"].(int) {
			t.Errorf("inccorect field Field2 value")
		}
		if new.Field3 != test.Start["Field3"].(bool) {
			t.Errorf("inccorect field Field3 value")

		}
		testWachSettings.Watch(context.TODO(), 1000*time.Microsecond)
		new.Field1 = test.result.Field1
		new.Field2 = test.result.Field2
		new.Field3 = test.result.Field3
		time.Sleep(5000 * time.Microsecond)
		if new.Field1 != test.Start["Field1"].(string) {
			t.Errorf("inccorect field Field1 value")
		}
		if new.Field2 != test.Start["Field2"].(int) {
			t.Errorf("inccorect field Field2 value")
		}
		if new.Field3 != test.Start["Field3"].(bool) {
			t.Errorf("inccorect field Field3 value")
		}
	}
}

type MockPreferences struct {
	mockMap map[string]interface{}
}

func (mp MockPreferences) SetBool(key string, value bool)     { mp.mockMap[key] = value }
func (mp MockPreferences) SetInt(key string, value int)       { mp.mockMap[key] = value }
func (mp MockPreferences) SetString(key string, value string) { mp.mockMap[key] = value }
func (mp MockPreferences) SetFloat(key string, value float64) { mp.mockMap[key] = value }

func (mp MockPreferences) Bool(key string) bool {
	if val, ok := mp.mockMap[key]; ok {
		return val.(bool)
	}
	return false
}

func (mp MockPreferences) BoolWithFallback(key string, fallback bool) bool {
	if val, ok := mp.mockMap[key]; ok {
		return val.(bool)
	}
	return fallback
}
func (mp MockPreferences) Float(key string) float64 {
	if val, ok := mp.mockMap[key]; ok {
		return val.(float64)
	}
	return 0
}
func (mp MockPreferences) FloatWithFallback(key string, fallback float64) float64 {
	if val, ok := mp.mockMap[key]; ok {
		return val.(float64)
	}
	return fallback
}
func (mp MockPreferences) Int(key string) int {
	if val, ok := mp.mockMap[key]; ok {
		return val.(int)
	}
	return 0
}
func (mp MockPreferences) IntWithFallback(key string, fallback int) int {
	if val, ok := mp.mockMap[key]; ok {
		return val.(int)
	}
	return fallback
}
func (mp MockPreferences) String(key string) string {
	if val, ok := mp.mockMap[key]; ok {
		return val.(string)
	}
	return ""
}
func (mp MockPreferences) StringWithFallback(key, fallback string) string {
	if val, ok := mp.mockMap[key]; ok {
		return val.(string)
	}
	return fallback
}
func (mp MockPreferences) AddChangeListener(func())  {}
func (mp MockPreferences) ChangeListeners() []func() { return []func(){} }
func (mp MockPreferences) RemoveValue(key string)    {}
