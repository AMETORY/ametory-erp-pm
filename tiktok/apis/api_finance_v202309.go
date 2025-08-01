/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apis

import (
    "bytes"
    "context"
    "io"
    "net/http"
    "net/url"
    "strings"
    "reflect"

    "tiktokshop/open/sdk_golang/models/finance/v202309"
)


// FinanceV202309APIService FinanceV202309API service
type FinanceV202309APIService service

type ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest struct {
    ctx context.Context
    ApiService *FinanceV202309APIService
    orderId string
    xTtsAccessToken *string
    contentType *string
    shopCipher *string
}

// 
func (r ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest) ContentType(contentType string) ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest {
    r.contentType = &contentType
    return r
}
// 
func (r ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest) ShopCipher(shopCipher string) ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest) Execute() (*finance_v202309.Finance202309GetTransactionsbyOrderResponse, *http.Response, error) {
    return r.ApiService.Finance202309OrdersOrderIdStatementTransactionsGetExecute(r)
}

/*
Finance202309OrdersOrderIdStatementTransactionsGet GetTransactionsbyOrder
**This API is currently exclusive to the following markets: US, UK.**
Retrieves the transactions associated with an order, including both order-level transactions and SKU-level detailed transactions. This covers all transactions related to sales, fees, commissions, shipping, taxes, adjustments, and refunds.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId The order ID in TikTok Shop.
@return ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest
*/
func (a *FinanceV202309APIService) Finance202309OrdersOrderIdStatementTransactionsGet(ctx context.Context, orderId string) ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest {
    return ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest{
        ApiService: a,
        ctx: ctx,
        orderId: orderId,
    }
}

