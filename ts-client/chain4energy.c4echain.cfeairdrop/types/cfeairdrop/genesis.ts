/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { CampaignAmountLeft, CampaignTotalAmount, Campaign, Mission, UserEntry } from "./claim";
import { Params } from "./params";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

/** GenesisState defines the cfeclaim module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  campaigns: Campaign[];
  userEntry: UserEntry[];
  missions: Mission[];
  claimClaimsLeft: CampaignAmountLeft[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  claimDistrubitions: CampaignTotalAmount[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    campaigns: [],
    userEntry: [],
    missions: [],
    claimClaimsLeft: [],
    claimDistrubitions: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.campaigns) {
      Campaign.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.userEntry) {
      UserEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.claimClaimsLeft) {
      CampaignAmountLeft.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.claimDistrubitions) {
      CampaignTotalAmount.encode(v!, writer.uint32(58).fork()).ldelim();
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
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.campaigns.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 3:
          message.userEntry.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 5:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        case 6:
          message.claimClaimsLeft.push(CampaignAmountLeft.decode(reader, reader.uint32()));
          break;
        case 7:
          message.claimDistrubitions.push(CampaignTotalAmount.decode(reader, reader.uint32()));
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
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      campaigns: Array.isArray(object?.campaigns) ? object.campaigns.map((e: any) => Campaign.fromJSON(e)) : [],
      userEntry: Array.isArray(object?.userEntry)
        ? object.userEntry.map((e: any) => UserEntry.fromJSON(e))
        : [],
      missions: Array.isArray(object?.missions) ? object.missions.map((e: any) => Mission.fromJSON(e)) : [],
      claimClaimsLeft: Array.isArray(object?.claimClaimsLeft)
        ? object.claimClaimsLeft.map((e: any) => CampaignAmountLeft.fromJSON(e))
        : [],
      claimDistrubitions: Array.isArray(object?.claimDistrubitions)
        ? object.claimDistrubitions.map((e: any) => CampaignTotalAmount.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.campaigns) {
      obj.campaigns = message.campaigns.map((e) => e ? Campaign.toJSON(e) : undefined);
    } else {
      obj.campaigns = [];
    }
    if (message.userEntry) {
      obj.userEntry = message.userEntry.map((e) => e ? UserEntry.toJSON(e) : undefined);
    } else {
      obj.userEntry = [];
    }
    if (message.missions) {
      obj.missions = message.missions.map((e) => e ? Mission.toJSON(e) : undefined);
    } else {
      obj.missions = [];
    }
    if (message.claimClaimsLeft) {
      obj.claimClaimsLeft = message.claimClaimsLeft.map((e) => e ? CampaignAmountLeft.toJSON(e) : undefined);
    } else {
      obj.claimClaimsLeft = [];
    }
    if (message.claimDistrubitions) {
      obj.claimDistrubitions = message.claimDistrubitions.map((e) =>
        e ? CampaignTotalAmount.toJSON(e) : undefined
      );
    } else {
      obj.claimDistrubitions = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.campaigns = object.campaigns?.map((e) => Campaign.fromPartial(e)) || [];
    message.userEntry = object.userEntry?.map((e) => UserEntry.fromPartial(e)) || [];
    message.missions = object.missions?.map((e) => Mission.fromPartial(e)) || [];
    message.claimClaimsLeft = object.claimClaimsLeft?.map((e) => CampaignAmountLeft.fromPartial(e)) || [];
    message.claimDistrubitions = object.claimDistrubitions?.map((e) => CampaignTotalAmount.fromPartial(e)) || [];
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
