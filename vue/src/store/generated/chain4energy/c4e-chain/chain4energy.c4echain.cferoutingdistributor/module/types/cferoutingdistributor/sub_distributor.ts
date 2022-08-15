/* eslint-disable */
import { DecCoin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

export interface Remains {
  account: Account | undefined;
  leftover_coin: DecCoin | undefined;
}

export interface RemainsList {
  remains: Remains[];
}

export interface RoutingDistributor {
  /** List contains distributors */
  sub_distributor: SubDistributor[];
  /** module account to load on start genesis */
  module_accounts: string[];
}

export interface SubDistributor {
  name: string;
  /** represent list of module account from which */
  sources: Account[];
  /** represent destinations */
  destination: Destination | undefined;
  order: number;
}

export interface Destination {
  account: Account | undefined;
  share: Share[];
  burn_share: BurnShare | undefined;
}

export interface BurnShare {
  percent: string;
}

export interface Share {
  name: string;
  percent: string;
  account: Account | undefined;
}

export interface Account {
  address: string;
  is_module_account: boolean;
  is_internal_account: boolean;
  is_main_collector: boolean;
}

const baseRemains: object = {};

export const Remains = {
  encode(message: Remains, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    if (message.leftover_coin !== undefined) {
      DecCoin.encode(message.leftover_coin, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Remains {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRemains } as Remains;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.leftover_coin = DecCoin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Remains {
    const message = { ...baseRemains } as Remains;
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account);
    } else {
      message.account = undefined;
    }
    if (object.leftover_coin !== undefined && object.leftover_coin !== null) {
      message.leftover_coin = DecCoin.fromJSON(object.leftover_coin);
    } else {
      message.leftover_coin = undefined;
    }
    return message;
  },

  toJSON(message: Remains): unknown {
    const obj: any = {};
    message.account !== undefined &&
      (obj.account = message.account
        ? Account.toJSON(message.account)
        : undefined);
    message.leftover_coin !== undefined &&
      (obj.leftover_coin = message.leftover_coin
        ? DecCoin.toJSON(message.leftover_coin)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Remains>): Remains {
    const message = { ...baseRemains } as Remains;
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account);
    } else {
      message.account = undefined;
    }
    if (object.leftover_coin !== undefined && object.leftover_coin !== null) {
      message.leftover_coin = DecCoin.fromPartial(object.leftover_coin);
    } else {
      message.leftover_coin = undefined;
    }
    return message;
  },
};

const baseRemainsList: object = {};

export const RemainsList = {
  encode(message: RemainsList, writer: Writer = Writer.create()): Writer {
    for (const v of message.remains) {
      Remains.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RemainsList {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRemainsList } as RemainsList;
    message.remains = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.remains.push(Remains.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemainsList {
    const message = { ...baseRemainsList } as RemainsList;
    message.remains = [];
    if (object.remains !== undefined && object.remains !== null) {
      for (const e of object.remains) {
        message.remains.push(Remains.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: RemainsList): unknown {
    const obj: any = {};
    if (message.remains) {
      obj.remains = message.remains.map((e) =>
        e ? Remains.toJSON(e) : undefined
      );
    } else {
      obj.remains = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<RemainsList>): RemainsList {
    const message = { ...baseRemainsList } as RemainsList;
    message.remains = [];
    if (object.remains !== undefined && object.remains !== null) {
      for (const e of object.remains) {
        message.remains.push(Remains.fromPartial(e));
      }
    }
    return message;
  },
};

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

const baseSubDistributor: object = { name: "", order: 0 };

export const SubDistributor = {
  encode(message: SubDistributor, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.sources) {
      Account.encode(v!, writer.uint32(18).fork()).ldelim();
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
          message.sources.push(Account.decode(reader, reader.uint32()));
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
        message.sources.push(Account.fromJSON(e));
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
      obj.sources = message.sources.map((e) =>
        e ? Account.toJSON(e) : undefined
      );
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
        message.sources.push(Account.fromPartial(e));
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

const baseDestination: object = {};

export const Destination = {
  encode(message: Destination, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.share) {
      Share.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.burn_share !== undefined) {
      BurnShare.encode(message.burn_share, writer.uint32(26).fork()).ldelim();
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
          message.account = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.share.push(Share.decode(reader, reader.uint32()));
          break;
        case 3:
          message.burn_share = BurnShare.decode(reader, reader.uint32());
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
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account);
    } else {
      message.account = undefined;
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromJSON(e));
      }
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = BurnShare.fromJSON(object.burn_share);
    } else {
      message.burn_share = undefined;
    }
    return message;
  },

  toJSON(message: Destination): unknown {
    const obj: any = {};
    message.account !== undefined &&
      (obj.account = message.account
        ? Account.toJSON(message.account)
        : undefined);
    if (message.share) {
      obj.share = message.share.map((e) => (e ? Share.toJSON(e) : undefined));
    } else {
      obj.share = [];
    }
    message.burn_share !== undefined &&
      (obj.burn_share = message.burn_share
        ? BurnShare.toJSON(message.burn_share)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Destination>): Destination {
    const message = { ...baseDestination } as Destination;
    message.share = [];
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account);
    } else {
      message.account = undefined;
    }
    if (object.share !== undefined && object.share !== null) {
      for (const e of object.share) {
        message.share.push(Share.fromPartial(e));
      }
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = BurnShare.fromPartial(object.burn_share);
    } else {
      message.burn_share = undefined;
    }
    return message;
  },
};

const baseBurnShare: object = { percent: "" };

export const BurnShare = {
  encode(message: BurnShare, writer: Writer = Writer.create()): Writer {
    if (message.percent !== "") {
      writer.uint32(10).string(message.percent);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): BurnShare {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseBurnShare } as BurnShare;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.percent = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BurnShare {
    const message = { ...baseBurnShare } as BurnShare;
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = String(object.percent);
    } else {
      message.percent = "";
    }
    return message;
  },

  toJSON(message: BurnShare): unknown {
    const obj: any = {};
    message.percent !== undefined && (obj.percent = message.percent);
    return obj;
  },

  fromPartial(object: DeepPartial<BurnShare>): BurnShare {
    const message = { ...baseBurnShare } as BurnShare;
    if (object.percent !== undefined && object.percent !== null) {
      message.percent = object.percent;
    } else {
      message.percent = "";
    }
    return message;
  },
};

const baseShare: object = { name: "", percent: "" };

export const Share = {
  encode(message: Share, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.percent !== "") {
      writer.uint32(18).string(message.percent);
    }
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(26).fork()).ldelim();
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
          message.percent = reader.string();
          break;
        case 3:
          message.account = Account.decode(reader, reader.uint32());
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
      message.percent = String(object.percent);
    } else {
      message.percent = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account);
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
        ? Account.toJSON(message.account)
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
      message.percent = "";
    }
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account);
    } else {
      message.account = undefined;
    }
    return message;
  },
};

