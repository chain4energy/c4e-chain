/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DecCoin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface State {
  account: Account | undefined;
  burn: boolean;
  coinsStates: DecCoin[];
}

export interface SubDistributor {
  name: string;
  sources: Account[];
  destination: Destination | undefined;
}

export interface Destination {
  account: Account | undefined;
  share: Share[];
  burnShare: BurnShare | undefined;
}

export interface BurnShare {
  /** float percent =1; */
  percent: string;
}

export interface Share {
  name: string;
  /** float percent = 2; */
  percent: string;
  account: Account | undefined;
}

export interface Account {
  id: string;
  type: string;
}

function createBaseState(): State {
  return { account: undefined, burn: false, coinsStates: [] };
}

export const State = {
  encode(message: State, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    if (message.burn === true) {
      writer.uint32(16).bool(message.burn);
    }
    for (const v of message.coinsStates) {
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
          message.coinsStates.push(DecCoin.decode(reader, reader.uint32()));
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
      coinsStates: Array.isArray(object?.coinsStates) ? object.coinsStates.map((e: any) => DecCoin.fromJSON(e)) : [],
    };
  },

  toJSON(message: State): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    message.burn !== undefined && (obj.burn = message.burn);
    if (message.coinsStates) {
      obj.coinsStates = message.coinsStates.map((e) => e ? DecCoin.toJSON(e) : undefined);
    } else {
      obj.coinsStates = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<State>, I>>(object: I): State {
    const message = createBaseState();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    message.burn = object.burn ?? false;
    message.coinsStates = object.coinsStates?.map((e) => DecCoin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSubDistributor(): SubDistributor {
  return { name: "", sources: [], destination: undefined };
}

export const SubDistributor = {
  encode(message: SubDistributor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.sources) {
      Account.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.destination !== undefined) {
      Destination.encode(message.destination, writer.uint32(26).fork()).ldelim();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      sources: Array.isArray(object?.sources) ? object.sources.map((e: any) => Account.fromJSON(e)) : [],
      destination: isSet(object.destination) ? Destination.fromJSON(object.destination) : undefined,
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
    message.destination !== undefined
      && (obj.destination = message.destination ? Destination.toJSON(message.destination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SubDistributor>, I>>(object: I): SubDistributor {
    const message = createBaseSubDistributor();
    message.name = object.name ?? "";
    message.sources = object.sources?.map((e) => Account.fromPartial(e)) || [];
    message.destination = (object.destination !== undefined && object.destination !== null)
      ? Destination.fromPartial(object.destination)
      : undefined;
    return message;
  },
};

function createBaseDestination(): Destination {
  return { account: undefined, share: [], burnShare: undefined };
}

export const Destination = {
  encode(message: Destination, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.share) {
      Share.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.burnShare !== undefined) {
      BurnShare.encode(message.burnShare, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Destination {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDestination();
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
          message.burnShare = BurnShare.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Destination {
    return {
      account: isSet(object.account) ? Account.fromJSON(object.account) : undefined,
      share: Array.isArray(object?.share) ? object.share.map((e: any) => Share.fromJSON(e)) : [],
      burnShare: isSet(object.burnShare) ? BurnShare.fromJSON(object.burnShare) : undefined,
    };
  },

  toJSON(message: Destination): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    if (message.share) {
      obj.share = message.share.map((e) => e ? Share.toJSON(e) : undefined);
    } else {
      obj.share = [];
    }
    message.burnShare !== undefined
      && (obj.burnShare = message.burnShare ? BurnShare.toJSON(message.burnShare) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Destination>, I>>(object: I): Destination {
    const message = createBaseDestination();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    message.share = object.share?.map((e) => Share.fromPartial(e)) || [];
    message.burnShare = (object.burnShare !== undefined && object.burnShare !== null)
      ? BurnShare.fromPartial(object.burnShare)
      : undefined;
    return message;
  },
};

function createBaseBurnShare(): BurnShare {
  return { percent: "" };
}

export const BurnShare = {
  encode(message: BurnShare, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.percent !== "") {
      writer.uint32(10).string(message.percent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BurnShare {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBurnShare();
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
    return { percent: isSet(object.percent) ? String(object.percent) : "" };
  },

  toJSON(message: BurnShare): unknown {
    const obj: any = {};
    message.percent !== undefined && (obj.percent = message.percent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BurnShare>, I>>(object: I): BurnShare {
    const message = createBaseBurnShare();
    message.percent = object.percent ?? "";
    return message;
  },
};

function createBaseShare(): Share {
  return { name: "", percent: "", account: undefined };
}

export const Share = {
  encode(message: Share, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Share {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseShare();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      percent: isSet(object.percent) ? String(object.percent) : "",
      account: isSet(object.account) ? Account.fromJSON(object.account) : undefined,
    };
  },

  toJSON(message: Share): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.percent !== undefined && (obj.percent = message.percent);
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Share>, I>>(object: I): Share {
    const message = createBaseShare();
    message.name = object.name ?? "";
    message.percent = object.percent ?? "";
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
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
