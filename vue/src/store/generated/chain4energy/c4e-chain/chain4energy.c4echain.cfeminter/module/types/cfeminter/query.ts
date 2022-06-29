/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cfeminter/params";
import { MinterState } from "../cfeminter/minter";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryInflationRequest {}

export interface QueryInflationResponse {
  inflation: string;
}

export interface QueryStateRequest {}

export interface QueryStateResponse {
  minter_state: MinterState | undefined;
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

const baseQueryInflationRequest: object = {};

export const QueryInflationRequest = {
  encode(_: QueryInflationRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryInflationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryInflationRequest } as QueryInflationRequest;
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

  fromJSON(_: any): QueryInflationRequest {
    const message = { ...baseQueryInflationRequest } as QueryInflationRequest;
    return message;
  },

  toJSON(_: QueryInflationRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryInflationRequest>): QueryInflationRequest {
    const message = { ...baseQueryInflationRequest } as QueryInflationRequest;
    return message;
  },
};

const baseQueryInflationResponse: object = { inflation: "" };

export const QueryInflationResponse = {
  encode(
    message: QueryInflationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.inflation !== "") {
      writer.uint32(10).string(message.inflation);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryInflationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryInflationResponse } as QueryInflationResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.inflation = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryInflationResponse {
    const message = { ...baseQueryInflationResponse } as QueryInflationResponse;
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = String(object.inflation);
    } else {
      message.inflation = "";
    }
    return message;
  },

  toJSON(message: QueryInflationResponse): unknown {
    const obj: any = {};
    message.inflation !== undefined && (obj.inflation = message.inflation);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryInflationResponse>
  ): QueryInflationResponse {
    const message = { ...baseQueryInflationResponse } as QueryInflationResponse;
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = object.inflation;
    } else {
      message.inflation = "";
    }
    return message;
  },
};

const baseQueryStateRequest: object = {};

export const QueryStateRequest = {
  encode(_: QueryStateRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryStateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryStateRequest } as QueryStateRequest;
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

  fromJSON(_: any): QueryStateRequest {
    const message = { ...baseQueryStateRequest } as QueryStateRequest;
    return message;
  },

  toJSON(_: QueryStateRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryStateRequest>): QueryStateRequest {
    const message = { ...baseQueryStateRequest } as QueryStateRequest;
    return message;
  },
};

const baseQueryStateResponse: object = {};

export const QueryStateResponse = {
  encode(
    message: QueryStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.minter_state !== undefined) {
      MinterState.encode(
        message.minter_state,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryStateResponse } as QueryStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.minter_state = MinterState.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryStateResponse {
    const message = { ...baseQueryStateResponse } as QueryStateResponse;
    if (object.minter_state !== undefined && object.minter_state !== null) {
      message.minter_state = MinterState.fromJSON(object.minter_state);
    } else {
      message.minter_state = undefined;
    }
    return message;
  },

  toJSON(message: QueryStateResponse): unknown {
    const obj: any = {};
    message.minter_state !== undefined &&
      (obj.minter_state = message.minter_state
        ? MinterState.toJSON(message.minter_state)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryStateResponse>): QueryStateResponse {
    const message = { ...baseQueryStateResponse } as QueryStateResponse;
    if (object.minter_state !== undefined && object.minter_state !== null) {
      message.minter_state = MinterState.fromPartial(object.minter_state);
    } else {
      message.minter_state = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Inflation items. */
  Inflation(request: QueryInflationRequest): Promise<QueryInflationResponse>;
  /** Queries a list of State items. */
  State(request: QueryStateRequest): Promise<QueryStateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeminter.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Inflation(request: QueryInflationRequest): Promise<QueryInflationResponse> {
    const data = QueryInflationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeminter.Query",
      "Inflation",
      data
    );
    return promise.then((data) =>
      QueryInflationResponse.decode(new Reader(data))
    );
  }

  State(request: QueryStateRequest): Promise<QueryStateResponse> {
    const data = QueryStateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeminter.Query",
      "State",
      data
    );
    return promise.then((data) => QueryStateResponse.decode(new Reader(data)));
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