// Execute executes the request
//  @return Finance202309GetTransactionsbyOrderResponse
func (a *FinanceV202309APIService) Finance202309OrdersOrderIdStatementTransactionsGetExecute(r ApiFinance202309OrdersOrderIdStatementTransactionsGetRequest) (*finance_v202309.Finance202309GetTransactionsbyOrderResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *finance_v202309.Finance202309GetTransactionsbyOrderResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceV202309APIService.Finance202309OrdersOrderIdStatementTransactionsGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/finance/202309/orders/{order_id}/statement_transactions"
    localVarPath = strings.Replace(localVarPath, "{"+"order_id"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFinance202309PaymentsGetRequest struct {
    ctx context.Context
    ApiService *FinanceV202309APIService
    sortField *string
    xTtsAccessToken *string
    contentType *string
    createTimeLt *int64
    pageSize *interface{}
    pageToken *string
    sortOrder *string
    createTimeGe *int64
    shopCipher *string
}

// The returned results will be sorted by the specified field. Only supports &#x60;create_time&#x60;.
func (r ApiFinance202309PaymentsGetRequest) SortField(sortField string) ApiFinance202309PaymentsGetRequest {
    r.sortField = &sortField
    return r
}
// 
func (r ApiFinance202309PaymentsGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiFinance202309PaymentsGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFinance202309PaymentsGetRequest) ContentType(contentType string) ApiFinance202309PaymentsGetRequest {
    r.contentType = &contentType
    return r
}
// Filter payments to show only those that occurred before the specified date and time. Unix timestamp. Refer to notes in &#x60;create_time_ge&#x60; for more usage information.
func (r ApiFinance202309PaymentsGetRequest) CreateTimeLt(createTimeLt int64) ApiFinance202309PaymentsGetRequest {
    r.createTimeLt = &createTimeLt
    return r
}
// The number of results to be returned per page.  Default: 20 Valid range: [1-100]
func (r ApiFinance202309PaymentsGetRequest) PageSize(pageSize interface{}) ApiFinance202309PaymentsGetRequest {
    r.pageSize = &pageSize
    return r
}
// An opaque token used to retrieve the next page of a paginated result set. Retrieve this value from the result of the &#x60;next_page_token&#x60; from a previous response. It is not needed for the first page.
func (r ApiFinance202309PaymentsGetRequest) PageToken(pageToken string) ApiFinance202309PaymentsGetRequest {
    r.pageToken = &pageToken
    return r
}
// The sort order for the &#x60;sort_field&#x60; parameter.  Default: ASC  Possible values: - ASC: Ascending order - DESC: Descending order
func (r ApiFinance202309PaymentsGetRequest) SortOrder(sortOrder string) ApiFinance202309PaymentsGetRequest {
    r.sortOrder = &sortOrder
    return r
}
// Filter payments to show only those that occurred on or after the specified date and time. Unix timestamp.  **Note:** &#x60;create_time_ge&#x60; and &#x60;create_time_lt&#x60; together constitute the creation time filter condition. - If &#x60;create_time_ge&#x60; is filled but &#x60;create_time_lt&#x60; is empty, &#x60;create_time_lt&#x60; will default to the current time. - If &#x60;create_time_lt&#x60; is filled but &#x60;create_time_ge&#x60; is empty, &#x60;create_time_ge&#x60; will default to the earliest shop time.
func (r ApiFinance202309PaymentsGetRequest) CreateTimeGe(createTimeGe int64) ApiFinance202309PaymentsGetRequest {
    r.createTimeGe = &createTimeGe
    return r
}
// 
func (r ApiFinance202309PaymentsGetRequest) ShopCipher(shopCipher string) ApiFinance202309PaymentsGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFinance202309PaymentsGetRequest) Execute() (*finance_v202309.Finance202309GetPaymentsResponse, *http.Response, error) {
    return r.ApiService.Finance202309PaymentsGetExecute(r)
}

/*
Finance202309PaymentsGet GetPayments
**This API is currently unavailable to SEA markets.**
Retrieves records of automated payments for a shop based on a specified date range.
Use the returned list to verify and reconcile payments with the transactions in the seller's bank account.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFinance202309PaymentsGetRequest
*/
func (a *FinanceV202309APIService) Finance202309PaymentsGet(ctx context.Context) ApiFinance202309PaymentsGetRequest {
    return ApiFinance202309PaymentsGetRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Finance202309GetPaymentsResponse
func (a *FinanceV202309APIService) Finance202309PaymentsGetExecute(r ApiFinance202309PaymentsGetRequest) (*finance_v202309.Finance202309GetPaymentsResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *finance_v202309.Finance202309GetPaymentsResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceV202309APIService.Finance202309PaymentsGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/finance/202309/payments"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.sortField == nil {
        return localVarReturnValue, nil, reportError("sortField is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.createTimeLt != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "create_time_lt", r.createTimeLt, "")
    }
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
    if r.sortOrder != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_order", r.sortOrder, "")
    }
    if r.createTimeGe != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "create_time_ge", r.createTimeGe, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFinance202309StatementsGetRequest struct {
    ctx context.Context
    ApiService *FinanceV202309APIService
    sortField *string
    xTtsAccessToken *string
    contentType *string
    statementTimeLt *int64
    paymentStatus *string
    pageSize *interface{}
    pageToken *string
    sortOrder *string
    statementTimeGe *int64
    shopCipher *string
}

// The returned results will be sorted by the specified field. Only supports &#x60;statement_time&#x60;.
func (r ApiFinance202309StatementsGetRequest) SortField(sortField string) ApiFinance202309StatementsGetRequest {
    r.sortField = &sortField
    return r
}
// 
func (r ApiFinance202309StatementsGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiFinance202309StatementsGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFinance202309StatementsGetRequest) ContentType(contentType string) ApiFinance202309StatementsGetRequest {
    r.contentType = &contentType
    return r
}
// Filter statements to show only those that are generated before the specified date and time. Unix timestamp. Refer to notes in &#x60;statement_time_ge&#x60; for more usage information.
func (r ApiFinance202309StatementsGetRequest) StatementTimeLt(statementTimeLt int64) ApiFinance202309StatementsGetRequest {
    r.statementTimeLt = &statementTimeLt
    return r
}
// Filter statements based on the payment status. Possible values: - PAID: Payment has been transferred to the seller. - FAILED: Payment transfer failed. - PROCESSING: Payment is currently being processed. Default: All statuses are returned.
func (r ApiFinance202309StatementsGetRequest) PaymentStatus(paymentStatus string) ApiFinance202309StatementsGetRequest {
    r.paymentStatus = &paymentStatus
    return r
}
// The number of results to be returned per page.  Default: 20 Valid range: [1-100]
func (r ApiFinance202309StatementsGetRequest) PageSize(pageSize interface{}) ApiFinance202309StatementsGetRequest {
    r.pageSize = &pageSize
    return r
}
// An opaque token used to retrieve the next page of a paginated result set. Retrieve this value from the result of the &#x60;next_page_token&#x60; from a previous response. It is not needed for the first page.
func (r ApiFinance202309StatementsGetRequest) PageToken(pageToken string) ApiFinance202309StatementsGetRequest {
    r.pageToken = &pageToken
    return r
}
// The sort order for the &#x60;sort_field&#x60; parameter.  Default: ASC  Possible values: - ASC: Ascending order - DESC: Descending order
func (r ApiFinance202309StatementsGetRequest) SortOrder(sortOrder string) ApiFinance202309StatementsGetRequest {
    r.sortOrder = &sortOrder
    return r
}
// Filter statements to show only those that are generated on or after the specified date and time. Unix timestamp.  **Note:** &#x60;statement_time_ge&#x60; and &#x60;statement_time_le&#x60; together constitute the creation time filter condition. - If &#x60;statement_time_ge&#x60; is filled but &#x60;statement_time_lt&#x60; is empty, &#x60;statement_time_lt&#x60; will default to the current time. - If &#x60;statement_time_lt&#x60; is filled but &#x60;statement_time_ge&#x60; is empty, &#x60;statement_time_ge&#x60; will default to the earliest shop time.  **Example:** As statements are generated daily at 00:00 UTC, to retrieve statements for the period from Oct 5 to Oct 10, configure the parameters as follows: - Set &#x60;statement_time_ge&#x60; to 00:00 on Oct 6  or any time on Oct 5 (excluding 00:00). - Set &#x60;statement_time_lt&#x60; to any time on Oct 11 (excluding 00:00).
func (r ApiFinance202309StatementsGetRequest) StatementTimeGe(statementTimeGe int64) ApiFinance202309StatementsGetRequest {
    r.statementTimeGe = &statementTimeGe
    return r
}
// 
func (r ApiFinance202309StatementsGetRequest) ShopCipher(shopCipher string) ApiFinance202309StatementsGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFinance202309StatementsGetRequest) Execute() (*finance_v202309.Finance202309GetStatementsResponse, *http.Response, error) {
    return r.ApiService.Finance202309StatementsGetExecute(r)
}

/*
Finance202309StatementsGet GetStatements
Retrieves the statements generated for a shop and the key statement information based on a specified date range or their payment status. Use this API to get an overview of your daily statements over a range of time, or to find out which statements have been paid or not. For the detailed transactions, refer to [Get Statement Transactions](650a6749defece02be67da87) or [Get Order Statement Transactions](650a6734defece02be67d724).
Applicable for all regions' sellers. Only data after 2023-07-01 is available.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFinance202309StatementsGetRequest
*/
func (a *FinanceV202309APIService) Finance202309StatementsGet(ctx context.Context) ApiFinance202309StatementsGetRequest {
    return ApiFinance202309StatementsGetRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Finance202309GetStatementsResponse
func (a *FinanceV202309APIService) Finance202309StatementsGetExecute(r ApiFinance202309StatementsGetRequest) (*finance_v202309.Finance202309GetStatementsResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *finance_v202309.Finance202309GetStatementsResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceV202309APIService.Finance202309StatementsGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/finance/202309/statements"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.sortField == nil {
        return localVarReturnValue, nil, reportError("sortField is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.statementTimeLt != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "statement_time_lt", r.statementTimeLt, "")
    }
    if r.paymentStatus != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "payment_status", r.paymentStatus, "")
    }
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
    if r.sortOrder != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_order", r.sortOrder, "")
    }
    if r.statementTimeGe != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "statement_time_ge", r.statementTimeGe, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest struct {
    ctx context.Context
    ApiService *FinanceV202309APIService
    statementId string
    sortField *string
    xTtsAccessToken *string
    contentType *string
    pageToken *string
    sortOrder *string
    pageSize *interface{}
    shopCipher *string
}

