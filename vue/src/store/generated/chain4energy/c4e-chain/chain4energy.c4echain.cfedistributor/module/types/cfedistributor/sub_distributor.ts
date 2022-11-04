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
  destinations: Destinations | undefined;
}

export interface Destinations {
  primary_share: Account | undefined;
  burn_share: string;
  shares: DestinationShare[];
}

export interface DestinationShare {
  name: string;
  share: string;
  destination: Account | undefined;
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
    if (message.destinations !== undefined) {
      Destinations.encode(
        message.destinations,
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
          message.destinations = Destinations.decode(reader, reader.uint32());
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
    if (object.destinations !== undefined && object.destinations !== null) {
      message.destinations = Destinations.fromJSON(object.destinations);
    } else {
      message.destinations = undefined;
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
    message.destinations !== undefined &&
      (obj.destinations = message.destinations
        ? Destinations.toJSON(message.destinations)
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
    if (object.destinations !== undefined && object.destinations !== null) {
      message.destinations = Destinations.fromPartial(object.destinations);
    } else {
      message.destinations = undefined;
    }
    return message;
  },
};

const baseDestinations: object = { burn_share: "" };

export const Destinations = {
  encode(message: Destinations, writer: Writer = Writer.create()): Writer {
    if (message.primary_share !== undefined) {
      Account.encode(message.primary_share, writer.uint32(10).fork()).ldelim();
    }
    if (message.burn_share !== "") {
      writer.uint32(18).string(message.burn_share);
    }
    for (const v of message.shares) {
      DestinationShare.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Destinations {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDestinations } as Destinations;
    message.shares = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.primary_share = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.burn_share = reader.string();
          break;
        case 3:
          message.shares.push(DestinationShare.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Destinations {
    const message = { ...baseDestinations } as Destinations;
    message.shares = [];
    if (object.primary_share !== undefined && object.primary_share !== null) {
      message.primary_share = Account.fromJSON(object.primary_share);
    } else {
      message.primary_share = undefined;
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = String(object.burn_share);
    } else {
      message.burn_share = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      for (const e of object.shares) {
        message.shares.push(DestinationShare.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Destinations): unknown {
    const obj: any = {};
    message.primary_share !== undefined &&
      (obj.primary_share = message.primary_share
        ? Account.toJSON(message.primary_share)
        : undefined);
    message.burn_share !== undefined && (obj.burn_share = message.burn_share);
    if (message.shares) {
      obj.shares = message.shares.map((e) =>
        e ? DestinationShare.toJSON(e) : undefined
      );
    } else {
      obj.shares = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Destinations>): Destinations {
    const message = { ...baseDestinations } as Destinations;
    message.shares = [];
    if (object.primary_share !== undefined && object.primary_share !== null) {
      message.primary_share = Account.fromPartial(object.primary_share);
    } else {
      message.primary_share = undefined;
    }
    if (object.burn_share !== undefined && object.burn_share !== null) {
      message.burn_share = object.burn_share;
    } else {
      message.burn_share = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      for (const e of object.shares) {
        message.shares.push(DestinationShare.fromPartial(e));
      }
    }
    return message;
  },
};

const baseDestinationShare: object = { name: "", share: "" };

export const DestinationShare = {
  encode(message: DestinationShare, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.share !== "") {
      writer.uint32(18).string(message.share);
    }
    if (message.destination !== undefined) {
      Account.encode(message.destination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DestinationShare {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDestinationShare } as DestinationShare;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.share = reader.string();
          break;
        case 3:
          message.destination = Account.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DestinationShare {
    const message = { ...baseDestinationShare } as DestinationShare;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.share !== undefined && object.share !== null) {
      message.share = String(object.share);
    } else {
      message.share = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromJSON(object.destination);
    } else {
      message.destination = undefined;
    }
    return message;
  },

  toJSON(message: DestinationShare): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.share !== undefined && (obj.share = message.share);
    message.destination !== undefined &&
      (obj.destination = message.destination
        ? Account.toJSON(message.destination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DestinationShare>): DestinationShare {
    const message = { ...baseDestinationShare } as DestinationShare;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.share !== undefined && object.share !== null) {
      message.share = object.share;
    } else {
      message.share = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromPartial(object.destination);
    } else {
      message.destination = undefined;
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
