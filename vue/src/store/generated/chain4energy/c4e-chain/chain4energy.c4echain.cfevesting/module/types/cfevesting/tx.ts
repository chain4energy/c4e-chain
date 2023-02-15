/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Duration } from "../google/protobuf/duration";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface MsgCreateVestingPool {
  /** TODO: rename to owner */
  creator: string;
  name: string;
  amount: string;
  duration: Duration | undefined;
  vesting_type: string;
}

export interface MsgCreateVestingPoolResponse {}

export interface MsgWithdrawAllAvailable {
  /** TODO: rename to owner */
  creator: string;
}

export interface MsgWithdrawAllAvailableResponse {
  withdrawn: string;
}

export interface MsgCreateVestingAccount {
  from_address: string;
  to_address: string;
  amount: Coin[];
  start_time: number;
  end_time: number;
}

export interface MsgCreateVestingAccountResponse {}

export interface MsgSendToVestingAccount {
  /** TODO: rename to owner */
  from_address: string;
  to_address: string;
  vesting_pool_name: string;
  amount: string;
  restart_vesting: boolean;
}

export interface MsgSendToVestingAccountResponse {}

export interface MsgSplitVesting {
  from_address: string;
  to_address: string;
  amount: Coin[];
}

export interface MsgSplitVestingResponse {}

export interface MsgMoveAvailableVesting {
  fromAddress: string;
  toAddress: string;
}

export interface MsgMoveAvailableVestingResponse {}

export interface MsgMoveAvailableVestingByDenoms {
  fromAddress: string;
  toAddress: string;
  denoms: string;
}

export interface MsgMoveAvailableVestingByDenomsResponse {}

const baseMsgCreateVestingPool: object = {
  creator: "",
  name: "",
  amount: "",
  vesting_type: "",
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
    if (message.vesting_type !== "") {
      writer.uint32(50).string(message.vesting_type);
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
          message.vesting_type = reader.string();
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
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = String(object.vesting_type);
    } else {
      message.vesting_type = "";
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
    message.vesting_type !== undefined &&
      (obj.vesting_type = message.vesting_type);
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
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = object.vesting_type;
    } else {
      message.vesting_type = "";
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

const baseMsgWithdrawAllAvailableResponse: object = { withdrawn: "" };

export const MsgWithdrawAllAvailableResponse = {
  encode(
    message: MsgWithdrawAllAvailableResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.withdrawn !== "") {
      writer.uint32(10).string(message.withdrawn);
    }
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
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = String(object.withdrawn);
    } else {
      message.withdrawn = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawAllAvailableResponse): unknown {
    const obj: any = {};
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgWithdrawAllAvailableResponse>
  ): MsgWithdrawAllAvailableResponse {
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = object.withdrawn;
    } else {
      message.withdrawn = "";
    }
    return message;
  },
};

const baseMsgCreateVestingAccount: object = {
  from_address: "",
  to_address: "",
  start_time: 0,
  end_time: 0,
};

