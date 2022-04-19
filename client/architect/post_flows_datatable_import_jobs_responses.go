// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostFlowsDatatableImportJobsReader is a Reader for the PostFlowsDatatableImportJobs structure.
type PostFlowsDatatableImportJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFlowsDatatableImportJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostFlowsDatatableImportJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostFlowsDatatableImportJobsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFlowsDatatableImportJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFlowsDatatableImportJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFlowsDatatableImportJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFlowsDatatableImportJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostFlowsDatatableImportJobsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostFlowsDatatableImportJobsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostFlowsDatatableImportJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostFlowsDatatableImportJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFlowsDatatableImportJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFlowsDatatableImportJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFlowsDatatableImportJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFlowsDatatableImportJobsOK creates a PostFlowsDatatableImportJobsOK with default headers values
func NewPostFlowsDatatableImportJobsOK() *PostFlowsDatatableImportJobsOK {
	return &PostFlowsDatatableImportJobsOK{}
}

/*PostFlowsDatatableImportJobsOK handles this case with default header values.

successful operation
*/
type PostFlowsDatatableImportJobsOK struct {
	Payload *models.DataTableImportJob
}

func (o *PostFlowsDatatableImportJobsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsOK  %+v", 200, o.Payload)
}

func (o *PostFlowsDatatableImportJobsOK) GetPayload() *models.DataTableImportJob {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTableImportJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsAccepted creates a PostFlowsDatatableImportJobsAccepted with default headers values
func NewPostFlowsDatatableImportJobsAccepted() *PostFlowsDatatableImportJobsAccepted {
	return &PostFlowsDatatableImportJobsAccepted{}
}

/*PostFlowsDatatableImportJobsAccepted handles this case with default header values.

Request Accepted
*/
type PostFlowsDatatableImportJobsAccepted struct {
	Payload *models.DataTableImportJob
}

func (o *PostFlowsDatatableImportJobsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostFlowsDatatableImportJobsAccepted) GetPayload() *models.DataTableImportJob {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTableImportJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsBadRequest creates a PostFlowsDatatableImportJobsBadRequest with default headers values
func NewPostFlowsDatatableImportJobsBadRequest() *PostFlowsDatatableImportJobsBadRequest {
	return &PostFlowsDatatableImportJobsBadRequest{}
}

/*PostFlowsDatatableImportJobsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostFlowsDatatableImportJobsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFlowsDatatableImportJobsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsUnauthorized creates a PostFlowsDatatableImportJobsUnauthorized with default headers values
func NewPostFlowsDatatableImportJobsUnauthorized() *PostFlowsDatatableImportJobsUnauthorized {
	return &PostFlowsDatatableImportJobsUnauthorized{}
}

/*PostFlowsDatatableImportJobsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostFlowsDatatableImportJobsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFlowsDatatableImportJobsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsForbidden creates a PostFlowsDatatableImportJobsForbidden with default headers values
func NewPostFlowsDatatableImportJobsForbidden() *PostFlowsDatatableImportJobsForbidden {
	return &PostFlowsDatatableImportJobsForbidden{}
}

/*PostFlowsDatatableImportJobsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostFlowsDatatableImportJobsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostFlowsDatatableImportJobsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsNotFound creates a PostFlowsDatatableImportJobsNotFound with default headers values
func NewPostFlowsDatatableImportJobsNotFound() *PostFlowsDatatableImportJobsNotFound {
	return &PostFlowsDatatableImportJobsNotFound{}
}

/*PostFlowsDatatableImportJobsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostFlowsDatatableImportJobsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostFlowsDatatableImportJobsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsRequestTimeout creates a PostFlowsDatatableImportJobsRequestTimeout with default headers values
func NewPostFlowsDatatableImportJobsRequestTimeout() *PostFlowsDatatableImportJobsRequestTimeout {
	return &PostFlowsDatatableImportJobsRequestTimeout{}
}

/*PostFlowsDatatableImportJobsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostFlowsDatatableImportJobsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostFlowsDatatableImportJobsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsRequestEntityTooLarge creates a PostFlowsDatatableImportJobsRequestEntityTooLarge with default headers values
func NewPostFlowsDatatableImportJobsRequestEntityTooLarge() *PostFlowsDatatableImportJobsRequestEntityTooLarge {
	return &PostFlowsDatatableImportJobsRequestEntityTooLarge{}
}

/*PostFlowsDatatableImportJobsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostFlowsDatatableImportJobsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostFlowsDatatableImportJobsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsUnsupportedMediaType creates a PostFlowsDatatableImportJobsUnsupportedMediaType with default headers values
func NewPostFlowsDatatableImportJobsUnsupportedMediaType() *PostFlowsDatatableImportJobsUnsupportedMediaType {
	return &PostFlowsDatatableImportJobsUnsupportedMediaType{}
}

/*PostFlowsDatatableImportJobsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostFlowsDatatableImportJobsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostFlowsDatatableImportJobsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsTooManyRequests creates a PostFlowsDatatableImportJobsTooManyRequests with default headers values
func NewPostFlowsDatatableImportJobsTooManyRequests() *PostFlowsDatatableImportJobsTooManyRequests {
	return &PostFlowsDatatableImportJobsTooManyRequests{}
}

/*PostFlowsDatatableImportJobsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostFlowsDatatableImportJobsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostFlowsDatatableImportJobsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsInternalServerError creates a PostFlowsDatatableImportJobsInternalServerError with default headers values
func NewPostFlowsDatatableImportJobsInternalServerError() *PostFlowsDatatableImportJobsInternalServerError {
	return &PostFlowsDatatableImportJobsInternalServerError{}
}

/*PostFlowsDatatableImportJobsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostFlowsDatatableImportJobsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFlowsDatatableImportJobsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsServiceUnavailable creates a PostFlowsDatatableImportJobsServiceUnavailable with default headers values
func NewPostFlowsDatatableImportJobsServiceUnavailable() *PostFlowsDatatableImportJobsServiceUnavailable {
	return &PostFlowsDatatableImportJobsServiceUnavailable{}
}

/*PostFlowsDatatableImportJobsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostFlowsDatatableImportJobsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFlowsDatatableImportJobsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFlowsDatatableImportJobsGatewayTimeout creates a PostFlowsDatatableImportJobsGatewayTimeout with default headers values
func NewPostFlowsDatatableImportJobsGatewayTimeout() *PostFlowsDatatableImportJobsGatewayTimeout {
	return &PostFlowsDatatableImportJobsGatewayTimeout{}
}

/*PostFlowsDatatableImportJobsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostFlowsDatatableImportJobsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostFlowsDatatableImportJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/flows/datatables/{datatableId}/import/jobs][%d] postFlowsDatatableImportJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFlowsDatatableImportJobsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostFlowsDatatableImportJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
