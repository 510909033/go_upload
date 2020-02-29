// 此文件由bo_builder工具自动生成，不要修改
package bo

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/510909033/bgf_bo"
	"github.com/510909033/bgf_database"
	"github.com/510909033/bgf_util/sizedwaitgroup"
)

type menuBO struct {
	bo   *bgf_bo.BO  `ignore:"" structs:"-"`
	hook bgf_bo.Hook `ignore:"" structs:"-"`

	//NewBO时传入的readOnly参数和forceLoad参数，很多php老代码的特殊场景需要
	readOnly  bool `ignore:"" structs:"-"`
	forceLoad bool `ignore:"" structs:"-"`

	// int(10) unsigned
	Id uint32 `model:"pk;column_name:id" json:"id" structs:"id"`

	// int(10) unsigned
	CategoryId uint32 `model:"column_name:category_id;column_default:0" json:"category_id" structs:"category_id"`

	// int(10) unsigned
	UserId uint32 `model:"column_name:user_id;column_default:0" json:"user_id" structs:"user_id"`

	// varchar(1000)
	Title string `model:"column_name:title;column_default:" json:"title" structs:"title"`

	// int(10) unsigned
	CreateTs uint32 `model:"column_name:create_ts;column_default:0" json:"create_ts" structs:"create_ts"`

	// int(10) unsigned
	UpdateTs uint32 `model:"column_name:update_ts;column_default:0" json:"update_ts" structs:"update_ts"`

	// varchar(1000)
	Extra string `model:"column_name:extra;column_default:" json:"extra" structs:"extra"`

	//扩展字段，这些字段不是数据库字段
	ext *menuBOExt `ignore:"" structs:"-"`
}

// 注册
func init() {
	bgf_database.RegisterModel(&menuBO{})
}

