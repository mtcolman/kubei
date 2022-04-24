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

// NewGetApplicationsParams creates a new GetApplicationsParams object
// with the default values initialized.
func NewGetApplicationsParams() GetApplicationsParams {

	var (
		// initialize parameters with default values

		sortDirDefault = string("ASC")
	)

	return GetApplicationsParams{
		SortDir: &sortDirDefault,
	}
}

// GetApplicationsParams contains all the bound params for the get applications operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetApplications
type GetApplicationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	ApplicationEnvsContainElements []string
	/*
	  In: query
	*/
	ApplicationEnvsDoesntContainElements []string
	/*
	  In: query
	*/
	ApplicationLabelsContainElements []string
	/*
	  In: query
	*/
	ApplicationLabelsDoesntContainElements []string
	/*
	  In: query
	*/
	ApplicationNameContains []string
	/*
	  In: query
	*/
	ApplicationNameEnd *string
	/*
	  In: query
	*/
	ApplicationNameIsNot []string
	/*
	  In: query
	*/
	ApplicationNameIs []string
	/*
	  In: query
	*/
	ApplicationNameStart *string
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
	/*
	  In: query
	*/
	ApplicationTypeIs []string
	/*
	  In: query
	*/
	CisDockerBenchmarkLevelGte *string
	/*
	  In: query
	*/
	CisDockerBenchmarkLevelLte *string
	/*current runtime scan system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	CurrentRuntimeScan *bool
	/*package ID system filter, not visible to the user. only one of applicationID, applicationResourceID, packageID, currentRuntimeScan is allowed
	  In: query
	*/
	PackageID *string
	/*greater than or equal
	  In: query
	*/
	PackagesGte *int64
	/*
	  In: query
	*/
	PackagesIsNot []int64
	/*
	  In: query
	*/
	PackagesIs []int64
	/*less than or equal
	  In: query
	*/
	PackagesLte *int64
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
// To ensure default values, the struct must have been initialized with NewGetApplicationsParams() beforehand.
func (o *GetApplicationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qApplicationEnvsContainElements, qhkApplicationEnvsContainElements, _ := qs.GetOK("applicationEnvs[containElements]")
	if err := o.bindApplicationEnvsContainElements(qApplicationEnvsContainElements, qhkApplicationEnvsContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationEnvsDoesntContainElements, qhkApplicationEnvsDoesntContainElements, _ := qs.GetOK("applicationEnvs[doesntContainElements]")
	if err := o.bindApplicationEnvsDoesntContainElements(qApplicationEnvsDoesntContainElements, qhkApplicationEnvsDoesntContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationLabelsContainElements, qhkApplicationLabelsContainElements, _ := qs.GetOK("applicationLabels[containElements]")
	if err := o.bindApplicationLabelsContainElements(qApplicationLabelsContainElements, qhkApplicationLabelsContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationLabelsDoesntContainElements, qhkApplicationLabelsDoesntContainElements, _ := qs.GetOK("applicationLabels[doesntContainElements]")
	if err := o.bindApplicationLabelsDoesntContainElements(qApplicationLabelsDoesntContainElements, qhkApplicationLabelsDoesntContainElements, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationNameContains, qhkApplicationNameContains, _ := qs.GetOK("applicationName[contains]")
	if err := o.bindApplicationNameContains(qApplicationNameContains, qhkApplicationNameContains, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationNameEnd, qhkApplicationNameEnd, _ := qs.GetOK("applicationName[end]")
	if err := o.bindApplicationNameEnd(qApplicationNameEnd, qhkApplicationNameEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationNameIsNot, qhkApplicationNameIsNot, _ := qs.GetOK("applicationName[isNot]")
	if err := o.bindApplicationNameIsNot(qApplicationNameIsNot, qhkApplicationNameIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationNameIs, qhkApplicationNameIs, _ := qs.GetOK("applicationName[is]")
	if err := o.bindApplicationNameIs(qApplicationNameIs, qhkApplicationNameIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qApplicationNameStart, qhkApplicationNameStart, _ := qs.GetOK("applicationName[start]")
	if err := o.bindApplicationNameStart(qApplicationNameStart, qhkApplicationNameStart, route.Formats); err != nil {
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

	qApplicationTypeIs, qhkApplicationTypeIs, _ := qs.GetOK("applicationType[is]")
	if err := o.bindApplicationTypeIs(qApplicationTypeIs, qhkApplicationTypeIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qCisDockerBenchmarkLevelGte, qhkCisDockerBenchmarkLevelGte, _ := qs.GetOK("cisDockerBenchmarkLevel[gte]")
	if err := o.bindCisDockerBenchmarkLevelGte(qCisDockerBenchmarkLevelGte, qhkCisDockerBenchmarkLevelGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCisDockerBenchmarkLevelLte, qhkCisDockerBenchmarkLevelLte, _ := qs.GetOK("cisDockerBenchmarkLevel[lte]")
	if err := o.bindCisDockerBenchmarkLevelLte(qCisDockerBenchmarkLevelLte, qhkCisDockerBenchmarkLevelLte, route.Formats); err != nil {
		res = append(res, err)
	}

	qCurrentRuntimeScan, qhkCurrentRuntimeScan, _ := qs.GetOK("currentRuntimeScan")
	if err := o.bindCurrentRuntimeScan(qCurrentRuntimeScan, qhkCurrentRuntimeScan, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackageID, qhkPackageID, _ := qs.GetOK("packageID")
	if err := o.bindPackageID(qPackageID, qhkPackageID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesGte, qhkPackagesGte, _ := qs.GetOK("packages[gte]")
	if err := o.bindPackagesGte(qPackagesGte, qhkPackagesGte, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesIsNot, qhkPackagesIsNot, _ := qs.GetOK("packages[isNot]")
	if err := o.bindPackagesIsNot(qPackagesIsNot, qhkPackagesIsNot, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesIs, qhkPackagesIs, _ := qs.GetOK("packages[is]")
	if err := o.bindPackagesIs(qPackagesIs, qhkPackagesIs, route.Formats); err != nil {
		res = append(res, err)
	}

	qPackagesLte, qhkPackagesLte, _ := qs.GetOK("packages[lte]")
	if err := o.bindPackagesLte(qPackagesLte, qhkPackagesLte, route.Formats); err != nil {
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

// bindApplicationEnvsContainElements binds and validates array parameter ApplicationEnvsContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationEnvsContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationEnvsContainElements string
	if len(rawData) > 0 {
		qvApplicationEnvsContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationEnvsContainElementsIC := swag.SplitByFormat(qvApplicationEnvsContainElements, "")
	if len(applicationEnvsContainElementsIC) == 0 {
		return nil
	}

	var applicationEnvsContainElementsIR []string
	for _, applicationEnvsContainElementsIV := range applicationEnvsContainElementsIC {
		applicationEnvsContainElementsI := applicationEnvsContainElementsIV

		applicationEnvsContainElementsIR = append(applicationEnvsContainElementsIR, applicationEnvsContainElementsI)
	}

	o.ApplicationEnvsContainElements = applicationEnvsContainElementsIR

	return nil
}

// bindApplicationEnvsDoesntContainElements binds and validates array parameter ApplicationEnvsDoesntContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationEnvsDoesntContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationEnvsDoesntContainElements string
	if len(rawData) > 0 {
		qvApplicationEnvsDoesntContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationEnvsDoesntContainElementsIC := swag.SplitByFormat(qvApplicationEnvsDoesntContainElements, "")
	if len(applicationEnvsDoesntContainElementsIC) == 0 {
		return nil
	}

	var applicationEnvsDoesntContainElementsIR []string
	for _, applicationEnvsDoesntContainElementsIV := range applicationEnvsDoesntContainElementsIC {
		applicationEnvsDoesntContainElementsI := applicationEnvsDoesntContainElementsIV

		applicationEnvsDoesntContainElementsIR = append(applicationEnvsDoesntContainElementsIR, applicationEnvsDoesntContainElementsI)
	}

	o.ApplicationEnvsDoesntContainElements = applicationEnvsDoesntContainElementsIR

	return nil
}

// bindApplicationLabelsContainElements binds and validates array parameter ApplicationLabelsContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationLabelsContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationLabelsContainElements string
	if len(rawData) > 0 {
		qvApplicationLabelsContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationLabelsContainElementsIC := swag.SplitByFormat(qvApplicationLabelsContainElements, "")
	if len(applicationLabelsContainElementsIC) == 0 {
		return nil
	}

	var applicationLabelsContainElementsIR []string
	for _, applicationLabelsContainElementsIV := range applicationLabelsContainElementsIC {
		applicationLabelsContainElementsI := applicationLabelsContainElementsIV

		applicationLabelsContainElementsIR = append(applicationLabelsContainElementsIR, applicationLabelsContainElementsI)
	}

	o.ApplicationLabelsContainElements = applicationLabelsContainElementsIR

	return nil
}

// bindApplicationLabelsDoesntContainElements binds and validates array parameter ApplicationLabelsDoesntContainElements from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationLabelsDoesntContainElements(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationLabelsDoesntContainElements string
	if len(rawData) > 0 {
		qvApplicationLabelsDoesntContainElements = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationLabelsDoesntContainElementsIC := swag.SplitByFormat(qvApplicationLabelsDoesntContainElements, "")
	if len(applicationLabelsDoesntContainElementsIC) == 0 {
		return nil
	}

	var applicationLabelsDoesntContainElementsIR []string
	for _, applicationLabelsDoesntContainElementsIV := range applicationLabelsDoesntContainElementsIC {
		applicationLabelsDoesntContainElementsI := applicationLabelsDoesntContainElementsIV

		applicationLabelsDoesntContainElementsIR = append(applicationLabelsDoesntContainElementsIR, applicationLabelsDoesntContainElementsI)
	}

	o.ApplicationLabelsDoesntContainElements = applicationLabelsDoesntContainElementsIR

	return nil
}

// bindApplicationNameContains binds and validates array parameter ApplicationNameContains from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationNameContains(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationNameContains string
	if len(rawData) > 0 {
		qvApplicationNameContains = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationNameContainsIC := swag.SplitByFormat(qvApplicationNameContains, "")
	if len(applicationNameContainsIC) == 0 {
		return nil
	}

	var applicationNameContainsIR []string
	for _, applicationNameContainsIV := range applicationNameContainsIC {
		applicationNameContainsI := applicationNameContainsIV

		applicationNameContainsIR = append(applicationNameContainsIR, applicationNameContainsI)
	}

	o.ApplicationNameContains = applicationNameContainsIR

	return nil
}

// bindApplicationNameEnd binds and validates parameter ApplicationNameEnd from query.
func (o *GetApplicationsParams) bindApplicationNameEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationNameEnd = &raw

	return nil
}

// bindApplicationNameIsNot binds and validates array parameter ApplicationNameIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationNameIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationNameIsNot string
	if len(rawData) > 0 {
		qvApplicationNameIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationNameIsNotIC := swag.SplitByFormat(qvApplicationNameIsNot, "")
	if len(applicationNameIsNotIC) == 0 {
		return nil
	}

	var applicationNameIsNotIR []string
	for _, applicationNameIsNotIV := range applicationNameIsNotIC {
		applicationNameIsNotI := applicationNameIsNotIV

		applicationNameIsNotIR = append(applicationNameIsNotIR, applicationNameIsNotI)
	}

	o.ApplicationNameIsNot = applicationNameIsNotIR

	return nil
}

// bindApplicationNameIs binds and validates array parameter ApplicationNameIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationNameIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationNameIs string
	if len(rawData) > 0 {
		qvApplicationNameIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationNameIsIC := swag.SplitByFormat(qvApplicationNameIs, "")
	if len(applicationNameIsIC) == 0 {
		return nil
	}

	var applicationNameIsIR []string
	for _, applicationNameIsIV := range applicationNameIsIC {
		applicationNameIsI := applicationNameIsIV

		applicationNameIsIR = append(applicationNameIsIR, applicationNameIsI)
	}

	o.ApplicationNameIs = applicationNameIsIR

	return nil
}

// bindApplicationNameStart binds and validates parameter ApplicationNameStart from query.
func (o *GetApplicationsParams) bindApplicationNameStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ApplicationNameStart = &raw

	return nil
}

// bindApplicationResourceID binds and validates parameter ApplicationResourceID from query.
func (o *GetApplicationsParams) bindApplicationResourceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) bindApplicationResourcesGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) bindApplicationResourcesIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) bindApplicationResourcesIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) bindApplicationResourcesLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindApplicationTypeIs binds and validates array parameter ApplicationTypeIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindApplicationTypeIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvApplicationTypeIs string
	if len(rawData) > 0 {
		qvApplicationTypeIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	applicationTypeIsIC := swag.SplitByFormat(qvApplicationTypeIs, "")
	if len(applicationTypeIsIC) == 0 {
		return nil
	}

	var applicationTypeIsIR []string
	for i, applicationTypeIsIV := range applicationTypeIsIC {
		applicationTypeIsI := applicationTypeIsIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "applicationType[is]", i), "query", applicationTypeIsI, []interface{}{"POD", "DIRECTORY", "LAMBDA"}, true); err != nil {
			return err
		}

		applicationTypeIsIR = append(applicationTypeIsIR, applicationTypeIsI)
	}

	o.ApplicationTypeIs = applicationTypeIsIR

	return nil
}

// bindCisDockerBenchmarkLevelGte binds and validates parameter CisDockerBenchmarkLevelGte from query.
func (o *GetApplicationsParams) bindCisDockerBenchmarkLevelGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CisDockerBenchmarkLevelGte = &raw

	if err := o.validateCisDockerBenchmarkLevelGte(formats); err != nil {
		return err
	}

	return nil
}

// validateCisDockerBenchmarkLevelGte carries on validations for parameter CisDockerBenchmarkLevelGte
func (o *GetApplicationsParams) validateCisDockerBenchmarkLevelGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[gte]", "query", *o.CisDockerBenchmarkLevelGte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCisDockerBenchmarkLevelLte binds and validates parameter CisDockerBenchmarkLevelLte from query.
func (o *GetApplicationsParams) bindCisDockerBenchmarkLevelLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CisDockerBenchmarkLevelLte = &raw

	if err := o.validateCisDockerBenchmarkLevelLte(formats); err != nil {
		return err
	}

	return nil
}

// validateCisDockerBenchmarkLevelLte carries on validations for parameter CisDockerBenchmarkLevelLte
func (o *GetApplicationsParams) validateCisDockerBenchmarkLevelLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("cisDockerBenchmarkLevel[lte]", "query", *o.CisDockerBenchmarkLevelLte, []interface{}{"INFO", "WARN", "FATAL"}, true); err != nil {
		return err
	}

	return nil
}

// bindCurrentRuntimeScan binds and validates parameter CurrentRuntimeScan from query.
func (o *GetApplicationsParams) bindCurrentRuntimeScan(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindPackageID binds and validates parameter PackageID from query.
func (o *GetApplicationsParams) bindPackageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.PackageID = &raw

	return nil
}

// bindPackagesGte binds and validates parameter PackagesGte from query.
func (o *GetApplicationsParams) bindPackagesGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("packages[gte]", "query", "int64", raw)
	}
	o.PackagesGte = &value

	return nil
}

// bindPackagesIsNot binds and validates array parameter PackagesIsNot from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindPackagesIsNot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackagesIsNot string
	if len(rawData) > 0 {
		qvPackagesIsNot = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packagesIsNotIC := swag.SplitByFormat(qvPackagesIsNot, "")
	if len(packagesIsNotIC) == 0 {
		return nil
	}

	var packagesIsNotIR []int64
	for i, packagesIsNotIV := range packagesIsNotIC {
		packagesIsNotI, err := swag.ConvertInt64(packagesIsNotIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "packages[isNot]", i), "query", "int64", packagesIsNotI)
		}

		packagesIsNotIR = append(packagesIsNotIR, packagesIsNotI)
	}

	o.PackagesIsNot = packagesIsNotIR

	return nil
}

// bindPackagesIs binds and validates array parameter PackagesIs from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetApplicationsParams) bindPackagesIs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvPackagesIs string
	if len(rawData) > 0 {
		qvPackagesIs = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	packagesIsIC := swag.SplitByFormat(qvPackagesIs, "")
	if len(packagesIsIC) == 0 {
		return nil
	}

	var packagesIsIR []int64
	for i, packagesIsIV := range packagesIsIC {
		packagesIsI, err := swag.ConvertInt64(packagesIsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "packages[is]", i), "query", "int64", packagesIsI)
		}

		packagesIsIR = append(packagesIsIR, packagesIsI)
	}

	o.PackagesIs = packagesIsIR

	return nil
}

