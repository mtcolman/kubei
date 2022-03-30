// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetPackagesParams creates a new GetPackagesParams object
// with the default values initialized.
func NewGetPackagesParams() GetPackagesParams {

	var (
		// initialize parameters with default values

		sortDirDefault = string("ASC")
	)

	return GetPackagesParams{
		SortDir: &sortDirDefault,
	}
}

// GetPackagesParams contains all the bound params for the get packages operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPackages
type GetPackagesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*application ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	ApplicationID *string
	/*application resource ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	ApplicationResourceID *string
	/*greater than or equal
	  In: query
	*/
	ApplicationResourcesGte *int64
	/*
	  In: query
	*/
	ApplicationResourcesIsNot []int64
	/*
	  In: query
	*/
	ApplicationResourcesIs []int64
	/*less than or equal
	  In: query
	*/
	ApplicationResourcesLte *int64
	/*greater than or equal
	  In: query
	*/
	ApplicationsGte *int64
	/*
	  In: query
	*/
	ApplicationsIsNot []int64
	/*
	  In: query
	*/
	ApplicationsIs []int64
	/*less than or equal
	  In: query
	*/
	ApplicationsLte *int64
	/*current runtime scan system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	CurrentRuntimeScan *bool
	/*
	  In: query
	*/
	LanguageContains []string
	/*
	  In: query
	*/
	LanguageEnd *string
	/*
	  In: query
	*/
	LanguageIsNot []string
	/*
	  In: query
	*/
	LanguageIs []string
	/*
	  In: query
	*/
	LanguageStart *string
	/*
	  In: query
	*/
	LicenseContains []string
	/*
	  In: query
	*/
	LicenseEnd *string
	/*
	  In: query
	*/
	LicenseIsNot []string
	/*
	  In: query
	*/
	LicenseIs []string
	/*
	  In: query
	*/
	LicenseStart *string
	/*
	  In: query
	*/
	PackageNameContains []string
	/*
	  In: query
	*/
	PackageNameEnd *string
	/*
	  In: query
	*/
	PackageNameIsNot []string
	/*
	  In: query
	*/
	PackageNameIs []string
	/*
	  In: query
	*/
	PackageNameStart *string
	/*
	  In: query
	*/
	PackageVersionContains []string
	/*
	  In: query
	*/
	PackageVersionEnd *string
	/*
	  In: query
	*/
	PackageVersionIsNot []string
	/*
	  In: query
	*/
	PackageVersionIs []string
	/*
	  In: query
	*/
	PackageVersionStart *string
	/*Page number of the query
	  Required: true
	  In: query
	*/
	Page int64
	/*Maximum items to return
	  Required: true
	  Maximum: 50
	  Minimum: 1
	  In: query
	*/
	PageSize int64
	/*Sorting direction
	  In: query
	  Default: "ASC"
	*/
	SortDir *string
	/*Sort key
	  Required: true
	  In: query
	*/
	SortKey string
	/*
	  In: query
	*/
	VulnerabilitySeverityGte *string
	/*
	  In: query
	*/
	VulnerabilitySeverityLte *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPackagesParams() beforehand.
