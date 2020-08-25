// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUsageQueryExecutionIDResultsReader is a Reader for the GetUsageQueryExecutionIDResults structure.
type GetUsageQueryExecutionIDResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsageQueryExecutionIDResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsageQueryExecutionIDResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsageQueryExecutionIDResultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsageQueryExecutionIDResultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsageQueryExecutionIDResultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsageQueryExecutionIDResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUsageQueryExecutionIDResultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsageQueryExecutionIDResultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsageQueryExecutionIDResultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsageQueryExecutionIDResultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUsageQueryExecutionIDResultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsageQueryExecutionIDResultsOK creates a GetUsageQueryExecutionIDResultsOK with default headers values
func NewGetUsageQueryExecutionIDResultsOK() *GetUsageQueryExecutionIDResultsOK {
	return &GetUsageQueryExecutionIDResultsOK{}
}

/*GetUsageQueryExecutionIDResultsOK handles this case with default header values.

successful operation
*/
type GetUsageQueryExecutionIDResultsOK struct {
	Payload *models.APIUsageQueryResult
}

func (o *GetUsageQueryExecutionIDResultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsOK  %+v", 200, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsOK) GetPayload() *models.APIUsageQueryResult {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIUsageQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsBadRequest creates a GetUsageQueryExecutionIDResultsBadRequest with default headers values
func NewGetUsageQueryExecutionIDResultsBadRequest() *GetUsageQueryExecutionIDResultsBadRequest {
	return &GetUsageQueryExecutionIDResultsBadRequest{}
}

/*GetUsageQueryExecutionIDResultsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUsageQueryExecutionIDResultsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsUnauthorized creates a GetUsageQueryExecutionIDResultsUnauthorized with default headers values
func NewGetUsageQueryExecutionIDResultsUnauthorized() *GetUsageQueryExecutionIDResultsUnauthorized {
	return &GetUsageQueryExecutionIDResultsUnauthorized{}
}

/*GetUsageQueryExecutionIDResultsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUsageQueryExecutionIDResultsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsForbidden creates a GetUsageQueryExecutionIDResultsForbidden with default headers values
func NewGetUsageQueryExecutionIDResultsForbidden() *GetUsageQueryExecutionIDResultsForbidden {
	return &GetUsageQueryExecutionIDResultsForbidden{}
}

/*GetUsageQueryExecutionIDResultsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUsageQueryExecutionIDResultsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsNotFound creates a GetUsageQueryExecutionIDResultsNotFound with default headers values
func NewGetUsageQueryExecutionIDResultsNotFound() *GetUsageQueryExecutionIDResultsNotFound {
	return &GetUsageQueryExecutionIDResultsNotFound{}
}

/*GetUsageQueryExecutionIDResultsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUsageQueryExecutionIDResultsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge creates a GetUsageQueryExecutionIDResultsRequestEntityTooLarge with default headers values
func NewGetUsageQueryExecutionIDResultsRequestEntityTooLarge() *GetUsageQueryExecutionIDResultsRequestEntityTooLarge {
	return &GetUsageQueryExecutionIDResultsRequestEntityTooLarge{}
}

/*GetUsageQueryExecutionIDResultsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUsageQueryExecutionIDResultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsUnsupportedMediaType creates a GetUsageQueryExecutionIDResultsUnsupportedMediaType with default headers values
func NewGetUsageQueryExecutionIDResultsUnsupportedMediaType() *GetUsageQueryExecutionIDResultsUnsupportedMediaType {
	return &GetUsageQueryExecutionIDResultsUnsupportedMediaType{}
}

/*GetUsageQueryExecutionIDResultsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUsageQueryExecutionIDResultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsTooManyRequests creates a GetUsageQueryExecutionIDResultsTooManyRequests with default headers values
func NewGetUsageQueryExecutionIDResultsTooManyRequests() *GetUsageQueryExecutionIDResultsTooManyRequests {
	return &GetUsageQueryExecutionIDResultsTooManyRequests{}
}

/*GetUsageQueryExecutionIDResultsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUsageQueryExecutionIDResultsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsInternalServerError creates a GetUsageQueryExecutionIDResultsInternalServerError with default headers values
func NewGetUsageQueryExecutionIDResultsInternalServerError() *GetUsageQueryExecutionIDResultsInternalServerError {
	return &GetUsageQueryExecutionIDResultsInternalServerError{}
}

/*GetUsageQueryExecutionIDResultsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUsageQueryExecutionIDResultsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsServiceUnavailable creates a GetUsageQueryExecutionIDResultsServiceUnavailable with default headers values
func NewGetUsageQueryExecutionIDResultsServiceUnavailable() *GetUsageQueryExecutionIDResultsServiceUnavailable {
	return &GetUsageQueryExecutionIDResultsServiceUnavailable{}
}

/*GetUsageQueryExecutionIDResultsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUsageQueryExecutionIDResultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageQueryExecutionIDResultsGatewayTimeout creates a GetUsageQueryExecutionIDResultsGatewayTimeout with default headers values
func NewGetUsageQueryExecutionIDResultsGatewayTimeout() *GetUsageQueryExecutionIDResultsGatewayTimeout {
	return &GetUsageQueryExecutionIDResultsGatewayTimeout{}
}

/*GetUsageQueryExecutionIDResultsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUsageQueryExecutionIDResultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/usage/query/{executionId}/results][%d] getUsageQueryExecutionIdResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsageQueryExecutionIDResultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}