// bindPackagesLte binds and validates parameter PackagesLte from query.
func (o *GetApplicationsParams) bindPackagesLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("packages[lte]", "query", "int64", raw)
	}
	o.PackagesLte = &value

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *GetApplicationsParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", o.PageSize, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", o.PageSize, 50, false); err != nil {
		return err
	}

	return nil
}

// bindSortDir binds and validates parameter SortDir from query.
func (o *GetApplicationsParams) bindSortDir(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetApplicationsParams()
		return nil
	}
	o.SortDir = &raw

	if err := o.validateSortDir(formats); err != nil {
		return err
	}

	return nil
}

// validateSortDir carries on validations for parameter SortDir
func (o *GetApplicationsParams) validateSortDir(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortDir", "query", *o.SortDir, []interface{}{"ASC", "DESC"}, true); err != nil {
		return err
	}

	return nil
}

// bindSortKey binds and validates parameter SortKey from query.
func (o *GetApplicationsParams) bindSortKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) validateSortKey(formats strfmt.Registry) error {

	if err := validate.EnumCase("sortKey", "query", o.SortKey, []interface{}{"applicationName", "applicationType", "vulnerabilities", "cisDockerBenchmarkResults", "applicationResources", "packages"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityGte binds and validates parameter VulnerabilitySeverityGte from query.
func (o *GetApplicationsParams) bindVulnerabilitySeverityGte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) validateVulnerabilitySeverityGte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[gte]", "query", *o.VulnerabilitySeverityGte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}

// bindVulnerabilitySeverityLte binds and validates parameter VulnerabilitySeverityLte from query.
func (o *GetApplicationsParams) bindVulnerabilitySeverityLte(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetApplicationsParams) validateVulnerabilitySeverityLte(formats strfmt.Registry) error {

	if err := validate.EnumCase("vulnerabilitySeverity[lte]", "query", *o.VulnerabilitySeverityLte, []interface{}{"CRITICAL", "HIGH", "MEDIUM", "LOW", "NEGLIGIBLE"}, true); err != nil {
		return err
	}

	return nil
}
