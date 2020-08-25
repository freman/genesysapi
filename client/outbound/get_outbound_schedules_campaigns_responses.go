// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundSchedulesCampaignsReader is a Reader for the GetOutboundSchedulesCampaigns structure.
type GetOutboundSchedulesCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesCampaignsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesCampaignsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesCampaignsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesCampaignsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesCampaignsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesCampaignsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesCampaignsOK creates a GetOutboundSchedulesCampaignsOK with default headers values
func NewGetOutboundSchedulesCampaignsOK() *GetOutboundSchedulesCampaignsOK {
	return &GetOutboundSchedulesCampaignsOK{}
}

/*GetOutboundSchedulesCampaignsOK handles this case with default header values.

successful operation
*/
type GetOutboundSchedulesCampaignsOK struct {
	Payload []*models.CampaignSchedule
}

func (o *GetOutboundSchedulesCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsOK) GetPayload() []*models.CampaignSchedule {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsBadRequest creates a GetOutboundSchedulesCampaignsBadRequest with default headers values
func NewGetOutboundSchedulesCampaignsBadRequest() *GetOutboundSchedulesCampaignsBadRequest {
	return &GetOutboundSchedulesCampaignsBadRequest{}
}

/*GetOutboundSchedulesCampaignsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesCampaignsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsUnauthorized creates a GetOutboundSchedulesCampaignsUnauthorized with default headers values
func NewGetOutboundSchedulesCampaignsUnauthorized() *GetOutboundSchedulesCampaignsUnauthorized {
	return &GetOutboundSchedulesCampaignsUnauthorized{}
}

/*GetOutboundSchedulesCampaignsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesCampaignsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsForbidden creates a GetOutboundSchedulesCampaignsForbidden with default headers values
func NewGetOutboundSchedulesCampaignsForbidden() *GetOutboundSchedulesCampaignsForbidden {
	return &GetOutboundSchedulesCampaignsForbidden{}
}

/*GetOutboundSchedulesCampaignsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesCampaignsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsNotFound creates a GetOutboundSchedulesCampaignsNotFound with default headers values
func NewGetOutboundSchedulesCampaignsNotFound() *GetOutboundSchedulesCampaignsNotFound {
	return &GetOutboundSchedulesCampaignsNotFound{}
}

/*GetOutboundSchedulesCampaignsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesCampaignsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsRequestEntityTooLarge creates a GetOutboundSchedulesCampaignsRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesCampaignsRequestEntityTooLarge() *GetOutboundSchedulesCampaignsRequestEntityTooLarge {
	return &GetOutboundSchedulesCampaignsRequestEntityTooLarge{}
}

/*GetOutboundSchedulesCampaignsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundSchedulesCampaignsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsUnsupportedMediaType creates a GetOutboundSchedulesCampaignsUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesCampaignsUnsupportedMediaType() *GetOutboundSchedulesCampaignsUnsupportedMediaType {
	return &GetOutboundSchedulesCampaignsUnsupportedMediaType{}
}

/*GetOutboundSchedulesCampaignsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesCampaignsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsTooManyRequests creates a GetOutboundSchedulesCampaignsTooManyRequests with default headers values
func NewGetOutboundSchedulesCampaignsTooManyRequests() *GetOutboundSchedulesCampaignsTooManyRequests {
	return &GetOutboundSchedulesCampaignsTooManyRequests{}
}

/*GetOutboundSchedulesCampaignsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundSchedulesCampaignsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsInternalServerError creates a GetOutboundSchedulesCampaignsInternalServerError with default headers values
func NewGetOutboundSchedulesCampaignsInternalServerError() *GetOutboundSchedulesCampaignsInternalServerError {
	return &GetOutboundSchedulesCampaignsInternalServerError{}
}

/*GetOutboundSchedulesCampaignsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesCampaignsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsServiceUnavailable creates a GetOutboundSchedulesCampaignsServiceUnavailable with default headers values
func NewGetOutboundSchedulesCampaignsServiceUnavailable() *GetOutboundSchedulesCampaignsServiceUnavailable {
	return &GetOutboundSchedulesCampaignsServiceUnavailable{}
}

/*GetOutboundSchedulesCampaignsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesCampaignsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesCampaignsGatewayTimeout creates a GetOutboundSchedulesCampaignsGatewayTimeout with default headers values
func NewGetOutboundSchedulesCampaignsGatewayTimeout() *GetOutboundSchedulesCampaignsGatewayTimeout {
	return &GetOutboundSchedulesCampaignsGatewayTimeout{}
}

/*GetOutboundSchedulesCampaignsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundSchedulesCampaignsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/campaigns][%d] getOutboundSchedulesCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}