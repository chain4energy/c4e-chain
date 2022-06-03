/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

export interface RoutingDistributor {
  /** List contains distributors */
  sub_distributor: SubDistributor[];
  /** module account to load on start genesis */
  module_accounts: string[];
}

export interface SubDistributor {
  name: string;
  /** represent list of module account from which */
  sources: string[];
  /** represent destinations */
  destination: Destination | undefined;
  order: number;
}

export interface Destination {
  default_share_account: account | undefined;
  share: Share[];
  /** can be null */
  burn_share: number;
}

export interface Share {
  name: string;
  percent: number;
  account: account | undefined;
}

export interface account {
  address: string;
  is_module_account: boolean;
}

const baseRoutingDistributor: object = { module_accounts: "" };

export const RoutingDistributor = {
  encode(
    message: RoutingDistributor,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.sub_distributor) {
      SubDistributor.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.module_accounts) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RoutingDistributor {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRoutingDistributor } as RoutingDistributor;
    message.sub_distributor = [];
    message.module_accounts = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sub_distributor.push(
            SubDistributor.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.module_accounts.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RoutingDistributor {
    const message = { ...baseRoutingDistributor } as RoutingDistributor;
    message.sub_distributor = [];
    message.module_accounts = [];
    if (
      object.sub_distributor !== undefined &&
      object.sub_distributor !== null
    ) {
      for (const e of object.sub_distributor) {
        message.sub_distributor.push(SubDistributor.fromJSON(e));
      }
    }
    if (
      object.module_accounts !== undefined &&
      object.module_accounts !== null
    ) {
      for (const e of object.module_accounts) {
        message.module_accounts.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: RoutingDistributor): unknown {
    const obj: any = {};
    if (message.sub_distributor) {
      obj.sub_distributor = message.sub_distributor.map((e) =>
        e ? SubDistributor.toJSON(e) : undefined
      );
    } else {
      obj.sub_distributor = [];
    }
    if (message.module_accounts) {
      obj.module_accounts = message.module_accounts.map((e) => e);
    } else {
      obj.module_accounts = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<RoutingDistributor>): RoutingDistributor {
    const message = { ...baseRoutingDistributor } as RoutingDistributor;
    message.sub_distributor = [];
    message.module_accounts = [];
    if (
      object.sub_distributor !== undefined &&
      object.sub_distributor !== null
    ) {
      for (const e of object.sub_distributor) {
        message.sub_distributor.push(SubDistributor.fromPartial(e));
      }
    }
    if (
      object.module_accounts !== undefined &&
      object.module_accounts !== null
    ) {
      for (const e of object.module_accounts) {
        message.module_accounts.push(e);
      }
    }
    return message;
  },
};

const baseSubDistributor: object = { name: "", sources: "", order: 0 };

export const SubDistributor = {
  encode(message: SubDistributor, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.sources) {
      writer.uint32(18).string(v!);
    }
    if (message.destination !== undefined) {
      Destination.encode(
        message.destination,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.order !== 0) {
      writer.uint32(32).int32(message.order);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubDistributor {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSubDistributor } as SubDistributor;
    message.sources = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.sources.push(reader.string());
          break;
        case 3:
          message.destination = Destination.decode(reader, reader.uint32());
          break;
        case 4:
          message.order = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubDistributor {
    const message = { ...baseSubDistributor } as SubDistributor;
    message.sources = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(String(e));
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Destination.fromJSON(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.order !== undefined && object.order !== null) {
      message.order = Number(object.order);
    } else {
      message.order = 0;
    }
    return message;
  },

  toJSON(message: SubDistributor): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.sources) {
      obj.sources = message.sources.map((e) => e);
    } else {
      obj.sources = [];
    }
    message.destination !== undefined &&
      (obj.destination = message.destination
        ? Destination.toJSON(message.destination)
        : undefined);
    message.order !== undefined && (obj.order = message.order);
    return obj;
  },

  fromPartial(object: DeepPartial<SubDistributor>): SubDistributor {
    const message = { ...baseSubDistributor } as SubDistributor;
    message.sources = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(e);
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Destination.fromPartial(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.order !== undefined && object.order !== null) {
      message.order = object.order;
    } else {
      message.order = 0;
    }
    return message;
  },
};

const baseDestination: object = { burn_share: 0 };

export const Destination = {
  encode(message: Destination, writer: Writer = Writer.create()): Writer {
    if (message.default_share_account !== undefined) {
      account
        .encode(message.default_share_account, writer.uint32(10).fork())
        .ldelim();
    }
    for (const v of message.share) {
      Share.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.burn_share !== 0) {
      writer.uint32(24).int32(message.burn_share);
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
          message.default_share_account = account.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.share.push(Share.decode(reader, reader.uint32()));
          break;
        case 3:
          message.burn_share = reader.int32();
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
      message.default_share_account = account.fromJSON(
        object.default_share_account
      );
    } else {
      message.default_share_account = undefined;
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromJSON(e));
      }
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = Number(object.burn_share);
    } else {
      message.burn_share = 0;
    }
    return message;
  },

  toJSON(message: Destination): unknown {
    const obj: any = {};
    message.default_share_account !== undefined &&
      (obj.default_share_account = message.default_share_account
        ? account.toJSON(message.default_share_account)
        : undefined);
    if (message.share) {
      obj.share = message.share.map((e) => (e ? Share.toJSON(e) : undefined));
    } else {
      obj.share = [];
    }
    message.burn_share !== undefined && (obj.burn_share = message.burn_share);
    return obj;
  },

  fromPartial(object: DeepPartial<Destination>): Destination {
    const message = { ...baseDestination } as Destination;
    message.share = [];
    if (
      object.default_share_account !== undefined &&
      object.default_share_account !== null
    ) {
      message.default_share_account = account.fromPartial(
        object.default_share_account
      );
    } else {
      message.default_share_account = undefined;
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromPartial(e));
      }
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = object.burn_share;
    } else {
      message.burn_share = 0;
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
      writer.uint32(16).int64(message.percent);
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
          message.percent = longToNumber(reader.int64() as Long);
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

const baseaccount: object = { address: "", is_module_account: false };

export const account = {
  encode(message: account, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.is_module_account === true) {
      writer.uint32(16).bool(message.is_module_account);
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
          message.is_module_account = reader.bool();
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
      object.is_module_account !== undefined &&
      object.is_module_account !== null
    ) {
      message.is_module_account = Boolean(object.is_module_account);
    } else {
      message.is_module_account = false;
    }
    return message;
  },

  toJSON(message: account): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.is_module_account !== undefined &&
      (obj.is_module_account = message.is_module_account);
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
      object.is_module_account !== undefined &&
      object.is_module_account !== null
    ) {
      message.is_module_account = object.is_module_account;
    } else {
      message.is_module_account = false;
    }
    return message;
  },
};

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
