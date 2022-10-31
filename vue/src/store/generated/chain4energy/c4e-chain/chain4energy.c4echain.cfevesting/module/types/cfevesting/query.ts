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
  vesting_types: GenesisVestingType[];
}

export interface QueryVestingPoolsRequest {
  address: string;
}

export interface QueryVestingPoolsResponse {
  delegable_address: string;
  vesting_pools: VestingPoolInfo[];
}

export interface VestingPoolInfo {
  name: string;
  vesting_type: string;
  lock_start: Date | undefined;
  lock_end: Date | undefined;
  withdrawable: string;
  initially_locked: Coin | undefined;
  currently_locked: string;
  sent_amount: string;
}

export interface QueryVestingsSummaryRequest {}

export interface QueryVestingsSummaryResponse {
  vesting_all_amount: string;
  vesting_in_pools_amount: string;
  vesting_in_accounts_amount: string;
  delegated_vesting_amount: string;
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
    for (const v of message.vesting_types) {
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
    message.vesting_types = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.vesting_types.push(
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
    message.vesting_types = [];
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(GenesisVestingType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryVestingTypeResponse): unknown {
    const obj: any = {};
    if (message.vesting_types) {
      obj.vesting_types = message.vesting_types.map((e) =>
        e ? GenesisVestingType.toJSON(e) : undefined
      );
    } else {
      obj.vesting_types = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingTypeResponse>
  ): QueryVestingTypeResponse {
    const message = {
      ...baseQueryVestingTypeResponse,
    } as QueryVestingTypeResponse;
    message.vesting_types = [];
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(GenesisVestingType.fromPartial(e));
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

const baseQueryVestingPoolsResponse: object = { delegable_address: "" };

export const QueryVestingPoolsResponse = {
  encode(
    message: QueryVestingPoolsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegable_address !== "") {
      writer.uint32(10).string(message.delegable_address);
    }
    for (const v of message.vesting_pools) {
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
    message.vesting_pools = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegable_address = reader.string();
          break;
        case 2:
          message.vesting_pools.push(
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
    message.vesting_pools = [];
    if (
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = String(object.delegable_address);
    } else {
      message.delegable_address = "";
    }
    if (object.vesting_pools !== undefined && object.vesting_pools !== null) {
      for (const e of object.vesting_pools) {
        message.vesting_pools.push(VestingPoolInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryVestingPoolsResponse): unknown {
    const obj: any = {};
    message.delegable_address !== undefined &&
      (obj.delegable_address = message.delegable_address);
    if (message.vesting_pools) {
      obj.vesting_pools = message.vesting_pools.map((e) =>
        e ? VestingPoolInfo.toJSON(e) : undefined
      );
    } else {
      obj.vesting_pools = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingPoolsResponse>
  ): QueryVestingPoolsResponse {
    const message = {
      ...baseQueryVestingPoolsResponse,
    } as QueryVestingPoolsResponse;
    message.vesting_pools = [];
    if (
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = object.delegable_address;
    } else {
      message.delegable_address = "";
    }
    if (object.vesting_pools !== undefined && object.vesting_pools !== null) {
      for (const e of object.vesting_pools) {
        message.vesting_pools.push(VestingPoolInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingPoolInfo: object = {
  name: "",
  vesting_type: "",
  withdrawable: "",
  currently_locked: "",
  sent_amount: "",
};

export const VestingPoolInfo = {
  encode(message: VestingPoolInfo, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.vesting_type !== "") {
      writer.uint32(26).string(message.vesting_type);
    }
    if (message.lock_start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_start),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.lock_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_end),
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.withdrawable !== "") {
      writer.uint32(50).string(message.withdrawable);
    }
    if (message.initially_locked !== undefined) {
      Coin.encode(message.initially_locked, writer.uint32(58).fork()).ldelim();
    }
    if (message.currently_locked !== "") {
      writer.uint32(66).string(message.currently_locked);
    }
    if (message.sent_amount !== "") {
      writer.uint32(74).string(message.sent_amount);
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
          message.vesting_type = reader.string();
          break;
        case 4:
          message.lock_start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.lock_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.withdrawable = reader.string();
          break;
        case 7:
          message.initially_locked = Coin.decode(reader, reader.uint32());
          break;
        case 8:
          message.currently_locked = reader.string();
          break;
        case 9:
          message.sent_amount = reader.string();
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
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = String(object.vesting_type);
    } else {
      message.vesting_type = "";
    }
    if (object.lock_start !== undefined && object.lock_start !== null) {
      message.lock_start = fromJsonTimestamp(object.lock_start);
    } else {
      message.lock_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = fromJsonTimestamp(object.lock_end);
    } else {
      message.lock_end = undefined;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = String(object.withdrawable);
    } else {
      message.withdrawable = "";
    }
    if (
      object.initially_locked !== undefined &&
      object.initially_locked !== null
    ) {
      message.initially_locked = Coin.fromJSON(object.initially_locked);
    } else {
      message.initially_locked = undefined;
    }
    if (
      object.currently_locked !== undefined &&
      object.currently_locked !== null
    ) {
      message.currently_locked = String(object.currently_locked);
    } else {
      message.currently_locked = "";
    }
    if (object.sent_amount !== undefined && object.sent_amount !== null) {
      message.sent_amount = String(object.sent_amount);
    } else {
      message.sent_amount = "";
    }
    return message;
  },

  toJSON(message: VestingPoolInfo): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.vesting_type !== undefined &&
      (obj.vesting_type = message.vesting_type);
    message.lock_start !== undefined &&
      (obj.lock_start =
        message.lock_start !== undefined
          ? message.lock_start.toISOString()
          : null);
    message.lock_end !== undefined &&
      (obj.lock_end =
        message.lock_end !== undefined ? message.lock_end.toISOString() : null);
    message.withdrawable !== undefined &&
      (obj.withdrawable = message.withdrawable);
    message.initially_locked !== undefined &&
      (obj.initially_locked = message.initially_locked
        ? Coin.toJSON(message.initially_locked)
        : undefined);
    message.currently_locked !== undefined &&
      (obj.currently_locked = message.currently_locked);
    message.sent_amount !== undefined &&
      (obj.sent_amount = message.sent_amount);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingPoolInfo>): VestingPoolInfo {
    const message = { ...baseVestingPoolInfo } as VestingPoolInfo;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = object.vesting_type;
    } else {
      message.vesting_type = "";
    }
    if (object.lock_start !== undefined && object.lock_start !== null) {
      message.lock_start = object.lock_start;
    } else {
      message.lock_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = object.lock_end;
    } else {
      message.lock_end = undefined;
    }
    if (object.withdrawable !== undefined && object.withdrawable !== null) {
      message.withdrawable = object.withdrawable;
    } else {
      message.withdrawable = "";
    }
    if (
      object.initially_locked !== undefined &&
      object.initially_locked !== null
    ) {
      message.initially_locked = Coin.fromPartial(object.initially_locked);
    } else {
      message.initially_locked = undefined;
    }
    if (
      object.currently_locked !== undefined &&
      object.currently_locked !== null
    ) {
      message.currently_locked = object.currently_locked;
    } else {
      message.currently_locked = "";
    }
    if (object.sent_amount !== undefined && object.sent_amount !== null) {
      message.sent_amount = object.sent_amount;
    } else {
      message.sent_amount = "";
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
  vesting_all_amount: "",
  vesting_in_pools_amount: "",
  vesting_in_accounts_amount: "",
  delegated_vesting_amount: "",
};

export const QueryVestingsSummaryResponse = {
  encode(
    message: QueryVestingsSummaryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.vesting_all_amount !== "") {
      writer.uint32(10).string(message.vesting_all_amount);
    }
    if (message.vesting_in_pools_amount !== "") {
      writer.uint32(18).string(message.vesting_in_pools_amount);
    }
    if (message.vesting_in_accounts_amount !== "") {
      writer.uint32(26).string(message.vesting_in_accounts_amount);
    }
    if (message.delegated_vesting_amount !== "") {
      writer.uint32(34).string(message.delegated_vesting_amount);
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
          message.vesting_all_amount = reader.string();
          break;
        case 2:
          message.vesting_in_pools_amount = reader.string();
          break;
        case 3:
          message.vesting_in_accounts_amount = reader.string();
          break;
        case 4:
          message.delegated_vesting_amount = reader.string();
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
      object.vesting_all_amount !== undefined &&
      object.vesting_all_amount !== null
    ) {
      message.vesting_all_amount = String(object.vesting_all_amount);
    } else {
      message.vesting_all_amount = "";
    }
    if (
      object.vesting_in_pools_amount !== undefined &&
      object.vesting_in_pools_amount !== null
    ) {
      message.vesting_in_pools_amount = String(object.vesting_in_pools_amount);
    } else {
      message.vesting_in_pools_amount = "";
    }
    if (
      object.vesting_in_accounts_amount !== undefined &&
      object.vesting_in_accounts_amount !== null
    ) {
      message.vesting_in_accounts_amount = String(
        object.vesting_in_accounts_amount
      );
    } else {
      message.vesting_in_accounts_amount = "";
    }
    if (
      object.delegated_vesting_amount !== undefined &&
      object.delegated_vesting_amount !== null
    ) {
      message.delegated_vesting_amount = String(
        object.delegated_vesting_amount
      );
    } else {
      message.delegated_vesting_amount = "";
    }
    return message;
  },

  toJSON(message: QueryVestingsSummaryResponse): unknown {
    const obj: any = {};
    message.vesting_all_amount !== undefined &&
      (obj.vesting_all_amount = message.vesting_all_amount);
    message.vesting_in_pools_amount !== undefined &&
      (obj.vesting_in_pools_amount = message.vesting_in_pools_amount);
    message.vesting_in_accounts_amount !== undefined &&
      (obj.vesting_in_accounts_amount = message.vesting_in_accounts_amount);
    message.delegated_vesting_amount !== undefined &&
      (obj.delegated_vesting_amount = message.delegated_vesting_amount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVestingsSummaryResponse>
  ): QueryVestingsSummaryResponse {
    const message = {
      ...baseQueryVestingsSummaryResponse,
    } as QueryVestingsSummaryResponse;
    if (
      object.vesting_all_amount !== undefined &&
      object.vesting_all_amount !== null
    ) {
      message.vesting_all_amount = object.vesting_all_amount;
    } else {
      message.vesting_all_amount = "";
    }
    if (
      object.vesting_in_pools_amount !== undefined &&
      object.vesting_in_pools_amount !== null
    ) {
      message.vesting_in_pools_amount = object.vesting_in_pools_amount;
    } else {
      message.vesting_in_pools_amount = "";
    }
    if (
      object.vesting_in_accounts_amount !== undefined &&
      object.vesting_in_accounts_amount !== null
    ) {
      message.vesting_in_accounts_amount = object.vesting_in_accounts_amount;
    } else {
      message.vesting_in_accounts_amount = "";
    }
    if (
      object.delegated_vesting_amount !== undefined &&
      object.delegated_vesting_amount !== null
    ) {
      message.delegated_vesting_amount = object.delegated_vesting_amount;
    } else {
      message.delegated_vesting_amount = "";
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
