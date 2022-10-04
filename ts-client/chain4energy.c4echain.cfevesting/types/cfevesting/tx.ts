/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Duration } from "../google/protobuf/duration";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface MsgCreateVestingPool {
  creator: string;
  name: string;
  amount: string;
  duration: Duration | undefined;
  vestingType: string;
}

export interface MsgCreateVestingPoolResponse {}

export interface MsgWithdrawAllAvailable {
  creator: string;
}

export interface MsgWithdrawAllAvailableResponse {}

export interface MsgCreateVestingAccount {
  fromAddress: string;
  toAddress: string;
  amount: Coin[];
  startTime: number;
  endTime: number;
}

export interface MsgCreateVestingAccountResponse {}

export interface MsgSendToVestingAccount {
  fromAddress: string;
  toAddress: string;
  vestingId: number;
  amount: string;
  restartVesting: boolean;
}

export interface MsgSendToVestingAccountResponse {}

const baseMsgCreateVestingPool: object = {
  creator: "",
  name: "",
  amount: "",
  vestingType: "",
};

export const MsgCreateVestingPool = {
  encode(
    message: MsgCreateVestingPool,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVestingPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateVestingPool } as MsgCreateVestingPool;
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
    const message = { ...baseMsgCreateVestingPool } as MsgCreateVestingPool;
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
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Duration.fromJSON(object.duration);
    } else {
      message.duration = undefined;
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    return message;
  },

  toJSON(message: MsgCreateVestingPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined &&
      (obj.duration = message.duration
        ? Duration.toJSON(message.duration)
        : undefined);
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateVestingPool>): MsgCreateVestingPool {
    const message = { ...baseMsgCreateVestingPool } as MsgCreateVestingPool;
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
      message.amount = "";
    }
    if (object.duration !== undefined && object.duration !== null) {
      message.duration = Duration.fromPartial(object.duration);
    } else {
      message.duration = undefined;
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    return message;
  },
};

const baseMsgCreateVestingPoolResponse: object = {};

export const MsgCreateVestingPoolResponse = {
  encode(
    _: MsgCreateVestingPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateVestingPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateVestingPoolResponse,
    } as MsgCreateVestingPoolResponse;
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
    const message = {
      ...baseMsgCreateVestingPoolResponse,
    } as MsgCreateVestingPoolResponse;
    return message;
  },

  toJSON(_: MsgCreateVestingPoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateVestingPoolResponse>
  ): MsgCreateVestingPoolResponse {
    const message = {
      ...baseMsgCreateVestingPoolResponse,
    } as MsgCreateVestingPoolResponse;
    return message;
  },
};

const baseMsgWithdrawAllAvailable: object = { creator: "" };