// Only support: order_create_time
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) SortField(sortField string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.sortField = &sortField
    return r
}
// 
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) ContentType(contentType string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.contentType = &contentType
    return r
}
// The default is empty string
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) PageToken(pageToken string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.pageToken = &pageToken
    return r
}
// The default is ASC, the developer can choose ASC or DESC
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) SortOrder(sortOrder string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.sortOrder = &sortOrder
    return r
}
// The default is 20. It must be a positive integer,the range is 1-100
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) PageSize(pageSize interface{}) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.pageSize = &pageSize
    return r
}
// 
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) ShopCipher(shopCipher string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) Execute() (*finance_v202309.Finance202309GetTransactionsbyStatementResponse, *http.Response, error) {
    return r.ApiService.Finance202309StatementsStatementIdStatementTransactionsGetExecute(r)
}

/*
Finance202309StatementsStatementIdStatementTransactionsGet GetTransactionsbyStatement
Only for UK and US local sellers. Get a list of transactions based on statement_id. We will return a list of orders. If you require the SKU level transaction details, pass in the order_id to Get Order Statement Transactions.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param statementId The unique id of statement
@return ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest
*/
func (a *FinanceV202309APIService) Finance202309StatementsStatementIdStatementTransactionsGet(ctx context.Context, statementId string) ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest {
    return ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest{
        ApiService: a,
        ctx: ctx,
        statementId: statementId,
    }
}

