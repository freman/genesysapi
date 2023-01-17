// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyActionmapsEstimatesJobReader is a Reader for the GetJourneyActionmapsEstimatesJob structure.
type GetJourneyActionmapsEstimatesJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActionmapsEstimatesJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActionmapsEstimatesJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetJourneyActionmapsEstimatesJobAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActionmapsEstimatesJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActionmapsEstimatesJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActionmapsEstimatesJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActionmapsEstimatesJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyActionmapsEstimatesJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActionmapsEstimatesJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActionmapsEstimatesJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActionmapsEstimatesJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActionmapsEstimatesJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActionmapsEstimatesJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActionmapsEstimatesJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActionmapsEstimatesJobOK creates a GetJourneyActionmapsEstimatesJobOK with default headers values
func NewGetJourneyActionmapsEstimatesJobOK() *GetJourneyActionmapsEstimatesJobOK {
	return &GetJourneyActionmapsEstimatesJobOK{}
}

/*
GetJourneyActionmapsEstimatesJobOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyActionmapsEstimatesJobOK struct {
	Payload string
}

// IsSuccess returns true when this get journey actionmaps estimates job o k response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey actionmaps estimates job o k response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job o k response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps estimates job o k response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job o k response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyActionmapsEstimatesJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobOK) GetPayload() string {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobAccepted creates a GetJourneyActionmapsEstimatesJobAccepted with default headers values
func NewGetJourneyActionmapsEstimatesJobAccepted() *GetJourneyActionmapsEstimatesJobAccepted {
	return &GetJourneyActionmapsEstimatesJobAccepted{}
}

/*
GetJourneyActionmapsEstimatesJobAccepted describes a response with status code 202, with default header values.

Accepted - Running query asynchronously
*/
type GetJourneyActionmapsEstimatesJobAccepted struct {
	Payload string
}

// IsSuccess returns true when this get journey actionmaps estimates job accepted response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey actionmaps estimates job accepted response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job accepted response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps estimates job accepted response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job accepted response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GetJourneyActionmapsEstimatesJobAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobAccepted  %+v", 202, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobAccepted) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobAccepted  %+v", 202, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobAccepted) GetPayload() string {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobBadRequest creates a GetJourneyActionmapsEstimatesJobBadRequest with default headers values
func NewGetJourneyActionmapsEstimatesJobBadRequest() *GetJourneyActionmapsEstimatesJobBadRequest {
	return &GetJourneyActionmapsEstimatesJobBadRequest{}
}

