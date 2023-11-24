package item_repository

import (
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
)

type Repository interface {
	GetItemsByCodes(itemCodes []any) ([]entity.Item, errs.Error)
}
