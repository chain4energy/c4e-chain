/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface NewVestingAccount {
  address: string;
}

export interface NewVestingPool {
  /** TODO: rename to owner */
  creator: string;
  name: string;
  amount: string;
  duration: string;
  vestingType: string;
}

export interface NewVestingPeriodFromVestingPool {
  /** TODO: rename to owner */
  owner_address: string;
  address: string;
  vesting_pool_name: string;
  amount: string;
  restart_vesting: string;
}

export interface WithdrawAvailable {
  /** TODO: rename to owner */
  owner_address: string;
  vesting_pool_name: string;
  amount: string;
}

const baseNewVestingAccount: object = { address: "" };

export const NewVestingAccount = {
  encode(message: NewVestingAccount, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NewVestingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNewVestingAccount } as NewVestingAccount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NewVestingAccount {
    const message = { ...baseNewVestingAccount } as NewVestingAccount;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: NewVestingAccount): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<NewVestingAccount>): NewVestingAccount {
    const message = { ...baseNewVestingAccount } as NewVestingAccount;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseNewVestingPool: object = {
  creator: "",
  name: "",
  amount: "",
  duration: "",
  vestingType: "",
};

export const NewVestingPool = {
  encode(message: NewVestingPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.duration !== "") {
      writer.uint32(34).string(message.duration);
    }
    if (message.vestingType !== "") {
      writer.uint32(42).string(message.vestingType);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NewVestingPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNewVestingPool } as NewVestingPool;
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
          message.amount = reader.string();
          break;
        case 4:
          message.duration = reader.string();
          break;
        case 5:
          message.vestingType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NewVestingPool {
    const message = { ...baseNewVestingPool } as NewVestingPool;
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
      message.duration = String(object.duration);
    } else {
      message.duration = "";
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    return message;
  },

  toJSON(message: NewVestingPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined && (obj.duration = message.duration);
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial(object: DeepPartial<NewVestingPool>): NewVestingPool {
    const message = { ...baseNewVestingPool } as NewVestingPool;
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
      message.duration = object.duration;
    } else {
      message.duration = "";
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    return message;
  },
};

const baseNewVestingPeriodFromVestingPool: object = {
  owner_address: "",
  address: "",
  vesting_pool_name: "",
  amount: "",
  restart_vesting: "",
};

export const NewVestingPeriodFromVestingPool = {
  encode(
    message: NewVestingPeriodFromVestingPool,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner_address !== "") {
      writer.uint32(10).string(message.owner_address);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.vesting_pool_name !== "") {
      writer.uint32(26).string(message.vesting_pool_name);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restart_vesting !== "") {
      writer.uint32(42).string(message.restart_vesting);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): NewVestingPeriodFromVestingPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseNewVestingPeriodFromVestingPool,
    } as NewVestingPeriodFromVestingPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner_address = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.vesting_pool_name = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.restart_vesting = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NewVestingPeriodFromVestingPool {
    const message = {
      ...baseNewVestingPeriodFromVestingPool,
    } as NewVestingPeriodFromVestingPool;
    if (object.owner_address !== undefined && object.owner_address !== null) {
      message.owner_address = String(object.owner_address);
    } else {
      message.owner_address = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
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
      message.restart_vesting = String(object.restart_vesting);
    } else {
      message.restart_vesting = "";
    }
    return message;
  },

  toJSON(message: NewVestingPeriodFromVestingPool): unknown {
    const obj: any = {};
    message.owner_address !== undefined &&
      (obj.owner_address = message.owner_address);
    message.address !== undefined && (obj.address = message.address);
    message.vesting_pool_name !== undefined &&
      (obj.vesting_pool_name = message.vesting_pool_name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restart_vesting !== undefined &&
      (obj.restart_vesting = message.restart_vesting);
    return obj;
  },

  fromPartial(
    object: DeepPartial<NewVestingPeriodFromVestingPool>
  ): NewVestingPeriodFromVestingPool {
    const message = {
      ...baseNewVestingPeriodFromVestingPool,
    } as NewVestingPeriodFromVestingPool;
    if (object.owner_address !== undefined && object.owner_address !== null) {
      message.owner_address = object.owner_address;
    } else {
      message.owner_address = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
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
      message.restart_vesting = "";
    }
    return message;
  },
};

const baseWithdrawAvailable: object = {
  owner_address: "",
  vesting_pool_name: "",
  amount: "",
};

export const WithdrawAvailable = {
  encode(message: WithdrawAvailable, writer: Writer = Writer.create()): Writer {
    if (message.owner_address !== "") {
      writer.uint32(10).string(message.owner_address);
    }
    if (message.vesting_pool_name !== "") {
      writer.uint32(18).string(message.vesting_pool_name);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): WithdrawAvailable {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWithdrawAvailable } as WithdrawAvailable;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner_address = reader.string();
          break;
        case 2:
          message.vesting_pool_name = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WithdrawAvailable {
    const message = { ...baseWithdrawAvailable } as WithdrawAvailable;
    if (object.owner_address !== undefined && object.owner_address !== null) {
      message.owner_address = String(object.owner_address);
    } else {
      message.owner_address = "";
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
    return message;
  },

  toJSON(message: WithdrawAvailable): unknown {
    const obj: any = {};
    message.owner_address !== undefined &&
      (obj.owner_address = message.owner_address);
    message.vesting_pool_name !== undefined &&
      (obj.vesting_pool_name = message.vesting_pool_name);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<WithdrawAvailable>): WithdrawAvailable {
    const message = { ...baseWithdrawAvailable } as WithdrawAvailable;
    if (object.owner_address !== undefined && object.owner_address !== null) {
      message.owner_address = object.owner_address;
    } else {
      message.owner_address = "";
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
    return message;
  },
};

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
