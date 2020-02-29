// 此文件由bo_builder工具自动生成，不要修改
package mock

import "github.com/510909033/bgf_database"

var _ bgf_database.NullString

type UploadBO struct {
	Id       uint32 `model:"pk;column_name:id" json:"id" structs:"id"`
	Name     string `model:"column_name:name;column_default:" json:"name" structs:"name"`
	CreateTs uint32 `model:"column_name:create_ts;column_default:0" json:"create_ts" structs:"create_ts"`
	UpdateTs uint32 `model:"column_name:update_ts;column_default:0" json:"update_ts" structs:"update_ts"`
	Extra    string `model:"column_name:extra;column_default:" json:"extra" structs:"extra"`
}

func NewUploadBO(aBO *UploadBO, readOnly bool) (*UploadBO, error) {
	panic("NewUploadBO还未实现")
}

func NewUploadBOWithForceLoad(aBO *UploadBO, readOnly bool) (*UploadBO, error) {
	panic("NewUploadBO还未实现")
}

func (bo *UploadBO) Save() (bool, int64, error) {
	panic("UploadBO的Save方法暂时还未实现")
}

func (bo *UploadBO) Insert() (int64, error) {
	panic("UploadBO的Insert方法暂时还未实现")
}

func (bo *UploadBO) Update() (int64, error) {
	panic("UploadBO的Update方法暂时还未实现")
}

func (bo *UploadBO) Delete() (int64, error) {
	panic("UploadBO的Delete方法暂时还未实现")
}

func (bo *UploadBO) IsNewRow() bool {
	panic("UploadBO的IsNewRow方法暂时还未实现")
}

func (bo *UploadBO) TableName() string {
	return "Upload"
}

func (bo *UploadBO) DBName() string {
	return "menu"
}

func (bo *UploadBO) DoCache() bool {
	return true
}

func (bo *UploadBO) IsCompatiblePHP() bool {
	return false
}

func (bo *UploadBO) IsCompatiblePHPOld() bool {
	return false
}

func (bo *UploadBO) IsAutoIncrement() bool {
	return true
}
