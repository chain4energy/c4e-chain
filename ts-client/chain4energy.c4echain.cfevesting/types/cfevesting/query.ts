/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import { Params } from "../cfevesting/params";
import { GenesisVestingType } from "../cfevesting/genesis";
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
  vestingTypes: GenesisVestingType[];
}

export interface QueryVestingPoolsRequest {
  address: string;
}

export interface QueryVestingPoolsResponse {
  delegableAddress: string;
  vestingPools: VestingPoolInfo[];
}

export interface VestingPoolInfo {
  name: string;
  vestingType: string;
  lockStart: Date | undefined;
  lockEnd: Date | undefined;
  withdrawable: string;
  initiallyLocked: Coin | undefined;
  currentlyLocked: string;
  sentAmount: string;
}

export interface QueryVestingsSummaryRequest {}

export interface QueryVestingsSummaryResponse {
  vestingAllAmount: string;
  vestingInPoolsAmount: string;
  vestingInAccountsAmount: string;
  delegatedVestingAmount: string;
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
    for (const v of message.vestingTypes) {
      GenesisVestingType.encode(v!, writer.uint32(18).fork()).ldelim();
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
    message.vestingTypes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.vestingTypes.push(
            GenesisVestingType.decode(reader, reader.uint32())
          );
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
    message.vestingTypes = [];
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(GenesisVestingType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryVestingTypeResponse): unknown {
    const obj: any = {};
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) =>
        e ? GenesisVestingType.toJSON(e) : undefined
      );
    } else {
      obj.vestingTypes = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingTypeResponse>
  ): QueryVestingTypeResponse {
    const message = {
      ...baseQueryVestingTypeResponse,
    } as QueryVestingTypeResponse;
    message.vestingTypes = [];
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(GenesisVestingType.fromPartial(e));
      }
    }
    return message;
  },
};

const baseQueryVestingPoolsRequest: object = { address: "" };

