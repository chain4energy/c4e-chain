/* eslint-disable */
import { Account } from "../cfedistributor/sub_distributor";
import { DecCoin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface Distribution {
  subdistributor: string;
  share_name: string;
  sources: Account[];
  destination: Account | undefined;
  amount: DecCoin[];
}

export interface DistributionBurn {
  subdistributor: string;
  sources: Account[];
  amount: DecCoin[];
}

const baseDistribution: object = { subdistributor: "", share_name: "" };

export const Distribution = {
  encode(message: Distribution, writer: Writer = Writer.create()): Writer {
    if (message.subdistributor !== "") {
      writer.uint32(10).string(message.subdistributor);
    }
    if (message.share_name !== "") {
      writer.uint32(18).string(message.share_name);
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

  decode(input: Reader | Uint8Array, length?: number): Distribution {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDistribution } as Distribution;
    message.sources = [];
    message.amount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subdistributor = reader.string();
          break;
        case 2:
          message.share_name = reader.string();
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
    const message = { ...baseDistribution } as Distribution;
    message.sources = [];
    message.amount = [];
    if (object.subdistributor !== undefined && object.subdistributor !== null) {
      message.subdistributor = String(object.subdistributor);
    } else {
      message.subdistributor = "";
    }
    if (object.share_name !== undefined && object.share_name !== null) {
      message.share_name = String(object.share_name);
    } else {
      message.share_name = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(Account.fromJSON(e));
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromJSON(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(DecCoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Distribution): unknown {
    const obj: any = {};
    message.subdistributor !== undefined &&
      (obj.subdistributor = message.subdistributor);
    message.share_name !== undefined && (obj.share_name = message.share_name);
    if (message.sources) {
      obj.sources = message.sources.map((e) =>
        e ? Account.toJSON(e) : undefined
      );
    } else {
      obj.sources = [];
    }
    message.destination !== undefined &&
      (obj.destination = message.destination
        ? Account.toJSON(message.destination)
        : undefined);
    if (message.amount) {
      obj.amount = message.amount.map((e) =>
        e ? DecCoin.toJSON(e) : undefined
      );
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Distribution>): Distribution {
    const message = { ...baseDistribution } as Distribution;
    message.sources = [];
    message.amount = [];
    if (object.subdistributor !== undefined && object.subdistributor !== null) {
      message.subdistributor = object.subdistributor;
    } else {
      message.subdistributor = "";
    }
    if (object.share_name !== undefined && object.share_name !== null) {
      message.share_name = object.share_name;
    } else {
      message.share_name = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(Account.fromPartial(e));
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromPartial(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(DecCoin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseDistributionBurn: object = { subdistributor: "" };

export const DistributionBurn = {
  encode(message: DistributionBurn, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): DistributionBurn {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDistributionBurn } as DistributionBurn;
    message.sources = [];
    message.amount = [];
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
    const message = { ...baseDistributionBurn } as DistributionBurn;
    message.sources = [];
    message.amount = [];
    if (object.subdistributor !== undefined && object.subdistributor !== null) {
      message.subdistributor = String(object.subdistributor);
    } else {
      message.subdistributor = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(Account.fromJSON(e));
      }
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(DecCoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: DistributionBurn): unknown {
    const obj: any = {};
    message.subdistributor !== undefined &&
      (obj.subdistributor = message.subdistributor);
    if (message.sources) {
      obj.sources = message.sources.map((e) =>
        e ? Account.toJSON(e) : undefined
      );
    } else {
      obj.sources = [];
    }
    if (message.amount) {
      obj.amount = message.amount.map((e) =>
        e ? DecCoin.toJSON(e) : undefined
      );
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DistributionBurn>): DistributionBurn {
    const message = { ...baseDistributionBurn } as DistributionBurn;
    message.sources = [];
    message.amount = [];
    if (object.subdistributor !== undefined && object.subdistributor !== null) {
      message.subdistributor = object.subdistributor;
    } else {
      message.subdistributor = "";
    }
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(Account.fromPartial(e));
      }
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(DecCoin.fromPartial(e));
      }
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
