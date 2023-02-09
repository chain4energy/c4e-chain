/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DecCoin } from "../../cosmos/base/v1beta1/coin";
import { Account } from "./sub_distributor";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface Distribution {
  subdistributor: string;
  shareName: string;
  sources: Account[];
  destination: Account | undefined;
  amount: DecCoin[];
}

export interface DistributionBurn {
  subdistributor: string;
  sources: Account[];
  amount: DecCoin[];
}

function createBaseDistribution(): Distribution {
  return { subdistributor: "", shareName: "", sources: [], destination: undefined, amount: [] };
}

export const Distribution = {
  encode(message: Distribution, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subdistributor !== "") {
      writer.uint32(10).string(message.subdistributor);
    }
    if (message.shareName !== "") {
      writer.uint32(18).string(message.shareName);
    }
    for (const v of message.sources) {
      Account.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.destination !== undefined) {
      Account.encode(message.destination, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.amount) {
      DecCoin.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Distribution {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistribution();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subdistributor = reader.string();
          break;
        case 2:
          message.shareName = reader.string();
          break;
        case 3:
          message.sources.push(Account.decode(reader, reader.uint32()));
          break;
        case 4:
          message.destination = Account.decode(reader, reader.uint32());
          break;
        case 5:
          message.amount.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Distribution {
    return {
      subdistributor: isSet(object.subdistributor) ? String(object.subdistributor) : "",
      shareName: isSet(object.shareName) ? String(object.shareName) : "",
      sources: Array.isArray(object?.sources) ? object.sources.map((e: any) => Account.fromJSON(e)) : [],
      destination: isSet(object.destination) ? Account.fromJSON(object.destination) : undefined,
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => DecCoin.fromJSON(e)) : [],
    };
  },

  toJSON(message: Distribution): unknown {
    const obj: any = {};
    message.subdistributor !== undefined && (obj.subdistributor = message.subdistributor);
    message.shareName !== undefined && (obj.shareName = message.shareName);
    if (message.sources) {
      obj.sources = message.sources.map((e) => e ? Account.toJSON(e) : undefined);
    } else {
      obj.sources = [];
    }
    message.destination !== undefined
      && (obj.destination = message.destination ? Account.toJSON(message.destination) : undefined);
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? DecCoin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Distribution>, I>>(object: I): Distribution {
    const message = createBaseDistribution();
    message.subdistributor = object.subdistributor ?? "";
    message.shareName = object.shareName ?? "";
    message.sources = object.sources?.map((e) => Account.fromPartial(e)) || [];
    message.destination = (object.destination !== undefined && object.destination !== null)
      ? Account.fromPartial(object.destination)
      : undefined;
    message.amount = object.amount?.map((e) => DecCoin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDistributionBurn(): DistributionBurn {
  return { subdistributor: "", sources: [], amount: [] };
}

export const DistributionBurn = {
  encode(message: DistributionBurn, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subdistributor !== "") {
      writer.uint32(10).string(message.subdistributor);
    }
    for (const v of message.sources) {
      Account.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.amount) {
      DecCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistributionBurn {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistributionBurn();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subdistributor = reader.string();
          break;
        case 2:
          message.sources.push(Account.decode(reader, reader.uint32()));
          break;
        case 3:
          message.amount.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DistributionBurn {
    return {
      subdistributor: isSet(object.subdistributor) ? String(object.subdistributor) : "",
      sources: Array.isArray(object?.sources) ? object.sources.map((e: any) => Account.fromJSON(e)) : [],
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => DecCoin.fromJSON(e)) : [],
    };
  },

  toJSON(message: DistributionBurn): unknown {
    const obj: any = {};
    message.subdistributor !== undefined && (obj.subdistributor = message.subdistributor);
    if (message.sources) {
      obj.sources = message.sources.map((e) => e ? Account.toJSON(e) : undefined);
    } else {
      obj.sources = [];
    }
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? DecCoin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DistributionBurn>, I>>(object: I): DistributionBurn {
    const message = createBaseDistributionBurn();
    message.subdistributor = object.subdistributor ?? "";
    message.sources = object.sources?.map((e) => Account.fromPartial(e)) || [];
    message.amount = object.amount?.map((e) => DecCoin.fromPartial(e)) || [];
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
