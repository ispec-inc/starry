package service

import (
	"context"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/model"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/query"
)

// CancelOrganization はクリニックをキャンセルするサービス
type CancelOrganization struct {
	Q *query.Query
}

// Do はクリニックをキャンセルする処理を行う
// active_organizationsのレコードを削除し、cancel_organizationsにレコードを追加することでキャンセルを表現する
func (c CancelOrganization) Do(ctx context.Context, id string) error {
	return c.Q.Transaction(func(tx *query.Query) error {
		organization, err := tx.Organization.WithContext(ctx).Where(c.Q.Organization.ID.Eq(id)).First()
		if err != nil {
			return err
		}

		if err := tx.Organization.Active.WithContext(ctx).Model(organization).Clear(); err != nil {
			return err
		}

		if err := tx.Organization.Cancel.WithContext(ctx).Model(organization).Append(&model.CanceledOrganization{}); err != nil {
			return err
		}

		return nil
	})
}