const baseAccount: object = {
  address: "",
  is_module_account: false,
  is_internal_account: false,
  is_main_collector: false,
};

export const Account = {
  encode(message: Account, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.is_module_account === true) {
      writer.uint32(16).bool(message.is_module_account);
    }
    if (message.is_internal_account === true) {
      writer.uint32(24).bool(message.is_internal_account);
    }
    if (message.is_main_collector === true) {
      writer.uint32(32).bool(message.is_main_collector);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Account {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccount } as Account;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.is_module_account = reader.bool();
          break;
        case 3:
          message.is_internal_account = reader.bool();
          break;
        case 4:
          message.is_main_collector = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Account {
    const message = { ...baseAccount } as Account;
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
    if (
      object.is_internal_account !== undefined &&
      object.is_internal_account !== null
    ) {
      message.is_internal_account = Boolean(object.is_internal_account);
    } else {
      message.is_internal_account = false;
    }
    if (
      object.is_main_collector !== undefined &&
      object.is_main_collector !== null
    ) {
      message.is_main_collector = Boolean(object.is_main_collector);
    } else {
      message.is_main_collector = false;
    }
    return message;
  },

  toJSON(message: Account): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.is_module_account !== undefined &&
      (obj.is_module_account = message.is_module_account);
    message.is_internal_account !== undefined &&
      (obj.is_internal_account = message.is_internal_account);
    message.is_main_collector !== undefined &&
      (obj.is_main_collector = message.is_main_collector);
    return obj;
  },

  fromPartial(object: DeepPartial<Account>): Account {
    const message = { ...baseAccount } as Account;
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
    if (
      object.is_internal_account !== undefined &&
      object.is_internal_account !== null
    ) {
      message.is_internal_account = object.is_internal_account;
    } else {
      message.is_internal_account = false;
    }
    if (
      object.is_main_collector !== undefined &&
      object.is_main_collector !== null
    ) {
      message.is_main_collector = object.is_main_collector;
    } else {
      message.is_main_collector = false;
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
