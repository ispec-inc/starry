package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func genOrganization(g *gen.Generator) {
	g.GenerateModel(
		"organizations",
		gen.FieldNewTag("id", `validate:"required,ulid"`),
		gen.FieldNewTag("organizational_detail", `validate:"required"`),
		gen.FieldRelate(
			field.HasOne,
			"OrganizationDetail",
			g.GenerateModel(
				"organization_details",
				gen.FieldNewTag("name", `validate:"required"`),
				gen.FieldNewTag("name_kana", `validate:"required,kana"`),
				gen.FieldNewTag("contact", `validate:"required,email"`),
			),
			&field.RelateConfig{}),
		gen.FieldRelate(
			field.HasOne,
			"PartOf",
			g.GenerateModel("part_of_organizations"),
			&field.RelateConfig{
				RelatePointer: true,
			}),
		gen.FieldRelate(
			field.HasOne,
			"Active",
			g.GenerateModel("active_organizations"),
			&field.RelateConfig{
				RelatePointer: true,
			}),
		gen.FieldRelate(
			field.HasOne,
			"Cancel",
			g.GenerateModel("canceled_organizations"),
			&field.RelateConfig{
				RelatePointer: true,
			}),
	)
}
