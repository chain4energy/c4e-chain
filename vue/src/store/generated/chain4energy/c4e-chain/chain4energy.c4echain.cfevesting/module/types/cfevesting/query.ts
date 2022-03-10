/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cfevesting/params";
import { VestingTypes } from "../cfevesting/vesting_types";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryVestingTypeRequest {}

export interface QueryVestingTypeResponse {
  vestingTypes: VestingTypes | undefined;
}

export interface QueryVestingRequest {
  address: string;
}

export interface QueryVestingResponse {
  delegableAddress: string;
  vestings: VestingInfo[];
}

export interface VestingInfo {
  vestingType: string;
  vestingStartHeight: number;
  lockEndHeight: number;
  vestingEndHeight: number;
  withdrawable: string;
  delegationAllowed: boolean;
  vested: Coin | undefined;
  currentVestedAmount: string;
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

const baseQueryVestingTypeRequest: object = {};

export const QueryVestingTypeRequest = {
  encode(_: QueryVestingTypeRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryVestingTypeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingTypeRequest,
    } as QueryVestingTypeRequest;
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

  fromJSON(_: any): QueryVestingTypeRequest {
    const message = {
      ...baseQueryVestingTypeRequest,
    } as QueryVestingTypeRequest;
    return message;
  },

  toJSON(_: QueryVestingTypeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryVestingTypeRequest>
  ): QueryVestingTypeRequest {
    const message = {
      ...baseQueryVestingTypeRequest,
    } as QueryVestingTypeRequest;
    return message;
  },
};

const baseQueryVestingTypeResponse: object = {};

export const QueryVestingTypeResponse = {
  encode(
    message: QueryVestingTypeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vestingTypes !== undefined) {
      VestingTypes.encode(
        message.vestingTypes,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVestingTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingTypeResponse,
    } as QueryVestingTypeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestingTypes = VestingTypes.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingTypeResponse {
    const message = {
      ...baseQueryVestingTypeResponse,
    } as QueryVestingTypeResponse;
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      message.vestingTypes = VestingTypes.fromJSON(object.vestingTypes);
    } else {
      message.vestingTypes = undefined;
    }
    return message;
  },

  toJSON(message: QueryVestingTypeResponse): unknown {
    const obj: any = {};
    message.vestingTypes !== undefined &&
      (obj.vestingTypes = message.vestingTypes
        ? VestingTypes.toJSON(message.vestingTypes)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingTypeResponse>
  ): QueryVestingTypeResponse {
    const message = {
      ...baseQueryVestingTypeResponse,
    } as QueryVestingTypeResponse;
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      message.vestingTypes = VestingTypes.fromPartial(object.vestingTypes);
    } else {
      message.vestingTypes = undefined;
    }
    return message;
  },
};

const baseQueryVestingRequest: object = { address: "" };

export const QueryVestingRequest = {
  encode(
    message: QueryVestingRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryVestingRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryVestingRequest } as QueryVestingRequest;
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

  fromJSON(object: any): QueryVestingRequest {
    const message = { ...baseQueryVestingRequest } as QueryVestingRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryVestingRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryVestingRequest>): QueryVestingRequest {
    const message = { ...baseQueryVestingRequest } as QueryVestingRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryVestingResponse: object = { delegableAddress: "" };

export const QueryVestingResponse = {
  encode(
    message: QueryVestingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegableAddress !== "") {
      writer.uint32(10).string(message.delegableAddress);
    }
    for (const v of message.vestings) {
      VestingInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryVestingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryVestingResponse } as QueryVestingResponse;
    message.vestings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegableAddress = reader.string();
          break;
        case 2:
          message.vestings.push(VestingInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingResponse {
    const message = { ...baseQueryVestingResponse } as QueryVestingResponse;
    message.vestings = [];
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = String(object.delegableAddress);
    } else {
      message.delegableAddress = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(VestingInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryVestingResponse): unknown {
    const obj: any = {};
    message.delegableAddress !== undefined &&
      (obj.delegableAddress = message.delegableAddress);
    if (message.vestings) {
      obj.vestings = message.vestings.map((e) =>
        e ? VestingInfo.toJSON(e) : undefined
      );
    } else {
      obj.vestings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QueryVestingResponse>): QueryVestingResponse {
    const message = { ...baseQueryVestingResponse } as QueryVestingResponse;
    message.vestings = [];
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = object.delegableAddress;
    } else {
      message.delegableAddress = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(VestingInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingInfo: object = {
  vestingType: "",
  vestingStartHeight: 0,
  lockEndHeight: 0,
  vestingEndHeight: 0,
  withdrawable: "",
  delegationAllowed: false,
  currentVestedAmount: "",
};

export const VestingInfo = {
  encode(message: VestingInfo, writer: Writer = Writer.create()): Writer {
    if (message.vestingType !== "") {
      writer.uint32(10).string(message.vestingType);
    }
    if (message.vestingStartHeight !== 0) {
      writer.uint32(16).int64(message.vestingStartHeight);
    }
    if (message.lockEndHeight !== 0) {
      writer.uint32(24).int64(message.lockEndHeight);
    }
    if (message.vestingEndHeight !== 0) {
      writer.uint32(32).int64(message.vestingEndHeight);
    }
    if (message.withdrawable !== "") {
      writer.uint32(42).string(message.withdrawable);
    }
    if (message.delegationAllowed === true) {
      writer.uint32(48).bool(message.delegationAllowed);
    }
    if (message.vested !== undefined) {
      Coin.encode(message.vested, writer.uint32(58).fork()).ldelim();
    }
    if (message.currentVestedAmount !== "") {
      writer.uint32(66).string(message.currentVestedAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingInfo } as VestingInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestingType = reader.string();
          break;
        case 2:
          message.vestingStartHeight = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.lockEndHeight = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.vestingEndHeight = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.withdrawable = reader.string();
          break;
        case 6:
          message.delegationAllowed = reader.bool();
          break;
        case 7:
          message.vested = Coin.decode(reader, reader.uint32());
          break;
        case 8:
          message.currentVestedAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingInfo {
    const message = { ...baseVestingInfo } as VestingInfo;
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    if (
      object.vestingStartHeight !== undefined &&
      object.vestingStartHeight !== null
    ) {
      message.vestingStartHeight = Number(object.vestingStartHeight);
    } else {
      message.vestingStartHeight = 0;
    }
    if (object.lockEndHeight !== undefined && object.lockEndHeight !== null) {
      message.lockEndHeight = Number(object.lockEndHeight);
    } else {
      message.lockEndHeight = 0;
    }
    if (
      object.vestingEndHeight !== undefined &&
      object.vestingEndHeight !== null
    ) {
      message.vestingEndHeight = Number(object.vestingEndHeight);
    } else {
      message.vestingEndHeight = 0;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = String(object.withdrawable);
    } else {
      message.withdrawable = "";
    }
    if (
      object.delegationAllowed !== undefined &&
      object.delegationAllowed !== null
    ) {
      message.delegationAllowed = Boolean(object.delegationAllowed);
    } else {
      message.delegationAllowed = false;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = Coin.fromJSON(object.vested);
    } else {
      message.vested = undefined;
    }
    if (
      object.currentVestedAmount !== undefined &&
      object.currentVestedAmount !== null
    ) {
      message.currentVestedAmount = String(object.currentVestedAmount);
    } else {
      message.currentVestedAmount = "";
    }
    return message;
  },

  toJSON(message: VestingInfo): unknown {
    const obj: any = {};
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    message.vestingStartHeight !== undefined &&
      (obj.vestingStartHeight = message.vestingStartHeight);
    message.lockEndHeight !== undefined &&
      (obj.lockEndHeight = message.lockEndHeight);
    message.vestingEndHeight !== undefined &&
      (obj.vestingEndHeight = message.vestingEndHeight);
    message.withdrawable !== undefined &&
      (obj.withdrawable = message.withdrawable);
    message.delegationAllowed !== undefined &&
      (obj.delegationAllowed = message.delegationAllowed);
    message.vested !== undefined &&
      (obj.vested = message.vested ? Coin.toJSON(message.vested) : undefined);
    message.currentVestedAmount !== undefined &&
      (obj.currentVestedAmount = message.currentVestedAmount);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingInfo>): VestingInfo {
    const message = { ...baseVestingInfo } as VestingInfo;
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    if (
      object.vestingStartHeight !== undefined &&
      object.vestingStartHeight !== null
    ) {
      message.vestingStartHeight = object.vestingStartHeight;
    } else {
      message.vestingStartHeight = 0;
    }
    if (object.lockEndHeight !== undefined && object.lockEndHeight !== null) {
      message.lockEndHeight = object.lockEndHeight;
    } else {
      message.lockEndHeight = 0;
    }
    if (
      object.vestingEndHeight !== undefined &&
      object.vestingEndHeight !== null
    ) {
      message.vestingEndHeight = object.vestingEndHeight;
    } else {
      message.vestingEndHeight = 0;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = object.withdrawable;
    } else {
      message.withdrawable = "";
    }
    if (
      object.delegationAllowed !== undefined &&
      object.delegationAllowed !== null
    ) {
      message.delegationAllowed = object.delegationAllowed;
    } else {
      message.delegationAllowed = false;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = Coin.fromPartial(object.vested);
    } else {
      message.vested = undefined;
    }
    if (
      object.currentVestedAmount !== undefined &&
      object.currentVestedAmount !== null
    ) {
      message.currentVestedAmount = object.currentVestedAmount;
    } else {
      message.currentVestedAmount = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of VestingType items. */
  VestingType(
    request: QueryVestingTypeRequest
  ): Promise<QueryVestingTypeResponse>;
  /** Queries a list of Vesting items. */
  Vesting(request: QueryVestingRequest): Promise<QueryVestingResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  VestingType(
    request: QueryVestingTypeRequest
  ): Promise<QueryVestingTypeResponse> {
    const data = QueryVestingTypeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Query",
      "VestingType",
      data
    );
    return promise.then((data) =>
      QueryVestingTypeResponse.decode(new Reader(data))
    );
  }

  Vesting(request: QueryVestingRequest): Promise<QueryVestingResponse> {
    const data = QueryVestingRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Query",
      "Vesting",
      data
    );
    return promise.then((data) =>
      QueryVestingResponse.decode(new Reader(data))
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
