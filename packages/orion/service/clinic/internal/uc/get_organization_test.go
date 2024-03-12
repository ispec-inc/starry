package uc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ispec-inc/starry/orion/app/test"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/adapter"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/registry"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/uc"
)

/*
test.PrepareDB(t, seeds) の内部で使用している、txdbの仕様上、トランザクションを使用する関数では、
テスト関数内で実行されるテストケースごとにトランザクションがコミットされるため、
テストケースごとにtest.PrepareDB(t, seeds)を実行する必要がある。
テスト関数単位でtest.PrepareDBを実行した場合、"SAVEPOINT tx_5 does not exist"のエラーが発生する。

また、テストケースごとにtest.PrepareDBを呼び出した場合、
同じIDをのデータを複数の`go routine`から参照するため、デッドロックの可能性がある。

対応方法
- 直列実行を行う: t.Parallel()を使用しない
- t.Parallel()を使用して、テストケースを並列実行する場合は、テストケース事にテストデータの生成を行う

以下に直列実行と並列実行のテストコードを記述する

*/

func Test_GetOrganization_Serial(t *testing.T) { // nolint
	// テストデータの準備
	orgs, err := organizationFactory(3)
	if err != nil {
		t.Fatal(err)
	}
	oEnts := adapter.OrganizationListToEntityList(orgs)
	seeds := []any{}
	seeds = append(seeds, &oEnts)

	type (
		give struct {
			input uc.GetOrganizationInput
		}
		want struct {
			output uc.GetOrganizationOutput
			err    error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] IDを指定して組織を取得できる",
			give: give{
				input: uc.GetOrganizationInput{
					ID: orgs[0].ID,
				},
			},
			want: want{
				output: uc.GetOrganizationOutput{
					Organization: orgs[0],
				},
			},
		},
		{
			name: "[NG] 存在しないIDを指定して組織を取得できない",
			give: give{
				input: uc.GetOrganizationInput{
					ID: organization.ID("invalid_id"),
				},
			},
			want: want{
				err: organization.ErrNotFound,
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests { // nolint
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, cleanupDB := test.PrepareDB(t, seeds)
			t.Cleanup(cleanupDB)

			rgst := registry.NewTest(ctx, db)

			got, err := uc.NewGetOrganization(rgst).Do(ctx, tt.give.input)
			if diff := cmp.Diff(tt.want.output, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
			if !errors.Is(err, tt.want.err) {
				t.Fatalf("expected %v to be %v", err, tt.want.err)
			}
		})
	}

}

func Test_GetOrganization_Parallel(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			input uc.GetOrganizationInput
		}
		want struct {
			output uc.GetOrganizationOutput
			err    error
		}
		testCase struct {
			seeds []any
			give  give
			want  want
		}
		caseFunc func() testCase
	)

	tests := []struct {
		name     string
		caseFunc caseFunc
	}{
		{
			name: "[OK] IDを指定して組織を取得できる",
			caseFunc: func() testCase {
				orgs, err := organizationFactory(3)
				if err != nil {
					t.Fatal(err)
				}
				oEnts := adapter.OrganizationListToEntityList(orgs)
				seeds := []any{}
				seeds = append(seeds, &oEnts)

				return testCase{
					seeds: seeds,
					give: give{
						input: uc.GetOrganizationInput{
							ID: orgs[0].ID,
						},
					},
					want: want{
						output: uc.GetOrganizationOutput{
							Organization: orgs[0],
						},
					},
				}
			},
		},
		{
			name: "[NG] 存在しないIDを指定して組織を取得できない",
			caseFunc: func() testCase {
				orgs, err := organizationFactory(3)
				if err != nil {
					t.Fatal(err)
				}
				oEnts := adapter.OrganizationListToEntityList(orgs)
				seeds := []any{}
				seeds = append(seeds, &oEnts)

				return testCase{
					seeds: seeds,
					give: give{
						input: uc.GetOrganizationInput{
							ID: organization.ID("invalid_id"),
						},
					},
					want: want{
						err: organization.ErrNotFound,
					},
				}
			},
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tc := tt.caseFunc()
			seeds, give, want := tc.seeds, tc.give, tc.want

			db, cleanupDB := test.PrepareDB(t, seeds)
			t.Cleanup(cleanupDB)

			rgst := registry.NewTest(ctx, db)

			got, err := uc.NewGetOrganization(rgst).Do(ctx, give.input)
			if diff := cmp.Diff(want.output, got); diff != "" {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}
			if !errors.Is(err, want.err) {
				t.Fatalf("expected %v to be %v", err, want.err)
			}
		})
	}
}

func organizationFactory(count int) (organization.List, error) {
	orgs := make(organization.List, 0, count)
	for i := 0; i < count; i++ {
		name, err := organization.NewName("鈴木歯科医院", "suzuki")
		if err != nil {
			return organization.List{}, err
		}

		typ, err := organization.NewType(1)
		if err != nil {
			return organization.List{}, err
		}

		phone, err := organization.NewPhoneNumber("09012345678")
		if err != nil {
			return organization.List{}, err
		}

		org, err := organization.New(name, typ, phone)
		if err != nil {
			return organization.List{}, err
		}

		orgs = append(orgs, org)

	}

	return orgs, nil
}
