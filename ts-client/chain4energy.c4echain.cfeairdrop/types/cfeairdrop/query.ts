/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Campaign, Mission, UserEntry } from "./airdrop";
import { Params } from "./params";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryUsersEntriesRequest {
  address: string;
}

export interface QueryUsersEntriesResponse {
  userEntry: UserEntry | undefined;
}

export interface QueryCampaignTotalAmountRequest {
  campaignId: number;
}

export interface QueryCampaignTotalAmountResponse {
  amount: Coin[];
}

export interface QueryCampaignAmountLeftRequest {
  campaignId: number;
}

export interface QueryCampaignAmountLeftResponse {
  amount: Coin[];
}

export interface QueryUsersEntriesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryUsersEntriesResponse {
  usersEntries: UserEntry[];
  pagination: PageResponse | undefined;
}

export interface QueryMissionRequest {
  campaignId: number;
  missionId: number;
}

export interface QueryMissionResponse {
  mission: Mission | undefined;
}

export interface QueryMissionsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryMissionsResponse {
  mission: Mission[];
  pagination: PageResponse | undefined;
}

export interface QueryCampaignsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryCampaignsResponse {
  campaign: Campaign[];
  pagination: PageResponse | undefined;
}

export interface QueryCampaignRequest {
  campaignId: number;
}

