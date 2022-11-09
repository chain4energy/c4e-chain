/* eslint-disable */
import { DecCoin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface State {
  account: Account | undefined;
  burn: boolean;
  coins_states: DecCoin[];
}

export interface SubDistributor {
  name: string;
  sources: Account[];
  destination: Destination | undefined;
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
  id: string;
  type: string;
}

const baseState: object = { burn: false };

export const State = {
  encode(message: State, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    if (message.burn === true) {
      writer.uint32(16).bool(message.burn);
    }
    for (const v of message.coins_states) {
      DecCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): State {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseState } as State;
    message.coins_states = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.burn = reader.bool();
          break;
        case 3:
          message.coins_states.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): State {
    const message = { ...baseState } as State;
    message.coins_states = [];
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account);
    } else {
      message.account = undefined;
    }
    if (object.burn !== undefined && object.burn !== null) {
      message.burn = Boolean(object.burn);
    } else {
      message.burn = false;
    }
    if (object.coins_states !== undefined && object.coins_states !== null) {
      for (const e of object.coins_states) {
        message.coins_states.push(DecCoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: State): unknown {
    const obj: any = {};
    message.account !== undefined &&
      (obj.account = message.account
        ? Account.toJSON(message.account)
        : undefined);
    message.burn !== undefined && (obj.burn = message.burn);
    if (message.coins_states) {
      obj.coins_states = message.coins_states.map((e) =>
        e ? DecCoin.toJSON(e) : undefined
      );
    } else {
      obj.coins_states = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<State>): State {
    const message = { ...baseState } as State;
    message.coins_states = [];
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account);
    } else {
      message.account = undefined;
    }
    if (object.burn !== undefined && object.burn !== null) {
      message.burn = object.burn;
    } else {
      message.burn = false;
    }
    if (object.coins_states !== undefined && object.coins_states !== null) {
      for (const e of object.coins_states) {
        message.coins_states.push(DecCoin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSubDistributor: object = { name: "" };

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

const baseAccount: object = { id: "", type: "" };

export const Account = {
  encode(message: Account, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
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
          message.id = reader.string();
          break;
        case 2:
          message.type = reader.string();
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
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    return message;
  },

  toJSON(message: Account): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial(object: DeepPartial<Account>): Account {
    const message = { ...baseAccount } as Account;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
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
