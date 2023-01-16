/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cfeairdrop/params";
import {
  UserAirdropEntries,
  InitialClaim,
  Mission,
  Campaign,
  AirdropEntry,
} from "../cfeairdrop/airdrop";
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

export interface QueryClaimRecordRequest {
  address: string;
}

export interface QueryClaimRecordResponse {
  userAirdropEntries: UserAirdropEntries | undefined;
}

export interface QueryClaimRecordsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryClaimRecordsResponse {
  usersAirdropEntries: UserAirdropEntries[];
  pagination: PageResponse | undefined;
}

export interface QueryInitialClaimRequest {
  campaignId: number;
}

export interface QueryInitialClaimResponse {
  initialClaim: InitialClaim | undefined;
}

export interface QueryInitialClaimsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryInitialClaimsResponse {
  initialClaim: InitialClaim[];
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

export interface QueryGetAirdropEntryRequest {
  id: number;
}

export interface QueryGetAirdropEntryResponse {
  AirdropEntry: AirdropEntry | undefined;
}

export interface QueryAllAirdropEntryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAirdropEntryResponse {
  AirdropEntry: AirdropEntry[];
  pagination: PageResponse | undefined;
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

const baseQueryClaimRecordRequest: object = { address: "" };

export const QueryClaimRecordRequest = {
  encode(
    message: QueryClaimRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryClaimRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryClaimRecordRequest,
    } as QueryClaimRecordRequest;
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

  fromJSON(object: any): QueryClaimRecordRequest {
    const message = {
      ...baseQueryClaimRecordRequest,
    } as QueryClaimRecordRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryClaimRecordRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryClaimRecordRequest>
  ): QueryClaimRecordRequest {
    const message = {
      ...baseQueryClaimRecordRequest,
    } as QueryClaimRecordRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryClaimRecordResponse: object = {};

export const QueryClaimRecordResponse = {
  encode(
    message: QueryClaimRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userAirdropEntries !== undefined) {
      UserAirdropEntries.encode(
        message.userAirdropEntries,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryClaimRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryClaimRecordResponse,
    } as QueryClaimRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAirdropEntries = UserAirdropEntries.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryClaimRecordResponse {
    const message = {
      ...baseQueryClaimRecordResponse,
    } as QueryClaimRecordResponse;
    if (
      object.userAirdropEntries !== undefined &&
      object.userAirdropEntries !== null
    ) {
      message.userAirdropEntries = UserAirdropEntries.fromJSON(
        object.userAirdropEntries
      );
    } else {
      message.userAirdropEntries = undefined;
    }
    return message;
  },

  toJSON(message: QueryClaimRecordResponse): unknown {
    const obj: any = {};
    message.userAirdropEntries !== undefined &&
      (obj.userAirdropEntries = message.userAirdropEntries
        ? UserAirdropEntries.toJSON(message.userAirdropEntries)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryClaimRecordResponse>
  ): QueryClaimRecordResponse {
    const message = {
      ...baseQueryClaimRecordResponse,
    } as QueryClaimRecordResponse;
    if (
      object.userAirdropEntries !== undefined &&
      object.userAirdropEntries !== null
    ) {
      message.userAirdropEntries = UserAirdropEntries.fromPartial(
        object.userAirdropEntries
      );
    } else {
      message.userAirdropEntries = undefined;
    }
    return message;
  },
};

const baseQueryClaimRecordsRequest: object = {};

export const QueryClaimRecordsRequest = {
  encode(
    message: QueryClaimRecordsRequest,
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
  ): QueryClaimRecordsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryClaimRecordsRequest,
    } as QueryClaimRecordsRequest;
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

  fromJSON(object: any): QueryClaimRecordsRequest {
    const message = {
      ...baseQueryClaimRecordsRequest,
    } as QueryClaimRecordsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryClaimRecordsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryClaimRecordsRequest>
  ): QueryClaimRecordsRequest {
    const message = {
      ...baseQueryClaimRecordsRequest,
    } as QueryClaimRecordsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryClaimRecordsResponse: object = {};

export const QueryClaimRecordsResponse = {
  encode(
    message: QueryClaimRecordsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.usersAirdropEntries) {
      UserAirdropEntries.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryClaimRecordsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryClaimRecordsResponse,
    } as QueryClaimRecordsResponse;
    message.usersAirdropEntries = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.usersAirdropEntries.push(
            UserAirdropEntries.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryClaimRecordsResponse {
    const message = {
      ...baseQueryClaimRecordsResponse,
    } as QueryClaimRecordsResponse;
    message.usersAirdropEntries = [];
    if (
      object.usersAirdropEntries !== undefined &&
      object.usersAirdropEntries !== null
    ) {
      for (const e of object.usersAirdropEntries) {
        message.usersAirdropEntries.push(UserAirdropEntries.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryClaimRecordsResponse): unknown {
    const obj: any = {};
    if (message.usersAirdropEntries) {
      obj.usersAirdropEntries = message.usersAirdropEntries.map((e) =>
        e ? UserAirdropEntries.toJSON(e) : undefined
      );
    } else {
      obj.usersAirdropEntries = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryClaimRecordsResponse>
  ): QueryClaimRecordsResponse {
    const message = {
      ...baseQueryClaimRecordsResponse,
    } as QueryClaimRecordsResponse;
    message.usersAirdropEntries = [];
    if (
      object.usersAirdropEntries !== undefined &&
      object.usersAirdropEntries !== null
    ) {
      for (const e of object.usersAirdropEntries) {
        message.usersAirdropEntries.push(UserAirdropEntries.fromPartial(e));
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

const baseQueryInitialClaimRequest: object = { campaignId: 0 };

export const QueryInitialClaimRequest = {
  encode(
    message: QueryInitialClaimRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryInitialClaimRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInitialClaimRequest,
    } as QueryInitialClaimRequest;
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

  fromJSON(object: any): QueryInitialClaimRequest {
    const message = {
      ...baseQueryInitialClaimRequest,
    } as QueryInitialClaimRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = Number(object.campaignId);
    } else {
      message.campaignId = 0;
    }
    return message;
  },

  toJSON(message: QueryInitialClaimRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInitialClaimRequest>
  ): QueryInitialClaimRequest {
    const message = {
      ...baseQueryInitialClaimRequest,
    } as QueryInitialClaimRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = 0;
    }
    return message;
  },
};

const baseQueryInitialClaimResponse: object = {};

export const QueryInitialClaimResponse = {
  encode(
    message: QueryInitialClaimResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.initialClaim !== undefined) {
      InitialClaim.encode(
        message.initialClaim,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryInitialClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInitialClaimResponse,
    } as QueryInitialClaimResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.initialClaim = InitialClaim.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInitialClaimResponse {
    const message = {
      ...baseQueryInitialClaimResponse,
    } as QueryInitialClaimResponse;
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      message.initialClaim = InitialClaim.fromJSON(object.initialClaim);
    } else {
      message.initialClaim = undefined;
    }
    return message;
  },

  toJSON(message: QueryInitialClaimResponse): unknown {
    const obj: any = {};
    message.initialClaim !== undefined &&
      (obj.initialClaim = message.initialClaim
        ? InitialClaim.toJSON(message.initialClaim)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInitialClaimResponse>
  ): QueryInitialClaimResponse {
    const message = {
      ...baseQueryInitialClaimResponse,
    } as QueryInitialClaimResponse;
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      message.initialClaim = InitialClaim.fromPartial(object.initialClaim);
    } else {
      message.initialClaim = undefined;
    }
    return message;
  },
};

const baseQueryInitialClaimsRequest: object = {};

export const QueryInitialClaimsRequest = {
  encode(
    message: QueryInitialClaimsRequest,
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
  ): QueryInitialClaimsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInitialClaimsRequest,
    } as QueryInitialClaimsRequest;
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

  fromJSON(object: any): QueryInitialClaimsRequest {
    const message = {
      ...baseQueryInitialClaimsRequest,
    } as QueryInitialClaimsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryInitialClaimsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInitialClaimsRequest>
  ): QueryInitialClaimsRequest {
    const message = {
      ...baseQueryInitialClaimsRequest,
    } as QueryInitialClaimsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryInitialClaimsResponse: object = {};

export const QueryInitialClaimsResponse = {
  encode(
    message: QueryInitialClaimsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.initialClaim) {
      InitialClaim.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryInitialClaimsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryInitialClaimsResponse,
    } as QueryInitialClaimsResponse;
    message.initialClaim = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.initialClaim.push(
            InitialClaim.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryInitialClaimsResponse {
    const message = {
      ...baseQueryInitialClaimsResponse,
    } as QueryInitialClaimsResponse;
    message.initialClaim = [];
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      for (const e of object.initialClaim) {
        message.initialClaim.push(InitialClaim.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryInitialClaimsResponse): unknown {
    const obj: any = {};
    if (message.initialClaim) {
      obj.initialClaim = message.initialClaim.map((e) =>
        e ? InitialClaim.toJSON(e) : undefined
      );
    } else {
      obj.initialClaim = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInitialClaimsResponse>
  ): QueryInitialClaimsResponse {
    const message = {
      ...baseQueryInitialClaimsResponse,
    } as QueryInitialClaimsResponse;
    message.initialClaim = [];
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      for (const e of object.initialClaim) {
        message.initialClaim.push(InitialClaim.fromPartial(e));
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

const baseQueryMissionRequest: object = { campaignId: 0, missionId: 0 };

export const QueryMissionRequest = {
  encode(
    message: QueryMissionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(16).uint64(message.missionId);
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
    const message = { ...baseQueryMissionRequest } as QueryMissionRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = Number(object.campaignId);
    } else {
      message.campaignId = 0;
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = Number(object.missionId);
    } else {
      message.missionId = 0;
    }
    return message;
  },

  toJSON(message: QueryMissionRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryMissionRequest>): QueryMissionRequest {
    const message = { ...baseQueryMissionRequest } as QueryMissionRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = 0;
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = object.missionId;
    } else {
      message.missionId = 0;
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

const baseQueryCampaignRequest: object = { campaignId: 0 };

export const QueryCampaignRequest = {
  encode(
    message: QueryCampaignRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
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
    const message = { ...baseQueryCampaignRequest } as QueryCampaignRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = Number(object.campaignId);
    } else {
      message.campaignId = 0;
    }
    return message;
  },

  toJSON(message: QueryCampaignRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCampaignRequest>): QueryCampaignRequest {
    const message = { ...baseQueryCampaignRequest } as QueryCampaignRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = 0;
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

const baseQueryGetAirdropEntryRequest: object = { id: 0 };

export const QueryGetAirdropEntryRequest = {
  encode(
    message: QueryGetAirdropEntryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAirdropEntryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAirdropEntryRequest,
    } as QueryGetAirdropEntryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAirdropEntryRequest {
    const message = {
      ...baseQueryGetAirdropEntryRequest,
    } as QueryGetAirdropEntryRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetAirdropEntryRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAirdropEntryRequest>
  ): QueryGetAirdropEntryRequest {
    const message = {
      ...baseQueryGetAirdropEntryRequest,
    } as QueryGetAirdropEntryRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetAirdropEntryResponse: object = {};

export const QueryGetAirdropEntryResponse = {
  encode(
    message: QueryGetAirdropEntryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.AirdropEntry !== undefined) {
      AirdropEntry.encode(
        message.AirdropEntry,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAirdropEntryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAirdropEntryResponse,
    } as QueryGetAirdropEntryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AirdropEntry = AirdropEntry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAirdropEntryResponse {
    const message = {
      ...baseQueryGetAirdropEntryResponse,
    } as QueryGetAirdropEntryResponse;
    if (object.AirdropEntry !== undefined && object.AirdropEntry !== null) {
      message.AirdropEntry = AirdropEntry.fromJSON(object.AirdropEntry);
    } else {
      message.AirdropEntry = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAirdropEntryResponse): unknown {
    const obj: any = {};
    message.AirdropEntry !== undefined &&
      (obj.AirdropEntry = message.AirdropEntry
        ? AirdropEntry.toJSON(message.AirdropEntry)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAirdropEntryResponse>
  ): QueryGetAirdropEntryResponse {
    const message = {
      ...baseQueryGetAirdropEntryResponse,
    } as QueryGetAirdropEntryResponse;
    if (object.AirdropEntry !== undefined && object.AirdropEntry !== null) {
      message.AirdropEntry = AirdropEntry.fromPartial(object.AirdropEntry);
    } else {
      message.AirdropEntry = undefined;
    }
    return message;
  },
};

const baseQueryAllAirdropEntryRequest: object = {};

export const QueryAllAirdropEntryRequest = {
  encode(
    message: QueryAllAirdropEntryRequest,
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
  ): QueryAllAirdropEntryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAirdropEntryRequest,
    } as QueryAllAirdropEntryRequest;
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

  fromJSON(object: any): QueryAllAirdropEntryRequest {
    const message = {
      ...baseQueryAllAirdropEntryRequest,
    } as QueryAllAirdropEntryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAirdropEntryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAirdropEntryRequest>
  ): QueryAllAirdropEntryRequest {
    const message = {
      ...baseQueryAllAirdropEntryRequest,
    } as QueryAllAirdropEntryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAirdropEntryResponse: object = {};

export const QueryAllAirdropEntryResponse = {
  encode(
    message: QueryAllAirdropEntryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.AirdropEntry) {
      AirdropEntry.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllAirdropEntryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAirdropEntryResponse,
    } as QueryAllAirdropEntryResponse;
    message.AirdropEntry = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AirdropEntry.push(
            AirdropEntry.decode(reader, reader.uint32())
          );
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

  fromJSON(object: any): QueryAllAirdropEntryResponse {
    const message = {
      ...baseQueryAllAirdropEntryResponse,
    } as QueryAllAirdropEntryResponse;
    message.AirdropEntry = [];
    if (object.AirdropEntry !== undefined && object.AirdropEntry !== null) {
      for (const e of object.AirdropEntry) {
        message.AirdropEntry.push(AirdropEntry.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAirdropEntryResponse): unknown {
    const obj: any = {};
    if (message.AirdropEntry) {
      obj.AirdropEntry = message.AirdropEntry.map((e) =>
        e ? AirdropEntry.toJSON(e) : undefined
      );
    } else {
      obj.AirdropEntry = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAirdropEntryResponse>
  ): QueryAllAirdropEntryResponse {
    const message = {
      ...baseQueryAllAirdropEntryResponse,
    } as QueryAllAirdropEntryResponse;
    message.AirdropEntry = [];
    if (object.AirdropEntry !== undefined && object.AirdropEntry !== null) {
      for (const e of object.AirdropEntry) {
        message.AirdropEntry.push(AirdropEntry.fromPartial(e));
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a UserAirdropEntries by index. */
  UserAirdropEntries(
    request: QueryClaimRecordRequest
  ): Promise<QueryClaimRecordResponse>;
  /** Queries a list of UserAirdropEntries items. */
  ClaimRecords(
    request: QueryClaimRecordsRequest
  ): Promise<QueryClaimRecordsResponse>;
  /** Queries a InitialClaim by index. */
  InitialClaim(
    request: QueryInitialClaimRequest
  ): Promise<QueryInitialClaimResponse>;
  /** Queries a list of InitialClaim items. */
  InitialClaims(
    request: QueryInitialClaimsRequest
  ): Promise<QueryInitialClaimsResponse>;
  /** Queries a Mission by index. */
  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse>;
  /** Queries a list of Mission items. */
  MissionAll(request: QueryMissionsRequest): Promise<QueryMissionsResponse>;
  /** Queries a list of Campaigns items. */
  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse>;
  /** Queries a list of Campaigns items. */
  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse>;
  /** Queries a AirdropEntry by id. */
  AirdropEntry(
    request: QueryGetAirdropEntryRequest
  ): Promise<QueryGetAirdropEntryResponse>;
  /** Queries a list of AirdropEntry items. */
  AirdropEntryAll(
    request: QueryAllAirdropEntryRequest
  ): Promise<QueryAllAirdropEntryResponse>;
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

  UserAirdropEntries(
    request: QueryClaimRecordRequest
  ): Promise<QueryClaimRecordResponse> {
    const data = QueryClaimRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "UserAirdropEntries",
      data
    );
    return promise.then((data) =>
      QueryClaimRecordResponse.decode(new Reader(data))
    );
  }

  ClaimRecords(
    request: QueryClaimRecordsRequest
  ): Promise<QueryClaimRecordsResponse> {
    const data = QueryClaimRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "ClaimRecords",
      data
    );
    return promise.then((data) =>
      QueryClaimRecordsResponse.decode(new Reader(data))
    );
  }

  InitialClaim(
    request: QueryInitialClaimRequest
  ): Promise<QueryInitialClaimResponse> {
    const data = QueryInitialClaimRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "InitialClaim",
      data
    );
    return promise.then((data) =>
      QueryInitialClaimResponse.decode(new Reader(data))
    );
  }

  InitialClaims(
    request: QueryInitialClaimsRequest
  ): Promise<QueryInitialClaimsResponse> {
    const data = QueryInitialClaimsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "InitialClaims",
      data
    );
    return promise.then((data) =>
      QueryInitialClaimsResponse.decode(new Reader(data))
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

  AirdropEntry(
    request: QueryGetAirdropEntryRequest
  ): Promise<QueryGetAirdropEntryResponse> {
    const data = QueryGetAirdropEntryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "AirdropEntry",
      data
    );
    return promise.then((data) =>
      QueryGetAirdropEntryResponse.decode(new Reader(data))
    );
  }

  AirdropEntryAll(
    request: QueryAllAirdropEntryRequest
  ): Promise<QueryAllAirdropEntryResponse> {
    const data = QueryAllAirdropEntryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "AirdropEntryAll",
      data
    );
    return promise.then((data) =>
      QueryAllAirdropEntryResponse.decode(new Reader(data))
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
