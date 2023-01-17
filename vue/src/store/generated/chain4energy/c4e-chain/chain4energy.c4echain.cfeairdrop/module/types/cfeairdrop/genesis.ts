/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cfeairdrop/params";
import {
  Campaign,
  UserAirdropEntries,
  Mission,
  AirdropEntry,
} from "../cfeairdrop/airdrop";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** GenesisState defines the cfeairdrop module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  campaigns: Campaign[];
  user_airdrop_entries: UserAirdropEntries[];
  missions: Mission[];
  airdropEntryList: AirdropEntry[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  airdropEntryCount: number;
}

const baseGenesisState: object = { airdropEntryCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.campaigns) {
      Campaign.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.user_airdrop_entries) {
      UserAirdropEntries.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.airdropEntryList) {
      AirdropEntry.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.airdropEntryCount !== 0) {
      writer.uint32(56).uint64(message.airdropEntryCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.user_airdrop_entries = [];
    message.missions = [];
    message.airdropEntryList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.campaigns.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 3:
          message.user_airdrop_entries.push(
            UserAirdropEntries.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        case 6:
          message.airdropEntryList.push(
            AirdropEntry.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.airdropEntryCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.user_airdrop_entries = [];
    message.missions = [];
    message.airdropEntryList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.campaigns !== undefined && object.campaigns !== null) {
      for (const e of object.campaigns) {
        message.campaigns.push(Campaign.fromJSON(e));
      }
    }
    if (
      object.user_airdrop_entries !== undefined &&
      object.user_airdrop_entries !== null
    ) {
      for (const e of object.user_airdrop_entries) {
        message.user_airdrop_entries.push(UserAirdropEntries.fromJSON(e));
      }
    }
    if (object.missions !== undefined && object.missions !== null) {
      for (const e of object.missions) {
        message.missions.push(Mission.fromJSON(e));
      }
    }
    if (
      object.airdropEntryList !== undefined &&
      object.airdropEntryList !== null
    ) {
      for (const e of object.airdropEntryList) {
        message.airdropEntryList.push(AirdropEntry.fromJSON(e));
      }
    }
    if (
      object.airdropEntryCount !== undefined &&
      object.airdropEntryCount !== null
    ) {
      message.airdropEntryCount = Number(object.airdropEntryCount);
    } else {
      message.airdropEntryCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.campaigns) {
      obj.campaigns = message.campaigns.map((e) =>
        e ? Campaign.toJSON(e) : undefined
      );
    } else {
      obj.campaigns = [];
    }
    if (message.user_airdrop_entries) {
      obj.user_airdrop_entries = message.user_airdrop_entries.map((e) =>
        e ? UserAirdropEntries.toJSON(e) : undefined
      );
    } else {
      obj.user_airdrop_entries = [];
    }
    if (message.missions) {
      obj.missions = message.missions.map((e) =>
        e ? Mission.toJSON(e) : undefined
      );
    } else {
      obj.missions = [];
    }
    if (message.airdropEntryList) {
      obj.airdropEntryList = message.airdropEntryList.map((e) =>
        e ? AirdropEntry.toJSON(e) : undefined
      );
    } else {
      obj.airdropEntryList = [];
    }
    message.airdropEntryCount !== undefined &&
      (obj.airdropEntryCount = message.airdropEntryCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.user_airdrop_entries = [];
    message.missions = [];
    message.airdropEntryList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.campaigns !== undefined && object.campaigns !== null) {
      for (const e of object.campaigns) {
        message.campaigns.push(Campaign.fromPartial(e));
      }
    }
    if (
      object.user_airdrop_entries !== undefined &&
      object.user_airdrop_entries !== null
    ) {
      for (const e of object.user_airdrop_entries) {
        message.user_airdrop_entries.push(UserAirdropEntries.fromPartial(e));
      }
    }
    if (object.missions !== undefined && object.missions !== null) {
      for (const e of object.missions) {
        message.missions.push(Mission.fromPartial(e));
      }
    }
    if (
      object.airdropEntryList !== undefined &&
      object.airdropEntryList !== null
    ) {
      for (const e of object.airdropEntryList) {
        message.airdropEntryList.push(AirdropEntry.fromPartial(e));
      }
    }
    if (
      object.airdropEntryCount !== undefined &&
      object.airdropEntryCount !== null
    ) {
      message.airdropEntryCount = object.airdropEntryCount;
    } else {
      message.airdropEntryCount = 0;
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
