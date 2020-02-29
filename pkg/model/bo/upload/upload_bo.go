// 此文件由bo_builder工具自动生成，不要修改
package upload

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/510909033/bgf_bo"
	"github.com/510909033/bgf_database"
	"github.com/510909033/bgf_util/sizedwaitgroup"
)

type UploadBO struct {
	bo   *bgf_bo.BO  `ignore:"" structs:"-"`
	hook bgf_bo.Hook `ignore:"" structs:"-"`

	//NewBO时传入的readOnly参数和forceLoad参数，很多php老代码的特殊场景需要
	readOnly  bool `ignore:"" structs:"-"`
	forceLoad bool `ignore:"" structs:"-"`

	// int(10) unsigned
	Id uint32 `model:"pk;column_name:id" json:"id" structs:"id"`

	// varchar(1000)
	Name string `model:"column_name:name;column_default:" json:"name" structs:"name"`

	// varchar(1000)
	Ext string `model:"column_name:ext;column_default:" json:"ext" structs:"ext"`

	// int(10) unsigned
	CreateTs uint32 `model:"column_name:create_ts;column_default:0" json:"create_ts" structs:"create_ts"`

	// int(10) unsigned
	UpdateTs uint32 `model:"column_name:update_ts;column_default:0" json:"update_ts" structs:"update_ts"`

	// varchar(1000)
	Extra string `model:"column_name:extra;column_default:" json:"extra" structs:"extra"`

	//扩展字段，这些字段不是数据库字段
	ext *UploadBOExt `ignore:"" structs:"-"`
}

// 注册
func init() {
	bgf_database.RegisterModel(&UploadBO{})
}

func NewUploadBO(aBO *UploadBO, readOnly bool) (*UploadBO, error) {
	aBO.ext = new(UploadBOExt)
	var i interface{} = aBO
	var hook bgf_bo.Hook
	var ok bool
	hook, ok = i.(bgf_bo.Hook)
	if ok {
		aBO.hook = hook
	} else {
		aBO.hook = bgf_bo.EmptyHook
	}
	bo := bgf_bo.NewBO(aBO, readOnly)
	bo.SetRealTableName(aBO.TableName())
	aBO.bo = bo
	aBO.readOnly = readOnly
	aBO.forceLoad = false
	aBO.hook.BeforeLoad()
	err := bo.Load(false)
	if err == bgf_bo.ERR_NO_DATA_IN_DB {
		return aBO, nil
	}
	if err != nil {
		return nil, err
	}
	aBO.hook.AfterLoad()
	return aBO, nil
}

func NewUploadBOWithForceLoad(aBO *UploadBO, readOnly bool) (*UploadBO, error) {
	aBO.ext = new(UploadBOExt)
	var i interface{} = aBO
	var hook bgf_bo.Hook
	var ok bool
	hook, ok = i.(bgf_bo.Hook)
	if ok {
		aBO.hook = hook
	} else {
		aBO.hook = bgf_bo.EmptyHook
	}
	bo := bgf_bo.NewBO(aBO, readOnly)
	bo.SetRealTableName(aBO.TableName())
	aBO.bo = bo
	aBO.readOnly = readOnly
	aBO.forceLoad = true
	aBO.hook.BeforeLoad()
	err := bo.Load(true)
	if err == bgf_bo.ERR_NO_DATA_IN_DB {
		return aBO, nil
	}
	if err != nil {
		return nil, err
	}
	aBO.hook.AfterLoad()
	return aBO, nil
}

// LoadMultiUploadBO方法提供并发加载UploadBO
// 此方法返回的结果没有顺序, 如果需要顺序，可以使用LoadUploadBOList方法
func LoadMultiUploadBO(m map[string]*UploadBO) (map[string]*UploadBO, error) {
	var ret = make(map[string]*UploadBO)
	if len(m) <= 0 {
		return ret, nil
	}
	var mu sync.Mutex
	swg := sizedwaitgroup.New(10)
	var err error = nil
	for key, bo := range m {
		swg.Add()
		go func(key string, bo *UploadBO) {
			defer func() {
				swg.Done()

				if e := recover(); e != nil {
					mu.Lock()
					if err == nil {
						err = errors.New(fmt.Sprintf("%v", e))
					}
					ret[key] = nil
					mu.Unlock()
				}
			}()

			var tmpErr error
			bo, tmpErr = NewUploadBO(bo, true)

			mu.Lock()
			if err == nil && tmpErr != nil {
				err = tmpErr
			}
			ret[key] = bo
			mu.Unlock()
		}(key, bo)
	}

	swg.Wait()

	return ret, err
}

// LoadUploadBOList方法提供并发加载UploadBO
func LoadUploadBOList(list []*UploadBO) ([]*UploadBO, error) {
	newList := make([]*UploadBO, 0)
	if len(list) <= 0 {
		return newList, nil
	}
	m := make(map[string]*UploadBO)
	for i, bo := range list {
		iStr := strconv.Itoa(i)
		m[iStr] = bo
	}
	m, err := LoadMultiUploadBO(m)
	if err != nil {
		return nil, err
	}

	//排序
	for i, _ := range list {
		iStr := strconv.Itoa(i)
		newList = append(newList, m[iStr])
	}

	return newList, nil
}

