package types

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Body_hasOptionalFields(t *testing.T) {
	t.Parallel()

	var body Body

	assert.False(t, body.hasOptionalFields(), "body has no optional field set")

	v := reflect.ValueOf(&body).Elem()
	typ := reflect.TypeOf(body)
	for i := 0; i < v.NumField(); i++ {
		fieldType := typ.Field(i)
		rlpTag := fieldType.Tag.Get("rlp")
		rlpTagFields := strings.Split(rlpTag, ",")
		if !slices.Contains(rlpTagFields, "optional") {
			continue
		}

		field := v.Field(i)
		before := field.Interface()
		switch field.Kind() {
		case reflect.Slice:
			slice := reflect.MakeSlice(field.Type(), 1, 1)
			field.Set(slice)
		default:
			t.Errorf("unexpected field kind %q for field %q", field.Kind(), fieldType.Name)
		}

		assert.Truef(t, body.hasOptionalFields(), "body has the optional field %q set", fieldType.Name)
		field.Set(reflect.ValueOf(before)) // reset the field
	}
}