/*
GetJourneyActionmapsEstimatesJobBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActionmapsEstimatesJobBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job bad request response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job bad request response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job bad request response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job bad request response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job bad request response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyActionmapsEstimatesJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobUnauthorized creates a GetJourneyActionmapsEstimatesJobUnauthorized with default headers values
func NewGetJourneyActionmapsEstimatesJobUnauthorized() *GetJourneyActionmapsEstimatesJobUnauthorized {
	return &GetJourneyActionmapsEstimatesJobUnauthorized{}
}

/*
GetJourneyActionmapsEstimatesJobUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActionmapsEstimatesJobUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job unauthorized response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job unauthorized response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job unauthorized response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job unauthorized response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job unauthorized response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyActionmapsEstimatesJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobForbidden creates a GetJourneyActionmapsEstimatesJobForbidden with default headers values
func NewGetJourneyActionmapsEstimatesJobForbidden() *GetJourneyActionmapsEstimatesJobForbidden {
	return &GetJourneyActionmapsEstimatesJobForbidden{}
}

/*
GetJourneyActionmapsEstimatesJobForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActionmapsEstimatesJobForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job forbidden response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job forbidden response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job forbidden response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job forbidden response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job forbidden response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyActionmapsEstimatesJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobNotFound creates a GetJourneyActionmapsEstimatesJobNotFound with default headers values
func NewGetJourneyActionmapsEstimatesJobNotFound() *GetJourneyActionmapsEstimatesJobNotFound {
	return &GetJourneyActionmapsEstimatesJobNotFound{}
}

/*
GetJourneyActionmapsEstimatesJobNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyActionmapsEstimatesJobNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job not found response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job not found response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job not found response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job not found response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job not found response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyActionmapsEstimatesJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobRequestTimeout creates a GetJourneyActionmapsEstimatesJobRequestTimeout with default headers values
func NewGetJourneyActionmapsEstimatesJobRequestTimeout() *GetJourneyActionmapsEstimatesJobRequestTimeout {
	return &GetJourneyActionmapsEstimatesJobRequestTimeout{}
}

/*
GetJourneyActionmapsEstimatesJobRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyActionmapsEstimatesJobRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job request timeout response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job request timeout response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job request timeout response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job request timeout response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job request timeout response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobRequestEntityTooLarge creates a GetJourneyActionmapsEstimatesJobRequestEntityTooLarge with default headers values
func NewGetJourneyActionmapsEstimatesJobRequestEntityTooLarge() *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge {
	return &GetJourneyActionmapsEstimatesJobRequestEntityTooLarge{}
}

/*
GetJourneyActionmapsEstimatesJobRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetJourneyActionmapsEstimatesJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job request entity too large response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job request entity too large response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job request entity too large response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job request entity too large response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job request entity too large response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobUnsupportedMediaType creates a GetJourneyActionmapsEstimatesJobUnsupportedMediaType with default headers values
func NewGetJourneyActionmapsEstimatesJobUnsupportedMediaType() *GetJourneyActionmapsEstimatesJobUnsupportedMediaType {
	return &GetJourneyActionmapsEstimatesJobUnsupportedMediaType{}
}

/*
GetJourneyActionmapsEstimatesJobUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActionmapsEstimatesJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job unsupported media type response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job unsupported media type response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job unsupported media type response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job unsupported media type response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job unsupported media type response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobTooManyRequests creates a GetJourneyActionmapsEstimatesJobTooManyRequests with default headers values
func NewGetJourneyActionmapsEstimatesJobTooManyRequests() *GetJourneyActionmapsEstimatesJobTooManyRequests {
	return &GetJourneyActionmapsEstimatesJobTooManyRequests{}
}

/*
GetJourneyActionmapsEstimatesJobTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyActionmapsEstimatesJobTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job too many requests response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job too many requests response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job too many requests response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps estimates job too many requests response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps estimates job too many requests response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobInternalServerError creates a GetJourneyActionmapsEstimatesJobInternalServerError with default headers values
func NewGetJourneyActionmapsEstimatesJobInternalServerError() *GetJourneyActionmapsEstimatesJobInternalServerError {
	return &GetJourneyActionmapsEstimatesJobInternalServerError{}
}

/*
GetJourneyActionmapsEstimatesJobInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActionmapsEstimatesJobInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job internal server error response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job internal server error response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job internal server error response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps estimates job internal server error response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps estimates job internal server error response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyActionmapsEstimatesJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobServiceUnavailable creates a GetJourneyActionmapsEstimatesJobServiceUnavailable with default headers values
func NewGetJourneyActionmapsEstimatesJobServiceUnavailable() *GetJourneyActionmapsEstimatesJobServiceUnavailable {
	return &GetJourneyActionmapsEstimatesJobServiceUnavailable{}
}

/*
GetJourneyActionmapsEstimatesJobServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActionmapsEstimatesJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job service unavailable response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job service unavailable response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job service unavailable response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps estimates job service unavailable response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps estimates job service unavailable response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsEstimatesJobGatewayTimeout creates a GetJourneyActionmapsEstimatesJobGatewayTimeout with default headers values
func NewGetJourneyActionmapsEstimatesJobGatewayTimeout() *GetJourneyActionmapsEstimatesJobGatewayTimeout {
	return &GetJourneyActionmapsEstimatesJobGatewayTimeout{}
}

/*
GetJourneyActionmapsEstimatesJobGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyActionmapsEstimatesJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps estimates job gateway timeout response has a 2xx status code
func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps estimates job gateway timeout response has a 3xx status code
func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps estimates job gateway timeout response has a 4xx status code
func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps estimates job gateway timeout response has a 5xx status code
func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps estimates job gateway timeout response a status code equal to that given
func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/estimates/jobs/{jobId}][%d] getJourneyActionmapsEstimatesJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsEstimatesJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}