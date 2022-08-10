/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "chain4energy.c4echain.cfeenergybank";

export interface MsgCreateTokenParams {
  creator: string;
  name: string;
  tradingCompany: string;
  burningTime: number;
  burningType: string;
  sendPrice: number;
}

export interface MsgCreateTokenParamsResponse {}

export interface MsgMintToken {
  creator: string;
  name: string;
  amount: number;
  userAddress: string;
}

export interface MsgMintTokenResponse {}

const baseMsgCreateTokenParams: object = {
  creator: "",
  name: "",
  tradingCompany: "",
  burningTime: 0,
  burningType: "",
  sendPrice: 0,
};

export const MsgCreateTokenParams = {
  encode(
    message: MsgCreateTokenParams,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.tradingCompany !== "") {
      writer.uint32(26).string(message.tradingCompany);
    }
    if (message.burningTime !== 0) {
      writer.uint32(32).uint64(message.burningTime);
    }
    if (message.burningType !== "") {
      writer.uint32(42).string(message.burningType);
    }
    if (message.sendPrice !== 0) {
      writer.uint32(48).uint64(message.sendPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateTokenParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateTokenParams } as MsgCreateTokenParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.tradingCompany = reader.string();
          break;
        case 4:
          message.burningTime = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.burningType = reader.string();
          break;
        case 6:
          message.sendPrice = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTokenParams {
    const message = { ...baseMsgCreateTokenParams } as MsgCreateTokenParams;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.tradingCompany !== undefined && object.tradingCompany !== null) {
      message.tradingCompany = String(object.tradingCompany);
    } else {
      message.tradingCompany = "";
    }
    if (object.burningTime !== undefined && object.burningTime !== null) {
      message.burningTime = Number(object.burningTime);
    } else {
      message.burningTime = 0;
    }
    if (object.burningType !== undefined && object.burningType !== null) {
      message.burningType = String(object.burningType);
    } else {
      message.burningType = "";
    }
    if (object.sendPrice !== undefined && object.sendPrice !== null) {
      message.sendPrice = Number(object.sendPrice);
    } else {
      message.sendPrice = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTokenParams): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.tradingCompany !== undefined &&
      (obj.tradingCompany = message.tradingCompany);
    message.burningTime !== undefined &&
      (obj.burningTime = message.burningTime);
    message.burningType !== undefined &&
      (obj.burningType = message.burningType);
    message.sendPrice !== undefined && (obj.sendPrice = message.sendPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateTokenParams>): MsgCreateTokenParams {
    const message = { ...baseMsgCreateTokenParams } as MsgCreateTokenParams;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.tradingCompany !== undefined && object.tradingCompany !== null) {
      message.tradingCompany = object.tradingCompany;
    } else {
      message.tradingCompany = "";
    }
    if (object.burningTime !== undefined && object.burningTime !== null) {
      message.burningTime = object.burningTime;
    } else {
      message.burningTime = 0;
    }
    if (object.burningType !== undefined && object.burningType !== null) {
      message.burningType = object.burningType;
    } else {
      message.burningType = "";
    }
    if (object.sendPrice !== undefined && object.sendPrice !== null) {
      message.sendPrice = object.sendPrice;
    } else {
      message.sendPrice = 0;
    }
    return message;
  },
};

const baseMsgCreateTokenParamsResponse: object = {};

export const MsgCreateTokenParamsResponse = {
  encode(
    _: MsgCreateTokenParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateTokenParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTokenParamsResponse,
    } as MsgCreateTokenParamsResponse;
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

  fromJSON(_: any): MsgCreateTokenParamsResponse {
    const message = {
      ...baseMsgCreateTokenParamsResponse,
    } as MsgCreateTokenParamsResponse;
    return message;
  },

  toJSON(_: MsgCreateTokenParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateTokenParamsResponse>
  ): MsgCreateTokenParamsResponse {
    const message = {
      ...baseMsgCreateTokenParamsResponse,
    } as MsgCreateTokenParamsResponse;
    return message;
  },
};

const baseMsgMintToken: object = {
  creator: "",
  name: "",
  amount: 0,
  userAddress: "",
};

export const MsgMintToken = {
  encode(message: MsgMintToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.amount !== 0) {
      writer.uint32(24).uint64(message.amount);
    }
    if (message.userAddress !== "") {
      writer.uint32(34).string(message.userAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintToken } as MsgMintToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.userAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintToken {
    const message = { ...baseMsgMintToken } as MsgMintToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    return message;
  },

  toJSON(message: MsgMintToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintToken>): MsgMintToken {
    const message = { ...baseMsgMintToken } as MsgMintToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    return message;
  },
};

const baseMsgMintTokenResponse: object = {};

export const MsgMintTokenResponse = {
  encode(_: MsgMintTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
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

  fromJSON(_: any): MsgMintTokenResponse {
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
    return message;
  },

  toJSON(_: MsgMintTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMintTokenResponse>): MsgMintTokenResponse {
    const message = { ...baseMsgMintTokenResponse } as MsgMintTokenResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateTokenParams(
    request: MsgCreateTokenParams
  ): Promise<MsgCreateTokenParamsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateTokenParams(
    request: MsgCreateTokenParams
  ): Promise<MsgCreateTokenParamsResponse> {
    const data = MsgCreateTokenParams.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Msg",
      "CreateTokenParams",
      data
    );
    return promise.then((data) =>
      MsgCreateTokenParamsResponse.decode(new Reader(data))
    );
  }

  MintToken(request: MsgMintToken): Promise<MsgMintTokenResponse> {
    const data = MsgMintToken.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Msg",
      "MintToken",
      data
    );
    return promise.then((data) =>
      MsgMintTokenResponse.decode(new Reader(data))
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
