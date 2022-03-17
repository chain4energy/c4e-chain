export interface CfevestingMsgBeginRedelegateResponse {
    /** @format date-time */
    completionTime?: string;
}
export declare type CfevestingMsgDelegateResponse = object;
export declare type CfevestingMsgSendVestingResponse = object;
export interface CfevestingMsgUndelegateResponse {
    /** @format date-time */
    completionTime?: string;
}
export declare type CfevestingMsgVestResponse = object;
export declare type CfevestingMsgWithdrawAllAvailableResponse = object;
export declare type CfevestingMsgWithdrawDelegatorRewardResponse = object;
/**
 * Params defines the parameters for the module.
 */
export interface CfevestingParams {
    denom?: string;
}
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CfevestingQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: CfevestingParams;
}
export interface CfevestingQueryVestingResponse {
    delegableAddress?: string;
    vestings?: CfevestingVestingInfo[];
}
export interface CfevestingQueryVestingTypeResponse {
    vestingTypes?: CfevestingVestingTypes;
}
export interface CfevestingVestingInfo {
    /** @format int32 */
    id?: number;
    vestingType?: string;
    /** @format int64 */
    vestingStartHeight?: string;
    /** @format int64 */
    lockEndHeight?: string;
    /** @format int64 */
    vestingEndHeight?: string;
    withdrawable?: string;
    delegationAllowed?: boolean;
    /**
     * Coin defines a token with a denomination and an amount.
     *
     * NOTE: The amount field is an Int which implements the custom method
     * signatures required by gogoproto.
     */
    vested?: V1Beta1Coin;
    currentVestedAmount?: string;
}
export interface CfevestingVestingType {
    name?: string;
    /** @format int64 */
    lockupPeriod?: string;
    /** @format int64 */
    vestingPeriod?: string;
    /** @format int64 */
    tokenReleasingPeriod?: string;
    delegationsAllowed?: boolean;
}
export interface CfevestingVestingTypes {
    vestingTypes?: CfevestingVestingType[];
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
    denom?: string;
    amount?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title cfevesting/account_vesting.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/chain4energy/vesting/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<CfevestingQueryParamsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryVesting
     * @summary Queries a list of Vesting items.
     * @request GET:/chain4energy/vesting/vesting/{address}
     */
    queryVesting: (address: string, params?: RequestParams) => Promise<HttpResponse<CfevestingQueryVestingResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryVestingType
     * @summary Queries a list of VestingType items.
     * @request GET:/chain4energy/vesting/vesting_type
     */
    queryVestingType: (params?: RequestParams) => Promise<HttpResponse<CfevestingQueryVestingTypeResponse, RpcStatus>>;
}
export {};
