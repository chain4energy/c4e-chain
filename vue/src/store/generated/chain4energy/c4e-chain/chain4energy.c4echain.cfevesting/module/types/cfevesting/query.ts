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

export interface QueryVestingRequest {
  address: string;
}

export interface QueryVestingResponse {
  delegable_address: string;
  vestings: VestingInfo[];
}

export interface VestingInfo {
  id: number;
  vesting_type: string;
  lock_start: Date | undefined;
  lock_end: Date | undefined;
  /** google.protobuf.Timestamp vesting_end = 5 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true]; */
  withdrawable: string;
  /** bool delegation_allowed = 7; */
  vested: Coin | undefined;
  current_vested_amount: string;
  sent_amount: string;
  transfer_allowed: boolean;
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

const baseQueryVestingResponse: object = { delegable_address: "" };

export const QueryVestingResponse = {
  encode(
    message: QueryVestingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegable_address !== "") {
      writer.uint32(10).string(message.delegable_address);
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
          message.delegable_address = reader.string();
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
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = String(object.delegable_address);
    } else {
      message.delegable_address = "";
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
    message.delegable_address !== undefined &&
      (obj.delegable_address = message.delegable_address);
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
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = object.delegable_address;
    } else {
      message.delegable_address = "";
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
  id: 0,
  vesting_type: "",
  withdrawable: "",
  current_vested_amount: "",
  sent_amount: "",
  transfer_allowed: false,
};

export const VestingInfo = {
  encode(message: VestingInfo, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.vesting_type !== "") {
      writer.uint32(18).string(message.vesting_type);
    }
    if (message.lock_start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_start),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lock_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_end),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.withdrawable !== "") {
      writer.uint32(50).string(message.withdrawable);
    }
    if (message.vested !== undefined) {
      Coin.encode(message.vested, writer.uint32(66).fork()).ldelim();
    }
    if (message.current_vested_amount !== "") {
      writer.uint32(74).string(message.current_vested_amount);
    }
    if (message.sent_amount !== "") {
      writer.uint32(82).string(message.sent_amount);
    }
    if (message.transfer_allowed === true) {
      writer.uint32(88).bool(message.transfer_allowed);
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
          message.id = reader.int32();
          break;
        case 2:
          message.vesting_type = reader.string();
          break;
        case 3:
          message.lock_start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.lock_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.withdrawable = reader.string();
          break;
        case 8:
          message.vested = Coin.decode(reader, reader.uint32());
          break;
        case 9:
          message.current_vested_amount = reader.string();
          break;
        case 10:
          message.sent_amount = reader.string();
          break;
        case 11:
          message.transfer_allowed = reader.bool();
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
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
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = Coin.fromJSON(object.vested);
    } else {
      message.vested = undefined;
    }
    if (
      object.current_vested_amount !== undefined &&
      object.current_vested_amount !== null
    ) {
      message.current_vested_amount = String(object.current_vested_amount);
    } else {
      message.current_vested_amount = "";
    }
    if (object.sent_amount !== undefined && object.sent_amount !== null) {
      message.sent_amount = String(object.sent_amount);
    } else {
      message.sent_amount = "";
    }
    if (
      object.transfer_allowed !== undefined &&
      object.transfer_allowed !== null
    ) {
      message.transfer_allowed = Boolean(object.transfer_allowed);
    } else {
      message.transfer_allowed = false;
    }
    return message;
  },

  toJSON(message: VestingInfo): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
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
    message.vested !== undefined &&
      (obj.vested = message.vested ? Coin.toJSON(message.vested) : undefined);
    message.current_vested_amount !== undefined &&
      (obj.current_vested_amount = message.current_vested_amount);
    message.sent_amount !== undefined &&
      (obj.sent_amount = message.sent_amount);
    message.transfer_allowed !== undefined &&
      (obj.transfer_allowed = message.transfer_allowed);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingInfo>): VestingInfo {
    const message = { ...baseVestingInfo } as VestingInfo;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
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
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = Coin.fromPartial(object.vested);
    } else {
      message.vested = undefined;
    }
    if (
      object.current_vested_amount !== undefined &&
      object.current_vested_amount !== null
    ) {
      message.current_vested_amount = object.current_vested_amount;
    } else {
      message.current_vested_amount = "";
    }
    if (object.sent_amount !== undefined && object.sent_amount !== null) {
      message.sent_amount = object.sent_amount;
    } else {
      message.sent_amount = "";
    }
    if (
      object.transfer_allowed !== undefined &&
      object.transfer_allowed !== null
    ) {
      message.transfer_allowed = object.transfer_allowed;
    } else {
      message.transfer_allowed = false;
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
