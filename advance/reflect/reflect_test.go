package reflect

import (
	"github.com/wonderivan/logger"
	"reflect"
	"testing"
)

func CheckType(value interface{}) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		logger.Info("float")
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		logger.Info("int")
	default:
		logger.Error("unkown")
	}
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
	logger.Alert(newVal)
}
func TestType(t *testing.T) {
	CheckType(16)
}
func TestInvoke(t *testing.T) {
	e := &Employee{"111", "ff", 33}
	if eid, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed")
	} else {
		t.Log(eid.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(22)})
}
