// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicinalProductPharmaceutical is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductPharmaceutical
type MedicinalProductPharmaceutical struct {
	Id                    *string                                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                                               `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	AdministrableDoseForm CodeableConcept                                       `bson:"administrableDoseForm" json:"administrableDoseForm"`
	UnitOfPresentation    *CodeableConcept                                      `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	Ingredient            []Reference                                           `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Device                []Reference                                           `bson:"device,omitempty" json:"device,omitempty"`
	Characteristics       []MedicinalProductPharmaceuticalCharacteristics       `bson:"characteristics,omitempty" json:"characteristics,omitempty"`
	RouteOfAdministration []MedicinalProductPharmaceuticalRouteOfAdministration `bson:"routeOfAdministration" json:"routeOfAdministration"`
}
type MedicinalProductPharmaceuticalCharacteristics struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `bson:"code" json:"code"`
	Status            *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
}
type MedicinalProductPharmaceuticalRouteOfAdministration struct {
	Id                        *string                                                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                                                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `bson:"code" json:"code"`
	FirstDose                 *Quantity                                                          `bson:"firstDose,omitempty" json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `bson:"maxSingleDose,omitempty" json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `bson:"maxDosePerDay,omitempty" json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `bson:"maxDosePerTreatmentPeriod,omitempty" json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `bson:"maxTreatmentPeriod,omitempty" json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies `bson:"targetSpecies,omitempty" json:"targetSpecies,omitempty"`
}
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                                                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `bson:"code" json:"code"`
	WithdrawalPeriod  []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod `bson:"withdrawalPeriod,omitempty" json:"withdrawalPeriod,omitempty"`
}
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                    *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `bson:"tissue" json:"tissue"`
	Value                 Quantity        `bson:"value" json:"value"`
	SupportingInformation *string         `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}
type OtherMedicinalProductPharmaceutical MedicinalProductPharmaceutical

// MarshalJSON marshals the given MedicinalProductPharmaceutical as JSON into a byte slice
func (r MedicinalProductPharmaceutical) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPharmaceutical
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPharmaceutical: OtherMedicinalProductPharmaceutical(r),
		ResourceType:                        "MedicinalProductPharmaceutical",
	})
}

// UnmarshalMedicinalProductPharmaceutical unmarshals a MedicinalProductPharmaceutical.
func UnmarshalMedicinalProductPharmaceutical(b []byte) (MedicinalProductPharmaceutical, error) {
	var medicinalProductPharmaceutical MedicinalProductPharmaceutical
	if err := json.Unmarshal(b, &medicinalProductPharmaceutical); err != nil {
		return medicinalProductPharmaceutical, err
	}
	return medicinalProductPharmaceutical, nil
}
