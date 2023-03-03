package fyneappsettings

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/gr4c2-2000/animali/pkg/common"

	"fyne.io/fyne/v2"
)

const STRING = "string"
const INT = "int"
const BOOL = "bool"
const FLOAT = "float"

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
	case STRING:
		f.FynePreferences.SetString(key, reflectField.String())
	case INT:
		f.FynePreferences.SetInt(key, int(reflectField.Int()))
	case BOOL:
		f.FynePreferences.SetBool(key, reflectField.Bool())
	case FLOAT:
		f.FynePreferences.SetFloat(key, reflectField.Float())
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPPORTED DATA TYPE"))
	}
}

func (f *Bridge[T]) readPreference(t string, key string, reflectField reflect.Value) {
	switch t {
	case STRING:
		reflectField.SetString(f.FynePreferences.String(key))
	case INT:
		reflectField.SetInt(int64(f.FynePreferences.Int(key)))
	case BOOL:
		reflectField.SetBool(f.FynePreferences.Bool(key))
	case FLOAT:
		reflectField.SetFloat(f.FynePreferences.Float(key))
	default:
		fyne.LogError("Not suported config data type", errors.New("NOT SUPPORTED DATA TYPE"))
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
