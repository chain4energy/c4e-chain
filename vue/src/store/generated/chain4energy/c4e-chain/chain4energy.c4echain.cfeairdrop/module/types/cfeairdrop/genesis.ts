/* eslint-disable */
import { Params } from "../cfeairdrop/params";
import {
  Campaign,
  UserAirdropEntries,
  Mission,
  AirdropClaimsLeft,
  AirdropDistrubitions,
} from "../cfeairdrop/airdrop";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** GenesisState defines the cfeairdrop module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  campaigns: Campaign[];
  user_airdrop_entries: UserAirdropEntries[];
  missions: Mission[];
  airdrop_claims_left: AirdropClaimsLeft[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  airdrop_distrubitions: AirdropDistrubitions[];
}

const baseGenesisState: object = {};

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
    for (const v of message.airdrop_claims_left) {
      AirdropClaimsLeft.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.airdrop_distrubitions) {
      AirdropDistrubitions.encode(v!, writer.uint32(58).fork()).ldelim();
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
    message.airdrop_claims_left = [];
    message.airdrop_distrubitions = [];
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
          message.airdrop_claims_left.push(
            AirdropClaimsLeft.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.airdrop_distrubitions.push(
            AirdropDistrubitions.decode(reader, reader.uint32())
          );
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
    message.airdrop_claims_left = [];
    message.airdrop_distrubitions = [];
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
      object.airdrop_claims_left !== undefined &&
      object.airdrop_claims_left !== null
    ) {
      for (const e of object.airdrop_claims_left) {
        message.airdrop_claims_left.push(AirdropClaimsLeft.fromJSON(e));
      }
    }
    if (
      object.airdrop_distrubitions !== undefined &&
      object.airdrop_distrubitions !== null
    ) {
      for (const e of object.airdrop_distrubitions) {
        message.airdrop_distrubitions.push(AirdropDistrubitions.fromJSON(e));
      }
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
    if (message.airdrop_claims_left) {
      obj.airdrop_claims_left = message.airdrop_claims_left.map((e) =>
        e ? AirdropClaimsLeft.toJSON(e) : undefined
      );
    } else {
      obj.airdrop_claims_left = [];
    }
    if (message.airdrop_distrubitions) {
      obj.airdrop_distrubitions = message.airdrop_distrubitions.map((e) =>
        e ? AirdropDistrubitions.toJSON(e) : undefined
      );
    } else {
      obj.airdrop_distrubitions = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.user_airdrop_entries = [];
    message.missions = [];
    message.airdrop_claims_left = [];
    message.airdrop_distrubitions = [];
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
      object.airdrop_claims_left !== undefined &&
      object.airdrop_claims_left !== null
    ) {
      for (const e of object.airdrop_claims_left) {
        message.airdrop_claims_left.push(AirdropClaimsLeft.fromPartial(e));
      }
    }
    if (
      object.airdrop_distrubitions !== undefined &&
      object.airdrop_distrubitions !== null
    ) {
      for (const e of object.airdrop_distrubitions) {
        message.airdrop_distrubitions.push(AirdropDistrubitions.fromPartial(e));
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