export interface QueryCampaignResponse {
  campaign: Campaign | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryUsersEntriesRequest(): QueryUsersEntriesRequest {
  return { address: "" };
}

export const QueryUsersEntriesRequest = {
  encode(message: QueryUsersEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryUsersEntriesRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesRequest>, I>>(
    object: I,
  ): QueryUsersEntriesRequest {
    const message = createBaseQueryUsersEntriesRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryUsersEntriesResponse(): QueryUsersEntriesResponse {
  return { userEntry: undefined };
}

export const QueryUsersEntriesResponse = {
  encode(message: QueryUsersEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userEntry !== undefined) {
      UserEntry.encode(message.userEntry, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userEntry = UserEntry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesResponse {
    return {
      userEntry: isSet(object.userEntry)
        ? UserEntry.fromJSON(object.userEntry)
        : undefined,
    };
  },

  toJSON(message: QueryUsersEntriesResponse): unknown {
    const obj: any = {};
    message.userEntry !== undefined && (obj.userEntry = message.userEntry
      ? UserEntry.toJSON(message.userEntry)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesResponse>, I>>(
    object: I,
  ): QueryUsersEntriesResponse {
    const message = createBaseQueryUsersEntriesResponse();
    message.userEntry = (object.userEntry !== undefined && object.userEntry !== null)
      ? UserEntry.fromPartial(object.userEntry)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignTotalAmountRequest(): QueryCampaignTotalAmountRequest {
  return { campaignId: 0 };
}

export const QueryCampaignTotalAmountRequest = {
  encode(message: QueryCampaignTotalAmountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignTotalAmountRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignTotalAmountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignTotalAmountRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryCampaignTotalAmountRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignTotalAmountRequest>, I>>(
    object: I,
  ): QueryCampaignTotalAmountRequest {
    const message = createBaseQueryCampaignTotalAmountRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryCampaignTotalAmountResponse(): QueryCampaignTotalAmountResponse {
  return { amount: [] };
}

export const QueryCampaignTotalAmountResponse = {
  encode(message: QueryCampaignTotalAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignTotalAmountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignTotalAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignTotalAmountResponse {
    return {
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryCampaignTotalAmountResponse): unknown {
    const obj: any = {};
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignTotalAmountResponse>, I>>(
    object: I,
  ): QueryCampaignTotalAmountResponse {
    const message = createBaseQueryCampaignTotalAmountResponse();
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryCampaignAmountLeftRequest(): QueryCampaignAmountLeftRequest {
  return { campaignId: 0 };
}

export const QueryCampaignAmountLeftRequest = {
  encode(message: QueryCampaignAmountLeftRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignAmountLeftRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignAmountLeftRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignAmountLeftRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryCampaignAmountLeftRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignAmountLeftRequest>, I>>(
    object: I,
  ): QueryCampaignAmountLeftRequest {
    const message = createBaseQueryCampaignAmountLeftRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryCampaignAmountLeftResponse(): QueryCampaignAmountLeftResponse {
  return { amount: [] };
}

export const QueryCampaignAmountLeftResponse = {
  encode(message: QueryCampaignAmountLeftResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignAmountLeftResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignAmountLeftResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignAmountLeftResponse {
    return {
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryCampaignAmountLeftResponse): unknown {
    const obj: any = {};
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignAmountLeftResponse>, I>>(
    object: I,
  ): QueryCampaignAmountLeftResponse {
    const message = createBaseQueryCampaignAmountLeftResponse();
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryUsersEntriesRequest(): QueryUsersEntriesRequest {
  return { pagination: undefined };
}

export const QueryUsersEntriesRequest = {
  encode(message: QueryUsersEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryUsersEntriesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesRequest>, I>>(
    object: I,
  ): QueryUsersEntriesRequest {
    const message = createBaseQueryUsersEntriesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryUsersEntriesResponse(): QueryUsersEntriesResponse {
  return { usersEntries: [], pagination: undefined };
}

export const QueryUsersEntriesResponse = {
  encode(message: QueryUsersEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.usersEntries) {
      UserEntry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.usersEntries.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesResponse {
    return {
      usersEntries: Array.isArray(object?.usersEntries)
        ? object.usersEntries.map((e: any) => UserEntry.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryUsersEntriesResponse): unknown {
    const obj: any = {};
    if (message.usersEntries) {
      obj.usersEntries = message.usersEntries.map((e) => e ? UserEntry.toJSON(e) : undefined);
    } else {
      obj.usersEntries = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesResponse>, I>>(
    object: I,
  ): QueryUsersEntriesResponse {
    const message = createBaseQueryUsersEntriesResponse();
    message.usersEntries = object.usersEntries?.map((e) => UserEntry.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionRequest(): QueryMissionRequest {
  return { campaignId: 0, missionId: 0 };
}

export const QueryMissionRequest = {
  encode(message: QueryMissionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(16).uint64(message.missionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.missionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionRequest {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      missionId: isSet(object.missionId) ? Number(object.missionId) : 0,
    };
  },

  toJSON(message: QueryMissionRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionRequest>, I>>(object: I): QueryMissionRequest {
    const message = createBaseQueryMissionRequest();
    message.campaignId = object.campaignId ?? 0;
    message.missionId = object.missionId ?? 0;
    return message;
  },
};

function createBaseQueryMissionResponse(): QueryMissionResponse {
  return { mission: undefined };
}

export const QueryMissionResponse = {
  encode(message: QueryMissionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mission !== undefined) {
      Mission.encode(message.mission, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mission = Mission.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionResponse {
    return { mission: isSet(object.mission) ? Mission.fromJSON(object.mission) : undefined };
  },

  toJSON(message: QueryMissionResponse): unknown {
    const obj: any = {};
    message.mission !== undefined && (obj.mission = message.mission ? Mission.toJSON(message.mission) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionResponse>, I>>(object: I): QueryMissionResponse {
    const message = createBaseQueryMissionResponse();
    message.mission = (object.mission !== undefined && object.mission !== null)
      ? Mission.fromPartial(object.mission)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionsRequest(): QueryMissionsRequest {
  return { pagination: undefined };
}

export const QueryMissionsRequest = {
  encode(message: QueryMissionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryMissionsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionsRequest>, I>>(object: I): QueryMissionsRequest {
    const message = createBaseQueryMissionsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionsResponse(): QueryMissionsResponse {
  return { mission: [], pagination: undefined };
}

export const QueryMissionsResponse = {
  encode(message: QueryMissionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.mission) {
      Mission.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mission.push(Mission.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionsResponse {
    return {
      mission: Array.isArray(object?.mission) ? object.mission.map((e: any) => Mission.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryMissionsResponse): unknown {
    const obj: any = {};
    if (message.mission) {
      obj.mission = message.mission.map((e) => e ? Mission.toJSON(e) : undefined);
    } else {
      obj.mission = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionsResponse>, I>>(object: I): QueryMissionsResponse {
    const message = createBaseQueryMissionsResponse();
    message.mission = object.mission?.map((e) => Mission.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignsRequest(): QueryCampaignsRequest {
  return { pagination: undefined };
}

export const QueryCampaignsRequest = {
  encode(message: QueryCampaignsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryCampaignsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignsRequest>, I>>(object: I): QueryCampaignsRequest {
    const message = createBaseQueryCampaignsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignsResponse(): QueryCampaignsResponse {
  return { campaign: [], pagination: undefined };
}

export const QueryCampaignsResponse = {
  encode(message: QueryCampaignsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.campaign) {
      Campaign.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignsResponse {
    return {
      campaign: Array.isArray(object?.campaign) ? object.campaign.map((e: any) => Campaign.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryCampaignsResponse): unknown {
    const obj: any = {};
    if (message.campaign) {
      obj.campaign = message.campaign.map((e) => e ? Campaign.toJSON(e) : undefined);
    } else {
      obj.campaign = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignsResponse>, I>>(object: I): QueryCampaignsResponse {
    const message = createBaseQueryCampaignsResponse();
    message.campaign = object.campaign?.map((e) => Campaign.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignRequest(): QueryCampaignRequest {
  return { campaignId: 0 };
}

export const QueryCampaignRequest = {
  encode(message: QueryCampaignRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryCampaignRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignRequest>, I>>(object: I): QueryCampaignRequest {
    const message = createBaseQueryCampaignRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryCampaignResponse(): QueryCampaignResponse {
  return { campaign: undefined };
}

export const QueryCampaignResponse = {
  encode(message: QueryCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaign !== undefined) {
      Campaign.encode(message.campaign, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign = Campaign.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignResponse {
    return { campaign: isSet(object.campaign) ? Campaign.fromJSON(object.campaign) : undefined };
  },

  toJSON(message: QueryCampaignResponse): unknown {
    const obj: any = {};
    message.campaign !== undefined && (obj.campaign = message.campaign ? Campaign.toJSON(message.campaign) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignResponse>, I>>(object: I): QueryCampaignResponse {
    const message = createBaseQueryCampaignResponse();
    message.campaign = (object.campaign !== undefined && object.campaign !== null)
      ? Campaign.fromPartial(object.campaign)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a UserEntry by index. */
  UserEntry(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse>;
  /** Queries a list of UserEntry items. */
  UsersEntries(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse>;
  /** Queries a Mission by index. */
  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse>;
  /** Queries a list of Mission items. */
  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse>;
  /** Queries a list of Campaigns items. */
  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse>;
  /** Queries a list of Campaigns items. */
  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse>;
  /** Queries a CampaignTotalAmount by campaignId. */
  CampaignTotalAmount(request: QueryCampaignTotalAmountRequest): Promise<QueryCampaignTotalAmountResponse>;
  /** Queries a CampaignTotalAmount by campaignId. */
  CampaignAmountLeft(request: QueryCampaignAmountLeftRequest): Promise<QueryCampaignAmountLeftResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.UserEntry = this.UserEntry.bind(this);
    this.UsersEntries = this.UsersEntries.bind(this);
    this.Mission = this.Mission.bind(this);
    this.MissionAll = this.MissionAll.bind(this);
    this.Campaigns = this.Campaigns.bind(this);
    this.Campaign = this.Campaign.bind(this);
    this.CampaignTotalAmount = this.CampaignTotalAmount.bind(this);
    this.CampaignAmountLeft = this.CampaignAmountLeft.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  UserEntry(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse> {
    const data = QueryUsersEntriesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "UserEntry", data);
    return promise.then((data) => QueryUsersEntriesResponse.decode(new _m0.Reader(data)));
  }

  UsersEntries(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse> {
    const data = QueryUsersEntriesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "UsersEntries", data);
    return promise.then((data) => QueryUsersEntriesResponse.decode(new _m0.Reader(data)));
  }

  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse> {
    const data = QueryMissionRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "Mission", data);
    return promise.then((data) => QueryMissionResponse.decode(new _m0.Reader(data)));
  }

  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse> {
    const data = QueryMissionsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "MissionAll", data);
    return promise.then((data) => QueryMissionsResponse.decode(new _m0.Reader(data)));
  }

  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse> {
    const data = QueryCampaignsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "Campaigns", data);
    return promise.then((data) => QueryCampaignsResponse.decode(new _m0.Reader(data)));
  }

  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse> {
    const data = QueryCampaignRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "Campaign", data);
    return promise.then((data) => QueryCampaignResponse.decode(new _m0.Reader(data)));
  }

  CampaignTotalAmount(request: QueryCampaignTotalAmountRequest): Promise<QueryCampaignTotalAmountResponse> {
    const data = QueryCampaignTotalAmountRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "CampaignTotalAmount", data);
    return promise.then((data) => QueryCampaignTotalAmountResponse.decode(new _m0.Reader(data)));
  }

  CampaignAmountLeft(request: QueryCampaignAmountLeftRequest): Promise<QueryCampaignAmountLeftResponse> {
    const data = QueryCampaignAmountLeftRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "CampaignAmountLeft", data);
    return promise.then((data) => QueryCampaignAmountLeftResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
