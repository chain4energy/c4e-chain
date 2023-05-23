/* eslint-disable */
import { Params } from "../cfeclaim/params";
import {
  Campaign,
  CampaignCurrentAmount,
  CampaignTotalAmount,
} from "../cfeclaim/campaign";
import { UserEntry } from "../cfeclaim/claim_record";
import { Mission } from "../cfeclaim/mission";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

/** GenesisState defines the cfeclaim module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  campaigns: Campaign[];
  users_entries: UserEntry[];
  missions: Mission[];
  campaigns_amount_left: CampaignCurrentAmount[];
  campaigns_total_amount: CampaignTotalAmount[];
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
    for (const v of message.users_entries) {
      UserEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.campaigns_amount_left) {
      CampaignCurrentAmount.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.campaigns_total_amount) {
      CampaignTotalAmount.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.users_entries = [];
    message.missions = [];
    message.campaigns_amount_left = [];
    message.campaigns_total_amount = [];
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
          message.users_entries.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 5:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        case 6:
          message.campaigns_amount_left.push(
            CampaignCurrentAmount.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.campaigns_total_amount.push(
            CampaignTotalAmount.decode(reader, reader.uint32())
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
    message.users_entries = [];
    message.missions = [];
    message.campaigns_amount_left = [];
    message.campaigns_total_amount = [];
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
    if (object.users_entries !== undefined && object.users_entries !== null) {
      for (const e of object.users_entries) {
        message.users_entries.push(UserEntry.fromJSON(e));
      }
    }
    if (object.missions !== undefined && object.missions !== null) {
      for (const e of object.missions) {
        message.missions.push(Mission.fromJSON(e));
      }
    }
    if (
      object.campaigns_amount_left !== undefined &&
      object.campaigns_amount_left !== null
    ) {
      for (const e of object.campaigns_amount_left) {
        message.campaigns_amount_left.push(CampaignCurrentAmount.fromJSON(e));
      }
    }
    if (
      object.campaigns_total_amount !== undefined &&
      object.campaigns_total_amount !== null
    ) {
      for (const e of object.campaigns_total_amount) {
        message.campaigns_total_amount.push(CampaignTotalAmount.fromJSON(e));
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
    if (message.users_entries) {
      obj.users_entries = message.users_entries.map((e) =>
        e ? UserEntry.toJSON(e) : undefined
      );
    } else {
      obj.users_entries = [];
    }
    if (message.missions) {
      obj.missions = message.missions.map((e) =>
        e ? Mission.toJSON(e) : undefined
      );
    } else {
      obj.missions = [];
    }
    if (message.campaigns_amount_left) {
      obj.campaigns_amount_left = message.campaigns_amount_left.map((e) =>
        e ? CampaignCurrentAmount.toJSON(e) : undefined
      );
    } else {
      obj.campaigns_amount_left = [];
    }
    if (message.campaigns_total_amount) {
      obj.campaigns_total_amount = message.campaigns_total_amount.map((e) =>
        e ? CampaignTotalAmount.toJSON(e) : undefined
      );
    } else {
      obj.campaigns_total_amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.campaigns = [];
    message.users_entries = [];
    message.missions = [];
    message.campaigns_amount_left = [];
    message.campaigns_total_amount = [];
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
    if (object.users_entries !== undefined && object.users_entries !== null) {
      for (const e of object.users_entries) {
        message.users_entries.push(UserEntry.fromPartial(e));
      }
    }
    if (object.missions !== undefined && object.missions !== null) {
      for (const e of object.missions) {
        message.missions.push(Mission.fromPartial(e));
      }
    }
    if (
      object.campaigns_amount_left !== undefined &&
      object.campaigns_amount_left !== null
    ) {
      for (const e of object.campaigns_amount_left) {
        message.campaigns_amount_left.push(CampaignCurrentAmount.fromPartial(e));
      }
    }
    if (
      object.campaigns_total_amount !== undefined &&
      object.campaigns_total_amount !== null
    ) {
      for (const e of object.campaigns_total_amount) {
        message.campaigns_total_amount.push(CampaignTotalAmount.fromPartial(e));
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
