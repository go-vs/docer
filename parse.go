package docer

import (
	"fmt"
	"reflect"
)

type ConfigOption func(p *Parser)

func TagAsName(tag string) ConfigOption {
	return func(p *Parser) {
		p.tagAsName = tag
	}
}

type Parser struct {
	memType   map[string]*Type
	types     []*Type
	tagAsName string
}

func NewParser(ops ...ConfigOption) *Parser {
	p := &Parser{
		memType: make(map[string]*Type),
		types:   make([]*Type, 0),
	}
	for _, op := range ops {
		op(p)
	}
	return p
}

func (p *Parser) parse(data any) []*Type {
	p.parseStruct(reflect.TypeOf(data))
	return p.types
}

func (p *Parser) parseField(t reflect.Type) []*Field {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	res := make([]*Field, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			fields := p.parseField(f.Type)
			res = append(res, fields...)
			continue
		}
		tag := f.Tag.Get(p.tagAsName)
		if tag == "" {
			continue
		}
		field := &Field{
			Name:        tag,
			Type:        "",
			Required:    false,
			Ref:         "",
			Description: "",
		}
		subT := f.Type
		k := subT.Kind()
		if k == reflect.Ptr {
			subT = subT.Elem()
			k = subT.Kind()
		}
		switch k {
		case reflect.Struct:
			sub := p.parseStruct(subT)
			field.Type = "object"
			field.Ref = sub.Name
		case reflect.Slice, reflect.Array:
			sub := p.parseStruct(subT.Elem())
			if sub != nil {
				field.Ref = sub.Name
				field.Type = "array of object"
			} else {
				field.Type = "array of " + subT.Elem().String()
			}
		default:
			field.Type = subT.String()
		}
		res = append(res, field)
	}
	return res
}

func (p *Parser) parseStruct(data reflect.Type) *Type {
	if data.Kind() == reflect.Ptr {
		data = data.Elem()
	}
	if data.Kind() != reflect.Struct {
		return nil
	}
	if t, ok := p.memType[data.Name()]; ok {
		return t
	}
	fmt.Println("type", data.Name(), data.Kind())
	t := &Type{
		Name:        data.Name(),
		DisplayName: "",
		Description: "",
		Fields:      make([]*Field, 0),
	}
	p.memType[data.Name()] = t
	p.types = append(p.types, t)
	t.Fields = p.parseField(data)
	return t
}
