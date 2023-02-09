/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Duration } from "../../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface MsgCreateVestingPool {
  /** TODO: rename to owner */
  creator: string;
  name: string;
  amount: string;
  duration: Duration | undefined;
  vestingType: string;
}

export interface MsgCreateVestingPoolResponse {
}

export interface MsgWithdrawAllAvailable {
  /** TODO: rename to owner */
  creator: string;
}

export interface MsgWithdrawAllAvailableResponse {
  withdrawn: string;
}

export interface MsgCreateVestingAccount {
  fromAddress: string;
  toAddress: string;
  amount: Coin[];
  startTime: number;
  endTime: number;
}

export interface MsgCreateVestingAccountResponse {
}

export interface MsgSendToVestingAccount {
  /** TODO: rename to owner */
  fromAddress: string;
  toAddress: string;
  vestingPoolName: string;
  amount: string;
  restartVesting: boolean;
}

export interface MsgSendToVestingAccountResponse {
}

function createBaseMsgCreateVestingPool(): MsgCreateVestingPool {
  return { creator: "", name: "", amount: "", duration: undefined, vestingType: "" };
}

export const MsgCreateVestingPool = {
  encode(message: MsgCreateVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.duration !== undefined) {
      Duration.encode(message.duration, writer.uint32(42).fork()).ldelim();
    }
    if (message.vestingType !== "") {
      writer.uint32(50).string(message.vestingType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVestingPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.duration = Duration.decode(reader, reader.uint32());
          break;
        case 6:
          message.vestingType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVestingPool {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      duration: isSet(object.duration) ? Duration.fromJSON(object.duration) : undefined,
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
    };
  },

  toJSON(message: MsgCreateVestingPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined && (obj.duration = message.duration ? Duration.toJSON(message.duration) : undefined);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVestingPool>, I>>(object: I): MsgCreateVestingPool {
    const message = createBaseMsgCreateVestingPool();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.amount = object.amount ?? "";
    message.duration = (object.duration !== undefined && object.duration !== null)
      ? Duration.fromPartial(object.duration)
      : undefined;
    message.vestingType = object.vestingType ?? "";
    return message;
  },
};

function createBaseMsgCreateVestingPoolResponse(): MsgCreateVestingPoolResponse {
  return {};
}

export const MsgCreateVestingPoolResponse = {
  encode(_: MsgCreateVestingPoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVestingPoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVestingPoolResponse();
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

  fromJSON(_: any): MsgCreateVestingPoolResponse {
    return {};
  },

  toJSON(_: MsgCreateVestingPoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVestingPoolResponse>, I>>(_: I): MsgCreateVestingPoolResponse {
    const message = createBaseMsgCreateVestingPoolResponse();
    return message;
  },
};

function createBaseMsgWithdrawAllAvailable(): MsgWithdrawAllAvailable {
  return { creator: "" };
}

export const MsgWithdrawAllAvailable = {
  encode(message: MsgWithdrawAllAvailable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawAllAvailable {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawAllAvailable();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawAllAvailable {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgWithdrawAllAvailable): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawAllAvailable>, I>>(object: I): MsgWithdrawAllAvailable {
    const message = createBaseMsgWithdrawAllAvailable();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgWithdrawAllAvailableResponse(): MsgWithdrawAllAvailableResponse {
  return { withdrawn: "" };
}

export const MsgWithdrawAllAvailableResponse = {
  encode(message: MsgWithdrawAllAvailableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.withdrawn !== "") {
      writer.uint32(10).string(message.withdrawn);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawAllAvailableResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawAllAvailableResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.withdrawn = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawAllAvailableResponse {
    return { withdrawn: isSet(object.withdrawn) ? String(object.withdrawn) : "" };
  },

  toJSON(message: MsgWithdrawAllAvailableResponse): unknown {
    const obj: any = {};
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawAllAvailableResponse>, I>>(
    object: I,
  ): MsgWithdrawAllAvailableResponse {
    const message = createBaseMsgWithdrawAllAvailableResponse();
    message.withdrawn = object.withdrawn ?? "";
    return message;
  },
};

function createBaseMsgCreateVestingAccount(): MsgCreateVestingAccount {
  return { fromAddress: "", toAddress: "", amount: [], startTime: 0, endTime: 0 };
}

export const MsgCreateVestingAccount = {
  encode(message: MsgCreateVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.startTime !== 0) {
      writer.uint32(32).int64(message.startTime);
    }
    if (message.endTime !== 0) {
      writer.uint32(40).int64(message.endTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVestingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVestingAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fromAddress = reader.string();
          break;
        case 2:
          message.toAddress = reader.string();
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.startTime = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.endTime = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVestingAccount {
    return {
      fromAddress: isSet(object.fromAddress) ? String(object.fromAddress) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
      startTime: isSet(object.startTime) ? Number(object.startTime) : 0,
      endTime: isSet(object.endTime) ? Number(object.endTime) : 0,
    };
  },

  toJSON(message: MsgCreateVestingAccount): unknown {
    const obj: any = {};
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    message.startTime !== undefined && (obj.startTime = Math.round(message.startTime));
    message.endTime !== undefined && (obj.endTime = Math.round(message.endTime));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVestingAccount>, I>>(object: I): MsgCreateVestingAccount {
    const message = createBaseMsgCreateVestingAccount();
    message.fromAddress = object.fromAddress ?? "";
    message.toAddress = object.toAddress ?? "";
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    message.startTime = object.startTime ?? 0;
    message.endTime = object.endTime ?? 0;
    return message;
  },
};

function createBaseMsgCreateVestingAccountResponse(): MsgCreateVestingAccountResponse {
  return {};
}

export const MsgCreateVestingAccountResponse = {
  encode(_: MsgCreateVestingAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVestingAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVestingAccountResponse();
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

  fromJSON(_: any): MsgCreateVestingAccountResponse {
    return {};
  },

  toJSON(_: MsgCreateVestingAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVestingAccountResponse>, I>>(_: I): MsgCreateVestingAccountResponse {
    const message = createBaseMsgCreateVestingAccountResponse();
    return message;
  },
};

function createBaseMsgSendToVestingAccount(): MsgSendToVestingAccount {
  return { fromAddress: "", toAddress: "", vestingPoolName: "", amount: "", restartVesting: false };
}

export const MsgSendToVestingAccount = {
  encode(message: MsgSendToVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(26).string(message.vestingPoolName);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restartVesting === true) {
      writer.uint32(40).bool(message.restartVesting);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendToVestingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendToVestingAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fromAddress = reader.string();
          break;
        case 2:
          message.toAddress = reader.string();
          break;
        case 3:
          message.vestingPoolName = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.restartVesting = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendToVestingAccount {
    return {
      fromAddress: isSet(object.fromAddress) ? String(object.fromAddress) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      restartVesting: isSet(object.restartVesting) ? Boolean(object.restartVesting) : false,
    };
  },

  toJSON(message: MsgSendToVestingAccount): unknown {
    const obj: any = {};
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restartVesting !== undefined && (obj.restartVesting = message.restartVesting);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendToVestingAccount>, I>>(object: I): MsgSendToVestingAccount {
    const message = createBaseMsgSendToVestingAccount();
    message.fromAddress = object.fromAddress ?? "";
    message.toAddress = object.toAddress ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.amount = object.amount ?? "";
    message.restartVesting = object.restartVesting ?? false;
    return message;
  },
};

function createBaseMsgSendToVestingAccountResponse(): MsgSendToVestingAccountResponse {
  return {};
}

export const MsgSendToVestingAccountResponse = {
  encode(_: MsgSendToVestingAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendToVestingAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendToVestingAccountResponse();
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

  fromJSON(_: any): MsgSendToVestingAccountResponse {
    return {};
  },

  toJSON(_: MsgSendToVestingAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendToVestingAccountResponse>, I>>(_: I): MsgSendToVestingAccountResponse {
    const message = createBaseMsgSendToVestingAccountResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVestingPool(request: MsgCreateVestingPool): Promise<MsgCreateVestingPoolResponse>;
  WithdrawAllAvailable(request: MsgWithdrawAllAvailable): Promise<MsgWithdrawAllAvailableResponse>;
  CreateVestingAccount(request: MsgCreateVestingAccount): Promise<MsgCreateVestingAccountResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendToVestingAccount(request: MsgSendToVestingAccount): Promise<MsgSendToVestingAccountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateVestingPool = this.CreateVestingPool.bind(this);
    this.WithdrawAllAvailable = this.WithdrawAllAvailable.bind(this);
    this.CreateVestingAccount = this.CreateVestingAccount.bind(this);
    this.SendToVestingAccount = this.SendToVestingAccount.bind(this);
  }
  CreateVestingPool(request: MsgCreateVestingPool): Promise<MsgCreateVestingPoolResponse> {
    const data = MsgCreateVestingPool.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "CreateVestingPool", data);
    return promise.then((data) => MsgCreateVestingPoolResponse.decode(new _m0.Reader(data)));
  }

  WithdrawAllAvailable(request: MsgWithdrawAllAvailable): Promise<MsgWithdrawAllAvailableResponse> {
    const data = MsgWithdrawAllAvailable.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "WithdrawAllAvailable", data);
    return promise.then((data) => MsgWithdrawAllAvailableResponse.decode(new _m0.Reader(data)));
  }

  CreateVestingAccount(request: MsgCreateVestingAccount): Promise<MsgCreateVestingAccountResponse> {
    const data = MsgCreateVestingAccount.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "CreateVestingAccount", data);
    return promise.then((data) => MsgCreateVestingAccountResponse.decode(new _m0.Reader(data)));
  }

  SendToVestingAccount(request: MsgSendToVestingAccount): Promise<MsgSendToVestingAccountResponse> {
    const data = MsgSendToVestingAccount.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "SendToVestingAccount", data);
    return promise.then((data) => MsgSendToVestingAccountResponse.decode(new _m0.Reader(data)));
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
