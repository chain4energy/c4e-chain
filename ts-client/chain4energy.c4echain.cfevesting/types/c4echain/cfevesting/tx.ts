/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Duration } from "../../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface MsgCreateVestingPool {
  owner: string;
  name: string;
  amount: string;
  duration: Duration | undefined;
  vestingType: string;
}

export interface MsgCreateVestingPoolResponse {
}

export interface MsgWithdrawAllAvailable {
  owner: string;
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
  owner: string;
  toAddress: string;
  vestingPoolName: string;
  amount: string;
  restartVesting: boolean;
}

export interface MsgSendToVestingAccountResponse {
}

export interface MsgSplitVesting {
  fromAddress: string;
  toAddress: string;
  amount: Coin[];
}

export interface MsgSplitVestingResponse {
}

export interface MsgMoveAvailableVesting {
  fromAddress: string;
  toAddress: string;
}

export interface MsgMoveAvailableVestingResponse {
}

export interface MsgMoveAvailableVestingByDenoms {
  fromAddress: string;
  toAddress: string;
  denoms: string[];
}

export interface MsgMoveAvailableVestingByDenomsResponse {
}

export interface MsgUpdateDenomParam {
  /** authority is the address of the governance account. */
  authority: string;
  denom: string;
}

export interface MsgUpdateDenomParamResponse {
}

function createBaseMsgCreateVestingPool(): MsgCreateVestingPool {
  return { owner: "", name: "", amount: "", duration: undefined, vestingType: "" };
}

export const MsgCreateVestingPool = {
  encode(message: MsgCreateVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
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
          message.owner = reader.string();
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
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      duration: isSet(object.duration) ? Duration.fromJSON(object.duration) : undefined,
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
    };
  },

  toJSON(message: MsgCreateVestingPool): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined && (obj.duration = message.duration ? Duration.toJSON(message.duration) : undefined);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVestingPool>, I>>(object: I): MsgCreateVestingPool {
    const message = createBaseMsgCreateVestingPool();
    message.owner = object.owner ?? "";
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
  return { owner: "" };
}

