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
 * @title c4echain/cfesignature/genesis.proto
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
