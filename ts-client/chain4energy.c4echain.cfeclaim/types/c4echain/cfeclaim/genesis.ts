/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Campaign } from "./campaign";
import { UserEntry } from "./claim_record";
import { Mission } from "./mission";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

/** GenesisState defines the cfeclaim module's genesis state. */
export interface GenesisState {
  campaigns: Campaign[];
  usersEntries: UserEntry[];
  missions: Mission[];
}

function createBaseGenesisState(): GenesisState {
  return { campaigns: [], usersEntries: [], missions: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.campaigns) {
      Campaign.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.usersEntries) {
      UserEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(42).fork()).ldelim();
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
        case 2:
          message.campaigns.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 3:
          message.usersEntries.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 5:
          message.missions.push(Mission.decode(reader, reader.uint32()));
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
      usersEntries: Array.isArray(object?.usersEntries)
        ? object.usersEntries.map((e: any) => UserEntry.fromJSON(e))
        : [],
      missions: Array.isArray(object?.missions) ? object.missions.map((e: any) => Mission.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.campaigns) {
      obj.campaigns = message.campaigns.map((e) => e ? Campaign.toJSON(e) : undefined);
    } else {
      obj.campaigns = [];
    }
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
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.campaigns = object.campaigns?.map((e) => Campaign.fromPartial(e)) || [];
    message.usersEntries = object.usersEntries?.map((e) => UserEntry.fromPartial(e)) || [];
    message.missions = object.missions?.map((e) => Mission.fromPartial(e)) || [];
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
