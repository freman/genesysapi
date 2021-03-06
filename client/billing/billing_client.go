// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the billing client
type API interface {
	/*
	   GetBillingReportsBillableusage gets a report of the billable license usages
	   Report is of the billable usages (e.g. licenses and devices utilized) for a given period. If response's status is InProgress, wait a few seconds, then try the same request again.
	*/
	GetBillingReportsBillableusage(ctx context.Context, params *GetBillingReportsBillableusageParams) (*GetBillingReportsBillableusageOK, error)
	/*
	   GetBillingTrusteebillingoverviewTrustorOrgID gets the billing overview for an organization that is managed by a partner
	   Tax Disclaimer: Prices returned by this API do not include applicable taxes. It is the responsibility of the customer to pay all taxes that are appropriate in their jurisdiction. See the PureCloud API Documentation in the Developer Center for more information about this API: https://developer.mypurecloud.com/api/rest/v2/
	*/
	GetBillingTrusteebillingoverviewTrustorOrgID(ctx context.Context, params *GetBillingTrusteebillingoverviewTrustorOrgIDParams) (*GetBillingTrusteebillingoverviewTrustorOrgIDOK, error)
}

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetBillingReportsBillableusage gets a report of the billable license usages

Report is of the billable usages (e.g. licenses and devices utilized) for a given period. If response's status is InProgress, wait a few seconds, then try the same request again.
*/
func (a *Client) GetBillingReportsBillableusage(ctx context.Context, params *GetBillingReportsBillableusageParams) (*GetBillingReportsBillableusageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBillingReportsBillableusage",
		Method:             "GET",
		PathPattern:        "/api/v2/billing/reports/billableusage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingReportsBillableusageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBillingReportsBillableusageOK), nil

}

/*
GetBillingTrusteebillingoverviewTrustorOrgID gets the billing overview for an organization that is managed by a partner

Tax Disclaimer: Prices returned by this API do not include applicable taxes. It is the responsibility of the customer to pay all taxes that are appropriate in their jurisdiction. See the PureCloud API Documentation in the Developer Center for more information about this API: https://developer.mypurecloud.com/api/rest/v2/
*/
func (a *Client) GetBillingTrusteebillingoverviewTrustorOrgID(ctx context.Context, params *GetBillingTrusteebillingoverviewTrustorOrgIDParams) (*GetBillingTrusteebillingoverviewTrustorOrgIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBillingTrusteebillingoverviewTrustorOrgId",
		Method:             "GET",
		PathPattern:        "/api/v2/billing/trusteebillingoverview/{trustorOrgId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBillingTrusteebillingoverviewTrustorOrgIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBillingTrusteebillingoverviewTrustorOrgIDOK), nil

}
