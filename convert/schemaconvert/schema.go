package schemaconvert

import (
	types "github.com/rancher/mapper"
	"github.com/rancher/mapper/convert"
)

func InternalToInternal(from interface{}, fromSchema *types.Schema, toSchema *types.Schema, target interface{}) error {
	data, err := convert.EncodeToMap(from)
	if err != nil {
		return err
	}
	fromSchema.Mapper.FromInternal(data)
	if err := toSchema.Mapper.ToInternal(data); err != nil {
		return err
	}
	return convert.ToObj(data, target)
}

func ToInternal(from interface{}, schema *types.Schema, target interface{}) error {
	data, err := convert.EncodeToMap(from)
	if err != nil {
		return err
	}
	if err := schema.Mapper.ToInternal(data); err != nil {
		return err
	}
	return convert.ToObj(data, target)
}
