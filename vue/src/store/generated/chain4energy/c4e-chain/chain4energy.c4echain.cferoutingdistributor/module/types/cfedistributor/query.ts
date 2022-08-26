/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cfedistributor/params";
import { State } from "../cfedistributor/sub_distributor";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryStatesRequest {}

export interface QueryStatesResponse {
  states: State[];
  coins_on_distributor_account: Coin[];
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

const baseQueryStatesRequest: object = {};

export const QueryStatesRequest = {
  encode(_: QueryStatesRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryStatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryStatesRequest } as QueryStatesRequest;
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

  fromJSON(_: any): QueryStatesRequest {
    const message = { ...baseQueryStatesRequest } as QueryStatesRequest;
    return message;
  },

  toJSON(_: QueryStatesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryStatesRequest>): QueryStatesRequest {
    const message = { ...baseQueryStatesRequest } as QueryStatesRequest;
    return message;
  },
};

const baseQueryStatesResponse: object = {};

export const QueryStatesResponse = {
  encode(
    message: QueryStatesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.states) {
      State.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.coins_on_distributor_account) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryStatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryStatesResponse } as QueryStatesResponse;
    message.states = [];
    message.coins_on_distributor_account = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.states.push(State.decode(reader, reader.uint32()));
          break;
        case 2:
          message.coins_on_distributor_account.push(
            Coin.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryStatesResponse {
    const message = { ...baseQueryStatesResponse } as QueryStatesResponse;
    message.states = [];
    message.coins_on_distributor_account = [];
    if (object.states !== undefined && object.states !== null) {
      for (const e of object.states) {
        message.states.push(State.fromJSON(e));
      }
    }
    if (
      object.coins_on_distributor_account !== undefined &&
      object.coins_on_distributor_account !== null
    ) {
      for (const e of object.coins_on_distributor_account) {
        message.coins_on_distributor_account.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryStatesResponse): unknown {
    const obj: any = {};
    if (message.states) {
      obj.states = message.states.map((e) => (e ? State.toJSON(e) : undefined));
    } else {
      obj.states = [];
    }
    if (message.coins_on_distributor_account) {
      obj.coins_on_distributor_account = message.coins_on_distributor_account.map(
        (e) => (e ? Coin.toJSON(e) : undefined)
      );
    } else {
      obj.coins_on_distributor_account = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QueryStatesResponse>): QueryStatesResponse {
    const message = { ...baseQueryStatesResponse } as QueryStatesResponse;
    message.states = [];
    message.coins_on_distributor_account = [];
    if (object.states !== undefined && object.states !== null) {
      for (const e of object.states) {
        message.states.push(State.fromPartial(e));
      }
    }
    if (
      object.coins_on_distributor_account !== undefined &&
      object.coins_on_distributor_account !== null
    ) {
      for (const e of object.coins_on_distributor_account) {
        message.coins_on_distributor_account.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of States items. */
  States(request: QueryStatesRequest): Promise<QueryStatesResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cferoutingdistributor.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  States(request: QueryStatesRequest): Promise<QueryStatesResponse> {
    const data = QueryStatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cferoutingdistributor.Query",
      "States",
      data
    );
    return promise.then((data) => QueryStatesResponse.decode(new Reader(data)));
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
