package models

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

func (w *WrapUpCodeMapping) validateMappingItemsEnum(path, location string, value string) *errors.Validation {
	return validate.EnumCase(path, location, value, wrapUpCodeMappingDefaultSetItemsEnum, true)
}