export const MsgWithdrawAllAvailable = {
  encode(
    message: MsgWithdrawAllAvailable,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawAllAvailable {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
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
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawAllAvailable): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgWithdrawAllAvailable>
  ): MsgWithdrawAllAvailable {
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgWithdrawAllAvailableResponse: object = {};

export const MsgWithdrawAllAvailableResponse = {
  encode(
    _: MsgWithdrawAllAvailableResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawAllAvailableResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
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

  fromJSON(_: any): MsgWithdrawAllAvailableResponse {
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    return message;
  },

  toJSON(_: MsgWithdrawAllAvailableResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawAllAvailableResponse>
  ): MsgWithdrawAllAvailableResponse {
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    return message;
  },
};

const baseMsgCreateVestingAccount: object = {
  fromAddress: "",
  toAddress: "",
  startTime: 0,
  endTime: 0,
};

export const MsgCreateVestingAccount = {
  encode(
    message: MsgCreateVestingAccount,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVestingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateVestingAccount,
    } as MsgCreateVestingAccount;
    message.amount = [];
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
    const message = {
      ...baseMsgCreateVestingAccount,
    } as MsgCreateVestingAccount;
    message.amount = [];
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = String(object.fromAddress);
    } else {
      message.fromAddress = "";
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = String(object.toAddress);
    } else {
      message.toAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    if (object.startTime !== undefined && object.startTime !== null) {
      message.startTime = Number(object.startTime);
    } else {
      message.startTime = 0;
    }
    if (object.endTime !== undefined && object.endTime !== null) {
      message.endTime = Number(object.endTime);
    } else {
      message.endTime = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateVestingAccount): unknown {
    const obj: any = {};
    message.fromAddress !== undefined &&
      (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    message.startTime !== undefined && (obj.startTime = message.startTime);
    message.endTime !== undefined && (obj.endTime = message.endTime);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVestingAccount>
  ): MsgCreateVestingAccount {
    const message = {
      ...baseMsgCreateVestingAccount,
    } as MsgCreateVestingAccount;
    message.amount = [];
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = object.fromAddress;
    } else {
      message.fromAddress = "";
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = object.toAddress;
    } else {
      message.toAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    if (object.startTime !== undefined && object.startTime !== null) {
      message.startTime = object.startTime;
    } else {
      message.startTime = 0;
    }
    if (object.endTime !== undefined && object.endTime !== null) {
      message.endTime = object.endTime;
    } else {
      message.endTime = 0;
    }
    return message;
  },
};

const baseMsgCreateVestingAccountResponse: object = {};

export const MsgCreateVestingAccountResponse = {
  encode(
    _: MsgCreateVestingAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateVestingAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateVestingAccountResponse,
    } as MsgCreateVestingAccountResponse;
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
    const message = {
      ...baseMsgCreateVestingAccountResponse,
    } as MsgCreateVestingAccountResponse;
    return message;
  },

  toJSON(_: MsgCreateVestingAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateVestingAccountResponse>
  ): MsgCreateVestingAccountResponse {
    const message = {
      ...baseMsgCreateVestingAccountResponse,
    } as MsgCreateVestingAccountResponse;
    return message;
  },
};

const baseMsgSendToVestingAccount: object = {
  fromAddress: "",
  toAddress: "",
  vestingId: 0,
  amount: "",
  restartVesting: false,
};

export const MsgSendToVestingAccount = {
  encode(
    message: MsgSendToVestingAccount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    if (message.vestingId !== 0) {
      writer.uint32(24).int32(message.vestingId);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restartVesting === true) {
      writer.uint32(40).bool(message.restartVesting);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendToVestingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendToVestingAccount,
    } as MsgSendToVestingAccount;
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
          message.vestingId = reader.int32();
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
    const message = {
      ...baseMsgSendToVestingAccount,
    } as MsgSendToVestingAccount;
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = String(object.fromAddress);
    } else {
      message.fromAddress = "";
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = String(object.toAddress);
    } else {
      message.toAddress = "";
    }
    if (object.vestingId !== undefined && object.vestingId !== null) {
      message.vestingId = Number(object.vestingId);
    } else {
      message.vestingId = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.restartVesting !== undefined && object.restartVesting !== null) {
      message.restartVesting = Boolean(object.restartVesting);
    } else {
      message.restartVesting = false;
    }
    return message;
  },

  toJSON(message: MsgSendToVestingAccount): unknown {
    const obj: any = {};
    message.fromAddress !== undefined &&
      (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    message.vestingId !== undefined && (obj.vestingId = message.vestingId);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restartVesting !== undefined &&
      (obj.restartVesting = message.restartVesting);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendToVestingAccount>
  ): MsgSendToVestingAccount {
    const message = {
      ...baseMsgSendToVestingAccount,
    } as MsgSendToVestingAccount;
    if (object.fromAddress !== undefined && object.fromAddress !== null) {
      message.fromAddress = object.fromAddress;
    } else {
      message.fromAddress = "";
    }
    if (object.toAddress !== undefined && object.toAddress !== null) {
      message.toAddress = object.toAddress;
    } else {
      message.toAddress = "";
    }
    if (object.vestingId !== undefined && object.vestingId !== null) {
      message.vestingId = object.vestingId;
    } else {
      message.vestingId = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.restartVesting !== undefined && object.restartVesting !== null) {
      message.restartVesting = object.restartVesting;
    } else {
      message.restartVesting = false;
    }
    return message;
  },
};

const baseMsgSendToVestingAccountResponse: object = {};

export const MsgSendToVestingAccountResponse = {
  encode(
    _: MsgSendToVestingAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendToVestingAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendToVestingAccountResponse,
    } as MsgSendToVestingAccountResponse;
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
    const message = {
      ...baseMsgSendToVestingAccountResponse,
    } as MsgSendToVestingAccountResponse;
    return message;
  },

  toJSON(_: MsgSendToVestingAccountResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendToVestingAccountResponse>
  ): MsgSendToVestingAccountResponse {
    const message = {
      ...baseMsgSendToVestingAccountResponse,
    } as MsgSendToVestingAccountResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVestingPool(
    request: MsgCreateVestingPool
  ): Promise<MsgCreateVestingPoolResponse>;
  WithdrawAllAvailable(
    request: MsgWithdrawAllAvailable
  ): Promise<MsgWithdrawAllAvailableResponse>;
  CreateVestingAccount(
    request: MsgCreateVestingAccount
  ): Promise<MsgCreateVestingAccountResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendToVestingAccount(
    request: MsgSendToVestingAccount
  ): Promise<MsgSendToVestingAccountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateVestingPool(
    request: MsgCreateVestingPool
  ): Promise<MsgCreateVestingPoolResponse> {
    const data = MsgCreateVestingPool.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "CreateVestingPool",
      data
    );
    return promise.then((data) =>
      MsgCreateVestingPoolResponse.decode(new Reader(data))
    );
  }

  WithdrawAllAvailable(
    request: MsgWithdrawAllAvailable
  ): Promise<MsgWithdrawAllAvailableResponse> {
    const data = MsgWithdrawAllAvailable.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "WithdrawAllAvailable",
      data
    );
    return promise.then((data) =>
      MsgWithdrawAllAvailableResponse.decode(new Reader(data))
    );
  }

  CreateVestingAccount(
    request: MsgCreateVestingAccount
  ): Promise<MsgCreateVestingAccountResponse> {
    const data = MsgCreateVestingAccount.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "CreateVestingAccount",
      data
    );
    return promise.then((data) =>
      MsgCreateVestingAccountResponse.decode(new Reader(data))
    );
  }

  SendToVestingAccount(
    request: MsgSendToVestingAccount
  ): Promise<MsgSendToVestingAccountResponse> {
    const data = MsgSendToVestingAccount.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "SendToVestingAccount",
      data
    );
    return promise.then((data) =>
      MsgSendToVestingAccountResponse.decode(new Reader(data))
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
