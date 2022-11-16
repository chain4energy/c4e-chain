/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cfeairdrop/params";
import { ClaimRecord, InitialClaim, Mission } from "../cfeairdrop/airdrop";
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

export interface QueryGetClaimRecordRequest {
  address: string;
}

export interface QueryGetClaimRecordResponse {
  claimRecord: ClaimRecord | undefined;
}

export interface QueryAllClaimRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClaimRecordResponse {
  claimRecord: ClaimRecord[];
  pagination: PageResponse | undefined;
}

export interface QueryGetInitialClaimRequest {
  campaignId: number;
}

export interface QueryGetInitialClaimResponse {
  initialClaim: InitialClaim | undefined;
}

export interface QueryAllInitialClaimRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllInitialClaimResponse {
  initialClaim: InitialClaim[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMissionRequest {
  campaignId: number;
  missionId: number;
}

export interface QueryGetMissionResponse {
  mission: Mission | undefined;
}

export interface QueryAllMissionRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMissionResponse {
  mission: Mission[];
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

const baseQueryGetClaimRecordRequest: object = { address: "" };

export const QueryGetClaimRecordRequest = {
  encode(
    message: QueryGetClaimRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimRecordRequest,
    } as QueryGetClaimRecordRequest;
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

  fromJSON(object: any): QueryGetClaimRecordRequest {
    const message = {
      ...baseQueryGetClaimRecordRequest,
    } as QueryGetClaimRecordRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetClaimRecordRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimRecordRequest>
  ): QueryGetClaimRecordRequest {
    const message = {
      ...baseQueryGetClaimRecordRequest,
    } as QueryGetClaimRecordRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetClaimRecordResponse: object = {};

export const QueryGetClaimRecordResponse = {
  encode(
    message: QueryGetClaimRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.claimRecord !== undefined) {
      ClaimRecord.encode(
        message.claimRecord,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimRecordResponse,
    } as QueryGetClaimRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimRecord = ClaimRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClaimRecordResponse {
    const message = {
      ...baseQueryGetClaimRecordResponse,
    } as QueryGetClaimRecordResponse;
    if (object.claimRecord !== undefined && object.claimRecord !== null) {
      message.claimRecord = ClaimRecord.fromJSON(object.claimRecord);
    } else {
      message.claimRecord = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClaimRecordResponse): unknown {
    const obj: any = {};
    message.claimRecord !== undefined &&
      (obj.claimRecord = message.claimRecord
        ? ClaimRecord.toJSON(message.claimRecord)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimRecordResponse>
  ): QueryGetClaimRecordResponse {
    const message = {
      ...baseQueryGetClaimRecordResponse,
    } as QueryGetClaimRecordResponse;
    if (object.claimRecord !== undefined && object.claimRecord !== null) {
      message.claimRecord = ClaimRecord.fromPartial(object.claimRecord);
    } else {
      message.claimRecord = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimRecordRequest: object = {};

export const QueryAllClaimRecordRequest = {
  encode(
    message: QueryAllClaimRecordRequest,
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
  ): QueryAllClaimRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimRecordRequest,
    } as QueryAllClaimRecordRequest;
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

  fromJSON(object: any): QueryAllClaimRecordRequest {
    const message = {
      ...baseQueryAllClaimRecordRequest,
    } as QueryAllClaimRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimRecordRequest>
  ): QueryAllClaimRecordRequest {
    const message = {
      ...baseQueryAllClaimRecordRequest,
    } as QueryAllClaimRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimRecordResponse: object = {};

export const QueryAllClaimRecordResponse = {
  encode(
    message: QueryAllClaimRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.claimRecord) {
      ClaimRecord.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllClaimRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimRecordResponse,
    } as QueryAllClaimRecordResponse;
    message.claimRecord = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimRecord.push(ClaimRecord.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllClaimRecordResponse {
    const message = {
      ...baseQueryAllClaimRecordResponse,
    } as QueryAllClaimRecordResponse;
    message.claimRecord = [];
    if (object.claimRecord !== undefined && object.claimRecord !== null) {
      for (const e of object.claimRecord) {
        message.claimRecord.push(ClaimRecord.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimRecordResponse): unknown {
    const obj: any = {};
    if (message.claimRecord) {
      obj.claimRecord = message.claimRecord.map((e) =>
        e ? ClaimRecord.toJSON(e) : undefined
      );
    } else {
      obj.claimRecord = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimRecordResponse>
  ): QueryAllClaimRecordResponse {
    const message = {
      ...baseQueryAllClaimRecordResponse,
    } as QueryAllClaimRecordResponse;
    message.claimRecord = [];
    if (object.claimRecord !== undefined && object.claimRecord !== null) {
      for (const e of object.claimRecord) {
        message.claimRecord.push(ClaimRecord.fromPartial(e));
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

const baseQueryGetInitialClaimRequest: object = { campaignId: 0 };

export const QueryGetInitialClaimRequest = {
  encode(
    message: QueryGetInitialClaimRequest,
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
  ): QueryGetInitialClaimRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInitialClaimRequest,
    } as QueryGetInitialClaimRequest;
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

  fromJSON(object: any): QueryGetInitialClaimRequest {
    const message = {
      ...baseQueryGetInitialClaimRequest,
    } as QueryGetInitialClaimRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = Number(object.campaignId);
    } else {
      message.campaignId = 0;
    }
    return message;
  },

  toJSON(message: QueryGetInitialClaimRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInitialClaimRequest>
  ): QueryGetInitialClaimRequest {
    const message = {
      ...baseQueryGetInitialClaimRequest,
    } as QueryGetInitialClaimRequest;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = 0;
    }
    return message;
  },
};

const baseQueryGetInitialClaimResponse: object = {};

export const QueryGetInitialClaimResponse = {
  encode(
    message: QueryGetInitialClaimResponse,
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
  ): QueryGetInitialClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInitialClaimResponse,
    } as QueryGetInitialClaimResponse;
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

  fromJSON(object: any): QueryGetInitialClaimResponse {
    const message = {
      ...baseQueryGetInitialClaimResponse,
    } as QueryGetInitialClaimResponse;
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      message.initialClaim = InitialClaim.fromJSON(object.initialClaim);
    } else {
      message.initialClaim = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInitialClaimResponse): unknown {
    const obj: any = {};
    message.initialClaim !== undefined &&
      (obj.initialClaim = message.initialClaim
        ? InitialClaim.toJSON(message.initialClaim)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInitialClaimResponse>
  ): QueryGetInitialClaimResponse {
    const message = {
      ...baseQueryGetInitialClaimResponse,
    } as QueryGetInitialClaimResponse;
    if (object.initialClaim !== undefined && object.initialClaim !== null) {
      message.initialClaim = InitialClaim.fromPartial(object.initialClaim);
    } else {
      message.initialClaim = undefined;
    }
    return message;
  },
};

const baseQueryAllInitialClaimRequest: object = {};

export const QueryAllInitialClaimRequest = {
  encode(
    message: QueryAllInitialClaimRequest,
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
  ): QueryAllInitialClaimRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInitialClaimRequest,
    } as QueryAllInitialClaimRequest;
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

  fromJSON(object: any): QueryAllInitialClaimRequest {
    const message = {
      ...baseQueryAllInitialClaimRequest,
    } as QueryAllInitialClaimRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInitialClaimRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInitialClaimRequest>
  ): QueryAllInitialClaimRequest {
    const message = {
      ...baseQueryAllInitialClaimRequest,
    } as QueryAllInitialClaimRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInitialClaimResponse: object = {};

export const QueryAllInitialClaimResponse = {
  encode(
    message: QueryAllInitialClaimResponse,
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
  ): QueryAllInitialClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInitialClaimResponse,
    } as QueryAllInitialClaimResponse;
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

  fromJSON(object: any): QueryAllInitialClaimResponse {
    const message = {
      ...baseQueryAllInitialClaimResponse,
    } as QueryAllInitialClaimResponse;
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

  toJSON(message: QueryAllInitialClaimResponse): unknown {
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
    object: DeepPartial<QueryAllInitialClaimResponse>
  ): QueryAllInitialClaimResponse {
    const message = {
      ...baseQueryAllInitialClaimResponse,
    } as QueryAllInitialClaimResponse;
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

const baseQueryGetMissionRequest: object = { campaignId: 0, missionId: 0 };

export const QueryGetMissionRequest = {
  encode(
    message: QueryGetMissionRequest,
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

  decode(input: Reader | Uint8Array, length?: number): QueryGetMissionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMissionRequest } as QueryGetMissionRequest;
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

  fromJSON(object: any): QueryGetMissionRequest {
    const message = { ...baseQueryGetMissionRequest } as QueryGetMissionRequest;
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

  toJSON(message: QueryGetMissionRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMissionRequest>
  ): QueryGetMissionRequest {
    const message = { ...baseQueryGetMissionRequest } as QueryGetMissionRequest;
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

const baseQueryGetMissionResponse: object = {};

export const QueryGetMissionResponse = {
  encode(
    message: QueryGetMissionResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mission !== undefined) {
      Mission.encode(message.mission, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetMissionResponse,
    } as QueryGetMissionResponse;
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

  fromJSON(object: any): QueryGetMissionResponse {
    const message = {
      ...baseQueryGetMissionResponse,
    } as QueryGetMissionResponse;
    if (object.mission !== undefined && object.mission !== null) {
      message.mission = Mission.fromJSON(object.mission);
    } else {
      message.mission = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMissionResponse): unknown {
    const obj: any = {};
    message.mission !== undefined &&
      (obj.mission = message.mission
        ? Mission.toJSON(message.mission)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMissionResponse>
  ): QueryGetMissionResponse {
    const message = {
      ...baseQueryGetMissionResponse,
    } as QueryGetMissionResponse;
    if (object.mission !== undefined && object.mission !== null) {
      message.mission = Mission.fromPartial(object.mission);
    } else {
      message.mission = undefined;
    }
    return message;
  },
};

const baseQueryAllMissionRequest: object = {};

export const QueryAllMissionRequest = {
  encode(
    message: QueryAllMissionRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMissionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMissionRequest } as QueryAllMissionRequest;
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

  fromJSON(object: any): QueryAllMissionRequest {
    const message = { ...baseQueryAllMissionRequest } as QueryAllMissionRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMissionRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMissionRequest>
  ): QueryAllMissionRequest {
    const message = { ...baseQueryAllMissionRequest } as QueryAllMissionRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMissionResponse: object = {};

export const QueryAllMissionResponse = {
  encode(
    message: QueryAllMissionResponse,
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

  decode(input: Reader | Uint8Array, length?: number): QueryAllMissionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllMissionResponse,
    } as QueryAllMissionResponse;
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

  fromJSON(object: any): QueryAllMissionResponse {
    const message = {
      ...baseQueryAllMissionResponse,
    } as QueryAllMissionResponse;
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

  toJSON(message: QueryAllMissionResponse): unknown {
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
    object: DeepPartial<QueryAllMissionResponse>
  ): QueryAllMissionResponse {
    const message = {
      ...baseQueryAllMissionResponse,
    } as QueryAllMissionResponse;
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a ClaimRecord by index. */
  ClaimRecord(
    request: QueryGetClaimRecordRequest
  ): Promise<QueryGetClaimRecordResponse>;
  /** Queries a list of ClaimRecord items. */
  ClaimRecordAll(
    request: QueryAllClaimRecordRequest
  ): Promise<QueryAllClaimRecordResponse>;
  /** Queries a InitialClaim by index. */
  InitialClaim(
    request: QueryGetInitialClaimRequest
  ): Promise<QueryGetInitialClaimResponse>;
  /** Queries a list of InitialClaim items. */
  InitialClaimAll(
    request: QueryAllInitialClaimRequest
  ): Promise<QueryAllInitialClaimResponse>;
  /** Queries a Mission by index. */
  Mission(request: QueryGetMissionRequest): Promise<QueryGetMissionResponse>;
  /** Queries a list of Mission items. */
  MissionAll(request: QueryAllMissionRequest): Promise<QueryAllMissionResponse>;
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

  ClaimRecord(
    request: QueryGetClaimRecordRequest
  ): Promise<QueryGetClaimRecordResponse> {
    const data = QueryGetClaimRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "ClaimRecord",
      data
    );
    return promise.then((data) =>
      QueryGetClaimRecordResponse.decode(new Reader(data))
    );
  }

  ClaimRecordAll(
    request: QueryAllClaimRecordRequest
  ): Promise<QueryAllClaimRecordResponse> {
    const data = QueryAllClaimRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "ClaimRecordAll",
      data
    );
    return promise.then((data) =>
      QueryAllClaimRecordResponse.decode(new Reader(data))
    );
  }

  InitialClaim(
    request: QueryGetInitialClaimRequest
  ): Promise<QueryGetInitialClaimResponse> {
    const data = QueryGetInitialClaimRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "InitialClaim",
      data
    );
    return promise.then((data) =>
      QueryGetInitialClaimResponse.decode(new Reader(data))
    );
  }

  InitialClaimAll(
    request: QueryAllInitialClaimRequest
  ): Promise<QueryAllInitialClaimResponse> {
    const data = QueryAllInitialClaimRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "InitialClaimAll",
      data
    );
    return promise.then((data) =>
      QueryAllInitialClaimResponse.decode(new Reader(data))
    );
  }

  Mission(request: QueryGetMissionRequest): Promise<QueryGetMissionResponse> {
    const data = QueryGetMissionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "Mission",
      data
    );
    return promise.then((data) =>
      QueryGetMissionResponse.decode(new Reader(data))
    );
  }

  MissionAll(
    request: QueryAllMissionRequest
  ): Promise<QueryAllMissionResponse> {
    const data = QueryAllMissionRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "MissionAll",
      data
    );
    return promise.then((data) =>
      QueryAllMissionResponse.decode(new Reader(data))
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
