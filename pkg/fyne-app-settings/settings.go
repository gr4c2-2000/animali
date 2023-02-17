package fyneappsettings

import (
	"Animali/pkg/common"
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"fyne.io/fyne/v2"
)

type Settings interface {
}
type Bridge[T Settings] struct {
	FynePreferences fyne.Preferences
	AppSettings     *T
	CheckSum        string
}

func InitBridge[T Settings](as *T, fa fyne.Preferences) *Bridge[T] {

	fas := Bridge[T]{AppSettings: as, FynePreferences: fa}
	fas.Read()
	fmt.Println(fas.AppSettings)
	fas.CheckSum = common.AsSha256(as)
	return &fas
}
func (f *Bridge[T]) Watch(ctx context.Context, tick time.Duration) {
	ticker := time.NewTicker(tick)
	go func() {
		for {
			select {
			case <-ticker.C:
				f.Persist()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (f *Bridge[T]) Persist() {
	if f.verify() {
		return
	}
	f.actionByReflect(f.persistPreference)

}

func (f *Bridge[T]) actionByReflect(Action func(t string, key string, reflectField reflect.Value)) {
	val := reflect.ValueOf(f.AppSettings).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).CanInterface() {
			Action(val.Type().Field(i).Type.String(), val.Type().Field(i).Name, val.Field(i))
		}
	}
}

func (f *Bridge[T]) Read() {
	f.actionByReflect(f.readPreference)
}

func (f *Bridge[T]) persistPreference(t string, key string, reflectField reflect.Value) {
	switch t {
	case "string":
		f.FynePreferences.SetString(key, reflectField.String())
	case "int":
		f.FynePreferences.SetInt(key, int(reflectField.Int()))
	case "bool":
		f.FynePreferences.SetBool(key, reflectField.Bool())
	case "float":
		f.FynePreferences.SetFloat(key, reflectField.Float())
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPORTED DATA TYPE"))
	}
}

func (f *Bridge[T]) readPreference(t string, key string, reflectField reflect.Value) {
	switch t {
	case "string":
		reflectField.SetString(f.FynePreferences.String(key))
	case "int":
		reflectField.SetInt(int64(f.FynePreferences.Int(key)))
	case "bool":
		reflectField.SetBool(f.FynePreferences.Bool(key))
	case "float":
		reflectField.SetFloat(f.FynePreferences.Float(key))
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPORTED DATA TYPE"))
	}
}

func (f *Bridge[T]) verify() bool {
	newSum := common.AsSha256(f.AppSettings)
	if newSum == f.CheckSum {
		return true
	}
	f.CheckSum = newSum
	return false
}
