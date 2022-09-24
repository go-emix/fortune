package system

import (
	"errors"
	"github.com/go-emix/fortune/backend/pkg/casbin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"gorm.io/gorm"
)

func RoleList() (rs []Role, err error) {
	err = common.DB.Model(Role{}).Find(&rs).Error
	return
}

func UpdateRoleFeatures(roleId int, fids []int) (err error) {
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = common.DB.First(&Role{Id: roleId}).Error
		if err != nil {
			return
		}
		if len(fids) == 0 {
			err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleMenu{}).Error
			if err != nil {
				return
			}
			err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleFeature{}).Error
			return
		}
		fs := make([]Feature, 0)
		err = common.DB.Model(Feature{}).Where("id in (?)", fids).Find(&fs).Error
		if err != nil {
			return
		}
		if len(fs) == 0 {
			return nil
		}
		// 菜单去重
		mids := make([]int, 0)
		midmp := make(map[int]int)
		for _, f := range fs {
			_, ok := midmp[f.Menu]
			if !ok {
				midmp[f.Menu] = 0
				mids = append(mids, f.Menu)
			}
		}
		err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleMenu{}).Error
		if err != nil {
			return
		}
		for _, m := range mids {
			err = common.DB.Create(&RoleMenu{Role: roleId, Menu: m}).Error
			if err != nil {
				return
			}
		}
		err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleFeature{}).Error
		if err != nil {
			return
		}
		for _, f := range fids {
			err = common.DB.Create(&RoleFeature{Role: roleId, Feature: f}).Error
			if err != nil {
				return
			}
		}
		return
	})
	return
}

func SetMenuEntity(rm []Feature) []Feature {
	ln := len(rm)
	for i := 0; i < ln; i++ {
		rm[i].MenuEntity = Menu{Id: rm[i].Menu}
		common.DB.Find(&rm[i].MenuEntity)
	}
	return rm
}

func FeatureList() (rm []Feature, err error) {
	defer func() {
		rm = SetMenuEntity(rm)
	}()
	err = common.DB.Model(Feature{}).Find(&rm).Error
	return
}

func ApiList() (rs []Api, err error) {
	err = common.DB.Model(Api{}).Find(&rs).Error
	return
}

func UpdateRoleApis(roleId int, aids []int) (err error) {
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		role := Role{Id: roleId}
		err = common.DB.First(&role).Error
		if err != nil {
			return
		}
		if len(aids) == 0 {
			err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleApi{}).Error
			if err != nil {
				return err
			}
			_, err = casbin.Enforcer.RemoveFilteredPolicy(0, role.Name)
			return
		}
		as := make([]Api, 0)
		err = common.DB.Model(Api{}).Where("id in (?)", aids).Find(&as).Error
		if err != nil {
			return
		}
		if len(as) == 0 {
			return nil
		}
		err = common.DB.Where("role=?", roleId).Unscoped().Delete(RoleApi{}).Error
		if err != nil {
			return
		}
		rules := make([][]string, 0)
		for _, a := range as {
			err = common.DB.Create(&RoleApi{Role: roleId, Api: a.Id}).Error
			if err != nil {
				return
			}
			rules = append(rules, []string{role.Name, a.Path, a.Method})
		}
		if err != nil {
			return
		}
		_, err = casbin.Enforcer.RemoveFilteredPolicy(0, role.Name)
		if err != nil {
			return err
		}
		_, err = casbin.Enforcer.AddPolicies(rules)
		return
	})
	return
}

func NewRole(name string) (err error) {
	var c int64
	common.DB.Model(Role{}).Where("name=?", name).Count(&c)
	if c != 0 {
		err = errors.New("role name already exist")
		return
	}
	err = common.DB.Create(&Role{Name: name}).Error
	return
}

func DeleteRole(rid int) (err error) {
	var r = Role{Id: rid}
	err = common.DB.Find(&r).Error
	if err != nil {
		return
	}
	if r.Name == "root" {
		err = errors.New("root role not delete")
		return
	}
	var c int64
	common.DB.Model(AdminRole{}).Where("role=?", r.Id).Count(&c)
	if c != 0 {
		err = errors.New("admin bound role not delete")
		return
	}
	err = common.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = common.DB.Unscoped().Delete(&r).Error
		if err != nil {
			return
		}
		err = common.DB.Where("role=?", r.Id).Unscoped().Delete(RoleMenu{}).Error
		if err != nil {
			return
		}
		err = common.DB.Where("role=?", r.Id).Unscoped().Delete(RoleApi{}).Error
		if err != nil {
			return
		}
		err = common.DB.Where("role=?", r.Id).Unscoped().Delete(RoleFeature{}).Error
		if err != nil {
			return
		}
		_, err = casbin.Enforcer.RemoveFilteredPolicy(0, r.Name)
		return
	})
	return
}
