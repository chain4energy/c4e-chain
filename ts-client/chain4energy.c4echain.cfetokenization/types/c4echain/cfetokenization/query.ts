/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { CertificateType } from "./certificate_type";
import { Params } from "./params";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetCertificateTypeRequest {
  id: number;
}

export interface QueryGetCertificateTypeResponse {
  CertificateType: CertificateType | undefined;
}

export interface QueryAllCertificateTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCertificateTypeResponse {
  CertificateType: CertificateType[];
  pagination: PageResponse | undefined;
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

function createBaseQueryGetCertificateTypeRequest(): QueryGetCertificateTypeRequest {
  return { id: 0 };
}

export const QueryGetCertificateTypeRequest = {
  encode(message: QueryGetCertificateTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertificateTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertificateTypeRequest();
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

  fromJSON(object: any): QueryGetCertificateTypeRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetCertificateTypeRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertificateTypeRequest>, I>>(
    object: I,
  ): QueryGetCertificateTypeRequest {
    const message = createBaseQueryGetCertificateTypeRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetCertificateTypeResponse(): QueryGetCertificateTypeResponse {
  return { CertificateType: undefined };
}

export const QueryGetCertificateTypeResponse = {
  encode(message: QueryGetCertificateTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.CertificateType !== undefined) {
      CertificateType.encode(message.CertificateType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertificateTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertificateTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CertificateType = CertificateType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCertificateTypeResponse {
    return {
      CertificateType: isSet(object.CertificateType) ? CertificateType.fromJSON(object.CertificateType) : undefined,
    };
  },

  toJSON(message: QueryGetCertificateTypeResponse): unknown {
    const obj: any = {};
    message.CertificateType !== undefined
      && (obj.CertificateType = message.CertificateType ? CertificateType.toJSON(message.CertificateType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertificateTypeResponse>, I>>(
    object: I,
  ): QueryGetCertificateTypeResponse {
    const message = createBaseQueryGetCertificateTypeResponse();
    message.CertificateType = (object.CertificateType !== undefined && object.CertificateType !== null)
      ? CertificateType.fromPartial(object.CertificateType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertificateTypeRequest(): QueryAllCertificateTypeRequest {
  return { pagination: undefined };
}

export const QueryAllCertificateTypeRequest = {
  encode(message: QueryAllCertificateTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertificateTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertificateTypeRequest();
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

  fromJSON(object: any): QueryAllCertificateTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllCertificateTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertificateTypeRequest>, I>>(
    object: I,
  ): QueryAllCertificateTypeRequest {
    const message = createBaseQueryAllCertificateTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertificateTypeResponse(): QueryAllCertificateTypeResponse {
  return { CertificateType: [], pagination: undefined };
}

export const QueryAllCertificateTypeResponse = {
  encode(message: QueryAllCertificateTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.CertificateType) {
      CertificateType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertificateTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertificateTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CertificateType.push(CertificateType.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllCertificateTypeResponse {
    return {
      CertificateType: Array.isArray(object?.CertificateType)
        ? object.CertificateType.map((e: any) => CertificateType.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllCertificateTypeResponse): unknown {
    const obj: any = {};
    if (message.CertificateType) {
      obj.CertificateType = message.CertificateType.map((e) => e ? CertificateType.toJSON(e) : undefined);
    } else {
      obj.CertificateType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertificateTypeResponse>, I>>(
    object: I,
  ): QueryAllCertificateTypeResponse {
    const message = createBaseQueryAllCertificateTypeResponse();
    message.CertificateType = object.CertificateType?.map((e) => CertificateType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of CertificateType items. */
  CertificateType(request: QueryGetCertificateTypeRequest): Promise<QueryGetCertificateTypeResponse>;
  CertificateTypeAll(request: QueryAllCertificateTypeRequest): Promise<QueryAllCertificateTypeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.CertificateType = this.CertificateType.bind(this);
    this.CertificateTypeAll = this.CertificateTypeAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  CertificateType(request: QueryGetCertificateTypeRequest): Promise<QueryGetCertificateTypeResponse> {
    const data = QueryGetCertificateTypeRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "CertificateType", data);
    return promise.then((data) => QueryGetCertificateTypeResponse.decode(new _m0.Reader(data)));
  }

  CertificateTypeAll(request: QueryAllCertificateTypeRequest): Promise<QueryAllCertificateTypeResponse> {
    const data = QueryAllCertificateTypeRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "CertificateTypeAll", data);
    return promise.then((data) => QueryAllCertificateTypeResponse.decode(new _m0.Reader(data)));
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
