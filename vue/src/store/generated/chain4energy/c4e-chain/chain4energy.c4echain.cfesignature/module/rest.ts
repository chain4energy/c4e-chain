/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface CfesignatureMsgCreateAccountResponse {
  accountNumber?: string;
}

export interface CfesignatureMsgPublishReferencePayloadLinkResponse {
  txTimestamp?: string;
}

export interface CfesignatureMsgStoreSignatureResponse {
  txId?: string;
  txTimestamp?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type CfesignatureParams = object;

export interface CfesignatureQueryCreateReferenceIdResponse {
  referenceId?: string;
}

export interface CfesignatureQueryCreateReferencePayloadLinkResponse {
  referenceKey?: string;
  referenceValue?: string;
}

export interface CfesignatureQueryCreateStorageKeyResponse {
  storageKey?: string;
}

export interface CfesignatureQueryGetAccountInfoResponse {
  accAddress?: string;
  pubKey?: string;
}

export interface CfesignatureQueryGetReferencePayloadLinkResponse {
  referencePayloadLinkValue?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CfesignatureQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CfesignatureParams;
}

export interface CfesignatureQueryVerifyReferencePayloadLinkResponse {
  isValid?: boolean;
}

export interface CfesignatureQueryVerifySignatureResponse {
  signature?: string;
  algorithm?: string;
  certificate?: string;
  timestamp?: string;
  valid?: string;
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

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

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

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title cfesignature/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCreateReferenceId
   * @summary Queries a list of CreateReferenceId items.
   * @request GET:/c4e/signature/v1beta1/create_reference_id/{creator}
   */
  queryCreateReferenceId = (creator: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryCreateReferenceIdResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/create_reference_id/${creator}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCreateReferencePayloadLink
   * @summary Queries a list of CreateReferencePayloadLink items.
   * @request GET:/c4e/signature/v1beta1/create_reference_payload_link/{referenceId}/{payloadHash}
   */
  queryCreateReferencePayloadLink = (referenceId: string, payloadHash: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryCreateReferencePayloadLinkResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/create_reference_payload_link/${referenceId}/${payloadHash}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCreateStorageKey
   * @summary Queries a list of CreateStorageKey items.
   * @request GET:/c4e/signature/v1beta1/create_storage_key/{targetAccAddress}/{referenceId}
   */
  queryCreateStorageKey = (targetAccAddress: string, referenceId: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryCreateStorageKeyResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/create_storage_key/${targetAccAddress}/${referenceId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetAccountInfo
   * @summary Queries a list of GetAccountInfo items.
   * @request GET:/c4e/signature/v1beta1/get_account_info/{accAddressString}
   */
  queryGetAccountInfo = (accAddressString: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryGetAccountInfoResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/get_account_info/${accAddressString}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryGetReferencePayloadLink
   * @summary Queries a list of GetReferencePayloadLink items.
   * @request GET:/c4e/signature/v1beta1/get_reference_payload_link/{referenceId}
   */
  queryGetReferencePayloadLink = (referenceId: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryGetReferencePayloadLinkResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/get_reference_payload_link/${referenceId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/c4e/signature/v1beta1/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CfesignatureQueryParamsResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVerifyReferencePayloadLink
   * @summary Queries a list of VerifyReferencePayloadLink items.
   * @request GET:/c4e/signature/v1beta1/verify_reference_payload_link/{referenceId}/{payloadHash}
   */
  queryVerifyReferencePayloadLink = (referenceId: string, payloadHash: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryVerifyReferencePayloadLinkResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/verify_reference_payload_link/${referenceId}/${payloadHash}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVerifySignature
   * @summary Queries a list of VerifySignature items.
   * @request GET:/c4e/signature/v1beta1/verify_signature/{referenceId}/{targetAccAddress}
   */
  queryVerifySignature = (referenceId: string, targetAccAddress: string, params: RequestParams = {}) =>
    this.request<CfesignatureQueryVerifySignatureResponse, RpcStatus>({
      path: `/c4e/signature/v1beta1/verify_signature/${referenceId}/${targetAccAddress}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
