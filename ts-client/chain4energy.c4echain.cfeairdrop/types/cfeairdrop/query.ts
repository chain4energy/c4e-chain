/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Campaign, Mission, UserAirdropEntries } from "./airdrop";
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

export interface QueryUserAirdropEntriesRequest {
  address: string;
}

export interface QueryUserAirdropEntriesResponse {
  userAirdropEntries: UserAirdropEntries | undefined;
}

export interface QueryAirdropDistrubitionsRequest {
  campaignId: number;
}

export interface QueryAirdropDistrubitionsResponse {
  airdropCoins: Coin[];
}

export interface QueryAirdropClaimsLeftRequest {
  campaignId: number;
}

export interface QueryAirdropClaimsLeftResponse {
  airdropCoins: Coin[];
}

export interface QueryUsersAirdropEntriesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryUsersAirdropEntriesResponse {
  usersAirdropEntries: UserAirdropEntries[];
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

function createBaseQueryUserAirdropEntriesRequest(): QueryUserAirdropEntriesRequest {
  return { address: "" };
}

export const QueryUserAirdropEntriesRequest = {
  encode(message: QueryUserAirdropEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserAirdropEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserAirdropEntriesRequest();
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

  fromJSON(object: any): QueryUserAirdropEntriesRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryUserAirdropEntriesRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUserAirdropEntriesRequest>, I>>(
    object: I,
  ): QueryUserAirdropEntriesRequest {
    const message = createBaseQueryUserAirdropEntriesRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryUserAirdropEntriesResponse(): QueryUserAirdropEntriesResponse {
  return { userAirdropEntries: undefined };
}

export const QueryUserAirdropEntriesResponse = {
  encode(message: QueryUserAirdropEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userAirdropEntries !== undefined) {
      UserAirdropEntries.encode(message.userAirdropEntries, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserAirdropEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserAirdropEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAirdropEntries = UserAirdropEntries.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUserAirdropEntriesResponse {
    return {
      userAirdropEntries: isSet(object.userAirdropEntries)
        ? UserAirdropEntries.fromJSON(object.userAirdropEntries)
        : undefined,
    };
  },

  toJSON(message: QueryUserAirdropEntriesResponse): unknown {
    const obj: any = {};
    message.userAirdropEntries !== undefined && (obj.userAirdropEntries = message.userAirdropEntries
      ? UserAirdropEntries.toJSON(message.userAirdropEntries)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUserAirdropEntriesResponse>, I>>(
    object: I,
  ): QueryUserAirdropEntriesResponse {
    const message = createBaseQueryUserAirdropEntriesResponse();
    message.userAirdropEntries = (object.userAirdropEntries !== undefined && object.userAirdropEntries !== null)
      ? UserAirdropEntries.fromPartial(object.userAirdropEntries)
      : undefined;
    return message;
  },
};

function createBaseQueryAirdropDistrubitionsRequest(): QueryAirdropDistrubitionsRequest {
  return { campaignId: 0 };
}

export const QueryAirdropDistrubitionsRequest = {
  encode(message: QueryAirdropDistrubitionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAirdropDistrubitionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAirdropDistrubitionsRequest();
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

  fromJSON(object: any): QueryAirdropDistrubitionsRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryAirdropDistrubitionsRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAirdropDistrubitionsRequest>, I>>(
    object: I,
  ): QueryAirdropDistrubitionsRequest {
    const message = createBaseQueryAirdropDistrubitionsRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryAirdropDistrubitionsResponse(): QueryAirdropDistrubitionsResponse {
  return { airdropCoins: [] };
}

export const QueryAirdropDistrubitionsResponse = {
  encode(message: QueryAirdropDistrubitionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.airdropCoins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAirdropDistrubitionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAirdropDistrubitionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropDistrubitionsResponse {
    return {
      airdropCoins: Array.isArray(object?.airdropCoins) ? object.airdropCoins.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryAirdropDistrubitionsResponse): unknown {
    const obj: any = {};
    if (message.airdropCoins) {
      obj.airdropCoins = message.airdropCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.airdropCoins = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAirdropDistrubitionsResponse>, I>>(
    object: I,
  ): QueryAirdropDistrubitionsResponse {
    const message = createBaseQueryAirdropDistrubitionsResponse();
    message.airdropCoins = object.airdropCoins?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryAirdropClaimsLeftRequest(): QueryAirdropClaimsLeftRequest {
  return { campaignId: 0 };
}

export const QueryAirdropClaimsLeftRequest = {
  encode(message: QueryAirdropClaimsLeftRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAirdropClaimsLeftRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAirdropClaimsLeftRequest();
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

  fromJSON(object: any): QueryAirdropClaimsLeftRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryAirdropClaimsLeftRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAirdropClaimsLeftRequest>, I>>(
    object: I,
  ): QueryAirdropClaimsLeftRequest {
    const message = createBaseQueryAirdropClaimsLeftRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryAirdropClaimsLeftResponse(): QueryAirdropClaimsLeftResponse {
  return { airdropCoins: [] };
}

export const QueryAirdropClaimsLeftResponse = {
  encode(message: QueryAirdropClaimsLeftResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.airdropCoins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAirdropClaimsLeftResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAirdropClaimsLeftResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropClaimsLeftResponse {
    return {
      airdropCoins: Array.isArray(object?.airdropCoins) ? object.airdropCoins.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryAirdropClaimsLeftResponse): unknown {
    const obj: any = {};
    if (message.airdropCoins) {
      obj.airdropCoins = message.airdropCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.airdropCoins = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAirdropClaimsLeftResponse>, I>>(
    object: I,
  ): QueryAirdropClaimsLeftResponse {
    const message = createBaseQueryAirdropClaimsLeftResponse();
    message.airdropCoins = object.airdropCoins?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryUsersAirdropEntriesRequest(): QueryUsersAirdropEntriesRequest {
  return { pagination: undefined };
}

export const QueryUsersAirdropEntriesRequest = {
  encode(message: QueryUsersAirdropEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersAirdropEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersAirdropEntriesRequest();
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

  fromJSON(object: any): QueryUsersAirdropEntriesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryUsersAirdropEntriesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersAirdropEntriesRequest>, I>>(
    object: I,
  ): QueryUsersAirdropEntriesRequest {
    const message = createBaseQueryUsersAirdropEntriesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryUsersAirdropEntriesResponse(): QueryUsersAirdropEntriesResponse {
  return { usersAirdropEntries: [], pagination: undefined };
}

export const QueryUsersAirdropEntriesResponse = {
  encode(message: QueryUsersAirdropEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.usersAirdropEntries) {
      UserAirdropEntries.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersAirdropEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersAirdropEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.usersAirdropEntries.push(UserAirdropEntries.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryUsersAirdropEntriesResponse {
    return {
      usersAirdropEntries: Array.isArray(object?.usersAirdropEntries)
        ? object.usersAirdropEntries.map((e: any) => UserAirdropEntries.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryUsersAirdropEntriesResponse): unknown {
    const obj: any = {};
    if (message.usersAirdropEntries) {
      obj.usersAirdropEntries = message.usersAirdropEntries.map((e) => e ? UserAirdropEntries.toJSON(e) : undefined);
    } else {
      obj.usersAirdropEntries = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersAirdropEntriesResponse>, I>>(
    object: I,
  ): QueryUsersAirdropEntriesResponse {
    const message = createBaseQueryUsersAirdropEntriesResponse();
    message.usersAirdropEntries = object.usersAirdropEntries?.map((e) => UserAirdropEntries.fromPartial(e)) || [];
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
  /** Queries a UserAirdropEntries by index. */
  UserAirdropEntries(request: QueryUserAirdropEntriesRequest): Promise<QueryUserAirdropEntriesResponse>;
  /** Queries a list of UserAirdropEntries items. */
  UsersAirdropEntries(request: QueryUsersAirdropEntriesRequest): Promise<QueryUsersAirdropEntriesResponse>;
  /** Queries a Mission by index. */
  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse>;
  /** Queries a list of Mission items. */
  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse>;
  /** Queries a list of Campaigns items. */
  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse>;
  /** Queries a list of Campaigns items. */
  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse>;
  /** Queries a AirdropDistrubitions by campaignId. */
  AirdropDistrubitions(request: QueryAirdropDistrubitionsRequest): Promise<QueryAirdropDistrubitionsResponse>;
  /** Queries a AirdropDistrubitions by campaignId. */
  AirdropClaimsLeft(request: QueryAirdropClaimsLeftRequest): Promise<QueryAirdropClaimsLeftResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.UserAirdropEntries = this.UserAirdropEntries.bind(this);
    this.UsersAirdropEntries = this.UsersAirdropEntries.bind(this);
    this.Mission = this.Mission.bind(this);
    this.MissionAll = this.MissionAll.bind(this);
    this.Campaigns = this.Campaigns.bind(this);
    this.Campaign = this.Campaign.bind(this);
    this.AirdropDistrubitions = this.AirdropDistrubitions.bind(this);
    this.AirdropClaimsLeft = this.AirdropClaimsLeft.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  UserAirdropEntries(request: QueryUserAirdropEntriesRequest): Promise<QueryUserAirdropEntriesResponse> {
    const data = QueryUserAirdropEntriesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "UserAirdropEntries", data);
    return promise.then((data) => QueryUserAirdropEntriesResponse.decode(new _m0.Reader(data)));
  }

  UsersAirdropEntries(request: QueryUsersAirdropEntriesRequest): Promise<QueryUsersAirdropEntriesResponse> {
    const data = QueryUsersAirdropEntriesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "UsersAirdropEntries", data);
    return promise.then((data) => QueryUsersAirdropEntriesResponse.decode(new _m0.Reader(data)));
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

  AirdropDistrubitions(request: QueryAirdropDistrubitionsRequest): Promise<QueryAirdropDistrubitionsResponse> {
    const data = QueryAirdropDistrubitionsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "AirdropDistrubitions", data);
    return promise.then((data) => QueryAirdropDistrubitionsResponse.decode(new _m0.Reader(data)));
  }

  AirdropClaimsLeft(request: QueryAirdropClaimsLeftRequest): Promise<QueryAirdropClaimsLeftResponse> {
    const data = QueryAirdropClaimsLeftRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Query", "AirdropClaimsLeft", data);
    return promise.then((data) => QueryAirdropClaimsLeftResponse.decode(new _m0.Reader(data)));
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
