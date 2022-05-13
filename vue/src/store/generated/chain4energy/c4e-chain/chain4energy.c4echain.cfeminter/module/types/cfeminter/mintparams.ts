/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** HalvingMinter represents the inflation parameters. */
export interface HalvingMinter {
  /** the number of coins produced from the first block */
  new_coins_mint: number;
  /** expected blocks per year */
  blocks_per_year: number;
  /** type of coin to mint */
  mint_denom: string;
}

const baseHalvingMinter: object = {
  new_coins_mint: 0,
  blocks_per_year: 0,
  mint_denom: "",
};

export const HalvingMinter = {
  encode(message: HalvingMinter, writer: Writer = Writer.create()): Writer {
    if (message.new_coins_mint !== 0) {
      writer.uint32(8).int64(message.new_coins_mint);
    }
    if (message.blocks_per_year !== 0) {
      writer.uint32(48).int64(message.blocks_per_year);
    }
    if (message.mint_denom !== "") {
      writer.uint32(26).string(message.mint_denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): HalvingMinter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHalvingMinter } as HalvingMinter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.new_coins_mint = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.blocks_per_year = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.mint_denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HalvingMinter {
    const message = { ...baseHalvingMinter } as HalvingMinter;
    if (object.new_coins_mint !== undefined && object.new_coins_mint !== null) {
      message.new_coins_mint = Number(object.new_coins_mint);
    } else {
      message.new_coins_mint = 0;
    }
    if (
      object.blocks_per_year !== undefined &&
      object.blocks_per_year !== null
    ) {
      message.blocks_per_year = Number(object.blocks_per_year);
    } else {
      message.blocks_per_year = 0;
    }
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = String(object.mint_denom);
    } else {
      message.mint_denom = "";
    }
    return message;
  },

  toJSON(message: HalvingMinter): unknown {
    const obj: any = {};
    message.new_coins_mint !== undefined &&
      (obj.new_coins_mint = message.new_coins_mint);
    message.blocks_per_year !== undefined &&
      (obj.blocks_per_year = message.blocks_per_year);
    message.mint_denom !== undefined && (obj.mint_denom = message.mint_denom);
    return obj;
  },

  fromPartial(object: DeepPartial<HalvingMinter>): HalvingMinter {
    const message = { ...baseHalvingMinter } as HalvingMinter;
    if (object.new_coins_mint !== undefined && object.new_coins_mint !== null) {
      message.new_coins_mint = object.new_coins_mint;
    } else {
      message.new_coins_mint = 0;
    }
    if (
      object.blocks_per_year !== undefined &&
      object.blocks_per_year !== null
    ) {
      message.blocks_per_year = object.blocks_per_year;
    } else {
      message.blocks_per_year = 0;
    }
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = object.mint_denom;
    } else {
      message.mint_denom = "";
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