export const MsgWithdrawAllAvailable = {
  encode(message: MsgWithdrawAllAvailable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
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
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawAllAvailable {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: MsgWithdrawAllAvailable): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgWithdrawAllAvailable>, I>>(object: I): MsgWithdrawAllAvailable {
    const message = createBaseMsgWithdrawAllAvailable();
    message.owner = object.owner ?? "";
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
  return { owner: "", toAddress: "", vestingPoolName: "", amount: "", restartVesting: false };
}

export const MsgSendToVestingAccount = {
  encode(message: MsgSendToVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
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
          message.owner = reader.string();
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
      owner: isSet(object.owner) ? String(object.owner) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      restartVesting: isSet(object.restartVesting) ? Boolean(object.restartVesting) : false,
    };
  },

  toJSON(message: MsgSendToVestingAccount): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restartVesting !== undefined && (obj.restartVesting = message.restartVesting);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendToVestingAccount>, I>>(object: I): MsgSendToVestingAccount {
    const message = createBaseMsgSendToVestingAccount();
    message.owner = object.owner ?? "";
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

function createBaseMsgSplitVesting(): MsgSplitVesting {
  return { fromAddress: "", toAddress: "", amount: [] };
}

export const MsgSplitVesting = {
  encode(message: MsgSplitVesting, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSplitVesting {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSplitVesting();
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSplitVesting {
    return {
      fromAddress: isSet(object.fromAddress) ? String(object.fromAddress) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgSplitVesting): unknown {
    const obj: any = {};
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSplitVesting>, I>>(object: I): MsgSplitVesting {
    const message = createBaseMsgSplitVesting();
    message.fromAddress = object.fromAddress ?? "";
    message.toAddress = object.toAddress ?? "";
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgSplitVestingResponse(): MsgSplitVestingResponse {
  return {};
}

export const MsgSplitVestingResponse = {
  encode(_: MsgSplitVestingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSplitVestingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSplitVestingResponse();
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

  fromJSON(_: any): MsgSplitVestingResponse {
    return {};
  },

  toJSON(_: MsgSplitVestingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSplitVestingResponse>, I>>(_: I): MsgSplitVestingResponse {
    const message = createBaseMsgSplitVestingResponse();
    return message;
  },
};

function createBaseMsgMoveAvailableVesting(): MsgMoveAvailableVesting {
  return { fromAddress: "", toAddress: "" };
}

export const MsgMoveAvailableVesting = {
  encode(message: MsgMoveAvailableVesting, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMoveAvailableVesting {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMoveAvailableVesting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.fromAddress = reader.string();
          break;
        case 2:
          message.toAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMoveAvailableVesting {
    return {
      fromAddress: isSet(object.fromAddress) ? String(object.fromAddress) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
    };
  },

  toJSON(message: MsgMoveAvailableVesting): unknown {
    const obj: any = {};
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMoveAvailableVesting>, I>>(object: I): MsgMoveAvailableVesting {
    const message = createBaseMsgMoveAvailableVesting();
    message.fromAddress = object.fromAddress ?? "";
    message.toAddress = object.toAddress ?? "";
    return message;
  },
};

function createBaseMsgMoveAvailableVestingResponse(): MsgMoveAvailableVestingResponse {
  return {};
}

export const MsgMoveAvailableVestingResponse = {
  encode(_: MsgMoveAvailableVestingResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMoveAvailableVestingResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMoveAvailableVestingResponse();
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

  fromJSON(_: any): MsgMoveAvailableVestingResponse {
    return {};
  },

  toJSON(_: MsgMoveAvailableVestingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMoveAvailableVestingResponse>, I>>(_: I): MsgMoveAvailableVestingResponse {
    const message = createBaseMsgMoveAvailableVestingResponse();
    return message;
  },
};

function createBaseMsgMoveAvailableVestingByDenoms(): MsgMoveAvailableVestingByDenoms {
  return { fromAddress: "", toAddress: "", denoms: [] };
}

export const MsgMoveAvailableVestingByDenoms = {
  encode(message: MsgMoveAvailableVestingByDenoms, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    for (const v of message.denoms) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMoveAvailableVestingByDenoms {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMoveAvailableVestingByDenoms();
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
          message.denoms.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMoveAvailableVestingByDenoms {
    return {
      fromAddress: isSet(object.fromAddress) ? String(object.fromAddress) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      denoms: Array.isArray(object?.denoms) ? object.denoms.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgMoveAvailableVestingByDenoms): unknown {
    const obj: any = {};
    message.fromAddress !== undefined && (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    if (message.denoms) {
      obj.denoms = message.denoms.map((e) => e);
    } else {
      obj.denoms = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMoveAvailableVestingByDenoms>, I>>(
    object: I,
  ): MsgMoveAvailableVestingByDenoms {
    const message = createBaseMsgMoveAvailableVestingByDenoms();
    message.fromAddress = object.fromAddress ?? "";
    message.toAddress = object.toAddress ?? "";
    message.denoms = object.denoms?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgMoveAvailableVestingByDenomsResponse(): MsgMoveAvailableVestingByDenomsResponse {
  return {};
}

export const MsgMoveAvailableVestingByDenomsResponse = {
  encode(_: MsgMoveAvailableVestingByDenomsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMoveAvailableVestingByDenomsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMoveAvailableVestingByDenomsResponse();
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

  fromJSON(_: any): MsgMoveAvailableVestingByDenomsResponse {
    return {};
  },

  toJSON(_: MsgMoveAvailableVestingByDenomsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMoveAvailableVestingByDenomsResponse>, I>>(
    _: I,
  ): MsgMoveAvailableVestingByDenomsResponse {
    const message = createBaseMsgMoveAvailableVestingByDenomsResponse();
    return message;
  },
};

function createBaseMsgUpdateDenomParam(): MsgUpdateDenomParam {
  return { authority: "", denom: "" };
}

export const MsgUpdateDenomParam = {
  encode(message: MsgUpdateDenomParam, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.denom !== "") {
      writer.uint32(18).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDenomParam {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDenomParam();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authority = reader.string();
          break;
        case 2:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateDenomParam {
    return {
      authority: isSet(object.authority) ? String(object.authority) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
    };
  },

  toJSON(message: MsgUpdateDenomParam): unknown {
    const obj: any = {};
    message.authority !== undefined && (obj.authority = message.authority);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDenomParam>, I>>(object: I): MsgUpdateDenomParam {
    const message = createBaseMsgUpdateDenomParam();
    message.authority = object.authority ?? "";
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseMsgUpdateDenomParamResponse(): MsgUpdateDenomParamResponse {
  return {};
}

export const MsgUpdateDenomParamResponse = {
  encode(_: MsgUpdateDenomParamResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateDenomParamResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateDenomParamResponse();
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

  fromJSON(_: any): MsgUpdateDenomParamResponse {
    return {};
  },

  toJSON(_: MsgUpdateDenomParamResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateDenomParamResponse>, I>>(_: I): MsgUpdateDenomParamResponse {
    const message = createBaseMsgUpdateDenomParamResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVestingPool(request: MsgCreateVestingPool): Promise<MsgCreateVestingPoolResponse>;
  WithdrawAllAvailable(request: MsgWithdrawAllAvailable): Promise<MsgWithdrawAllAvailableResponse>;
  CreateVestingAccount(request: MsgCreateVestingAccount): Promise<MsgCreateVestingAccountResponse>;
  SendToVestingAccount(request: MsgSendToVestingAccount): Promise<MsgSendToVestingAccountResponse>;
  SplitVesting(request: MsgSplitVesting): Promise<MsgSplitVestingResponse>;
  MoveAvailableVesting(request: MsgMoveAvailableVesting): Promise<MsgMoveAvailableVestingResponse>;
  MoveAvailableVestingByDenoms(
    request: MsgMoveAvailableVestingByDenoms,
  ): Promise<MsgMoveAvailableVestingByDenomsResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateDenomParam(request: MsgUpdateDenomParam): Promise<MsgUpdateDenomParamResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateVestingPool = this.CreateVestingPool.bind(this);
    this.WithdrawAllAvailable = this.WithdrawAllAvailable.bind(this);
    this.CreateVestingAccount = this.CreateVestingAccount.bind(this);
    this.SendToVestingAccount = this.SendToVestingAccount.bind(this);
    this.SplitVesting = this.SplitVesting.bind(this);
    this.MoveAvailableVesting = this.MoveAvailableVesting.bind(this);
    this.MoveAvailableVestingByDenoms = this.MoveAvailableVestingByDenoms.bind(this);
    this.UpdateDenomParam = this.UpdateDenomParam.bind(this);
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

  SplitVesting(request: MsgSplitVesting): Promise<MsgSplitVestingResponse> {
    const data = MsgSplitVesting.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "SplitVesting", data);
    return promise.then((data) => MsgSplitVestingResponse.decode(new _m0.Reader(data)));
  }

  MoveAvailableVesting(request: MsgMoveAvailableVesting): Promise<MsgMoveAvailableVestingResponse> {
    const data = MsgMoveAvailableVesting.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "MoveAvailableVesting", data);
    return promise.then((data) => MsgMoveAvailableVestingResponse.decode(new _m0.Reader(data)));
  }

  MoveAvailableVestingByDenoms(
    request: MsgMoveAvailableVestingByDenoms,
  ): Promise<MsgMoveAvailableVestingByDenomsResponse> {
    const data = MsgMoveAvailableVestingByDenoms.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "MoveAvailableVestingByDenoms", data);
    return promise.then((data) => MsgMoveAvailableVestingByDenomsResponse.decode(new _m0.Reader(data)));
  }

  UpdateDenomParam(request: MsgUpdateDenomParam): Promise<MsgUpdateDenomParamResponse> {
    const data = MsgUpdateDenomParam.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Msg", "UpdateDenomParam", data);
    return promise.then((data) => MsgUpdateDenomParamResponse.decode(new _m0.Reader(data)));
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
