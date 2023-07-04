/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Campaign } from "./campaign";
import { UserEntry } from "./claim_record";
import { Mission } from "./mission";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

/** GenesisState defines the cfeclaim module's genesis state. */
export interface GenesisState {
  campaigns: Campaign[];
  campaignCount: number;
  usersEntries: UserEntry[];
  missions: Mission[];
  missionCounts: MissionCount[];
}

export interface MissionCount {
  campaignId: number;
  count: number;
}

function createBaseGenesisState(): GenesisState {
  return { campaigns: [], campaignCount: 0, usersEntries: [], missions: [], missionCounts: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.campaigns) {
      Campaign.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.campaignCount !== 0) {
      writer.uint32(16).uint64(message.campaignCount);
    }
    for (const v of message.usersEntries) {
      UserEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.missionCounts) {
      MissionCount.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaigns.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 2:
          message.campaignCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.usersEntries.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 4:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        case 5:
          message.missionCounts.push(MissionCount.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      campaigns: Array.isArray(object?.campaigns) ? object.campaigns.map((e: any) => Campaign.fromJSON(e)) : [],
      campaignCount: isSet(object.campaignCount) ? Number(object.campaignCount) : 0,
      usersEntries: Array.isArray(object?.usersEntries)
        ? object.usersEntries.map((e: any) => UserEntry.fromJSON(e))
        : [],
      missions: Array.isArray(object?.missions) ? object.missions.map((e: any) => Mission.fromJSON(e)) : [],
      missionCounts: Array.isArray(object?.missionCounts)
        ? object.missionCounts.map((e: any) => MissionCount.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.campaigns) {
      obj.campaigns = message.campaigns.map((e) => e ? Campaign.toJSON(e) : undefined);
    } else {
      obj.campaigns = [];
    }
    message.campaignCount !== undefined && (obj.campaignCount = Math.round(message.campaignCount));
    if (message.usersEntries) {
      obj.usersEntries = message.usersEntries.map((e) => e ? UserEntry.toJSON(e) : undefined);
    } else {
      obj.usersEntries = [];
    }
    if (message.missions) {
      obj.missions = message.missions.map((e) => e ? Mission.toJSON(e) : undefined);
    } else {
      obj.missions = [];
    }
    if (message.missionCounts) {
      obj.missionCounts = message.missionCounts.map((e) => e ? MissionCount.toJSON(e) : undefined);
    } else {
      obj.missionCounts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.campaigns = object.campaigns?.map((e) => Campaign.fromPartial(e)) || [];
    message.campaignCount = object.campaignCount ?? 0;
    message.usersEntries = object.usersEntries?.map((e) => UserEntry.fromPartial(e)) || [];
    message.missions = object.missions?.map((e) => Mission.fromPartial(e)) || [];
    message.missionCounts = object.missionCounts?.map((e) => MissionCount.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMissionCount(): MissionCount {
  return { campaignId: 0, count: 0 };
}

export const MissionCount = {
  encode(message: MissionCount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.count !== 0) {
      writer.uint32(16).uint64(message.count);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MissionCount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMissionCount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.count = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MissionCount {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      count: isSet(object.count) ? Number(object.count) : 0,
    };
  },

  toJSON(message: MissionCount): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.count !== undefined && (obj.count = Math.round(message.count));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MissionCount>, I>>(object: I): MissionCount {
    const message = createBaseMissionCount();
    message.campaignId = object.campaignId ?? 0;
    message.count = object.count ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