// Execute executes the request
//  @return Finance202309GetTransactionsbyStatementResponse
func (a *FinanceV202309APIService) Finance202309StatementsStatementIdStatementTransactionsGetExecute(r ApiFinance202309StatementsStatementIdStatementTransactionsGetRequest) (*finance_v202309.Finance202309GetTransactionsbyStatementResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *finance_v202309.Finance202309GetTransactionsbyStatementResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceV202309APIService.Finance202309StatementsStatementIdStatementTransactionsGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/finance/202309/statements/{statement_id}/statement_transactions"
    localVarPath = strings.Replace(localVarPath, "{"+"statement_id"+"}", url.PathEscape(parameterValueToString(r.statementId, "statementId")), -1)

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.sortField == nil {
        return localVarReturnValue, nil, reportError("sortField is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
    if r.sortOrder != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "sort_order", r.sortOrder, "")
    }
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFinance202309WithdrawalsGetRequest struct {
    ctx context.Context
    ApiService *FinanceV202309APIService
    types *[]string
    xTtsAccessToken *string
    contentType *string
    createTimeLt *int64
    pageSize *int32
    pageToken *string
    createTimeGe *int64
    shopCipher *string
}

// The type of transaction. Possible values: - WITHDRAW：The action of the seller to receive the settlement amount to the bank card through the action of withdrawal - SETTLE：The platform settles the amount to the seller - TRANSFER：Platform subsidies or deductions due to platform policies - REVERSE：Withdrawal failure due to incorrect bank card 
func (r ApiFinance202309WithdrawalsGetRequest) Types(types []string) ApiFinance202309WithdrawalsGetRequest {
    r.types = &types
    return r
}
// 
func (r ApiFinance202309WithdrawalsGetRequest) XTtsAccessToken(xTtsAccessToken string) ApiFinance202309WithdrawalsGetRequest {
    r.xTtsAccessToken = &xTtsAccessToken
    return r
}
// Allowed type: application/json
func (r ApiFinance202309WithdrawalsGetRequest) ContentType(contentType string) ApiFinance202309WithdrawalsGetRequest {
    r.contentType = &contentType
    return r
}
// Unix timestamp representing the end of transactions time range one wants to request
func (r ApiFinance202309WithdrawalsGetRequest) CreateTimeLt(createTimeLt int64) ApiFinance202309WithdrawalsGetRequest {
    r.createTimeLt = &createTimeLt
    return r
}
// The default is 20, it must be positive integer,the range is 1-100
func (r ApiFinance202309WithdrawalsGetRequest) PageSize(pageSize int32) ApiFinance202309WithdrawalsGetRequest {
    r.pageSize = &pageSize
    return r
}
// The next page token
func (r ApiFinance202309WithdrawalsGetRequest) PageToken(pageToken string) ApiFinance202309WithdrawalsGetRequest {
    r.pageToken = &pageToken
    return r
}
// Unix timestamp representing the start of transactions time range one wants to request
func (r ApiFinance202309WithdrawalsGetRequest) CreateTimeGe(createTimeGe int64) ApiFinance202309WithdrawalsGetRequest {
    r.createTimeGe = &createTimeGe
    return r
}
// 
func (r ApiFinance202309WithdrawalsGetRequest) ShopCipher(shopCipher string) ApiFinance202309WithdrawalsGetRequest {
    r.shopCipher = &shopCipher
    return r
}
func (r ApiFinance202309WithdrawalsGetRequest) Execute() (*finance_v202309.Finance202309GetWithdrawalsResponse, *http.Response, error) {
    return r.ApiService.Finance202309WithdrawalsGetExecute(r)
}

/*
Finance202309WithdrawalsGet GetWithdrawals
Get the list of the withdrawal records (when Seller's withdraw money from TikTokShop) based on the specified date range. 

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFinance202309WithdrawalsGetRequest
*/
func (a *FinanceV202309APIService) Finance202309WithdrawalsGet(ctx context.Context) ApiFinance202309WithdrawalsGetRequest {
    return ApiFinance202309WithdrawalsGetRequest{
        ApiService: a,
        ctx: ctx,
    }
}

// Execute executes the request
//  @return Finance202309GetWithdrawalsResponse
func (a *FinanceV202309APIService) Finance202309WithdrawalsGetExecute(r ApiFinance202309WithdrawalsGetRequest) (*finance_v202309.Finance202309GetWithdrawalsResponse, *http.Response, error) {
    var (
        localVarHTTPMethod   = http.MethodGet
        localVarPostBody     interface{}
        formFiles            []formFile
        localVarReturnValue  *finance_v202309.Finance202309GetWithdrawalsResponse
    )

    localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinanceV202309APIService.Finance202309WithdrawalsGet")
    if err != nil {
        return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
    }

    localVarPath := localBasePath + "/finance/202309/withdrawals"

    localVarHeaderParams := make(map[string]string)
    localVarQueryParams := url.Values{}
    localVarFormParams := url.Values{}
    if r.types == nil {
        return localVarReturnValue, nil, reportError("types is required and must be specified")
    }
    if r.xTtsAccessToken == nil {
        return localVarReturnValue, nil, reportError("xTtsAccessToken is required and must be specified")
    }
    if r.contentType == nil {
        return localVarReturnValue, nil, reportError("contentType is required and must be specified")
    }

    if r.createTimeLt != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "create_time_lt", r.createTimeLt, "")
    }
    {
        t := *r.types
        if reflect.TypeOf(t).Kind() == reflect.Slice {
            s := reflect.ValueOf(t)
            for i := 0; i < s.Len(); i++ {
                parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i).Interface(), "multi")
            }
        } else {
            parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
        }
    }
    if r.pageSize != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
    }
    if r.pageToken != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "")
    }
    if r.createTimeGe != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "create_time_ge", r.createTimeGe, "")
    }
    if r.shopCipher != nil {
    parameterAddToHeaderOrQuery(localVarQueryParams, "shop_cipher", r.shopCipher, "")
    }
    // to determine the Content-Type header
    localVarHTTPContentTypes := []string{}

    // set Content-Type header
    localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
    if localVarHTTPContentType != "" {
        localVarHeaderParams["Content-Type"] = localVarHTTPContentType
    }

    // to determine the Accept header
    localVarHTTPHeaderAccepts := []string{"application/json"}

    // set Accept header
    localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
    if localVarHTTPHeaderAccept != "" {
        localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
    }
    parameterAddToHeaderOrQuery(localVarHeaderParams, "x-tts-access-token", r.xTtsAccessToken, "")
    parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
    req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
    if err != nil {
        return localVarReturnValue, nil, err
    }

    localVarHTTPResponse, err := a.client.callAPI(req)
    if err != nil || localVarHTTPResponse == nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
    localVarHTTPResponse.Body.Close()
    localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
    if err != nil {
        return localVarReturnValue, localVarHTTPResponse, err
    }

    if localVarHTTPResponse.StatusCode >= 300 {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: localVarHTTPResponse.Status,
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
    if err != nil {
        newErr := &GenericOpenAPIError{
            body:  localVarBody,
            error: err.Error(),
        }
        return localVarReturnValue, localVarHTTPResponse, newErr
    }

    return localVarReturnValue, localVarHTTPResponse, nil
}
