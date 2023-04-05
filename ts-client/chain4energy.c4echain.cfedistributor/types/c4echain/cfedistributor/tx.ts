/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { SubDistributor } from "./sub_distributor";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface MsgUpdateParams {
  /** authority is the address of the governance account. */
  authority: string;
  subDistributors: SubDistributor[];
}

export interface MsgUpdateParamsResponse {
}

export interface MsgUpdateSubDistributorParam {
  /** authority is the address of the governance account. */
  authority: string;
  subDistributor: SubDistributor | undefined;
}

export interface MsgUpdateSubDistributorParamResponse {
}

export interface MsgUpdateSubDistributorDestinationShareParam {
  authority: string;
  subDistributorName: string;
  destinationName: string;
  share: string;
}

export interface MsgUpdateSubDistributorDestinationShareParamResponse {
}

export interface MsgUpdateSubDistributorBurnShareParam {
  authority: string;
  subDistributorName: string;
  burnShare: string;
}

export interface MsgUpdateSubDistributorBurnShareParamResponse {
}

function createBaseMsgUpdateParams(): MsgUpdateParams {
  return { authority: "", subDistributors: [] };
}

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    for (const v of message.subDistributors) {
      SubDistributor.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.subDistributors.push(SubDistributor.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateParams {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      subDistributors: Array.isArray(object?.subDistributors)
        ? object.subDistributors.map((e: any) => SubDistributor.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    if (message.subDistributors) {
      obj.subDistributors = message.subDistributors.map((e) => e ? SubDistributor.toJSON(e) : undefined);
    } else {
      obj.subDistributors = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(object: I): MsgUpdateParams {
    const message = createBaseMsgUpdateParams();
    message.authority = object.authority ?? "";
    message.subDistributors = object.subDistributors?.map((e) => SubDistributor.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgUpdateParamsResponse(): MsgUpdateParamsResponse {
  return {};
}

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParamsResponse();
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

  fromJSON(_: any): MsgUpdateParamsResponse {
    return {};
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(_: I): MsgUpdateParamsResponse {
    const message = createBaseMsgUpdateParamsResponse();
    return message;
  },
};

function createBaseMsgUpdateSubDistributorParam(): MsgUpdateSubDistributorParam {
  return { authority: "", subDistributor: undefined };
}

export const MsgUpdateSubDistributorParam = {
  encode(message: MsgUpdateSubDistributorParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.subDistributor !== undefined) {
      SubDistributor.encode(message.subDistributor, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorParam();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.subDistributor = SubDistributor.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSubDistributorParam {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      subDistributor: isSet(object.subDistributor) ? SubDistributor.fromJSON(object.subDistributor) : undefined,
    };
  },

  toJSON(message: MsgUpdateSubDistributorParam): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    message.subDistributor !== undefined
      && (obj.subDistributor = message.subDistributor ? SubDistributor.toJSON(message.subDistributor) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorParam>, I>>(object: I): MsgUpdateSubDistributorParam {
    const message = createBaseMsgUpdateSubDistributorParam();
    message.authority = object.authority ?? "";
    message.subDistributor = (object.subDistributor !== undefined && object.subDistributor !== null)
      ? SubDistributor.fromPartial(object.subDistributor)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateSubDistributorParamResponse(): MsgUpdateSubDistributorParamResponse {
  return {};
}

export const MsgUpdateSubDistributorParamResponse = {
  encode(_: MsgUpdateSubDistributorParamResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorParamResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorParamResponse();
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

  fromJSON(_: any): MsgUpdateSubDistributorParamResponse {
    return {};
  },

  toJSON(_: MsgUpdateSubDistributorParamResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorParamResponse>, I>>(
    _: I,
  ): MsgUpdateSubDistributorParamResponse {
    const message = createBaseMsgUpdateSubDistributorParamResponse();
    return message;
  },
};

function createBaseMsgUpdateSubDistributorDestinationShareParam(): MsgUpdateSubDistributorDestinationShareParam {
  return { authority: "", subDistributorName: "", destinationName: "", share: "" };
}

export const MsgUpdateSubDistributorDestinationShareParam = {
  encode(message: MsgUpdateSubDistributorDestinationShareParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.subDistributorName !== "") {
      writer.uint32(18).string(message.subDistributorName);
    }
    if (message.destinationName !== "") {
      writer.uint32(26).string(message.destinationName);
    }
    if (message.share !== "") {
      writer.uint32(34).string(message.share);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorDestinationShareParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorDestinationShareParam();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.subDistributorName = reader.string();
          break;
        case 3:
          message.destinationName = reader.string();
          break;
        case 4:
          message.share = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSubDistributorDestinationShareParam {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      subDistributorName: isSet(object.subDistributorName) ? String(object.subDistributorName) : "",
      destinationName: isSet(object.destinationName) ? String(object.destinationName) : "",
      share: isSet(object.share) ? String(object.share) : "",
    };
  },

  toJSON(message: MsgUpdateSubDistributorDestinationShareParam): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    message.subDistributorName !== undefined && (obj.subDistributorName = message.subDistributorName);
    message.destinationName !== undefined && (obj.destinationName = message.destinationName);
    message.share !== undefined && (obj.share = message.share);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorDestinationShareParam>, I>>(
    object: I,
  ): MsgUpdateSubDistributorDestinationShareParam {
    const message = createBaseMsgUpdateSubDistributorDestinationShareParam();
    message.authority = object.authority ?? "";
    message.subDistributorName = object.subDistributorName ?? "";
    message.destinationName = object.destinationName ?? "";
    message.share = object.share ?? "";
    return message;
  },
};

function createBaseMsgUpdateSubDistributorDestinationShareParamResponse(): MsgUpdateSubDistributorDestinationShareParamResponse {
  return {};
}

export const MsgUpdateSubDistributorDestinationShareParamResponse = {
  encode(
    _: MsgUpdateSubDistributorDestinationShareParamResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorDestinationShareParamResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorDestinationShareParamResponse();
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

  fromJSON(_: any): MsgUpdateSubDistributorDestinationShareParamResponse {
    return {};
  },

  toJSON(_: MsgUpdateSubDistributorDestinationShareParamResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorDestinationShareParamResponse>, I>>(
    _: I,
  ): MsgUpdateSubDistributorDestinationShareParamResponse {
    const message = createBaseMsgUpdateSubDistributorDestinationShareParamResponse();
    return message;
  },
};

function createBaseMsgUpdateSubDistributorBurnShareParam(): MsgUpdateSubDistributorBurnShareParam {
  return { authority: "", subDistributorName: "", burnShare: "" };
}

export const MsgUpdateSubDistributorBurnShareParam = {
  encode(message: MsgUpdateSubDistributorBurnShareParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.subDistributorName !== "") {
      writer.uint32(18).string(message.subDistributorName);
    }
    if (message.burnShare !== "") {
      writer.uint32(26).string(message.burnShare);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorBurnShareParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorBurnShareParam();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.subDistributorName = reader.string();
          break;
        case 3:
          message.burnShare = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSubDistributorBurnShareParam {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      subDistributorName: isSet(object.subDistributorName) ? String(object.subDistributorName) : "",
      burnShare: isSet(object.burnShare) ? String(object.burnShare) : "",
    };
  },

  toJSON(message: MsgUpdateSubDistributorBurnShareParam): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    message.subDistributorName !== undefined && (obj.subDistributorName = message.subDistributorName);
    message.burnShare !== undefined && (obj.burnShare = message.burnShare);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorBurnShareParam>, I>>(
    object: I,
  ): MsgUpdateSubDistributorBurnShareParam {
    const message = createBaseMsgUpdateSubDistributorBurnShareParam();
    message.authority = object.authority ?? "";
    message.subDistributorName = object.subDistributorName ?? "";
    message.burnShare = object.burnShare ?? "";
    return message;
  },
};

function createBaseMsgUpdateSubDistributorBurnShareParamResponse(): MsgUpdateSubDistributorBurnShareParamResponse {
  return {};
}

export const MsgUpdateSubDistributorBurnShareParamResponse = {
  encode(_: MsgUpdateSubDistributorBurnShareParamResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateSubDistributorBurnShareParamResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateSubDistributorBurnShareParamResponse();
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

  fromJSON(_: any): MsgUpdateSubDistributorBurnShareParamResponse {
    return {};
  },

  toJSON(_: MsgUpdateSubDistributorBurnShareParamResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateSubDistributorBurnShareParamResponse>, I>>(
    _: I,
  ): MsgUpdateSubDistributorBurnShareParamResponse {
    const message = createBaseMsgUpdateSubDistributorBurnShareParamResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
  UpdateSubDistributorParam(request: MsgUpdateSubDistributorParam): Promise<MsgUpdateSubDistributorParamResponse>;
  UpdateSubDistributorDestinationShareParam(
    request: MsgUpdateSubDistributorDestinationShareParam,
  ): Promise<MsgUpdateSubDistributorDestinationShareParamResponse>;
  UpdateSubDistributorBurnShareParam(
    request: MsgUpdateSubDistributorBurnShareParam,
  ): Promise<MsgUpdateSubDistributorBurnShareParamResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.UpdateParams = this.UpdateParams.bind(this);
    this.UpdateSubDistributorParam = this.UpdateSubDistributorParam.bind(this);
    this.UpdateSubDistributorDestinationShareParam = this.UpdateSubDistributorDestinationShareParam.bind(this);
    this.UpdateSubDistributorBurnShareParam = this.UpdateSubDistributorBurnShareParam.bind(this);
  }
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfedistributor.Msg", "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(new _m0.Reader(data)));
  }

  UpdateSubDistributorParam(request: MsgUpdateSubDistributorParam): Promise<MsgUpdateSubDistributorParamResponse> {
    const data = MsgUpdateSubDistributorParam.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfedistributor.Msg", "UpdateSubDistributorParam", data);
    return promise.then((data) => MsgUpdateSubDistributorParamResponse.decode(new _m0.Reader(data)));
  }

  UpdateSubDistributorDestinationShareParam(
    request: MsgUpdateSubDistributorDestinationShareParam,
  ): Promise<MsgUpdateSubDistributorDestinationShareParamResponse> {
    const data = MsgUpdateSubDistributorDestinationShareParam.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfedistributor.Msg",
      "UpdateSubDistributorDestinationShareParam",
      data,
    );
    return promise.then((data) => MsgUpdateSubDistributorDestinationShareParamResponse.decode(new _m0.Reader(data)));
  }

  UpdateSubDistributorBurnShareParam(
    request: MsgUpdateSubDistributorBurnShareParam,
  ): Promise<MsgUpdateSubDistributorBurnShareParamResponse> {
    const data = MsgUpdateSubDistributorBurnShareParam.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfedistributor.Msg",
      "UpdateSubDistributorBurnShareParam",
      data,
    );
    return promise.then((data) => MsgUpdateSubDistributorBurnShareParamResponse.decode(new _m0.Reader(data)));
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
