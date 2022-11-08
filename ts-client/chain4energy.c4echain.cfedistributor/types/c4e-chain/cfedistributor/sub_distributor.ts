/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DecCoin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface State {
  account: Account | undefined;
  burn: boolean;
  remains: DecCoin[];
}

export interface SubDistributor {
  name: string;
  sources: Account[];
  destinations: Destinations | undefined;
}

export interface Destinations {
  primaryShare: Account | undefined;
  burnShare: string;
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

function createBaseState(): State {
  return { account: undefined, burn: false, remains: [] };
}

export const State = {
  encode(message: State, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    if (message.burn === true) {
      writer.uint32(16).bool(message.burn);
    }
    for (const v of message.remains) {
      DecCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): State {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseState();
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
          message.remains.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): State {
    return {
      account: isSet(object.account) ? Account.fromJSON(object.account) : undefined,
      burn: isSet(object.burn) ? Boolean(object.burn) : false,
      remains: Array.isArray(object?.remains) ? object.remains.map((e: any) => DecCoin.fromJSON(e)) : [],
    };
  },

  toJSON(message: State): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    message.burn !== undefined && (obj.burn = message.burn);
    if (message.remains) {
      obj.remains = message.remains.map((e) => e ? DecCoin.toJSON(e) : undefined);
    } else {
      obj.remains = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<State>, I>>(object: I): State {
    const message = createBaseState();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    message.burn = object.burn ?? false;
    message.remains = object.remains?.map((e) => DecCoin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSubDistributor(): SubDistributor {
  return { name: "", sources: [], destinations: undefined };
}

export const SubDistributor = {
  encode(message: SubDistributor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.sources) {
      Account.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.destinations !== undefined) {
      Destinations.encode(message.destinations, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubDistributor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubDistributor();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      sources: Array.isArray(object?.sources) ? object.sources.map((e: any) => Account.fromJSON(e)) : [],
      destinations: isSet(object.destinations) ? Destinations.fromJSON(object.destinations) : undefined,
    };
  },

  toJSON(message: SubDistributor): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.sources) {
      obj.sources = message.sources.map((e) => e ? Account.toJSON(e) : undefined);
    } else {
      obj.sources = [];
    }
    message.destinations !== undefined
      && (obj.destinations = message.destinations ? Destinations.toJSON(message.destinations) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SubDistributor>, I>>(object: I): SubDistributor {
    const message = createBaseSubDistributor();
    message.name = object.name ?? "";
    message.sources = object.sources?.map((e) => Account.fromPartial(e)) || [];
    message.destinations = (object.destinations !== undefined && object.destinations !== null)
      ? Destinations.fromPartial(object.destinations)
      : undefined;
    return message;
  },
};

function createBaseDestinations(): Destinations {
  return { primaryShare: undefined, burnShare: "", shares: [] };
}

export const Destinations = {
  encode(message: Destinations, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.primaryShare !== undefined) {
      Account.encode(message.primaryShare, writer.uint32(10).fork()).ldelim();
    }
    if (message.burnShare !== "") {
      writer.uint32(18).string(message.burnShare);
    }
    for (const v of message.shares) {
      DestinationShare.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Destinations {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDestinations();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.primaryShare = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.burnShare = reader.string();
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
    return {
      primaryShare: isSet(object.primaryShare) ? Account.fromJSON(object.primaryShare) : undefined,
      burnShare: isSet(object.burnShare) ? String(object.burnShare) : "",
      shares: Array.isArray(object?.shares) ? object.shares.map((e: any) => DestinationShare.fromJSON(e)) : [],
    };
  },

  toJSON(message: Destinations): unknown {
    const obj: any = {};
    message.primaryShare !== undefined
      && (obj.primaryShare = message.primaryShare ? Account.toJSON(message.primaryShare) : undefined);
    message.burnShare !== undefined && (obj.burnShare = message.burnShare);
    if (message.shares) {
      obj.shares = message.shares.map((e) => e ? DestinationShare.toJSON(e) : undefined);
    } else {
      obj.shares = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Destinations>, I>>(object: I): Destinations {
    const message = createBaseDestinations();
    message.primaryShare = (object.primaryShare !== undefined && object.primaryShare !== null)
      ? Account.fromPartial(object.primaryShare)
      : undefined;
    message.burnShare = object.burnShare ?? "";
    message.shares = object.shares?.map((e) => DestinationShare.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDestinationShare(): DestinationShare {
  return { name: "", share: "", destination: undefined };
}

export const DestinationShare = {
  encode(message: DestinationShare, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): DestinationShare {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDestinationShare();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      share: isSet(object.share) ? String(object.share) : "",
      destination: isSet(object.destination) ? Account.fromJSON(object.destination) : undefined,
    };
  },

  toJSON(message: DestinationShare): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.share !== undefined && (obj.share = message.share);
    message.destination !== undefined
      && (obj.destination = message.destination ? Account.toJSON(message.destination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DestinationShare>, I>>(object: I): DestinationShare {
    const message = createBaseDestinationShare();
    message.name = object.name ?? "";
    message.share = object.share ?? "";
    message.destination = (object.destination !== undefined && object.destination !== null)
      ? Account.fromPartial(object.destination)
      : undefined;
    return message;
  },
};

function createBaseAccount(): Account {
  return { id: "", type: "" };
}

export const Account = {
  encode(message: Account, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.type !== "") {
      writer.uint32(18).string(message.type);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Account {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccount();
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
    return { id: isSet(object.id) ? String(object.id) : "", type: isSet(object.type) ? String(object.type) : "" };
  },

  toJSON(message: Account): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.type !== undefined && (obj.type = message.type);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Account>, I>>(object: I): Account {
    const message = createBaseAccount();
    message.id = object.id ?? "";
    message.type = object.type ?? "";
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
