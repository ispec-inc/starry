package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func genClinic(g *gen.Generator) {

	g.GenerateModel(
		"clinics",
		gen.FieldRelate(
			field.HasOne,
			"ClinicDetail",
			g.GenerateModel("clinic_details"),
			&field.RelateConfig{}),
	)
}

func genClinicalCase(g *gen.Generator) {
	g.GenerateModel(
		"clinical_cases",
		gen.FieldRelate(
			field.HasOne,
			"AttendingDoctor",
			g.GenerateModel("clinical_case_attending_doctors"),
			&field.RelateConfig{
				GORMTag: "foreignKey:clinical_case_attending_doctors",
			}),
		gen.FieldRelate(
			field.HasOne,
			"Patient",
			g.GenerateModel("clinical_case_patients"),
			&field.RelateConfig{
				GORMTag: "foreignKey:clinical_case_patients",
			}),
	)
}
