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

export interface CfetokenizationCertificate {
  /** @format uint64 */
  id?: string;

  /** @format uint64 */
  certyficate_type_id?: string;

  /** @format uint64 */
  power?: string;
  device_address?: string;
  measurements?: string[];
  allowed_authorities?: string[];
  authority?: string;
  certificate_status?: CfetokenizationCertificateStatus;

  /** @format date-time */
  valid_until?: string;
}

export interface CfetokenizationCertificateOffer {
  /** @format uint64 */
  id?: string;

  /** @format uint64 */
  certificate_id?: string;
  owner?: string;
  buyer?: string;
  price?: V1Beta1Coin[];
  authorizer?: string;

  /** @format uint64 */
  power?: string;
}

export enum CfetokenizationCertificateStatus {
  UNKNOWN_CERTIFICATE_STATUS = "UNKNOWN_CERTIFICATE_STATUS",
  VALID = "VALID",
  INVALID = "INVALID",
  ON_MARKETPLACE = "ON_MARKETPLACE",
  BURNED = "BURNED",
}

export interface CfetokenizationCertificateType {
  /** @format uint64 */
  id?: string;
  name?: string;
  description?: string;
}

export interface CfetokenizationDevice {
  device_address?: string;
  measurements?: CfetokenizationMeasurement[];

  /** @format uint64 */
  active_power_sum?: string;

  /** @format uint64 */
  reverse_power_sum?: string;

  /** @format uint64 */
  used_active_power?: string;

  /** @format uint64 */
  fulfilled_reverse_power?: string;
}

export interface CfetokenizationMeasurement {
  /** @format uint64 */
  id?: string;

  /** @format date-time */
  timestamp?: string;

  /** @format uint64 */
  active_power?: string;
  used_for_certificate?: boolean;

  /** @format uint64 */
  reverse_power?: string;
  metadata?: string;
}

export type CfetokenizationMsgAcceptDeviceResponse = object;

export type CfetokenizationMsgAddCertificateToMarketplaceResponse = object;

export type CfetokenizationMsgAddMeasurementResponse = object;

export type CfetokenizationMsgAssignDeviceToUserResponse = object;

export type CfetokenizationMsgAuthorizeCertificateResponse = object;

export type CfetokenizationMsgBurnCertificateResponse = object;

export type CfetokenizationMsgBuyCertificateResponse = object;

export type CfetokenizationMsgCreateUserCertificatesResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type CfetokenizationParams = object;

export interface CfetokenizationQueryAllCertificateTypeResponse {
  CertificateType?: CfetokenizationCertificateType[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfetokenizationQueryAllUserCertificatesResponse {
  UserCertificates?: CfetokenizationUserCertificates[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfetokenizationQueryAllUserDevicesResponse {
  UserDevices?: CfetokenizationUserDevices[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfetokenizationQueryDeviceAllResponse {
  devices?: CfetokenizationDevice[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfetokenizationQueryDeviceResponse {
  device?: CfetokenizationDevice;
}

export interface CfetokenizationQueryGetCertificateTypeResponse {
  CertificateType?: CfetokenizationCertificateType;
}

export interface CfetokenizationQueryGetUserCertificatesResponse {
  UserCertificates?: CfetokenizationUserCertificates;
}

export interface CfetokenizationQueryGetUserDevicesResponse {
  UserDevices?: CfetokenizationUserDevices;
}

export interface CfetokenizationQueryMarketplaceCertificateResponse {
  marketplace_certificate?: CfetokenizationCertificateOffer;
}

export interface CfetokenizationQueryMarketplaceCertificatesAllResponse {
  marketplace_certificates?: CfetokenizationCertificateOffer[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CfetokenizationQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CfetokenizationParams;
}

export interface CfetokenizationUserCertificates {
  owner?: string;
  certificates?: CfetokenizationCertificate[];
}

export interface CfetokenizationUserDevice {
  device_address?: string;
  name?: string;
  location?: string;
}

export interface CfetokenizationUserDevices {
  owner?: string;
  devices?: CfetokenizationUserDevice[];
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

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title c4echain/cfetokenization/certificate_type.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCertificateTypeAll
   * @request GET:/c4e/tokenization/v1beta1/certificate_type
   */
  queryCertificateTypeAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfetokenizationQueryAllCertificateTypeResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/certificate_type`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCertificateType
   * @summary Queries a list of CertificateType items.
   * @request GET:/c4e/tokenization/v1beta1/certificate_type/{id}
   */
  queryCertificateType = (id: string, params: RequestParams = {}) =>
    this.request<CfetokenizationQueryGetCertificateTypeResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/certificate_type/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDevice
   * @request GET:/c4e/tokenization/v1beta1/device/{device_address}
   */
  queryDevice = (deviceAddress: string, params: RequestParams = {}) =>
    this.request<CfetokenizationQueryDeviceResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/device/${deviceAddress}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDeviceAll
   * @request GET:/c4e/tokenization/v1beta1/devices
   */
  queryDeviceAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfetokenizationQueryDeviceAllResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/devices`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMarketplaceCertificate
   * @request GET:/c4e/tokenization/v1beta1/marketplace_certificate/{id}
   */
  queryMarketplaceCertificate = (id: string, params: RequestParams = {}) =>
    this.request<CfetokenizationQueryMarketplaceCertificateResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/marketplace_certificate/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMarketplaceCertificatesAll
   * @request GET:/c4e/tokenization/v1beta1/marketplace_certificates
   */
  queryMarketplaceCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfetokenizationQueryMarketplaceCertificatesAllResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/marketplace_certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/c4e/tokenization/v1beta1/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CfetokenizationQueryParamsResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserCertificatesAll
   * @request GET:/c4e/tokenization/v1beta1/user_certificates
   */
  queryUserCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfetokenizationQueryAllUserCertificatesResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/user_certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserCertificates
   * @summary Queries a list of UserCertificates items.
   * @request GET:/c4e/tokenization/v1beta1/user_certificates/{owner}
   */
  queryUserCertificates = (owner: string, params: RequestParams = {}) =>
    this.request<CfetokenizationQueryGetUserCertificatesResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/user_certificates/${owner}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserDevicesAll
   * @request GET:/c4e/tokenization/v1beta1/user_devices
   */
  queryUserDevicesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfetokenizationQueryAllUserDevicesResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/user_devices`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserDevices
   * @summary Queries a list of UserDevices items.
   * @request GET:/c4e/tokenization/v1beta1/user_devices/{owner}
   */
  queryUserDevices = (owner: string, params: RequestParams = {}) =>
    this.request<CfetokenizationQueryGetUserDevicesResponse, RpcStatus>({
      path: `/c4e/tokenization/v1beta1/user_devices/${owner}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