export const MsgCreateVestingAccount = {
  encode(
    message: MsgCreateVestingAccount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.from_address !== "") {
      writer.uint32(10).string(message.from_address);
    }
    if (message.to_address !== "") {
      writer.uint32(18).string(message.to_address);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.start_time !== 0) {
      writer.uint32(32).int64(message.start_time);
    }
    if (message.end_time !== 0) {
      writer.uint32(40).int64(message.end_time);
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
          message.from_address = reader.string();
          break;
        case 2:
          message.to_address = reader.string();
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.start_time = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.end_time = longToNumber(reader.int64() as Long);
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
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = String(object.from_address);
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = String(object.to_address);
    } else {
      message.to_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = Number(object.start_time);
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = Number(object.end_time);
    } else {
      message.end_time = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateVestingAccount): unknown {
    const obj: any = {};
    message.from_address !== undefined &&
      (obj.from_address = message.from_address);
    message.to_address !== undefined && (obj.to_address = message.to_address);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    message.start_time !== undefined && (obj.start_time = message.start_time);
    message.end_time !== undefined && (obj.end_time = message.end_time);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateVestingAccount>
  ): MsgCreateVestingAccount {
    const message = {
      ...baseMsgCreateVestingAccount,
    } as MsgCreateVestingAccount;
    message.amount = [];
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = object.from_address;
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = object.to_address;
    } else {
      message.to_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = 0;
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
  from_address: "",
  to_address: "",
  vesting_pool_name: "",
  amount: "",
  restart_vesting: false,
};

export const MsgSendToVestingAccount = {
  encode(
    message: MsgSendToVestingAccount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.from_address !== "") {
      writer.uint32(10).string(message.from_address);
    }
    if (message.to_address !== "") {
      writer.uint32(18).string(message.to_address);
    }
    if (message.vesting_pool_name !== "") {
      writer.uint32(26).string(message.vesting_pool_name);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restart_vesting === true) {
      writer.uint32(40).bool(message.restart_vesting);
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
          message.from_address = reader.string();
          break;
        case 2:
          message.to_address = reader.string();
          break;
        case 3:
          message.vesting_pool_name = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.restart_vesting = reader.bool();
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
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = String(object.from_address);
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = String(object.to_address);
    } else {
      message.to_address = "";
    }
    if (
      object.vesting_pool_name !== undefined &&
      object.vesting_pool_name !== null
    ) {
      message.vesting_pool_name = String(object.vesting_pool_name);
    } else {
      message.vesting_pool_name = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (
      object.restart_vesting !== undefined &&
      object.restart_vesting !== null
    ) {
      message.restart_vesting = Boolean(object.restart_vesting);
    } else {
      message.restart_vesting = false;
    }
    return message;
  },

  toJSON(message: MsgSendToVestingAccount): unknown {
    const obj: any = {};
    message.from_address !== undefined &&
      (obj.from_address = message.from_address);
    message.to_address !== undefined && (obj.to_address = message.to_address);
    message.vesting_pool_name !== undefined &&
      (obj.vesting_pool_name = message.vesting_pool_name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restart_vesting !== undefined &&
      (obj.restart_vesting = message.restart_vesting);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSendToVestingAccount>
  ): MsgSendToVestingAccount {
    const message = {
      ...baseMsgSendToVestingAccount,
    } as MsgSendToVestingAccount;
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = object.from_address;
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = object.to_address;
    } else {
      message.to_address = "";
    }
    if (
      object.vesting_pool_name !== undefined &&
      object.vesting_pool_name !== null
    ) {
      message.vesting_pool_name = object.vesting_pool_name;
    } else {
      message.vesting_pool_name = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (
      object.restart_vesting !== undefined &&
      object.restart_vesting !== null
    ) {
      message.restart_vesting = object.restart_vesting;
    } else {
      message.restart_vesting = false;
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

const baseMsgSplitVesting: object = { from_address: "", to_address: "" };

export const MsgSplitVesting = {
  encode(message: MsgSplitVesting, writer: Writer = Writer.create()): Writer {
    if (message.from_address !== "") {
      writer.uint32(10).string(message.from_address);
    }
    if (message.to_address !== "") {
      writer.uint32(18).string(message.to_address);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSplitVesting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSplitVesting } as MsgSplitVesting;
    message.amount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.from_address = reader.string();
          break;
        case 2:
          message.to_address = reader.string();
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
    const message = { ...baseMsgSplitVesting } as MsgSplitVesting;
    message.amount = [];
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = String(object.from_address);
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = String(object.to_address);
    } else {
      message.to_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgSplitVesting): unknown {
    const obj: any = {};
    message.from_address !== undefined &&
      (obj.from_address = message.from_address);
    message.to_address !== undefined && (obj.to_address = message.to_address);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSplitVesting>): MsgSplitVesting {
    const message = { ...baseMsgSplitVesting } as MsgSplitVesting;
    message.amount = [];
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = object.from_address;
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = object.to_address;
    } else {
      message.to_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgSplitVestingResponse: object = {};

export const MsgSplitVestingResponse = {
  encode(_: MsgSplitVestingResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSplitVestingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSplitVestingResponse,
    } as MsgSplitVestingResponse;
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
    const message = {
      ...baseMsgSplitVestingResponse,
    } as MsgSplitVestingResponse;
    return message;
  },

  toJSON(_: MsgSplitVestingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSplitVestingResponse>
  ): MsgSplitVestingResponse {
    const message = {
      ...baseMsgSplitVestingResponse,
    } as MsgSplitVestingResponse;
    return message;
  },
};

const baseMsgMoveAvailableVesting: object = { fromAddress: "", toAddress: "" };

export const MsgMoveAvailableVesting = {
  encode(
    message: MsgMoveAvailableVesting,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMoveAvailableVesting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgMoveAvailableVesting,
    } as MsgMoveAvailableVesting;
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
    const message = {
      ...baseMsgMoveAvailableVesting,
    } as MsgMoveAvailableVesting;
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
    return message;
  },

  toJSON(message: MsgMoveAvailableVesting): unknown {
    const obj: any = {};
    message.fromAddress !== undefined &&
      (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgMoveAvailableVesting>
  ): MsgMoveAvailableVesting {
    const message = {
      ...baseMsgMoveAvailableVesting,
    } as MsgMoveAvailableVesting;
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
    return message;
  },
};

const baseMsgMoveAvailableVestingResponse: object = {};

export const MsgMoveAvailableVestingResponse = {
  encode(
    _: MsgMoveAvailableVestingResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgMoveAvailableVestingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgMoveAvailableVestingResponse,
    } as MsgMoveAvailableVestingResponse;
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
    const message = {
      ...baseMsgMoveAvailableVestingResponse,
    } as MsgMoveAvailableVestingResponse;
    return message;
  },

  toJSON(_: MsgMoveAvailableVestingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgMoveAvailableVestingResponse>
  ): MsgMoveAvailableVestingResponse {
    const message = {
      ...baseMsgMoveAvailableVestingResponse,
    } as MsgMoveAvailableVestingResponse;
    return message;
  },
};

const baseMsgMoveAvailableVestingByDenoms: object = {
  fromAddress: "",
  toAddress: "",
  denoms: "",
};

export const MsgMoveAvailableVestingByDenoms = {
  encode(
    message: MsgMoveAvailableVestingByDenoms,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.fromAddress !== "") {
      writer.uint32(10).string(message.fromAddress);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    if (message.denoms !== "") {
      writer.uint32(26).string(message.denoms);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgMoveAvailableVestingByDenoms {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgMoveAvailableVestingByDenoms,
    } as MsgMoveAvailableVestingByDenoms;
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
          message.denoms = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMoveAvailableVestingByDenoms {
    const message = {
      ...baseMsgMoveAvailableVestingByDenoms,
    } as MsgMoveAvailableVestingByDenoms;
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
    if (object.denoms !== undefined && object.denoms !== null) {
      message.denoms = String(object.denoms);
    } else {
      message.denoms = "";
    }
    return message;
  },

  toJSON(message: MsgMoveAvailableVestingByDenoms): unknown {
    const obj: any = {};
    message.fromAddress !== undefined &&
      (obj.fromAddress = message.fromAddress);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    message.denoms !== undefined && (obj.denoms = message.denoms);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgMoveAvailableVestingByDenoms>
  ): MsgMoveAvailableVestingByDenoms {
    const message = {
      ...baseMsgMoveAvailableVestingByDenoms,
    } as MsgMoveAvailableVestingByDenoms;
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
    if (object.denoms !== undefined && object.denoms !== null) {
      message.denoms = object.denoms;
    } else {
      message.denoms = "";
    }
    return message;
  },
};

const baseMsgMoveAvailableVestingByDenomsResponse: object = {};

export const MsgMoveAvailableVestingByDenomsResponse = {
  encode(
    _: MsgMoveAvailableVestingByDenomsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgMoveAvailableVestingByDenomsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgMoveAvailableVestingByDenomsResponse,
    } as MsgMoveAvailableVestingByDenomsResponse;
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
    const message = {
      ...baseMsgMoveAvailableVestingByDenomsResponse,
    } as MsgMoveAvailableVestingByDenomsResponse;
    return message;
  },

  toJSON(_: MsgMoveAvailableVestingByDenomsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgMoveAvailableVestingByDenomsResponse>
  ): MsgMoveAvailableVestingByDenomsResponse {
    const message = {
      ...baseMsgMoveAvailableVestingByDenomsResponse,
    } as MsgMoveAvailableVestingByDenomsResponse;
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
  SendToVestingAccount(
    request: MsgSendToVestingAccount
  ): Promise<MsgSendToVestingAccountResponse>;
  SplitVesting(request: MsgSplitVesting): Promise<MsgSplitVestingResponse>;
  MoveAvailableVesting(
    request: MsgMoveAvailableVesting
  ): Promise<MsgMoveAvailableVestingResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  MoveAvailableVestingByDenoms(
    request: MsgMoveAvailableVestingByDenoms
  ): Promise<MsgMoveAvailableVestingByDenomsResponse>;
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

  SplitVesting(request: MsgSplitVesting): Promise<MsgSplitVestingResponse> {
    const data = MsgSplitVesting.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "SplitVesting",
      data
    );
    return promise.then((data) =>
      MsgSplitVestingResponse.decode(new Reader(data))
    );
  }

  MoveAvailableVesting(
    request: MsgMoveAvailableVesting
  ): Promise<MsgMoveAvailableVestingResponse> {
    const data = MsgMoveAvailableVesting.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "MoveAvailableVesting",
      data
    );
    return promise.then((data) =>
      MsgMoveAvailableVestingResponse.decode(new Reader(data))
    );
  }

  MoveAvailableVestingByDenoms(
    request: MsgMoveAvailableVestingByDenoms
  ): Promise<MsgMoveAvailableVestingByDenomsResponse> {
    const data = MsgMoveAvailableVestingByDenoms.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "MoveAvailableVestingByDenoms",
      data
    );
    return promise.then((data) =>
      MsgMoveAvailableVestingByDenomsResponse.decode(new Reader(data))
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
