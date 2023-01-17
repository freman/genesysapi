// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsReportingReportIDMetadataReader is a Reader for the GetAnalyticsReportingReportIDMetadata structure.
type GetAnalyticsReportingReportIDMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsReportingReportIDMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsReportingReportIDMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsReportingReportIDMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsReportingReportIDMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsReportingReportIDMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsReportingReportIDMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsReportingReportIDMetadataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsReportingReportIDMetadataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsReportingReportIDMetadataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsReportingReportIDMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsReportingReportIDMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsReportingReportIDMetadataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsReportingReportIDMetadataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsReportingReportIDMetadataOK creates a GetAnalyticsReportingReportIDMetadataOK with default headers values
func NewGetAnalyticsReportingReportIDMetadataOK() *GetAnalyticsReportingReportIDMetadataOK {
	return &GetAnalyticsReportingReportIDMetadataOK{}
}

/*
GetAnalyticsReportingReportIDMetadataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsReportingReportIDMetadataOK struct {
	Payload *models.ReportMetaData
}

// IsSuccess returns true when this get analytics reporting report Id metadata o k response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics reporting report Id metadata o k response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata o k response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting report Id metadata o k response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata o k response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsReportingReportIDMetadataOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataOK) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataOK) GetPayload() *models.ReportMetaData {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportMetaData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataBadRequest creates a GetAnalyticsReportingReportIDMetadataBadRequest with default headers values
func NewGetAnalyticsReportingReportIDMetadataBadRequest() *GetAnalyticsReportingReportIDMetadataBadRequest {
	return &GetAnalyticsReportingReportIDMetadataBadRequest{}
}

/*
GetAnalyticsReportingReportIDMetadataBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsReportingReportIDMetadataBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata bad request response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata bad request response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata bad request response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata bad request response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata bad request response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsReportingReportIDMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataUnauthorized creates a GetAnalyticsReportingReportIDMetadataUnauthorized with default headers values
func NewGetAnalyticsReportingReportIDMetadataUnauthorized() *GetAnalyticsReportingReportIDMetadataUnauthorized {
	return &GetAnalyticsReportingReportIDMetadataUnauthorized{}
}

/*
GetAnalyticsReportingReportIDMetadataUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsReportingReportIDMetadataUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata unauthorized response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata unauthorized response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata unauthorized response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata unauthorized response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata unauthorized response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataForbidden creates a GetAnalyticsReportingReportIDMetadataForbidden with default headers values
func NewGetAnalyticsReportingReportIDMetadataForbidden() *GetAnalyticsReportingReportIDMetadataForbidden {
	return &GetAnalyticsReportingReportIDMetadataForbidden{}
}

/*
GetAnalyticsReportingReportIDMetadataForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsReportingReportIDMetadataForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata forbidden response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata forbidden response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata forbidden response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata forbidden response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata forbidden response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsReportingReportIDMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataNotFound creates a GetAnalyticsReportingReportIDMetadataNotFound with default headers values
func NewGetAnalyticsReportingReportIDMetadataNotFound() *GetAnalyticsReportingReportIDMetadataNotFound {
	return &GetAnalyticsReportingReportIDMetadataNotFound{}
}

/*
GetAnalyticsReportingReportIDMetadataNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsReportingReportIDMetadataNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata not found response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata not found response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata not found response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata not found response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata not found response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsReportingReportIDMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataRequestTimeout creates a GetAnalyticsReportingReportIDMetadataRequestTimeout with default headers values
func NewGetAnalyticsReportingReportIDMetadataRequestTimeout() *GetAnalyticsReportingReportIDMetadataRequestTimeout {
	return &GetAnalyticsReportingReportIDMetadataRequestTimeout{}
}

/*
GetAnalyticsReportingReportIDMetadataRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsReportingReportIDMetadataRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata request timeout response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata request timeout response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata request timeout response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata request timeout response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata request timeout response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataRequestEntityTooLarge creates a GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge with default headers values
func NewGetAnalyticsReportingReportIDMetadataRequestEntityTooLarge() *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge {
	return &GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge{}
}

/*
GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata request entity too large response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata request entity too large response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata request entity too large response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata request entity too large response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata request entity too large response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataUnsupportedMediaType creates a GetAnalyticsReportingReportIDMetadataUnsupportedMediaType with default headers values
func NewGetAnalyticsReportingReportIDMetadataUnsupportedMediaType() *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType {
	return &GetAnalyticsReportingReportIDMetadataUnsupportedMediaType{}
}

/*
GetAnalyticsReportingReportIDMetadataUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsReportingReportIDMetadataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata unsupported media type response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata unsupported media type response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata unsupported media type response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata unsupported media type response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata unsupported media type response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataTooManyRequests creates a GetAnalyticsReportingReportIDMetadataTooManyRequests with default headers values
func NewGetAnalyticsReportingReportIDMetadataTooManyRequests() *GetAnalyticsReportingReportIDMetadataTooManyRequests {
	return &GetAnalyticsReportingReportIDMetadataTooManyRequests{}
}

/*
GetAnalyticsReportingReportIDMetadataTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsReportingReportIDMetadataTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata too many requests response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata too many requests response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata too many requests response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics reporting report Id metadata too many requests response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics reporting report Id metadata too many requests response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataInternalServerError creates a GetAnalyticsReportingReportIDMetadataInternalServerError with default headers values
func NewGetAnalyticsReportingReportIDMetadataInternalServerError() *GetAnalyticsReportingReportIDMetadataInternalServerError {
	return &GetAnalyticsReportingReportIDMetadataInternalServerError{}
}

/*
GetAnalyticsReportingReportIDMetadataInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsReportingReportIDMetadataInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata internal server error response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata internal server error response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata internal server error response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting report Id metadata internal server error response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting report Id metadata internal server error response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataServiceUnavailable creates a GetAnalyticsReportingReportIDMetadataServiceUnavailable with default headers values
func NewGetAnalyticsReportingReportIDMetadataServiceUnavailable() *GetAnalyticsReportingReportIDMetadataServiceUnavailable {
	return &GetAnalyticsReportingReportIDMetadataServiceUnavailable{}
}

/*
GetAnalyticsReportingReportIDMetadataServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsReportingReportIDMetadataServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata service unavailable response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata service unavailable response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata service unavailable response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting report Id metadata service unavailable response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting report Id metadata service unavailable response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsReportingReportIDMetadataGatewayTimeout creates a GetAnalyticsReportingReportIDMetadataGatewayTimeout with default headers values
func NewGetAnalyticsReportingReportIDMetadataGatewayTimeout() *GetAnalyticsReportingReportIDMetadataGatewayTimeout {
	return &GetAnalyticsReportingReportIDMetadataGatewayTimeout{}
}

/*
GetAnalyticsReportingReportIDMetadataGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsReportingReportIDMetadataGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics reporting report Id metadata gateway timeout response has a 2xx status code
func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics reporting report Id metadata gateway timeout response has a 3xx status code
func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics reporting report Id metadata gateway timeout response has a 4xx status code
func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics reporting report Id metadata gateway timeout response has a 5xx status code
func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics reporting report Id metadata gateway timeout response a status code equal to that given
func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/reporting/{reportId}/metadata][%d] getAnalyticsReportingReportIdMetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsReportingReportIDMetadataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
