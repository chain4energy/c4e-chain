/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface NewVestingAccount {
  address: string;
}

export interface NewVestingPool {
  creator: string;
  name: string;
  amount: string;
  duration: string;
  vestingType: string;
}

export interface NewVestingPeriodFromVestingPool {
  ownerAddress: string;
  address: string;
  vestingPoolName: string;
  amount: string;
  restartVesting: string;
}

export interface WithdrawAvailable {
  ownerAddress: string;
  vestingPoolName: string;
  amount: string;
}

function createBaseNewVestingAccount(): NewVestingAccount {
  return { address: "" };
}

export const NewVestingAccount = {
  encode(message: NewVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NewVestingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNewVestingAccount();
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
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: NewVestingAccount): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NewVestingAccount>, I>>(object: I): NewVestingAccount {
    const message = createBaseNewVestingAccount();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseNewVestingPool(): NewVestingPool {
  return { creator: "", name: "", amount: "", duration: "", vestingType: "" };
}

export const NewVestingPool = {
  encode(message: NewVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): NewVestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNewVestingPool();
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
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      duration: isSet(object.duration) ? String(object.duration) : "",
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
    };
  },

  toJSON(message: NewVestingPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined && (obj.duration = message.duration);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NewVestingPool>, I>>(object: I): NewVestingPool {
    const message = createBaseNewVestingPool();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.amount = object.amount ?? "";
    message.duration = object.duration ?? "";
    message.vestingType = object.vestingType ?? "";
    return message;
  },
};

function createBaseNewVestingPeriodFromVestingPool(): NewVestingPeriodFromVestingPool {
  return { ownerAddress: "", address: "", vestingPoolName: "", amount: "", restartVesting: "" };
}

export const NewVestingPeriodFromVestingPool = {
  encode(message: NewVestingPeriodFromVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ownerAddress !== "") {
      writer.uint32(10).string(message.ownerAddress);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(34).string(message.vestingPoolName);
    }
    if (message.amount !== "") {
      writer.uint32(42).string(message.amount);
    }
    if (message.restartVesting !== "") {
      writer.uint32(50).string(message.restartVesting);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NewVestingPeriodFromVestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNewVestingPeriodFromVestingPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ownerAddress = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 4:
          message.vestingPoolName = reader.string();
          break;
        case 5:
          message.amount = reader.string();
          break;
        case 6:
          message.restartVesting = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NewVestingPeriodFromVestingPool {
    return {
      ownerAddress: isSet(object.ownerAddress) ? String(object.ownerAddress) : "",
      address: isSet(object.address) ? String(object.address) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      restartVesting: isSet(object.restartVesting) ? String(object.restartVesting) : "",
    };
  },

  toJSON(message: NewVestingPeriodFromVestingPool): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.address !== undefined && (obj.address = message.address);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restartVesting !== undefined && (obj.restartVesting = message.restartVesting);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NewVestingPeriodFromVestingPool>, I>>(
    object: I,
  ): NewVestingPeriodFromVestingPool {
    const message = createBaseNewVestingPeriodFromVestingPool();
    message.ownerAddress = object.ownerAddress ?? "";
    message.address = object.address ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.amount = object.amount ?? "";
    message.restartVesting = object.restartVesting ?? "";
    return message;
  },
};

function createBaseWithdrawAvailable(): WithdrawAvailable {
  return { ownerAddress: "", vestingPoolName: "", amount: "" };
}

export const WithdrawAvailable = {
  encode(message: WithdrawAvailable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ownerAddress !== "") {
      writer.uint32(10).string(message.ownerAddress);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(34).string(message.vestingPoolName);
    }
    if (message.amount !== "") {
      writer.uint32(42).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WithdrawAvailable {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWithdrawAvailable();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ownerAddress = reader.string();
          break;
        case 4:
          message.vestingPoolName = reader.string();
          break;
        case 5:
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
    return {
      ownerAddress: isSet(object.ownerAddress) ? String(object.ownerAddress) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: WithdrawAvailable): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WithdrawAvailable>, I>>(object: I): WithdrawAvailable {
    const message = createBaseWithdrawAvailable();
    message.ownerAddress = object.ownerAddress ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

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
