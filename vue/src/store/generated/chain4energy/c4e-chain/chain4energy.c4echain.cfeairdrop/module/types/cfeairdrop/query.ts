/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cfeairdrop/params";
import { ClaimRecordXX } from "../cfeairdrop/claim_record_xx";
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

export interface QueryGetClaimRecordXXRequest {
  index: string;
}

export interface QueryGetClaimRecordXXResponse {
  claimRecordXX: ClaimRecordXX | undefined;
}

export interface QueryAllClaimRecordXXRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClaimRecordXXResponse {
  claimRecordXX: ClaimRecordXX[];
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

const baseQueryGetClaimRecordXXRequest: object = { index: "" };

export const QueryGetClaimRecordXXRequest = {
  encode(
    message: QueryGetClaimRecordXXRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimRecordXXRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimRecordXXRequest,
    } as QueryGetClaimRecordXXRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClaimRecordXXRequest {
    const message = {
      ...baseQueryGetClaimRecordXXRequest,
    } as QueryGetClaimRecordXXRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetClaimRecordXXRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimRecordXXRequest>
  ): QueryGetClaimRecordXXRequest {
    const message = {
      ...baseQueryGetClaimRecordXXRequest,
    } as QueryGetClaimRecordXXRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetClaimRecordXXResponse: object = {};

export const QueryGetClaimRecordXXResponse = {
  encode(
    message: QueryGetClaimRecordXXResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.claimRecordXX !== undefined) {
      ClaimRecordXX.encode(
        message.claimRecordXX,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimRecordXXResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimRecordXXResponse,
    } as QueryGetClaimRecordXXResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimRecordXX = ClaimRecordXX.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClaimRecordXXResponse {
    const message = {
      ...baseQueryGetClaimRecordXXResponse,
    } as QueryGetClaimRecordXXResponse;
    if (object.claimRecordXX !== undefined && object.claimRecordXX !== null) {
      message.claimRecordXX = ClaimRecordXX.fromJSON(object.claimRecordXX);
    } else {
      message.claimRecordXX = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClaimRecordXXResponse): unknown {
    const obj: any = {};
    message.claimRecordXX !== undefined &&
      (obj.claimRecordXX = message.claimRecordXX
        ? ClaimRecordXX.toJSON(message.claimRecordXX)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimRecordXXResponse>
  ): QueryGetClaimRecordXXResponse {
    const message = {
      ...baseQueryGetClaimRecordXXResponse,
    } as QueryGetClaimRecordXXResponse;
    if (object.claimRecordXX !== undefined && object.claimRecordXX !== null) {
      message.claimRecordXX = ClaimRecordXX.fromPartial(object.claimRecordXX);
    } else {
      message.claimRecordXX = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimRecordXXRequest: object = {};

export const QueryAllClaimRecordXXRequest = {
  encode(
    message: QueryAllClaimRecordXXRequest,
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
  ): QueryAllClaimRecordXXRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimRecordXXRequest,
    } as QueryAllClaimRecordXXRequest;
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

  fromJSON(object: any): QueryAllClaimRecordXXRequest {
    const message = {
      ...baseQueryAllClaimRecordXXRequest,
    } as QueryAllClaimRecordXXRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimRecordXXRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimRecordXXRequest>
  ): QueryAllClaimRecordXXRequest {
    const message = {
      ...baseQueryAllClaimRecordXXRequest,
    } as QueryAllClaimRecordXXRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimRecordXXResponse: object = {};

export const QueryAllClaimRecordXXResponse = {
  encode(
    message: QueryAllClaimRecordXXResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.claimRecordXX) {
      ClaimRecordXX.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllClaimRecordXXResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimRecordXXResponse,
    } as QueryAllClaimRecordXXResponse;
    message.claimRecordXX = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimRecordXX.push(
            ClaimRecordXX.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllClaimRecordXXResponse {
    const message = {
      ...baseQueryAllClaimRecordXXResponse,
    } as QueryAllClaimRecordXXResponse;
    message.claimRecordXX = [];
    if (object.claimRecordXX !== undefined && object.claimRecordXX !== null) {
      for (const e of object.claimRecordXX) {
        message.claimRecordXX.push(ClaimRecordXX.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimRecordXXResponse): unknown {
    const obj: any = {};
    if (message.claimRecordXX) {
      obj.claimRecordXX = message.claimRecordXX.map((e) =>
        e ? ClaimRecordXX.toJSON(e) : undefined
      );
    } else {
      obj.claimRecordXX = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimRecordXXResponse>
  ): QueryAllClaimRecordXXResponse {
    const message = {
      ...baseQueryAllClaimRecordXXResponse,
    } as QueryAllClaimRecordXXResponse;
    message.claimRecordXX = [];
    if (object.claimRecordXX !== undefined && object.claimRecordXX !== null) {
      for (const e of object.claimRecordXX) {
        message.claimRecordXX.push(ClaimRecordXX.fromPartial(e));
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
  /** Queries a ClaimRecordXX by index. */
  ClaimRecordXX(
    request: QueryGetClaimRecordXXRequest
  ): Promise<QueryGetClaimRecordXXResponse>;
  /** Queries a list of ClaimRecordXX items. */
  ClaimRecordXXAll(
    request: QueryAllClaimRecordXXRequest
  ): Promise<QueryAllClaimRecordXXResponse>;
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

  ClaimRecordXX(
    request: QueryGetClaimRecordXXRequest
  ): Promise<QueryGetClaimRecordXXResponse> {
    const data = QueryGetClaimRecordXXRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "ClaimRecordXX",
      data
    );
    return promise.then((data) =>
      QueryGetClaimRecordXXResponse.decode(new Reader(data))
    );
  }

  ClaimRecordXXAll(
    request: QueryAllClaimRecordXXRequest
  ): Promise<QueryAllClaimRecordXXResponse> {
    const data = QueryAllClaimRecordXXRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Query",
      "ClaimRecordXXAll",
      data
    );
    return promise.then((data) =>
      QueryAllClaimRecordXXResponse.decode(new Reader(data))
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
