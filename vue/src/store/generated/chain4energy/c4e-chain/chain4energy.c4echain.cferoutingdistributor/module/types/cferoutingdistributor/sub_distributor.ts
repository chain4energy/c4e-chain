/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

/**
 * DenomUnit represents a struct that describes a given
 * denomination unit of the basic token.
 */
export interface DenomUnit {
  /** represent list of module account from which */
  sources: string[];
  /** denom represents the string name of the given denom unit (e.g uatom). */
  denom: string;
  /**
   * exponent represents power of 10 exponent that one must
   * raise the base_denom to in order to equal the given DenomUnit's denom
   * 1 denom = 10^exponent base_denom
   * (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
   * exponent = 6, thus: 1 atom = 10^6 uatom).
   */
  exponent: number;
}

export interface Destination {
  default_share_account: string;
  share: Share[];
}

export interface Share {
  name: string;
  percent: number;
  account: account | undefined;
}

export interface account {
  address: string;
  isModuleAccount: boolean;
}

const baseDenomUnit: object = { sources: "", denom: "", exponent: 0 };

export const DenomUnit = {
  encode(message: DenomUnit, writer: Writer = Writer.create()): Writer {
    for (const v of message.sources) {
      writer.uint32(26).string(v!);
    }
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.exponent !== 0) {
      writer.uint32(16).int32(message.exponent);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DenomUnit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDenomUnit } as DenomUnit;
    message.sources = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 3:
          message.sources.push(reader.string());
          break;
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.exponent = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DenomUnit {
    const message = { ...baseDenomUnit } as DenomUnit;
    message.sources = [];
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(String(e));
      }
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.exponent !== undefined && object.exponent !== null) {
      message.exponent = Number(object.exponent);
    } else {
      message.exponent = 0;
    }
    return message;
  },

  toJSON(message: DenomUnit): unknown {
    const obj: any = {};
    if (message.sources) {
      obj.sources = message.sources.map((e) => e);
    } else {
      obj.sources = [];
    }
    message.denom !== undefined && (obj.denom = message.denom);
    message.exponent !== undefined && (obj.exponent = message.exponent);
    return obj;
  },

  fromPartial(object: DeepPartial<DenomUnit>): DenomUnit {
    const message = { ...baseDenomUnit } as DenomUnit;
    message.sources = [];
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(e);
      }
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.exponent !== undefined && object.exponent !== null) {
      message.exponent = object.exponent;
    } else {
      message.exponent = 0;
    }
    return message;
  },
};

const baseDestination: object = { default_share_account: "" };

export const Destination = {
  encode(message: Destination, writer: Writer = Writer.create()): Writer {
    if (message.default_share_account !== "") {
      writer.uint32(10).string(message.default_share_account);
    }
    for (const v of message.share) {
      Share.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Destination {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDestination } as Destination;
    message.share = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.default_share_account = reader.string();
          break;
        case 2:
          message.share.push(Share.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Destination {
    const message = { ...baseDestination } as Destination;
    message.share = [];
    if (
      object.default_share_account !== undefined &&
      object.default_share_account !== null
    ) {
      message.default_share_account = String(object.default_share_account);
    } else {
      message.default_share_account = "";
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Destination): unknown {
    const obj: any = {};
    message.default_share_account !== undefined &&
      (obj.default_share_account = message.default_share_account);
    if (message.share) {
      obj.share = message.share.map((e) => (e ? Share.toJSON(e) : undefined));
    } else {
      obj.share = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Destination>): Destination {
    const message = { ...baseDestination } as Destination;
    message.share = [];
    if (
      object.default_share_account !== undefined &&
      object.default_share_account !== null
    ) {
      message.default_share_account = object.default_share_account;
    } else {
      message.default_share_account = "";
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromPartial(e));
      }
    }
    return message;
  },
};

const baseShare: object = { name: "", percent: 0 };

export const Share = {
  encode(message: Share, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.percent !== 0) {
      writer.uint32(16).int32(message.percent);
    }
    if (message.account !== undefined) {
      account.encode(message.account, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Share {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShare } as Share;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.percent = reader.int32();
          break;
        case 3:
          message.account = account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Share {
    const message = { ...baseShare } as Share;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = Number(object.percent);
    } else {
      message.percent = 0;
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = account.fromJSON(object.account);
    } else {
      message.account = undefined;
    }
    return message;
  },

  toJSON(message: Share): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.percent !== undefined && (obj.percent = message.percent);
    message.account !== undefined &&
      (obj.account = message.account
        ? account.toJSON(message.account)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Share>): Share {
    const message = { ...baseShare } as Share;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = object.percent;
    } else {
      message.percent = 0;
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = account.fromPartial(object.account);
    } else {
      message.account = undefined;
    }
    return message;
  },
};

const baseaccount: object = { address: "", isModuleAccount: false };

export const account = {
  encode(message: account, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.isModuleAccount === true) {
      writer.uint32(16).bool(message.isModuleAccount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): account {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseaccount } as account;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.isModuleAccount = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): account {
    const message = { ...baseaccount } as account;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (
      object.isModuleAccount !== undefined &&
      object.isModuleAccount !== null
    ) {
      message.isModuleAccount = Boolean(object.isModuleAccount);
    } else {
      message.isModuleAccount = false;
    }
    return message;
  },

  toJSON(message: account): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.isModuleAccount !== undefined &&
      (obj.isModuleAccount = message.isModuleAccount);
    return obj;
  },

  fromPartial(object: DeepPartial<account>): account {
    const message = { ...baseaccount } as account;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (
      object.isModuleAccount !== undefined &&
      object.isModuleAccount !== null
    ) {
      message.isModuleAccount = object.isModuleAccount;
    } else {
      message.isModuleAccount = false;
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