func NewmenuBO(aBO *menuBO, readOnly bool) (*menuBO, error) {
	aBO.ext = new(menuBOExt)
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

func NewmenuBOWithForceLoad(aBO *menuBO, readOnly bool) (*menuBO, error) {
	aBO.ext = new(menuBOExt)
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

// LoadMultimenuBO方法提供并发加载menuBO
// 此方法返回的结果没有顺序, 如果需要顺序，可以使用LoadmenuBOList方法
func LoadMultimenuBO(m map[string]*menuBO) (map[string]*menuBO, error) {
	var ret = make(map[string]*menuBO)
	if len(m) <= 0 {
		return ret, nil
	}
	var mu sync.Mutex
	swg := sizedwaitgroup.New(10)
	var err error = nil
	for key, bo := range m {
		swg.Add()
		go func(key string, bo *menuBO) {
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
			bo, tmpErr = NewmenuBO(bo, true)

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

// LoadmenuBOList方法提供并发加载menuBO
func LoadmenuBOList(list []*menuBO) ([]*menuBO, error) {
	newList := make([]*menuBO, 0)
	if len(list) <= 0 {
		return newList, nil
	}
	m := make(map[string]*menuBO)
	for i, bo := range list {
		iStr := strconv.Itoa(i)
		m[iStr] = bo
	}
	m, err := LoadMultimenuBO(m)
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

// LoadMultimenuBO2功能与LoadMultimenuBO类似，但可以设定并发取bo的并发数量，而且返回每一个key对应的错误，使用方需要判断每一个错误，如果为nil，说明取得的对应bo数据是正确的。
func LoadMultimenuBO2(m map[string]*menuBO, concurrentNum int) (map[string]*menuBO, map[string]error) {
	var ret = make(map[string]*menuBO)
	var errs = make(map[string]error)

	if len(m) <= 0 {
		return ret, nil
	}
	var mu sync.Mutex
	swg := sizedwaitgroup.New(concurrentNum)
	for key, bo := range m {
		swg.Add()
		go func(key string, bo *menuBO) {
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
			bo, tmpErr = NewmenuBO(bo, true)

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

// LoadmenuBOList2功能与LoadmenuBOList类似，但可以设定并发取bo的并发数量，而且返回每一个index对应的错误，使用方需要判断每一个错误，如果未nil，说明取得的对应bo数据是正确的。
func LoadmenuBOList2(list []*menuBO, concurrentNum int) ([]*menuBO, []error) {
	newList := make([]*menuBO, 0)
	errs := make([]error, 0)

	if len(list) <= 0 {
		return newList, nil
	}
	m := make(map[string]*menuBO)
	for i, bo := range list {
		iStr := strconv.Itoa(i)
		m[iStr] = bo
	}
	m, mErr := LoadMultimenuBO2(m, concurrentNum)

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
func (bo *menuBO) Save() (bool, int64, error) {
	b, i, err := bo.bo.Save()
	return b, i, err
}

// Insert方法的第一个参数在IsAutoIncrement方法为true时，返回的新插入行的自增id的值
func (bo *menuBO) Insert() (int64, error) {
	bo.hook.BeforeInsert()
	i, err := bo.bo.Insert()
	bo.hook.AfterInsert()
	return i, err
}

// Update方法的第一个参数返回更新了几行，在没有错误时总是返回1
func (bo *menuBO) Update() (int64, error) {
	bo.hook.BeforeUpdate()
	i, err := bo.bo.Update()
	bo.hook.AfterUpdate()
	return i, err
}

// Delete方法的第一个参数返回删除了几行，在没有错误时总是返回1
func (bo *menuBO) Delete() (int64, error) {
	bo.hook.BeforeDelete()
	i, err := bo.bo.Delete()
	bo.hook.AfterDelete()
	return i, err
}

func (bo *menuBO) IsNewRow() bool {
	return bo.bo.IsNewRow()
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

func (bo *menuBO) GetExt() *menuBOExt {
	return bo.ext
}

// 这里是meta信息，放到const()中，为了方便IDE收起来
const (
// meta:{"Name":"menuBO","NameSelector":"menuSelector","FileName":"menu_bo","FileNameSelector":"menu_selector","TableName":"menu","RealTableName":"menu","DBName":"menu","DoCache":true,"IsCompatiblePHP":false,"IsCompatiblePHPOld":false,"IsAutoIncrement":true,"TagJsonFlag":true,"Columns":[{"Name":"Id","DataType":"uint32","Tag":"`model:\"pk;column_name:id\" json:\"id\" structs:\"id\"`","Comment":"// int(10) unsigned"},{"Name":"CategoryId","DataType":"uint32","Tag":"`model:\"column_name:category_id;column_default:0\" json:\"category_id\" structs:\"category_id\"`","Comment":"// int(10) unsigned"},{"Name":"UserId","DataType":"uint32","Tag":"`model:\"column_name:user_id;column_default:0\" json:\"user_id\" structs:\"user_id\"`","Comment":"// int(10) unsigned"},{"Name":"Title","DataType":"string","Tag":"`model:\"column_name:title;column_default:\" json:\"title\" structs:\"title\"`","Comment":"// varchar(1000)"},{"Name":"CreateTs","DataType":"uint32","Tag":"`model:\"column_name:create_ts;column_default:0\" json:\"create_ts\" structs:\"create_ts\"`","Comment":"// int(10) unsigned"},{"Name":"UpdateTs","DataType":"uint32","Tag":"`model:\"column_name:update_ts;column_default:0\" json:\"update_ts\" structs:\"update_ts\"`","Comment":"// int(10) unsigned"},{"Name":"Extra","DataType":"string","Tag":"`model:\"column_name:extra;column_default:\" json:\"extra\" structs:\"extra\"`","Comment":"// varchar(1000)"}],"Ignore":"`ignore:\"\" structs:\"-\"`","FullPackage":"baotian0506.com/app/go_upload/pkg/mode/bo","HookFlag":false,"Meta":"","IsSharding":false,"PackageName":"bo"}
)
