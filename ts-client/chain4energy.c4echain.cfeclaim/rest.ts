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

export interface CfeclaimCampaign {
  /** @format uint64 */
  id?: string;
  owner?: string;
  name?: string;
  description?: string;
  campaignType?: CfeclaimCampaignType;
  removable_claim_records?: boolean;
  feegrant_amount?: string;
  initial_claim_free_amount?: string;
  free?: string;
  enabled?: boolean;

  /** @format date-time */
  start_time?: string;

  /** @format date-time */
  end_time?: string;

  /** period of locked coins from claim */
  lockup_period?: string;

  /** period of vesting coins after lockup period */
  vesting_period?: string;
  vestingPoolName?: string;
  campaign_total_amount?: V1Beta1Coin[];
  campaign_current_amount?: V1Beta1Coin[];
}

export enum CfeclaimCampaignType {
  CAMPAIGN_TYPE_UNSPECIFIED = "CAMPAIGN_TYPE_UNSPECIFIED",
  DEFAULT = "DEFAULT",
  VESTING_POOL = "VESTING_POOL",
}

export interface CfeclaimClaimRecord {
  /** @format uint64 */
  campaign_id?: string;
  address?: string;
  amount?: V1Beta1Coin[];
  completedMissions?: string[];
  claimedMissions?: string[];
}

export interface CfeclaimClaimRecordEntry {
  /** @format uint64 */
  campaign_id?: string;
  user_entry_address?: string;
  amount?: V1Beta1Coin[];
}

export interface CfeclaimMission {
  /** @format uint64 */
  id?: string;

  /** @format uint64 */
  campaign_id?: string;
  name?: string;
  description?: string;
  missionType?: CfeclaimMissionType;
  weight?: string;

  /** @format date-time */
  claim_start_date?: string;
}

export enum CfeclaimMissionType {
  MISSION_TYPE_UNSPECIFIED = "MISSION_TYPE_UNSPECIFIED",
  INITIAL_CLAIM = "INITIAL_CLAIM",
  DELEGATE = "DELEGATE",
  VOTE = "VOTE",
  CLAIM = "CLAIM",
}

export type CfeclaimMsgAddClaimRecordsResponse = object;

export interface CfeclaimMsgAddMissionResponse {
  /** @format uint64 */
  mission_id?: string;
}

export interface CfeclaimMsgClaimResponse {
  amount?: V1Beta1Coin[];
}

export type CfeclaimMsgCloseCampaignResponse = object;

export interface CfeclaimMsgCreateCampaignResponse {
  /** @format uint64 */
  campaign_id?: string;
}

export type CfeclaimMsgDeleteClaimRecordResponse = object;

export type CfeclaimMsgEnableCampaignResponse = object;

export interface CfeclaimMsgInitialClaimResponse {
  amount?: V1Beta1Coin[];
}

export type CfeclaimMsgRemoveCampaignResponse = object;

export interface CfeclaimQueryCampaignMissionsResponse {
  missions?: CfeclaimMission[];
}

export interface CfeclaimQueryCampaignResponse {
  campaign?: CfeclaimCampaign;
}

export interface CfeclaimQueryCampaignsResponse {
  campaigns?: CfeclaimCampaign[];

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

export interface CfeclaimQueryMissionResponse {
  mission?: CfeclaimMission;
}

export interface CfeclaimQueryMissionsResponse {
  missions?: CfeclaimMission[];

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

export interface CfeclaimQueryUserEntryResponse {
  user_entry?: CfeclaimUserEntry;
}

export interface CfeclaimQueryUsersEntriesResponse {
  users_entries?: CfeclaimUserEntry[];

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

export interface CfeclaimUserEntry {
  address?: string;
  claim_records?: CfeclaimClaimRecord[];
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
 * @title c4echain/cfeclaim/campaign.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaign
   * @summary Queries a Campaign by id.
   * @request GET:/c4e/claim/v1beta1/campaign/{campaign_id}
   */
  queryCampaign = (campaignId: string, params: RequestParams = {}) =>
    this.request<CfeclaimQueryCampaignResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/campaign/${campaignId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaigns
   * @summary Queries a list of all Campaigns items.
   * @request GET:/c4e/claim/v1beta1/campaigns
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
    this.request<CfeclaimQueryCampaignsResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/campaigns`,
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
   * @summary Queries a Mission by campaign id and mission id.
   * @request GET:/c4e/claim/v1beta1/mission/{campaign_id}/{mission_id}
   */
  queryMission = (campaignId: string, missionId: string, params: RequestParams = {}) =>
    this.request<CfeclaimQueryMissionResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/mission/${campaignId}/${missionId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMissions
   * @summary Queries a list of all Missions items.
   * @request GET:/c4e/claim/v1beta1/missions
   */
  queryMissions = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeclaimQueryMissionsResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/missions`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCampaignMissions
   * @summary Queries a list of Mission items for a given campaign.
   * @request GET:/c4e/claim/v1beta1/missions/{campaign_id}
   */
  queryCampaignMissions = (campaignId: string, params: RequestParams = {}) =>
    this.request<CfeclaimQueryCampaignMissionsResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/missions/${campaignId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserEntry
   * @summary Queries a UserEntry by index.
   * @request GET:/c4e/claim/v1beta1/user_entry/{address}
   */
  queryUserEntry = (address: string, params: RequestParams = {}) =>
    this.request<CfeclaimQueryUserEntryResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/user_entry/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUsersEntries
   * @summary Queries a list of all UserEntry items.
   * @request GET:/c4e/claim/v1beta1/users_entries
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
    this.request<CfeclaimQueryUsersEntriesResponse, RpcStatus>({
      path: `/c4e/claim/v1beta1/users_entries`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });
}
