// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Organization{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Organization{}) fail: %s", err)
	}
}

func Test_organizationQuery(t *testing.T) {
	organization := newOrganization(db)
	organization = *organization.As(organization.TableName())
	_do := organization.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(organization.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <organizations> fail:", err)
		return
	}

	_, ok := organization.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from organization success")
	}

	err = _do.Create(&model.Organization{})
	if err != nil {
		t.Error("create item in table <organizations> fail:", err)
	}

	err = _do.Save(&model.Organization{})
	if err != nil {
		t.Error("create item in table <organizations> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Organization{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <organizations> fail:", err)
	}

	_, err = _do.Select(organization.ALL).Take()
	if err != nil {
		t.Error("Take() on table <organizations> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <organizations> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <organizations> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <organizations> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Organization{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <organizations> fail:", err)
	}

	_, err = _do.Select(organization.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <organizations> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <organizations> fail:", err)
	}

	_, err = _do.Select(organization.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <organizations> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <organizations> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <organizations> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <organizations> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Organization{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <organizations> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <organizations> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <organizations> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <organizations> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <organizations> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <organizations> fail:", err)
	}
}
