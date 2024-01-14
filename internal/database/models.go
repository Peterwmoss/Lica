package database

import (
	"github.com/uptrace/bun"
)

type (
	Product struct {
		bun.BaseModel `bun:"table:products,alias:p"`

		Id         int64     `bun:",pk,autoincrement"`
		Name       string    `bun:",unique:products_unique,notnull"`
		CategoryId int64     `bun:",notnull"`
		Category   *Category `bun:"rel:has-one,join:category_id=id"`
		UserId     *int64    `bun:",unique:products_unique"`
		User       *User     `bun:"rel:has-one,join:user_id=id"`
		IsCustom   bool      `bun:",notnull,default:false"`
	}

	List struct {
		bun.BaseModel `bun:"table:lists,alias:l"`

		Id        int64       `bun:",pk,autoincrement"`
		Name      string      `bun:",notnull"`
		UserId    int64       `bun:",notnull"`
		User      *User       `bun:"rel:has-one,join:user_id=id"`
		ListItems []*ListItem `bun:"rel:has-many,join:id=list_id"`
	}

	ListItem struct {
		bun.BaseModel `bun:"table:list_items,alias:li"`

		Id        int64    `bun:",pk,autoincrement"`
		Unit      *string  ``
		Amount    float32    `bun:",notnull,default:1.0"`
		ListId    int64    `bun:",notnull"`
		List      *List    `bun:"rel:has-one,join:list_id=id"`
		ProductId int64    `bun:",notnull"`
		Product   *Product `bun:"rel:has-one,join:product_id=id"`
	}

	User struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		Id    int64  `bun:",pk,autoincrement"`
		Email string `bun:",unique,notnull"`
	}

	Category struct {
		bun.BaseModel `bun:"table:categories,alias:c"`

		Id       int64           `bun:",pk,autoincrement"`
		Name     string          `bun:",unique:user_category-unique,notnull"`
		UserId   *int64          `bun:",unique:user_category_unique"`
		User     *User           `bun:"rel:has-one,join:user_id=id"`
		IsCustom bool            `bun:",notnull,default:false"`
    Orders   []CategoryOrder `bun:"rel:has-many,join:id=category_id,join:user_id=user_id"`
	}

	CategoryOrder struct {
		bun.BaseModel `bun:"table:category_orders,alias:co"`

		CategoryId int64     `bun:",pk"`
		Category   *Category `bun:"rel:has-one,join:category_id=id"`
		UserId     int64     `bun:",pk"`
		User       *User     `bun:"rel:has-one,join:user_id=id"`
		Order      int16     `bun:",notnull"`
	}
)
