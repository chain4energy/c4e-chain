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

export enum CfeairdropCampaignCloseAction {
  CLOSE_ACTION_UNSPECIFIED = "CLOSE_ACTION_UNSPECIFIED",
  SEND_TO_COMMUNITY_POOL = "SEND_TO_COMMUNITY_POOL",
  BURN = "BURN",
  SEND_TO_OWNER = "SEND_TO_OWNER",
}

export interface CfeairdropclaimRecord {
  /** @format uint64 */
  campaign_id?: string;
  address?: string;
  airdrop_coins?: V1Beta1Coin[];
  completedMissions?: string[];
  claimedMissions?: string[];
}

export interface CfeairdropCampaign {
  /** @format uint64 */
  id?: string;
  owner?: string;
  name?: string;
  description?: string;
  allow_feegrant?: boolean;
  initial_claim_free_amount?: string;
  enabled?: boolean;

  /** @format date-time */
  start_time?: string;

  /** @format date-time */
  end_time?: string;

  /** period of locked coins from claim */
  lockup_period?: string;

  /** period of vesting coins after lockup period */
  vesting_period?: string;
}

export interface CfeairdropMission {
  /** @format uint64 */
  id?: string;

  /** @format uint64 */
  campaign_id?: string;
  name?: string;
  description?: string;
  missionType?: CfeairdropMissionType;
  weight?: string;

  /** @format date-time */
  claim_start_date?: string;
}

export enum CfeairdropMissionType {
  MISSION_TYPE_UNSPECIFIED = "MISSION_TYPE_UNSPECIFIED",
  INITIAL_CLAIM = "INITIAL_CLAIM",
  DELEGATION = "DELEGATION",
  VOTE = "VOTE",
  CLAIM = "CLAIM",
}

export type CfeairdropMsgAddClaimRecordsResponse = object;

export type CfeairdropMsgAddMissionToAidropCampaignResponse = object;

export type CfeairdropMsgClaimResponse = object;

export type CfeairdropMsgCloseCampaignResponse = object;

export type CfeairdropMsgCreateCampaignResponse = object;

export type CfeairdropMsgDeleteClaimRecordResponse = object;

export type CfeairdropMsgEditCampaignResponse = object;

export type CfeairdropMsgInitialClaimResponse = object;

export type CfeairdropMsgStartCampaignResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type CfeairdropParams = object;

export interface CfeairdropQueryCampaignAmountLeftResponse {
  airdrop_coins?: V1Beta1Coin[];
}

export interface CfeairdropQueryCampaignTotalAmountResponse {
  airdrop_coins?: V1Beta1Coin[];
}

export interface CfeairdropQueryCampaignResponse {
  campaign?: CfeairdropCampaign;
}

export interface CfeairdropQueryCampaignsResponse {
  campaign?: CfeairdropCampaign[];

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

export interface CfeairdropQueryMissionResponse {
  mission?: CfeairdropMission;
}

export interface CfeairdropQueryMissionsResponse {
  mission?: CfeairdropMission[];

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
export interface CfeairdropQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CfeairdropParams;
}

export interface CfeairdropQueryUsersEntriesResponse {
  userEntry?: CfeairdropUsersEntries;
}

export interface CfeairdropQueryUsersEntriesResponse {
  usersEntries?: CfeairdropUsersEntries[];

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

export interface CfeairdropUsersEntries {
  address?: string;
  claim_address?: string;
  airdrop_entries?: CfeairdropclaimRecord[];
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
   * query the next page most efficiently
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
 * @title cfeairdrop/airdrop.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaignAmountLeft
   * @summary Queries a CampaignTotalAmount by campaignId.
   * @request GET:/c4e/airdrop/v1beta1/airdrop_claims_left/{campaign_id}
   */
  queryCampaignAmountLeft = (campaignId: string, params: RequestParams = {}) =>
    this.request<CfeairdropQueryCampaignAmountLeftResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/airdrop_claims_left/${campaignId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaignTotalAmount
   * @summary Queries a CampaignTotalAmount by campaignId.
   * @request GET:/c4e/airdrop/v1beta1/airdrop_distributions/{campaign_id}
   */
  queryCampaignTotalAmount = (campaignId: string, params: RequestParams = {}) =>
    this.request<CfeairdropQueryCampaignTotalAmountResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/airdrop_distributions/${campaignId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaign
   * @summary Queries a list of Campaigns items.
   * @request GET:/c4e/airdrop/v1beta1/campaign/{campaign_id}
   */
  queryCampaign = (campaignId: string, params: RequestParams = {}) =>
    this.request<CfeairdropQueryCampaignResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/campaign/${campaignId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaigns
   * @summary Queries a list of Campaigns items.
   * @request GET:/c4e/airdrop/v1beta1/campaigns
   */
  queryCampaigns = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeairdropQueryCampaignsResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/campaigns`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMissionAll
   * @summary Queries a list of Mission items.
   * @request GET:/c4e/airdrop/v1beta1/mission
   */
  queryMissionAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeairdropQueryMissionsResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/mission`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMission
   * @summary Queries a Mission by index.
   * @request GET:/c4e/airdrop/v1beta1/mission/{campaign_id}/{mission_id}
   */
  queryMission = (campaignId: string, missionId: string, params: RequestParams = {}) =>
    this.request<CfeairdropQueryMissionResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/mission/${campaignId}/${missionId}`,
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
   * @request GET:/c4e/airdrop/v1beta1/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CfeairdropQueryParamsResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUsersEntries
   * @summary Queries a UserEntry by index.
   * @request GET:/c4e/airdrop/v1beta1/user_airdrop_entries/{address}
   */
  queryUsersEntries = (address: string, params: RequestParams = {}) =>
    this.request<CfeairdropQueryUsersEntriesResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/user_airdrop_entries/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUsersEntries
   * @summary Queries a list of UserEntry items.
   * @request GET:/c4e/airdrop/v1beta1/users_airdrop_entries
   */
  queryUsersEntries = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeairdropQueryUsersEntriesResponse, RpcStatus>({
      path: `/c4e/airdrop/v1beta1/users_airdrop_entries`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
