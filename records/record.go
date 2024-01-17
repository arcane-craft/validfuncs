package records

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/arcane-craft/validfuncs"
	"github.com/arcane-craft/validfuncs/generic"
	"github.com/arcane-craft/validfuncs/utils"
)

type Naming func(reflect.StructField) string

func NamingByAttr() Naming {
	return func(field reflect.StructField) string {
		return field.Name
	}
}

func NamingByTag(key string, postProcs ...func(string) string) Naming {
	return func(field reflect.StructField) string {
		content := field.Tag.Get(key)
		for _, p := range postProcs {
			content = p(content)
		}
		return content
	}
}

func NamingByJSONTag() Naming {
	return NamingByTag("json", func(s string) string {
		parts := strings.Split(s, ",")
		if len(parts) <= 0 {
			return ""
		}
		return strings.TrimSpace(parts[0])
	})
}

func NamingByStr(name string) Naming {
	return func(reflect.StructField) string {
		return name
	}
}

func Field[R, F any](naming Naming, record *R, field *F, expr validfuncs.Expr[F], left ...validfuncs.Expr[F]) validfuncs.Expr[R] {
	if len(left) > 0 {
		expr = generic.All(expr, left[0], left[1:]...)
	}
	offset := uintptr(unsafe.Pointer(field)) - uintptr(unsafe.Pointer(record))
	recordType := reflect.TypeFor[R]()
	var fieldName string
	for i := range recordType.NumField() {
		structField := recordType.Field(i)
		if structField.Offset == offset {
			fieldName = naming(structField)
			break
		}
	}
	return validfuncs.NewExpr(
		func() (validfuncs.Instruction[R], error) {
			instrs, err := expr.Interpret()
			if err != nil {
				return nil, err
			}
			return validfuncs.NewInstr[R](
				func(ctx *validfuncs.Context, v R) bool {
					return validfuncs.WithSelector(fieldName, instrs).Exec(ctx, *(*F)(unsafe.Add(unsafe.Pointer(&v), offset)))
				},
				func() string {
					return instrs.String()
				},
			), nil
		}, func() string {
			return utils.StringifyExpr("Field", fieldName, expr)
		},
	)
}

func StructField[R, F any](record *R, field *F, expr validfuncs.Expr[F], left ...validfuncs.Expr[F]) validfuncs.Expr[R] {
	return Field(NamingByAttr(), record, field, expr, left...)
}

func JSONField[R, F any](record *R, field *F, expr validfuncs.Expr[F], left ...validfuncs.Expr[F]) validfuncs.Expr[R] {
	return Field(NamingByJSONTag(), record, field, expr, left...)
}

func CustomField[R, F any](name string, record *R, field *F, expr validfuncs.Expr[F], left ...validfuncs.Expr[F]) validfuncs.Expr[R] {
	return Field(NamingByStr(name), record, field, expr, left...)
}