export const QueryVestingPoolsRequest = {
  encode(
    message: QueryVestingPoolsRequest,
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
  ): QueryVestingPoolsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingPoolsRequest,
    } as QueryVestingPoolsRequest;
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

  fromJSON(object: any): QueryVestingPoolsRequest {
    const message = {
      ...baseQueryVestingPoolsRequest,
    } as QueryVestingPoolsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryVestingPoolsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingPoolsRequest>
  ): QueryVestingPoolsRequest {
    const message = {
      ...baseQueryVestingPoolsRequest,
    } as QueryVestingPoolsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryVestingPoolsResponse: object = { delegableAddress: "" };

export const QueryVestingPoolsResponse = {
  encode(
    message: QueryVestingPoolsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegableAddress !== "") {
      writer.uint32(10).string(message.delegableAddress);
    }
    for (const v of message.vestingPools) {
      VestingPoolInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVestingPoolsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingPoolsResponse,
    } as QueryVestingPoolsResponse;
    message.vestingPools = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegableAddress = reader.string();
          break;
        case 2:
          message.vestingPools.push(
            VestingPoolInfo.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingPoolsResponse {
    const message = {
      ...baseQueryVestingPoolsResponse,
    } as QueryVestingPoolsResponse;
    message.vestingPools = [];
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = String(object.delegableAddress);
    } else {
      message.delegableAddress = "";
    }
    if (object.vestingPools !== undefined && object.vestingPools !== null) {
      for (const e of object.vestingPools) {
        message.vestingPools.push(VestingPoolInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryVestingPoolsResponse): unknown {
    const obj: any = {};
    message.delegableAddress !== undefined &&
      (obj.delegableAddress = message.delegableAddress);
    if (message.vestingPools) {
      obj.vestingPools = message.vestingPools.map((e) =>
        e ? VestingPoolInfo.toJSON(e) : undefined
      );
    } else {
      obj.vestingPools = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingPoolsResponse>
  ): QueryVestingPoolsResponse {
    const message = {
      ...baseQueryVestingPoolsResponse,
    } as QueryVestingPoolsResponse;
    message.vestingPools = [];
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = object.delegableAddress;
    } else {
      message.delegableAddress = "";
    }
    if (object.vestingPools !== undefined && object.vestingPools !== null) {
      for (const e of object.vestingPools) {
        message.vestingPools.push(VestingPoolInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingPoolInfo: object = {
  name: "",
  vestingType: "",
  withdrawable: "",
  currentlyLocked: "",
  sentAmount: "",
};

export const VestingPoolInfo = {
  encode(message: VestingPoolInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.vestingType !== "") {
      writer.uint32(26).string(message.vestingType);
    }
    if (message.lockStart !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockStart),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.lockEnd !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockEnd),
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.withdrawable !== "") {
      writer.uint32(50).string(message.withdrawable);
    }
    if (message.initiallyLocked !== undefined) {
      Coin.encode(message.initiallyLocked, writer.uint32(58).fork()).ldelim();
    }
    if (message.currentlyLocked !== "") {
      writer.uint32(66).string(message.currentlyLocked);
    }
    if (message.sentAmount !== "") {
      writer.uint32(74).string(message.sentAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingPoolInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingPoolInfo } as VestingPoolInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.vestingType = reader.string();
          break;
        case 4:
          message.lockStart = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.lockEnd = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.withdrawable = reader.string();
          break;
        case 7:
          message.initiallyLocked = Coin.decode(reader, reader.uint32());
          break;
        case 8:
          message.currentlyLocked = reader.string();
          break;
        case 9:
          message.sentAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingPoolInfo {
    const message = { ...baseVestingPoolInfo } as VestingPoolInfo;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    if (object.lockStart !== undefined && object.lockStart !== null) {
      message.lockStart = fromJsonTimestamp(object.lockStart);
    } else {
      message.lockStart = undefined;
    }
    if (object.lockEnd !== undefined && object.lockEnd !== null) {
      message.lockEnd = fromJsonTimestamp(object.lockEnd);
    } else {
      message.lockEnd = undefined;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = String(object.withdrawable);
    } else {
      message.withdrawable = "";
    }
    if (
      object.initiallyLocked !== undefined &&
      object.initiallyLocked !== null
    ) {
      message.initiallyLocked = Coin.fromJSON(object.initiallyLocked);
    } else {
      message.initiallyLocked = undefined;
    }
    if (
      object.currentlyLocked !== undefined &&
      object.currentlyLocked !== null
    ) {
      message.currentlyLocked = String(object.currentlyLocked);
    } else {
      message.currentlyLocked = "";
    }
    if (object.sentAmount !== undefined && object.sentAmount !== null) {
      message.sentAmount = String(object.sentAmount);
    } else {
      message.sentAmount = "";
    }
    return message;
  },

  toJSON(message: VestingPoolInfo): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    message.lockStart !== undefined &&
      (obj.lockStart =
        message.lockStart !== undefined
          ? message.lockStart.toISOString()
          : null);
    message.lockEnd !== undefined &&
      (obj.lockEnd =
        message.lockEnd !== undefined ? message.lockEnd.toISOString() : null);
    message.withdrawable !== undefined &&
      (obj.withdrawable = message.withdrawable);
    message.initiallyLocked !== undefined &&
      (obj.initiallyLocked = message.initiallyLocked
        ? Coin.toJSON(message.initiallyLocked)
        : undefined);
    message.currentlyLocked !== undefined &&
      (obj.currentlyLocked = message.currentlyLocked);
    message.sentAmount !== undefined && (obj.sentAmount = message.sentAmount);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingPoolInfo>): VestingPoolInfo {
    const message = { ...baseVestingPoolInfo } as VestingPoolInfo;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    if (object.lockStart !== undefined && object.lockStart !== null) {
      message.lockStart = object.lockStart;
    } else {
      message.lockStart = undefined;
    }
    if (object.lockEnd !== undefined && object.lockEnd !== null) {
      message.lockEnd = object.lockEnd;
    } else {
      message.lockEnd = undefined;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = object.withdrawable;
    } else {
      message.withdrawable = "";
    }
    if (
      object.initiallyLocked !== undefined &&
      object.initiallyLocked !== null
    ) {
      message.initiallyLocked = Coin.fromPartial(object.initiallyLocked);
    } else {
      message.initiallyLocked = undefined;
    }
    if (
      object.currentlyLocked !== undefined &&
      object.currentlyLocked !== null
    ) {
      message.currentlyLocked = object.currentlyLocked;
    } else {
      message.currentlyLocked = "";
    }
    if (object.sentAmount !== undefined && object.sentAmount !== null) {
      message.sentAmount = object.sentAmount;
    } else {
      message.sentAmount = "";
    }
    return message;
  },
};

const baseQueryVestingsSummaryRequest: object = {};

export const QueryVestingsSummaryRequest = {
  encode(
    _: QueryVestingsSummaryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVestingsSummaryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingsSummaryRequest,
    } as QueryVestingsSummaryRequest;
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

  fromJSON(_: any): QueryVestingsSummaryRequest {
    const message = {
      ...baseQueryVestingsSummaryRequest,
    } as QueryVestingsSummaryRequest;
    return message;
  },

  toJSON(_: QueryVestingsSummaryRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryVestingsSummaryRequest>
  ): QueryVestingsSummaryRequest {
    const message = {
      ...baseQueryVestingsSummaryRequest,
    } as QueryVestingsSummaryRequest;
    return message;
  },
};

const baseQueryVestingsSummaryResponse: object = {
  vestingAllAmount: "",
  vestingInPoolsAmount: "",
  vestingInAccountsAmount: "",
  delegatedVestingAmount: "",
};

export const QueryVestingsSummaryResponse = {
  encode(
    message: QueryVestingsSummaryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vestingAllAmount !== "") {
      writer.uint32(10).string(message.vestingAllAmount);
    }
    if (message.vestingInPoolsAmount !== "") {
      writer.uint32(18).string(message.vestingInPoolsAmount);
    }
    if (message.vestingInAccountsAmount !== "") {
      writer.uint32(26).string(message.vestingInAccountsAmount);
    }
    if (message.delegatedVestingAmount !== "") {
      writer.uint32(34).string(message.delegatedVestingAmount);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVestingsSummaryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVestingsSummaryResponse,
    } as QueryVestingsSummaryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestingAllAmount = reader.string();
          break;
        case 2:
          message.vestingInPoolsAmount = reader.string();
          break;
        case 3:
          message.vestingInAccountsAmount = reader.string();
          break;
        case 4:
          message.delegatedVestingAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingsSummaryResponse {
    const message = {
      ...baseQueryVestingsSummaryResponse,
    } as QueryVestingsSummaryResponse;
    if (
      object.vestingAllAmount !== undefined &&
      object.vestingAllAmount !== null
    ) {
      message.vestingAllAmount = String(object.vestingAllAmount);
    } else {
      message.vestingAllAmount = "";
    }
    if (
      object.vestingInPoolsAmount !== undefined &&
      object.vestingInPoolsAmount !== null
    ) {
      message.vestingInPoolsAmount = String(object.vestingInPoolsAmount);
    } else {
      message.vestingInPoolsAmount = "";
    }
    if (
      object.vestingInAccountsAmount !== undefined &&
      object.vestingInAccountsAmount !== null
    ) {
      message.vestingInAccountsAmount = String(object.vestingInAccountsAmount);
    } else {
      message.vestingInAccountsAmount = "";
    }
    if (
      object.delegatedVestingAmount !== undefined &&
      object.delegatedVestingAmount !== null
    ) {
      message.delegatedVestingAmount = String(object.delegatedVestingAmount);
    } else {
      message.delegatedVestingAmount = "";
    }
    return message;
  },

  toJSON(message: QueryVestingsSummaryResponse): unknown {
    const obj: any = {};
    message.vestingAllAmount !== undefined &&
      (obj.vestingAllAmount = message.vestingAllAmount);
    message.vestingInPoolsAmount !== undefined &&
      (obj.vestingInPoolsAmount = message.vestingInPoolsAmount);
    message.vestingInAccountsAmount !== undefined &&
      (obj.vestingInAccountsAmount = message.vestingInAccountsAmount);
    message.delegatedVestingAmount !== undefined &&
      (obj.delegatedVestingAmount = message.delegatedVestingAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingsSummaryResponse>
  ): QueryVestingsSummaryResponse {
    const message = {
      ...baseQueryVestingsSummaryResponse,
    } as QueryVestingsSummaryResponse;
    if (
      object.vestingAllAmount !== undefined &&
      object.vestingAllAmount !== null
    ) {
      message.vestingAllAmount = object.vestingAllAmount;
    } else {
      message.vestingAllAmount = "";
    }
    if (
      object.vestingInPoolsAmount !== undefined &&
      object.vestingInPoolsAmount !== null
    ) {
      message.vestingInPoolsAmount = object.vestingInPoolsAmount;
    } else {
      message.vestingInPoolsAmount = "";
    }
    if (
      object.vestingInAccountsAmount !== undefined &&
      object.vestingInAccountsAmount !== null
    ) {
      message.vestingInAccountsAmount = object.vestingInAccountsAmount;
    } else {
      message.vestingInAccountsAmount = "";
    }
    if (
      object.delegatedVestingAmount !== undefined &&
      object.delegatedVestingAmount !== null
    ) {
      message.delegatedVestingAmount = object.delegatedVestingAmount;
    } else {
      message.delegatedVestingAmount = "";
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
  VestingPools(
    request: QueryVestingPoolsRequest
  ): Promise<QueryVestingPoolsResponse>;
  /** Queries a summary of the entire vesting. */
  VestingsSummary(
    request: QueryVestingsSummaryRequest
  ): Promise<QueryVestingsSummaryResponse>;
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

  VestingPools(
    request: QueryVestingPoolsRequest
  ): Promise<QueryVestingPoolsResponse> {
    const data = QueryVestingPoolsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Query",
      "VestingPools",
      data
    );
    return promise.then((data) =>
      QueryVestingPoolsResponse.decode(new Reader(data))
    );
  }

  VestingsSummary(
    request: QueryVestingsSummaryRequest
  ): Promise<QueryVestingsSummaryResponse> {
    const data = QueryVestingsSummaryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Query",
      "VestingsSummary",
      data
    );
    return promise.then((data) =>
      QueryVestingsSummaryResponse.decode(new Reader(data))
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}
