/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cfeairdrop/params";
import { UserEntry, Mission, Campaign } from "../cfeairdrop/airdrop";
import { Coin } from "../cosmos/base/v1beta1/coin";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryUserEntryRequest {
  address: string;
}

export interface QueryUserEntryResponse {
  user_entry: UserEntry | undefined;
}

export interface QueryAirdropDistrubitionsRequest {
  campaign_id: number;
}

export interface QueryAirdropDistrubitionsResponse {
  airdrop_coins: Coin[];
}

export interface QueryAirdropClaimsLeftRequest {
  campaign_id: number;
}

export interface QueryAirdropClaimsLeftResponse {
  airdrop_coins: Coin[];
}

export interface QueryUsersEntriesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryUsersEntriesResponse {
  users_entries: UserEntry[];
  pagination: PageResponse | undefined;
}

export interface QueryMissionRequest {
  campaign_id: number;
  mission_id: number;
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
  campaign_id: number;
}

export interface QueryCampaignResponse {
  campaign: Campaign | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryUserEntryRequest: object = { address: "" };

export const QueryUserEntryRequest = {
  encode(
    message: QueryUserEntryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryUserEntryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryUserEntryRequest } as QueryUserEntryRequest;
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

  fromJSON(object: any): QueryUserEntryRequest {
    const message = { ...baseQueryUserEntryRequest } as QueryUserEntryRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryUserEntryRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryUserEntryRequest>
  ): QueryUserEntryRequest {
    const message = { ...baseQueryUserEntryRequest } as QueryUserEntryRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryUserEntryResponse: object = {};

export const QueryUserEntryResponse = {
  encode(
    message: QueryUserEntryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.user_entry !== undefined) {
      UserEntry.encode(message.user_entry, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryUserEntryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryUserEntryResponse } as QueryUserEntryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user_entry = UserEntry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUserEntryResponse {
    const message = { ...baseQueryUserEntryResponse } as QueryUserEntryResponse;
    if (object.user_entry !== undefined && object.user_entry !== null) {
      message.user_entry = UserEntry.fromJSON(object.user_entry);
    } else {
      message.user_entry = undefined;
    }
    return message;
  },

  toJSON(message: QueryUserEntryResponse): unknown {
    const obj: any = {};
    message.user_entry !== undefined &&
      (obj.user_entry = message.user_entry
        ? UserEntry.toJSON(message.user_entry)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryUserEntryResponse>
  ): QueryUserEntryResponse {
    const message = { ...baseQueryUserEntryResponse } as QueryUserEntryResponse;
    if (object.user_entry !== undefined && object.user_entry !== null) {
      message.user_entry = UserEntry.fromPartial(object.user_entry);
    } else {
      message.user_entry = undefined;
    }
    return message;
  },
};

const baseQueryAirdropDistrubitionsRequest: object = { campaign_id: 0 };

export const QueryAirdropDistrubitionsRequest = {
  encode(
    message: QueryAirdropDistrubitionsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAirdropDistrubitionsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAirdropDistrubitionsRequest,
    } as QueryAirdropDistrubitionsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropDistrubitionsRequest {
    const message = {
      ...baseQueryAirdropDistrubitionsRequest,
    } as QueryAirdropDistrubitionsRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    return message;
  },

  toJSON(message: QueryAirdropDistrubitionsRequest): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAirdropDistrubitionsRequest>
  ): QueryAirdropDistrubitionsRequest {
    const message = {
      ...baseQueryAirdropDistrubitionsRequest,
    } as QueryAirdropDistrubitionsRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    return message;
  },
};

const baseQueryAirdropDistrubitionsResponse: object = {};

export const QueryAirdropDistrubitionsResponse = {
  encode(
    message: QueryAirdropDistrubitionsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.airdrop_coins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAirdropDistrubitionsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAirdropDistrubitionsResponse,
    } as QueryAirdropDistrubitionsResponse;
    message.airdrop_coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdrop_coins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropDistrubitionsResponse {
    const message = {
      ...baseQueryAirdropDistrubitionsResponse,
    } as QueryAirdropDistrubitionsResponse;
    message.airdrop_coins = [];
    if (object.airdrop_coins !== undefined && object.airdrop_coins !== null) {
      for (const e of object.airdrop_coins) {
        message.airdrop_coins.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryAirdropDistrubitionsResponse): unknown {
    const obj: any = {};
    if (message.airdrop_coins) {
      obj.airdrop_coins = message.airdrop_coins.map((e) =>
        e ? Coin.toJSON(e) : undefined
      );
    } else {
      obj.airdrop_coins = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAirdropDistrubitionsResponse>
  ): QueryAirdropDistrubitionsResponse {
    const message = {
      ...baseQueryAirdropDistrubitionsResponse,
    } as QueryAirdropDistrubitionsResponse;
    message.airdrop_coins = [];
    if (object.airdrop_coins !== undefined && object.airdrop_coins !== null) {
      for (const e of object.airdrop_coins) {
        message.airdrop_coins.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryAirdropClaimsLeftRequest: object = { campaign_id: 0 };

export const QueryAirdropClaimsLeftRequest = {
  encode(
    message: QueryAirdropClaimsLeftRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAirdropClaimsLeftRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAirdropClaimsLeftRequest,
    } as QueryAirdropClaimsLeftRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropClaimsLeftRequest {
    const message = {
      ...baseQueryAirdropClaimsLeftRequest,
    } as QueryAirdropClaimsLeftRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    return message;
  },

  toJSON(message: QueryAirdropClaimsLeftRequest): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAirdropClaimsLeftRequest>
  ): QueryAirdropClaimsLeftRequest {
    const message = {
      ...baseQueryAirdropClaimsLeftRequest,
    } as QueryAirdropClaimsLeftRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    return message;
  },
};

const baseQueryAirdropClaimsLeftResponse: object = {};

export const QueryAirdropClaimsLeftResponse = {
  encode(
    message: QueryAirdropClaimsLeftResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.airdrop_coins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAirdropClaimsLeftResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAirdropClaimsLeftResponse,
    } as QueryAirdropClaimsLeftResponse;
    message.airdrop_coins = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdrop_coins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAirdropClaimsLeftResponse {
    const message = {
      ...baseQueryAirdropClaimsLeftResponse,
    } as QueryAirdropClaimsLeftResponse;
    message.airdrop_coins = [];
    if (object.airdrop_coins !== undefined && object.airdrop_coins !== null) {
      for (const e of object.airdrop_coins) {
        message.airdrop_coins.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryAirdropClaimsLeftResponse): unknown {
    const obj: any = {};
    if (message.airdrop_coins) {
      obj.airdrop_coins = message.airdrop_coins.map((e) =>
        e ? Coin.toJSON(e) : undefined
      );
    } else {
      obj.airdrop_coins = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAirdropClaimsLeftResponse>
  ): QueryAirdropClaimsLeftResponse {
    const message = {
      ...baseQueryAirdropClaimsLeftResponse,
    } as QueryAirdropClaimsLeftResponse;
    message.airdrop_coins = [];
    if (object.airdrop_coins !== undefined && object.airdrop_coins !== null) {
      for (const e of object.airdrop_coins) {
        message.airdrop_coins.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryUsersEntriesRequest: object = {};

export const QueryUsersEntriesRequest = {
  encode(
    message: QueryUsersEntriesRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryUsersEntriesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryUsersEntriesRequest,
    } as QueryUsersEntriesRequest;
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
    const message = {
      ...baseQueryUsersEntriesRequest,
    } as QueryUsersEntriesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryUsersEntriesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryUsersEntriesRequest>
  ): QueryUsersEntriesRequest {
    const message = {
      ...baseQueryUsersEntriesRequest,
    } as QueryUsersEntriesRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryUsersEntriesResponse: object = {};

export const QueryUsersEntriesResponse = {
  encode(
    message: QueryUsersEntriesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.users_entries) {
      UserEntry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryUsersEntriesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryUsersEntriesResponse,
    } as QueryUsersEntriesResponse;
    message.users_entries = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users_entries.push(UserEntry.decode(reader, reader.uint32()));
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
    const message = {
      ...baseQueryUsersEntriesResponse,
    } as QueryUsersEntriesResponse;
    message.users_entries = [];
    if (object.users_entries !== undefined && object.users_entries !== null) {
      for (const e of object.users_entries) {
        message.users_entries.push(UserEntry.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryUsersEntriesResponse): unknown {
    const obj: any = {};
    if (message.users_entries) {
      obj.users_entries = message.users_entries.map((e) =>
        e ? UserEntry.toJSON(e) : undefined
      );
    } else {
      obj.users_entries = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryUsersEntriesResponse>
  ): QueryUsersEntriesResponse {
    const message = {
      ...baseQueryUsersEntriesResponse,
    } as QueryUsersEntriesResponse;
    message.users_entries = [];
    if (object.users_entries !== undefined && object.users_entries !== null) {
      for (const e of object.users_entries) {
        message.users_entries.push(UserEntry.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMissionRequest: object = { campaign_id: 0, mission_id: 0 };

export const QueryMissionRequest = {
  encode(
    message: QueryMissionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    if (message.mission_id !== 0) {
      writer.uint32(16).uint64(message.mission_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMissionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMissionRequest } as QueryMissionRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.mission_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionRequest {
    const message = { ...baseQueryMissionRequest } as QueryMissionRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = Number(object.mission_id);
    } else {
      message.mission_id = 0;
    }
    return message;
  },

  toJSON(message: QueryMissionRequest): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.mission_id !== undefined && (obj.mission_id = message.mission_id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMissionRequest>): QueryMissionRequest {
    const message = { ...baseQueryMissionRequest } as QueryMissionRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = object.mission_id;
    } else {
      message.mission_id = 0;
    }
    return message;
  },
};

const baseQueryMissionResponse: object = {};

export const QueryMissionResponse = {
  encode(
    message: QueryMissionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mission !== undefined) {
      Mission.encode(message.mission, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMissionResponse } as QueryMissionResponse;
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
    const message = { ...baseQueryMissionResponse } as QueryMissionResponse;
    if (object.mission !== undefined && object.mission !== null) {
      message.mission = Mission.fromJSON(object.mission);
    } else {
      message.mission = undefined;
    }
    return message;
  },

  toJSON(message: QueryMissionResponse): unknown {
    const obj: any = {};
    message.mission !== undefined &&
      (obj.mission = message.mission
        ? Mission.toJSON(message.mission)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMissionResponse>): QueryMissionResponse {
    const message = { ...baseQueryMissionResponse } as QueryMissionResponse;
    if (object.mission !== undefined && object.mission !== null) {
      message.mission = Mission.fromPartial(object.mission);
    } else {
      message.mission = undefined;
    }
    return message;
  },
};

const baseQueryMissionsRequest: object = {};

export const QueryMissionsRequest = {
  encode(
    message: QueryMissionsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMissionsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMissionsRequest } as QueryMissionsRequest;
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
    const message = { ...baseQueryMissionsRequest } as QueryMissionsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMissionsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMissionsRequest>): QueryMissionsRequest {
    const message = { ...baseQueryMissionsRequest } as QueryMissionsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryMissionsResponse: object = {};

export const QueryMissionsResponse = {
  encode(
    message: QueryMissionsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.mission) {
      Mission.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryMissionsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryMissionsResponse } as QueryMissionsResponse;
    message.mission = [];
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
    const message = { ...baseQueryMissionsResponse } as QueryMissionsResponse;
    message.mission = [];
    if (object.mission !== undefined && object.mission !== null) {
      for (const e of object.mission) {
        message.mission.push(Mission.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryMissionsResponse): unknown {
    const obj: any = {};
    if (message.mission) {
      obj.mission = message.mission.map((e) =>
        e ? Mission.toJSON(e) : undefined
      );
    } else {
      obj.mission = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryMissionsResponse>
  ): QueryMissionsResponse {
    const message = { ...baseQueryMissionsResponse } as QueryMissionsResponse;
    message.mission = [];
    if (object.mission !== undefined && object.mission !== null) {
      for (const e of object.mission) {
        message.mission.push(Mission.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCampaignsRequest: object = {};

export const QueryCampaignsRequest = {
  encode(
    message: QueryCampaignsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCampaignsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCampaignsRequest } as QueryCampaignsRequest;
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
    const message = { ...baseQueryCampaignsRequest } as QueryCampaignsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCampaignsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCampaignsRequest>
  ): QueryCampaignsRequest {
    const message = { ...baseQueryCampaignsRequest } as QueryCampaignsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCampaignsResponse: object = {};

export const QueryCampaignsResponse = {
  encode(
    message: QueryCampaignsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.campaign) {
      Campaign.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCampaignsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCampaignsResponse } as QueryCampaignsResponse;
    message.campaign = [];
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
    const message = { ...baseQueryCampaignsResponse } as QueryCampaignsResponse;
    message.campaign = [];
    if (object.campaign !== undefined && object.campaign !== null) {
      for (const e of object.campaign) {
        message.campaign.push(Campaign.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryCampaignsResponse): unknown {
    const obj: any = {};
    if (message.campaign) {
      obj.campaign = message.campaign.map((e) =>
        e ? Campaign.toJSON(e) : undefined
      );
    } else {
      obj.campaign = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCampaignsResponse>
  ): QueryCampaignsResponse {
    const message = { ...baseQueryCampaignsResponse } as QueryCampaignsResponse;
    message.campaign = [];
    if (object.campaign !== undefined && object.campaign !== null) {
      for (const e of object.campaign) {
        message.campaign.push(Campaign.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCampaignRequest: object = { campaign_id: 0 };

export const QueryCampaignRequest = {
  encode(
    message: QueryCampaignRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCampaignRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCampaignRequest } as QueryCampaignRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignRequest {
    const message = { ...baseQueryCampaignRequest } as QueryCampaignRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    return message;
  },

  toJSON(message: QueryCampaignRequest): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCampaignRequest>): QueryCampaignRequest {
    const message = { ...baseQueryCampaignRequest } as QueryCampaignRequest;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    return message;
  },
};

const baseQueryCampaignResponse: object = {};

export const QueryCampaignResponse = {
  encode(
    message: QueryCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign !== undefined) {
      Campaign.encode(message.campaign, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCampaignResponse } as QueryCampaignResponse;
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
    const message = { ...baseQueryCampaignResponse } as QueryCampaignResponse;
    if (object.campaign !== undefined && object.campaign !== null) {
      message.campaign = Campaign.fromJSON(object.campaign);
    } else {
      message.campaign = undefined;
    }
    return message;
  },

  toJSON(message: QueryCampaignResponse): unknown {
    const obj: any = {};
    message.campaign !== undefined &&
      (obj.campaign = message.campaign
        ? Campaign.toJSON(message.campaign)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCampaignResponse>
  ): QueryCampaignResponse {
    const message = { ...baseQueryCampaignResponse } as QueryCampaignResponse;
    if (object.campaign !== undefined && object.campaign !== null) {
      message.campaign = Campaign.fromPartial(object.campaign);
    } else {
      message.campaign = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a UserEntry by index. */
  UserEntry(request: QueryUserEntryRequest): Promise<QueryUserEntryResponse>;
  /** Queries a list of UserEntry items. */
  UsersEntries(
    request: QueryUsersEntriesRequest
  ): Promise<QueryUsersEntriesResponse>;
  /** Queries a Mission by index. */
  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse>;
  /** Queries a list of Mission items. */
  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse>;
  /** Queries a list of Campaigns items. */
  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse>;
  /** Queries a list of Campaigns items. */
  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse>;
  /** Queries a AirdropDistrubitions by campaignId. */
  AirdropDistrubitions(
    request: QueryAirdropDistrubitionsRequest
  ): Promise<QueryAirdropDistrubitionsResponse>;
  /** Queries a AirdropDistrubitions by campaignId. */
  AirdropClaimsLeft(
    request: QueryAirdropClaimsLeftRequest
  ): Promise<QueryAirdropClaimsLeftResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  UserEntry(request: QueryUserEntryRequest): Promise<QueryUserEntryResponse> {
    const data = QueryUserEntryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "UserEntry",
      data
    );
    return promise.then((data) =>
      QueryUserEntryResponse.decode(new Reader(data))
    );
  }

  UsersEntries(
    request: QueryUsersEntriesRequest
  ): Promise<QueryUsersEntriesResponse> {
    const data = QueryUsersEntriesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "UsersEntries",
      data
    );
    return promise.then((data) =>
      QueryUsersEntriesResponse.decode(new Reader(data))
    );
  }

  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse> {
    const data = QueryMissionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "Mission",
      data
    );
    return promise.then((data) =>
      QueryMissionResponse.decode(new Reader(data))
    );
  }

  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse> {
    const data = QueryMissionsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "MissionAll",
      data
    );
    return promise.then((data) =>
      QueryMissionsResponse.decode(new Reader(data))
    );
  }

  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse> {
    const data = QueryCampaignsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "Campaigns",
      data
    );
    return promise.then((data) =>
      QueryCampaignsResponse.decode(new Reader(data))
    );
  }

  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse> {
    const data = QueryCampaignRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "Campaign",
      data
    );
    return promise.then((data) =>
      QueryCampaignResponse.decode(new Reader(data))
    );
  }

  AirdropDistrubitions(
    request: QueryAirdropDistrubitionsRequest
  ): Promise<QueryAirdropDistrubitionsResponse> {
    const data = QueryAirdropDistrubitionsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "AirdropDistrubitions",
      data
    );
    return promise.then((data) =>
      QueryAirdropDistrubitionsResponse.decode(new Reader(data))
    );
  }

  AirdropClaimsLeft(
    request: QueryAirdropClaimsLeftRequest
  ): Promise<QueryAirdropClaimsLeftResponse> {
    const data = QueryAirdropClaimsLeftRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "AirdropClaimsLeft",
      data
    );
    return promise.then((data) =>
      QueryAirdropClaimsLeftResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
