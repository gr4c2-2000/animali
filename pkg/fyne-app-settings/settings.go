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

type FyneAppSettings[T comparable] struct {
	FyneApp     fyne.App
	AppSettings *T
	CheckSum    string
}

func InitFyneAppSettings[T comparable](as *T, fa fyne.App) *FyneAppSettings[T] {
	fas := FyneAppSettings[T]{AppSettings: as, FyneApp: fa}
	fas.Read()
	fmt.Println(fas.AppSettings)
	fas.CheckSum = common.AsSha256(as)
	return &fas
}
func (f *FyneAppSettings[T]) Watch(ctx context.Context, tick time.Duration) {
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

func (f *FyneAppSettings[T]) Persist() {
	if f.verify() {
		return
	}
	f.actionByReflect(f.persistPreference)

}

func (f *FyneAppSettings[T]) actionByReflect(Action func(t string, key string, reflectField reflect.Value)) {
	val := reflect.ValueOf(f.AppSettings).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).CanInterface() {
			Action(val.Type().Field(i).Type.String(), val.Type().Field(i).Name, val.Field(i))
		}
	}
}

func (f *FyneAppSettings[T]) Read() {
	f.actionByReflect(f.readPreference)
}

func (f *FyneAppSettings[T]) persistPreference(t string, key string, reflectField reflect.Value) {
	switch t {
	case "string":
		f.FyneApp.Preferences().SetString(key, reflectField.String())
	case "int":
		f.FyneApp.Preferences().SetInt(key, int(reflectField.Int()))
	case "bool":
		f.FyneApp.Preferences().SetBool(key, reflectField.Bool())
	case "float":
		f.FyneApp.Preferences().SetFloat(key, reflectField.Float())
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPORTED DATA TYPE"))
	}
}

func (f *FyneAppSettings[T]) readPreference(t string, key string, reflectField reflect.Value) {
	switch t {
	case "string":
		reflectField.SetString(f.FyneApp.Preferences().String(key))
	case "int":
		reflectField.SetInt(int64(f.FyneApp.Preferences().Int(key)))
	case "bool":
		reflectField.SetBool(f.FyneApp.Preferences().Bool(key))
	case "float":
		reflectField.SetFloat(f.FyneApp.Preferences().Float(key))
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPORTED DATA TYPE"))
	}
}

func (f *FyneAppSettings[T]) verify() bool {
	newSum := common.AsSha256(f.AppSettings)
	if newSum == f.CheckSum {
		return true
	}
	f.CheckSum = newSum
	return false
}
