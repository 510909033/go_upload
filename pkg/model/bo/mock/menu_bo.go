// 此文件由bo_builder工具自动生成，不要修改
package mock

import "github.com/510909033/bgf_database"

var _ bgf_database.NullString

type menuBO struct {
	Id         uint32 `model:"pk;column_name:id" json:"id" structs:"id"`
	CategoryId uint32 `model:"column_name:category_id;column_default:0" json:"category_id" structs:"category_id"`
	UserId     uint32 `model:"column_name:user_id;column_default:0" json:"user_id" structs:"user_id"`
	Title      string `model:"column_name:title;column_default:" json:"title" structs:"title"`
	CreateTs   uint32 `model:"column_name:create_ts;column_default:0" json:"create_ts" structs:"create_ts"`
	UpdateTs   uint32 `model:"column_name:update_ts;column_default:0" json:"update_ts" structs:"update_ts"`
	Extra      string `model:"column_name:extra;column_default:" json:"extra" structs:"extra"`
}

func NewmenuBO(aBO *menuBO, readOnly bool) (*menuBO, error) {
	panic("NewmenuBO还未实现")
}

func NewmenuBOWithForceLoad(aBO *menuBO, readOnly bool) (*menuBO, error) {
	panic("NewmenuBO还未实现")
}

func (bo *menuBO) Save() (bool, int64, error) {
	panic("menuBO的Save方法暂时还未实现")
}

func (bo *menuBO) Insert() (int64, error) {
	panic("menuBO的Insert方法暂时还未实现")
}

func (bo *menuBO) Update() (int64, error) {
	panic("menuBO的Update方法暂时还未实现")
}

func (bo *menuBO) Delete() (int64, error) {
	panic("menuBO的Delete方法暂时还未实现")
}

func (bo *menuBO) IsNewRow() bool {
	panic("menuBO的IsNewRow方法暂时还未实现")
}

func (bo *menuBO) TableName() string {
	return "menu"
}

func (bo *menuBO) DBName() string {
	return "menu"
}

func (bo *menuBO) DoCache() bool {
	return true
}

func (bo *menuBO) IsCompatiblePHP() bool {
	return false
}

func (bo *menuBO) IsCompatiblePHPOld() bool {
	return false
}

func (bo *menuBO) IsAutoIncrement() bool {
	return true
}
