/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";
import { GenesisVestingType } from "./genesis";
import { Params } from "./params";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryVestingTypeRequest {
}

export interface QueryVestingTypeResponse {
  vestingTypes: GenesisVestingType[];
}

export interface QueryVestingPoolsRequest {
  owner: string;
}

export interface QueryVestingPoolsResponse {
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

export interface QueryVestingsSummaryRequest {
}

export interface QueryVestingsSummaryResponse {
  vestingAllAmount: string;
  vestingInPoolsAmount: string;
  vestingInAccountsAmount: string;
  delegatedVestingAmount: string;
}

/** this line is used by starport scaffolding # 3 */
export interface QueryGenesisVestingsSummaryRequest {
}

export interface QueryGenesisVestingsSummaryResponse {
  vestingAllAmount: string;
  vestingInPoolsAmount: string;
  vestingInAccountsAmount: string;
  delegatedVestingAmount: string;
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

function createBaseQueryVestingTypeRequest(): QueryVestingTypeRequest {
  return {};
}

export const QueryVestingTypeRequest = {
  encode(_: QueryVestingTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingTypeRequest();
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
    return {};
  },

  toJSON(_: QueryVestingTypeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingTypeRequest>, I>>(_: I): QueryVestingTypeRequest {
    const message = createBaseQueryVestingTypeRequest();
    return message;
  },
};

function createBaseQueryVestingTypeResponse(): QueryVestingTypeResponse {
  return { vestingTypes: [] };
}

export const QueryVestingTypeResponse = {
  encode(message: QueryVestingTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vestingTypes) {
      GenesisVestingType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.vestingTypes.push(GenesisVestingType.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingTypeResponse {
    return {
      vestingTypes: Array.isArray(object?.vestingTypes)
        ? object.vestingTypes.map((e: any) => GenesisVestingType.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryVestingTypeResponse): unknown {
    const obj: any = {};
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) => e ? GenesisVestingType.toJSON(e) : undefined);
    } else {
      obj.vestingTypes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingTypeResponse>, I>>(object: I): QueryVestingTypeResponse {
    const message = createBaseQueryVestingTypeResponse();
    message.vestingTypes = object.vestingTypes?.map((e) => GenesisVestingType.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryVestingPoolsRequest(): QueryVestingPoolsRequest {
  return { owner: "" };
}

export const QueryVestingPoolsRequest = {
  encode(message: QueryVestingPoolsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingPoolsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingPoolsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingPoolsRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryVestingPoolsRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingPoolsRequest>, I>>(object: I): QueryVestingPoolsRequest {
    const message = createBaseQueryVestingPoolsRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryVestingPoolsResponse(): QueryVestingPoolsResponse {
  return { vestingPools: [] };
}

export const QueryVestingPoolsResponse = {
  encode(message: QueryVestingPoolsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vestingPools) {
      VestingPoolInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingPoolsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingPoolsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.vestingPools.push(VestingPoolInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVestingPoolsResponse {
    return {
      vestingPools: Array.isArray(object?.vestingPools)
        ? object.vestingPools.map((e: any) => VestingPoolInfo.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryVestingPoolsResponse): unknown {
    const obj: any = {};
    if (message.vestingPools) {
      obj.vestingPools = message.vestingPools.map((e) => e ? VestingPoolInfo.toJSON(e) : undefined);
    } else {
      obj.vestingPools = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingPoolsResponse>, I>>(object: I): QueryVestingPoolsResponse {
    const message = createBaseQueryVestingPoolsResponse();
    message.vestingPools = object.vestingPools?.map((e) => VestingPoolInfo.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVestingPoolInfo(): VestingPoolInfo {
  return {
    name: "",
    vestingType: "",
    lockStart: undefined,
    lockEnd: undefined,
    withdrawable: "",
    initiallyLocked: undefined,
    currentlyLocked: "",
    sentAmount: "",
  };
}

export const VestingPoolInfo = {
  encode(message: VestingPoolInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.vestingType !== "") {
      writer.uint32(18).string(message.vestingType);
    }
    if (message.lockStart !== undefined) {
      Timestamp.encode(toTimestamp(message.lockStart), writer.uint32(26).fork()).ldelim();
    }
    if (message.lockEnd !== undefined) {
      Timestamp.encode(toTimestamp(message.lockEnd), writer.uint32(34).fork()).ldelim();
    }
    if (message.withdrawable !== "") {
      writer.uint32(42).string(message.withdrawable);
    }
    if (message.initiallyLocked !== undefined) {
      Coin.encode(message.initiallyLocked, writer.uint32(50).fork()).ldelim();
    }
    if (message.currentlyLocked !== "") {
      writer.uint32(58).string(message.currentlyLocked);
    }
    if (message.sentAmount !== "") {
      writer.uint32(66).string(message.sentAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingPoolInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingPoolInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.vestingType = reader.string();
          break;
        case 3:
          message.lockStart = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 4:
          message.lockEnd = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 5:
          message.withdrawable = reader.string();
          break;
        case 6:
          message.initiallyLocked = Coin.decode(reader, reader.uint32());
          break;
        case 7:
          message.currentlyLocked = reader.string();
          break;
        case 8:
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
      lockStart: isSet(object.lockStart) ? fromJsonTimestamp(object.lockStart) : undefined,
      lockEnd: isSet(object.lockEnd) ? fromJsonTimestamp(object.lockEnd) : undefined,
      withdrawable: isSet(object.withdrawable) ? String(object.withdrawable) : "",
      initiallyLocked: isSet(object.initiallyLocked) ? Coin.fromJSON(object.initiallyLocked) : undefined,
      currentlyLocked: isSet(object.currentlyLocked) ? String(object.currentlyLocked) : "",
      sentAmount: isSet(object.sentAmount) ? String(object.sentAmount) : "",
    };
  },

  toJSON(message: VestingPoolInfo): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    message.lockStart !== undefined && (obj.lockStart = message.lockStart.toISOString());
    message.lockEnd !== undefined && (obj.lockEnd = message.lockEnd.toISOString());
    message.withdrawable !== undefined && (obj.withdrawable = message.withdrawable);
    message.initiallyLocked !== undefined
      && (obj.initiallyLocked = message.initiallyLocked ? Coin.toJSON(message.initiallyLocked) : undefined);
    message.currentlyLocked !== undefined && (obj.currentlyLocked = message.currentlyLocked);
    message.sentAmount !== undefined && (obj.sentAmount = message.sentAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingPoolInfo>, I>>(object: I): VestingPoolInfo {
    const message = createBaseVestingPoolInfo();
    message.name = object.name ?? "";
    message.vestingType = object.vestingType ?? "";
    message.lockStart = object.lockStart ?? undefined;
    message.lockEnd = object.lockEnd ?? undefined;
    message.withdrawable = object.withdrawable ?? "";
    message.initiallyLocked = (object.initiallyLocked !== undefined && object.initiallyLocked !== null)
      ? Coin.fromPartial(object.initiallyLocked)
      : undefined;
    message.currentlyLocked = object.currentlyLocked ?? "";
    message.sentAmount = object.sentAmount ?? "";
    return message;
  },
};

function createBaseQueryVestingsSummaryRequest(): QueryVestingsSummaryRequest {
  return {};
}

export const QueryVestingsSummaryRequest = {
  encode(_: QueryVestingsSummaryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingsSummaryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingsSummaryRequest();
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
    return {};
  },

  toJSON(_: QueryVestingsSummaryRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingsSummaryRequest>, I>>(_: I): QueryVestingsSummaryRequest {
    const message = createBaseQueryVestingsSummaryRequest();
    return message;
  },
};

function createBaseQueryVestingsSummaryResponse(): QueryVestingsSummaryResponse {
  return { vestingAllAmount: "", vestingInPoolsAmount: "", vestingInAccountsAmount: "", delegatedVestingAmount: "" };
}

export const QueryVestingsSummaryResponse = {
  encode(message: QueryVestingsSummaryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVestingsSummaryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVestingsSummaryResponse();
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
    return {
      vestingAllAmount: isSet(object.vestingAllAmount) ? String(object.vestingAllAmount) : "",
      vestingInPoolsAmount: isSet(object.vestingInPoolsAmount) ? String(object.vestingInPoolsAmount) : "",
      vestingInAccountsAmount: isSet(object.vestingInAccountsAmount) ? String(object.vestingInAccountsAmount) : "",
      delegatedVestingAmount: isSet(object.delegatedVestingAmount) ? String(object.delegatedVestingAmount) : "",
    };
  },

  toJSON(message: QueryVestingsSummaryResponse): unknown {
    const obj: any = {};
    message.vestingAllAmount !== undefined && (obj.vestingAllAmount = message.vestingAllAmount);
    message.vestingInPoolsAmount !== undefined && (obj.vestingInPoolsAmount = message.vestingInPoolsAmount);
    message.vestingInAccountsAmount !== undefined && (obj.vestingInAccountsAmount = message.vestingInAccountsAmount);
    message.delegatedVestingAmount !== undefined && (obj.delegatedVestingAmount = message.delegatedVestingAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVestingsSummaryResponse>, I>>(object: I): QueryVestingsSummaryResponse {
    const message = createBaseQueryVestingsSummaryResponse();
    message.vestingAllAmount = object.vestingAllAmount ?? "";
    message.vestingInPoolsAmount = object.vestingInPoolsAmount ?? "";
    message.vestingInAccountsAmount = object.vestingInAccountsAmount ?? "";
    message.delegatedVestingAmount = object.delegatedVestingAmount ?? "";
    return message;
  },
};

function createBaseQueryGenesisVestingsSummaryRequest(): QueryGenesisVestingsSummaryRequest {
  return {};
}

export const QueryGenesisVestingsSummaryRequest = {
  encode(_: QueryGenesisVestingsSummaryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGenesisVestingsSummaryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGenesisVestingsSummaryRequest();
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

  fromJSON(_: any): QueryGenesisVestingsSummaryRequest {
    return {};
  },

  toJSON(_: QueryGenesisVestingsSummaryRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGenesisVestingsSummaryRequest>, I>>(
    _: I,
  ): QueryGenesisVestingsSummaryRequest {
    const message = createBaseQueryGenesisVestingsSummaryRequest();
    return message;
  },
};

function createBaseQueryGenesisVestingsSummaryResponse(): QueryGenesisVestingsSummaryResponse {
  return { vestingAllAmount: "", vestingInPoolsAmount: "", vestingInAccountsAmount: "", delegatedVestingAmount: "" };
}

export const QueryGenesisVestingsSummaryResponse = {
  encode(message: QueryGenesisVestingsSummaryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGenesisVestingsSummaryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGenesisVestingsSummaryResponse();
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

  fromJSON(object: any): QueryGenesisVestingsSummaryResponse {
    return {
      vestingAllAmount: isSet(object.vestingAllAmount) ? String(object.vestingAllAmount) : "",
      vestingInPoolsAmount: isSet(object.vestingInPoolsAmount) ? String(object.vestingInPoolsAmount) : "",
      vestingInAccountsAmount: isSet(object.vestingInAccountsAmount) ? String(object.vestingInAccountsAmount) : "",
      delegatedVestingAmount: isSet(object.delegatedVestingAmount) ? String(object.delegatedVestingAmount) : "",
    };
  },

  toJSON(message: QueryGenesisVestingsSummaryResponse): unknown {
    const obj: any = {};
    message.vestingAllAmount !== undefined && (obj.vestingAllAmount = message.vestingAllAmount);
    message.vestingInPoolsAmount !== undefined && (obj.vestingInPoolsAmount = message.vestingInPoolsAmount);
    message.vestingInAccountsAmount !== undefined && (obj.vestingInAccountsAmount = message.vestingInAccountsAmount);
    message.delegatedVestingAmount !== undefined && (obj.delegatedVestingAmount = message.delegatedVestingAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGenesisVestingsSummaryResponse>, I>>(
    object: I,
  ): QueryGenesisVestingsSummaryResponse {
    const message = createBaseQueryGenesisVestingsSummaryResponse();
    message.vestingAllAmount = object.vestingAllAmount ?? "";
    message.vestingInPoolsAmount = object.vestingInPoolsAmount ?? "";
    message.vestingInAccountsAmount = object.vestingInAccountsAmount ?? "";
    message.delegatedVestingAmount = object.delegatedVestingAmount ?? "";
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of VestingType items. */
  VestingType(request: QueryVestingTypeRequest): Promise<QueryVestingTypeResponse>;
  /** Queries a list of Vesting items. */
  VestingPools(request: QueryVestingPoolsRequest): Promise<QueryVestingPoolsResponse>;
  /** Queries a summary of the entire vesting. */
  VestingsSummary(request: QueryVestingsSummaryRequest): Promise<QueryVestingsSummaryResponse>;
  /** Queries a list of GenesisVestingsSummary items. */
  GenesisVestingsSummary(request: QueryGenesisVestingsSummaryRequest): Promise<QueryGenesisVestingsSummaryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.VestingType = this.VestingType.bind(this);
    this.VestingPools = this.VestingPools.bind(this);
    this.VestingsSummary = this.VestingsSummary.bind(this);
    this.GenesisVestingsSummary = this.GenesisVestingsSummary.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  VestingType(request: QueryVestingTypeRequest): Promise<QueryVestingTypeResponse> {
    const data = QueryVestingTypeRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "VestingType", data);
    return promise.then((data) => QueryVestingTypeResponse.decode(new _m0.Reader(data)));
  }

  VestingPools(request: QueryVestingPoolsRequest): Promise<QueryVestingPoolsResponse> {
    const data = QueryVestingPoolsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "VestingPools", data);
    return promise.then((data) => QueryVestingPoolsResponse.decode(new _m0.Reader(data)));
  }

  VestingsSummary(request: QueryVestingsSummaryRequest): Promise<QueryVestingsSummaryResponse> {
    const data = QueryVestingsSummaryRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "VestingsSummary", data);
    return promise.then((data) => QueryVestingsSummaryResponse.decode(new _m0.Reader(data)));
  }

  GenesisVestingsSummary(request: QueryGenesisVestingsSummaryRequest): Promise<QueryGenesisVestingsSummaryResponse> {
    const data = QueryGenesisVestingsSummaryRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "GenesisVestingsSummary", data);
    return promise.then((data) => QueryGenesisVestingsSummaryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