// LoadMultiUploadBO2功能与LoadMultiUploadBO类似，但可以设定并发取bo的并发数量，而且返回每一个key对应的错误，使用方需要判断每一个错误，如果为nil，说明取得的对应bo数据是正确的。
func LoadMultiUploadBO2(m map[string]*UploadBO, concurrentNum int) (map[string]*UploadBO, map[string]error) {
	var ret = make(map[string]*UploadBO)
	var errs = make(map[string]error)

	if len(m) <= 0 {
		return ret, nil
	}
	var mu sync.Mutex
	swg := sizedwaitgroup.New(concurrentNum)
	for key, bo := range m {
		swg.Add()
		go func(key string, bo *UploadBO) {
			defer func() {
				swg.Done()

				if e := recover(); e != nil {
					mu.Lock()
					err := errors.New(fmt.Sprintf("%v", e))
					ret[key] = nil
					errs[key] = err
					mu.Unlock()
				}
			}()

			var tmpErr error
			bo, tmpErr = NewUploadBO(bo, true)

			mu.Lock()
			if tmpErr != nil {
				ret[key] = nil
				errs[key] = tmpErr
			} else {
				ret[key] = bo
				errs[key] = nil
			}
			mu.Unlock()
		}(key, bo)
	}

	swg.Wait()

	return ret, errs
}

// LoadUploadBOList2功能与LoadUploadBOList类似，但可以设定并发取bo的并发数量，而且返回每一个index对应的错误，使用方需要判断每一个错误，如果未nil，说明取得的对应bo数据是正确的。
func LoadUploadBOList2(list []*UploadBO, concurrentNum int) ([]*UploadBO, []error) {
	newList := make([]*UploadBO, 0)
	errs := make([]error, 0)

	if len(list) <= 0 {
		return newList, nil
	}
	m := make(map[string]*UploadBO)
	for i, bo := range list {
		iStr := strconv.Itoa(i)
		m[iStr] = bo
	}
	m, mErr := LoadMultiUploadBO2(m, concurrentNum)

	//排序
	for i, _ := range list {
		iStr := strconv.Itoa(i)
		newList = append(newList, m[iStr])
		errs = append(errs, mErr[iStr])
	}

	return newList, errs
}

// Save方法根据IsNewRow方法分别调用Insert或者Update方法
// 返回值:
//   第一个参数为ture，说明是底层调用的是Insert方法，否则调用的是Update方法
//   第二个和第三个返回值是Insert/Update方法的返回值
func (bo *UploadBO) Save() (bool, int64, error) {
	b, i, err := bo.bo.Save()
	return b, i, err
}

// Insert方法的第一个参数在IsAutoIncrement方法为true时，返回的新插入行的自增id的值
func (bo *UploadBO) Insert() (int64, error) {
	bo.hook.BeforeInsert()
	i, err := bo.bo.Insert()
	bo.hook.AfterInsert()
	return i, err
}

// Update方法的第一个参数返回更新了几行，在没有错误时总是返回1
func (bo *UploadBO) Update() (int64, error) {
	bo.hook.BeforeUpdate()
	i, err := bo.bo.Update()
	bo.hook.AfterUpdate()
	return i, err
}

// Delete方法的第一个参数返回删除了几行，在没有错误时总是返回1
func (bo *UploadBO) Delete() (int64, error) {
	bo.hook.BeforeDelete()
	i, err := bo.bo.Delete()
	bo.hook.AfterDelete()
	return i, err
}

func (bo *UploadBO) IsNewRow() bool {
	return bo.bo.IsNewRow()
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

func (bo *UploadBO) GetExt() *UploadBOExt {
	return bo.ext
}

// 这里是meta信息，放到const()中，为了方便IDE收起来
const (
// meta:{"Name":"UploadBO","NameSelector":"UploadSelector","FileName":"upload_bo","FileNameSelector":"upload_selector","TableName":"Upload","RealTableName":"Upload","DBName":"menu","DoCache":true,"IsCompatiblePHP":false,"IsCompatiblePHPOld":false,"IsAutoIncrement":true,"TagJsonFlag":true,"Columns":[{"Name":"Id","DataType":"uint32","Tag":"`model:\"pk;column_name:id\" json:\"id\" structs:\"id\"`","Comment":"// int(10) unsigned"},{"Name":"Name","DataType":"string","Tag":"`model:\"column_name:name;column_default:\" json:\"name\" structs:\"name\"`","Comment":"// varchar(1000)"},{"Name":"CreateTs","DataType":"uint32","Tag":"`model:\"column_name:create_ts;column_default:0\" json:\"create_ts\" structs:\"create_ts\"`","Comment":"// int(10) unsigned"},{"Name":"UpdateTs","DataType":"uint32","Tag":"`model:\"column_name:update_ts;column_default:0\" json:\"update_ts\" structs:\"update_ts\"`","Comment":"// int(10) unsigned"},{"Name":"Extra","DataType":"string","Tag":"`model:\"column_name:extra;column_default:\" json:\"extra\" structs:\"extra\"`","Comment":"// varchar(1000)"}],"Ignore":"`ignore:\"\" structs:\"-\"`","FullPackage":"gopkg.babytree-inc.com/app/go_upload/pkg/model/bo/upload","HookFlag":false,"Meta":"","IsSharding":false,"PackageName":"upload"}
)

//ALTER TABLE `Upload`
//ADD COLUMN `ext`  varchar(1000) NOT NULL DEFAULT '' AFTER `name`;