func (o *GetPackagesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplicationID, qhkApplicationID, _ := qs.GetOK("applicationID")
	if err := o.bindApplicationID(qApplicationID, qhkApplicationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourceID, qhkApplicationResourceID, _ := qs.GetOK("applicationResourceID")
	if err := o.bindApplicationResourceID(qApplicationResourceID, qhkApplicationResourceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesGte, qhkApplicationResourcesGte, _ := qs.GetOK("applicationResources[gte]")
	if err := o.bindApplicationResourcesGte(qApplicationResourcesGte, qhkApplicationResourcesGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesIsNot, qhkApplicationResourcesIsNot, _ := qs.GetOK("applicationResources[isNot]")
	if err := o.bindApplicationResourcesIsNot(qApplicationResourcesIsNot, qhkApplicationResourcesIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesIs, qhkApplicationResourcesIs, _ := qs.GetOK("applicationResources[is]")
	if err := o.bindApplicationResourcesIs(qApplicationResourcesIs, qhkApplicationResourcesIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationResourcesLte, qhkApplicationResourcesLte, _ := qs.GetOK("applicationResources[lte]")
	if err := o.bindApplicationResourcesLte(qApplicationResourcesLte, qhkApplicationResourcesLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsGte, qhkApplicationsGte, _ := qs.GetOK("applications[gte]")
	if err := o.bindApplicationsGte(qApplicationsGte, qhkApplicationsGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsIsNot, qhkApplicationsIsNot, _ := qs.GetOK("applications[isNot]")
	if err := o.bindApplicationsIsNot(qApplicationsIsNot, qhkApplicationsIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsIs, qhkApplicationsIs, _ := qs.GetOK("applications[is]")
	if err := o.bindApplicationsIs(qApplicationsIs, qhkApplicationsIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationsLte, qhkApplicationsLte, _ := qs.GetOK("applications[lte]")
	if err := o.bindApplicationsLte(qApplicationsLte, qhkApplicationsLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCurrentRuntimeScan, qhkCurrentRuntimeScan, _ := qs.GetOK("currentRuntimeScan")
	if err := o.bindCurrentRuntimeScan(qCurrentRuntimeScan, qhkCurrentRuntimeScan, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguageContains, qhkLanguageContains, _ := qs.GetOK("language[contains]")
	if err := o.bindLanguageContains(qLanguageContains, qhkLanguageContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguageEnd, qhkLanguageEnd, _ := qs.GetOK("language[end]")
	if err := o.bindLanguageEnd(qLanguageEnd, qhkLanguageEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguageIsNot, qhkLanguageIsNot, _ := qs.GetOK("language[isNot]")
	if err := o.bindLanguageIsNot(qLanguageIsNot, qhkLanguageIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguageIs, qhkLanguageIs, _ := qs.GetOK("language[is]")
	if err := o.bindLanguageIs(qLanguageIs, qhkLanguageIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qLanguageStart, qhkLanguageStart, _ := qs.GetOK("language[start]")
	if err := o.bindLanguageStart(qLanguageStart, qhkLanguageStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qLicenseContains, qhkLicenseContains, _ := qs.GetOK("license[contains]")
	if err := o.bindLicenseContains(qLicenseContains, qhkLicenseContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qLicenseEnd, qhkLicenseEnd, _ := qs.GetOK("license[end]")
	if err := o.bindLicenseEnd(qLicenseEnd, qhkLicenseEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qLicenseIsNot, qhkLicenseIsNot, _ := qs.GetOK("license[isNot]")
	if err := o.bindLicenseIsNot(qLicenseIsNot, qhkLicenseIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qLicenseIs, qhkLicenseIs, _ := qs.GetOK("license[is]")
	if err := o.bindLicenseIs(qLicenseIs, qhkLicenseIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qLicenseStart, qhkLicenseStart, _ := qs.GetOK("license[start]")
	if err := o.bindLicenseStart(qLicenseStart, qhkLicenseStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageNameContains, qhkPackageNameContains, _ := qs.GetOK("packageName[contains]")
	if err := o.bindPackageNameContains(qPackageNameContains, qhkPackageNameContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageNameEnd, qhkPackageNameEnd, _ := qs.GetOK("packageName[end]")
	if err := o.bindPackageNameEnd(qPackageNameEnd, qhkPackageNameEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageNameIsNot, qhkPackageNameIsNot, _ := qs.GetOK("packageName[isNot]")
	if err := o.bindPackageNameIsNot(qPackageNameIsNot, qhkPackageNameIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageNameIs, qhkPackageNameIs, _ := qs.GetOK("packageName[is]")
	if err := o.bindPackageNameIs(qPackageNameIs, qhkPackageNameIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageNameStart, qhkPackageNameStart, _ := qs.GetOK("packageName[start]")
	if err := o.bindPackageNameStart(qPackageNameStart, qhkPackageNameStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageVersionContains, qhkPackageVersionContains, _ := qs.GetOK("packageVersion[contains]")
	if err := o.bindPackageVersionContains(qPackageVersionContains, qhkPackageVersionContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageVersionEnd, qhkPackageVersionEnd, _ := qs.GetOK("packageVersion[end]")
	if err := o.bindPackageVersionEnd(qPackageVersionEnd, qhkPackageVersionEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageVersionIsNot, qhkPackageVersionIsNot, _ := qs.GetOK("packageVersion[isNot]")
	if err := o.bindPackageVersionIsNot(qPackageVersionIsNot, qhkPackageVersionIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageVersionIs, qhkPackageVersionIs, _ := qs.GetOK("packageVersion[is]")
	if err := o.bindPackageVersionIs(qPackageVersionIs, qhkPackageVersionIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageVersionStart, qhkPackageVersionStart, _ := qs.GetOK("packageVersion[start]")
	if err := o.bindPackageVersionStart(qPackageVersionStart, qhkPackageVersionStart, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortDir, qhkSortDir, _ := qs.GetOK("sortDir")
	if err := o.bindSortDir(qSortDir, qhkSortDir, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortKey, qhkSortKey, _ := qs.GetOK("sortKey")
	if err := o.bindSortKey(qSortKey, qhkSortKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qVulnerabilitySeverityGte, qhkVulnerabilitySeverityGte, _ := qs.GetOK("vulnerabilitySeverity[gte]")
	if err := o.bindVulnerabilitySeverityGte(qVulnerabilitySeverityGte, qhkVulnerabilitySeverityGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qVulnerabilitySeverityLte, qhkVulnerabilitySeverityLte, _ := qs.GetOK("vulnerabilitySeverity[lte]")
	if err := o.bindVulnerabilitySeverityLte(qVulnerabilitySeverityLte, qhkVulnerabilitySeverityLte, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindApplicationID binds and validates parameter ApplicationID from query.
func (o *GetPackagesParams) bindApplicationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationID = &raw

	return nil
}

// bindApplicationResourceID binds and validates parameter ApplicationResourceID from query.
func (o *GetPackagesParams) bindApplicationResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationResourceID = &raw

	return nil
}

// bindApplicationResourcesGte binds and validates parameter ApplicationResourcesGte from query.
func (o *GetPackagesParams) bindApplicationResourcesGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applicationResources[gte]", "query", "int64", raw)
	}
	o.ApplicationResourcesGte = &value

	return nil
}

// bindApplicationResourcesIsNot binds and validates array parameter ApplicationResourcesIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindApplicationResourcesIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationResourcesIsNot string
	if len(rawData) > 0 {
		qvApplicationResourcesIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationResourcesIsNotIC := swag.SplitByFormat(qvApplicationResourcesIsNot, "")
	if len(applicationResourcesIsNotIC) == 0 {
		return nil
	}

	var applicationResourcesIsNotIR []int64
	for i, applicationResourcesIsNotIV := range applicationResourcesIsNotIC {
		applicationResourcesIsNotI, err := swag.ConvertInt64(applicationResourcesIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applicationResources[isNot]", i), "query", "int64", applicationResourcesIsNotI)
		}

		applicationResourcesIsNotIR = append(applicationResourcesIsNotIR, applicationResourcesIsNotI)
	}

	o.ApplicationResourcesIsNot = applicationResourcesIsNotIR

	return nil
}

// bindApplicationResourcesIs binds and validates array parameter ApplicationResourcesIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindApplicationResourcesIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationResourcesIs string
	if len(rawData) > 0 {
		qvApplicationResourcesIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationResourcesIsIC := swag.SplitByFormat(qvApplicationResourcesIs, "")
	if len(applicationResourcesIsIC) == 0 {
		return nil
	}

	var applicationResourcesIsIR []int64
	for i, applicationResourcesIsIV := range applicationResourcesIsIC {
		applicationResourcesIsI, err := swag.ConvertInt64(applicationResourcesIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applicationResources[is]", i), "query", "int64", applicationResourcesIsI)
		}

		applicationResourcesIsIR = append(applicationResourcesIsIR, applicationResourcesIsI)
	}

	o.ApplicationResourcesIs = applicationResourcesIsIR

	return nil
}

// bindApplicationResourcesLte binds and validates parameter ApplicationResourcesLte from query.
func (o *GetPackagesParams) bindApplicationResourcesLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applicationResources[lte]", "query", "int64", raw)
	}
	o.ApplicationResourcesLte = &value

	return nil
}

// bindApplicationsGte binds and validates parameter ApplicationsGte from query.
func (o *GetPackagesParams) bindApplicationsGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applications[gte]", "query", "int64", raw)
	}
	o.ApplicationsGte = &value

	return nil
}

// bindApplicationsIsNot binds and validates array parameter ApplicationsIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindApplicationsIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationsIsNot string
	if len(rawData) > 0 {
		qvApplicationsIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationsIsNotIC := swag.SplitByFormat(qvApplicationsIsNot, "")
	if len(applicationsIsNotIC) == 0 {
		return nil
	}

	var applicationsIsNotIR []int64
	for i, applicationsIsNotIV := range applicationsIsNotIC {
		applicationsIsNotI, err := swag.ConvertInt64(applicationsIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applications[isNot]", i), "query", "int64", applicationsIsNotI)
		}

		applicationsIsNotIR = append(applicationsIsNotIR, applicationsIsNotI)
	}

	o.ApplicationsIsNot = applicationsIsNotIR

	return nil
}

// bindApplicationsIs binds and validates array parameter ApplicationsIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindApplicationsIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationsIs string
	if len(rawData) > 0 {
		qvApplicationsIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationsIsIC := swag.SplitByFormat(qvApplicationsIs, "")
	if len(applicationsIsIC) == 0 {
		return nil
	}

	var applicationsIsIR []int64
	for i, applicationsIsIV := range applicationsIsIC {
		applicationsIsI, err := swag.ConvertInt64(applicationsIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "applications[is]", i), "query", "int64", applicationsIsI)
		}

		applicationsIsIR = append(applicationsIsIR, applicationsIsI)
	}

	o.ApplicationsIs = applicationsIsIR

	return nil
}

// bindApplicationsLte binds and validates parameter ApplicationsLte from query.
func (o *GetPackagesParams) bindApplicationsLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("applications[lte]", "query", "int64", raw)
	}
	o.ApplicationsLte = &value

	return nil
}

// bindCurrentRuntimeScan binds and validates parameter CurrentRuntimeScan from query.
func (o *GetPackagesParams) bindCurrentRuntimeScan(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("currentRuntimeScan", "query", "bool", raw)
	}
	o.CurrentRuntimeScan = &value

	return nil
}

// bindLanguageContains binds and validates array parameter LanguageContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLanguageContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLanguageContains string
	if len(rawData) > 0 {
		qvLanguageContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	languageContainsIC := swag.SplitByFormat(qvLanguageContains, "")
	if len(languageContainsIC) == 0 {
		return nil
	}

	var languageContainsIR []string
	for _, languageContainsIV := range languageContainsIC {
		languageContainsI := languageContainsIV

		languageContainsIR = append(languageContainsIR, languageContainsI)
	}

	o.LanguageContains = languageContainsIR

	return nil
}

// bindLanguageEnd binds and validates parameter LanguageEnd from query.
func (o *GetPackagesParams) bindLanguageEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LanguageEnd = &raw

	return nil
}

// bindLanguageIsNot binds and validates array parameter LanguageIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLanguageIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLanguageIsNot string
	if len(rawData) > 0 {
		qvLanguageIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	languageIsNotIC := swag.SplitByFormat(qvLanguageIsNot, "")
	if len(languageIsNotIC) == 0 {
		return nil
	}

	var languageIsNotIR []string
	for _, languageIsNotIV := range languageIsNotIC {
		languageIsNotI := languageIsNotIV

		languageIsNotIR = append(languageIsNotIR, languageIsNotI)
	}

	o.LanguageIsNot = languageIsNotIR

	return nil
}

// bindLanguageIs binds and validates array parameter LanguageIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLanguageIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLanguageIs string
	if len(rawData) > 0 {
		qvLanguageIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	languageIsIC := swag.SplitByFormat(qvLanguageIs, "")
	if len(languageIsIC) == 0 {
		return nil
	}

	var languageIsIR []string
	for _, languageIsIV := range languageIsIC {
		languageIsI := languageIsIV

		languageIsIR = append(languageIsIR, languageIsI)
	}

	o.LanguageIs = languageIsIR

	return nil
}

// bindLanguageStart binds and validates parameter LanguageStart from query.
func (o *GetPackagesParams) bindLanguageStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LanguageStart = &raw

	return nil
}

// bindLicenseContains binds and validates array parameter LicenseContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLicenseContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLicenseContains string
	if len(rawData) > 0 {
		qvLicenseContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	licenseContainsIC := swag.SplitByFormat(qvLicenseContains, "")
	if len(licenseContainsIC) == 0 {
		return nil
	}

	var licenseContainsIR []string
	for _, licenseContainsIV := range licenseContainsIC {
		licenseContainsI := licenseContainsIV

		licenseContainsIR = append(licenseContainsIR, licenseContainsI)
	}

	o.LicenseContains = licenseContainsIR

	return nil
}

// bindLicenseEnd binds and validates parameter LicenseEnd from query.
func (o *GetPackagesParams) bindLicenseEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LicenseEnd = &raw

	return nil
}

// bindLicenseIsNot binds and validates array parameter LicenseIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLicenseIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLicenseIsNot string
	if len(rawData) > 0 {
		qvLicenseIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	licenseIsNotIC := swag.SplitByFormat(qvLicenseIsNot, "")
	if len(licenseIsNotIC) == 0 {
		return nil
	}

	var licenseIsNotIR []string
	for _, licenseIsNotIV := range licenseIsNotIC {
		licenseIsNotI := licenseIsNotIV

		licenseIsNotIR = append(licenseIsNotIR, licenseIsNotI)
	}

	o.LicenseIsNot = licenseIsNotIR

	return nil
}

// bindLicenseIs binds and validates array parameter LicenseIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindLicenseIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvLicenseIs string
	if len(rawData) > 0 {
		qvLicenseIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	licenseIsIC := swag.SplitByFormat(qvLicenseIs, "")
	if len(licenseIsIC) == 0 {
		return nil
	}

	var licenseIsIR []string
	for _, licenseIsIV := range licenseIsIC {
		licenseIsI := licenseIsIV

		licenseIsIR = append(licenseIsIR, licenseIsI)
	}

	o.LicenseIs = licenseIsIR

	return nil
}

// bindLicenseStart binds and validates parameter LicenseStart from query.
func (o *GetPackagesParams) bindLicenseStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LicenseStart = &raw

	return nil
}

// bindPackageNameContains binds and validates array parameter PackageNameContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageNameContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageNameContains string
	if len(rawData) > 0 {
		qvPackageNameContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageNameContainsIC := swag.SplitByFormat(qvPackageNameContains, "")
	if len(packageNameContainsIC) == 0 {
		return nil
	}

	var packageNameContainsIR []string
	for _, packageNameContainsIV := range packageNameContainsIC {
		packageNameContainsI := packageNameContainsIV

		packageNameContainsIR = append(packageNameContainsIR, packageNameContainsI)
	}

	o.PackageNameContains = packageNameContainsIR

	return nil
}

// bindPackageNameEnd binds and validates parameter PackageNameEnd from query.
func (o *GetPackagesParams) bindPackageNameEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageNameEnd = &raw

	return nil
}

// bindPackageNameIsNot binds and validates array parameter PackageNameIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageNameIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageNameIsNot string
	if len(rawData) > 0 {
		qvPackageNameIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageNameIsNotIC := swag.SplitByFormat(qvPackageNameIsNot, "")
	if len(packageNameIsNotIC) == 0 {
		return nil
	}

	var packageNameIsNotIR []string
	for _, packageNameIsNotIV := range packageNameIsNotIC {
		packageNameIsNotI := packageNameIsNotIV

		packageNameIsNotIR = append(packageNameIsNotIR, packageNameIsNotI)
	}

	o.PackageNameIsNot = packageNameIsNotIR

	return nil
}

// bindPackageNameIs binds and validates array parameter PackageNameIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageNameIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageNameIs string
	if len(rawData) > 0 {
		qvPackageNameIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageNameIsIC := swag.SplitByFormat(qvPackageNameIs, "")
	if len(packageNameIsIC) == 0 {
		return nil
	}

	var packageNameIsIR []string
	for _, packageNameIsIV := range packageNameIsIC {
		packageNameIsI := packageNameIsIV

		packageNameIsIR = append(packageNameIsIR, packageNameIsI)
	}

	o.PackageNameIs = packageNameIsIR

	return nil
}

// bindPackageNameStart binds and validates parameter PackageNameStart from query.
func (o *GetPackagesParams) bindPackageNameStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageNameStart = &raw

	return nil
}

// bindPackageVersionContains binds and validates array parameter PackageVersionContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageVersionContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageVersionContains string
	if len(rawData) > 0 {
		qvPackageVersionContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageVersionContainsIC := swag.SplitByFormat(qvPackageVersionContains, "")
	if len(packageVersionContainsIC) == 0 {
		return nil
	}

	var packageVersionContainsIR []string
	for _, packageVersionContainsIV := range packageVersionContainsIC {
		packageVersionContainsI := packageVersionContainsIV

		packageVersionContainsIR = append(packageVersionContainsIR, packageVersionContainsI)
	}

	o.PackageVersionContains = packageVersionContainsIR

	return nil
}

// bindPackageVersionEnd binds and validates parameter PackageVersionEnd from query.
func (o *GetPackagesParams) bindPackageVersionEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageVersionEnd = &raw

	return nil
}

// bindPackageVersionIsNot binds and validates array parameter PackageVersionIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageVersionIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageVersionIsNot string
	if len(rawData) > 0 {
		qvPackageVersionIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageVersionIsNotIC := swag.SplitByFormat(qvPackageVersionIsNot, "")
	if len(packageVersionIsNotIC) == 0 {
		return nil
	}

	var packageVersionIsNotIR []string
	for _, packageVersionIsNotIV := range packageVersionIsNotIC {
		packageVersionIsNotI := packageVersionIsNotIV

		packageVersionIsNotIR = append(packageVersionIsNotIR, packageVersionIsNotI)
	}

	o.PackageVersionIsNot = packageVersionIsNotIR

	return nil
}

// bindPackageVersionIs binds and validates array parameter PackageVersionIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetPackagesParams) bindPackageVersionIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackageVersionIs string
	if len(rawData) > 0 {
		qvPackageVersionIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packageVersionIsIC := swag.SplitByFormat(qvPackageVersionIs, "")
	if len(packageVersionIsIC) == 0 {
		return nil
	}

	var packageVersionIsIR []string
	for _, packageVersionIsIV := range packageVersionIsIC {
		packageVersionIsI := packageVersionIsIV

		packageVersionIsIR = append(packageVersionIsIR, packageVersionIsI)
	}

	o.PackageVersionIs = packageVersionIsIR

	return nil
}

// bindPackageVersionStart binds and validates parameter PackageVersionStart from query.
func (o *GetPackagesParams) bindPackageVersionStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageVersionStart = &raw

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetPackagesParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("page", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("page", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetPackagesParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("pageSize", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("pageSize", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *GetPackagesParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", o.PageSize, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", o.PageSize, 50, false); err != nil {
		return err
	}

	return nil
}

// bindSortDir binds and validates parameter SortDir from query.
func (o *GetPackagesParams) bindSortDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetPackagesParams()
		return nil
	}
	o.SortDir = &raw

	if err := o.validateSortDir(formats); err != nil {
		return err
	}

	return nil
}

// validateSortDir carries on validations for parameter SortDir
func (o *GetPackagesParams) validateSortDir(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortDir", "query", *o.SortDir, []interface{}{"ASC", "DESC"}, true); err != nil {
		return err
	}

	return nil
}

// bindSortKey binds and validates parameter SortKey from query.
func (o *GetPackagesParams) bindSortKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("sortKey", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("sortKey", "query", raw); err != nil {
		return err
	}
	o.SortKey = raw

	if err := o.validateSortKey(formats); err != nil {
		return err
	}

	return nil
}

// validateSortKey carries on validations for parameter SortKey
func (o *GetPackagesParams) validateSortKey(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortKey", "query", o.SortKey, []interface{}{"packageName", "applications", "applicationResources", "language", "vulnerabilities", "version", "license"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityGte binds and validates parameter VulnerabilitySeverityGte from query.
func (o *GetPackagesParams) bindVulnerabilitySeverityGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.VulnerabilitySeverityGte = &raw

	if err := o.validateVulnerabilitySeverityGte(formats); err != nil {
		return err
	}

	return nil
}

// validateVulnerabilitySeverityGte carries on validations for parameter VulnerabilitySeverityGte
func (o *GetPackagesParams) validateVulnerabilitySeverityGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[gte]", "query", *o.VulnerabilitySeverityGte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityLte binds and validates parameter VulnerabilitySeverityLte from query.
func (o *GetPackagesParams) bindVulnerabilitySeverityLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.VulnerabilitySeverityLte = &raw

	if err := o.validateVulnerabilitySeverityLte(formats); err != nil {
		return err
	}

	return nil
}

// validateVulnerabilitySeverityLte carries on validations for parameter VulnerabilitySeverityLte
func (o *GetPackagesParams) validateVulnerabilitySeverityLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[lte]", "query", *o.VulnerabilitySeverityLte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}
