/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export enum CampaignType {
  CAMPAIGN_TYPE_UNSPECIFIED = 0,
  /** DEFAULT - campaign must be enabled to add new claim records */
  DEFAULT = 1,
  /** VESTING_POOL - campaign don't have to be enabled (started) to add new claim records and it is */
  VESTING_POOL = 2,
  UNRECOGNIZED = -1,
}

export function campaignTypeFromJSON(object: any): CampaignType {
  switch (object) {
    case 0:
    case "CAMPAIGN_TYPE_UNSPECIFIED":
      return CampaignType.CAMPAIGN_TYPE_UNSPECIFIED;
    case 1:
    case "DEFAULT":
      return CampaignType.DEFAULT;
    case 2:
    case "VESTING_POOL":
      return CampaignType.VESTING_POOL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CampaignType.UNRECOGNIZED;
  }
}

export function campaignTypeToJSON(object: CampaignType): string {
  switch (object) {
    case CampaignType.CAMPAIGN_TYPE_UNSPECIFIED:
      return "CAMPAIGN_TYPE_UNSPECIFIED";
    case CampaignType.DEFAULT:
      return "DEFAULT";
    case CampaignType.VESTING_POOL:
      return "VESTING_POOL";
    case CampaignType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Campaign {
  id: number;
  owner: string;
  name: string;
  description: string;
  campaignType: CampaignType;
  removableClaimRecords: boolean;
  feegrantAmount: string;
  initialClaimFreeAmount: string;
  free: string;
  enabled: boolean;
  startTime: Date | undefined;
  endTime:
    | Date
    | undefined;
  /** period of locked coins from claim */
  lockupPeriod:
    | Duration
    | undefined;
  /** period of vesting coins after lockup period */
  vestingPeriod: Duration | undefined;
  vestingPoolName: string;
  campaignTotalAmount: Coin[];
  campaignCurrentAmount: Coin[];
}

function createBaseCampaign(): Campaign {
  return {
    id: 0,
    owner: "",
    name: "",
    description: "",
    campaignType: 0,
    removableClaimRecords: false,
    feegrantAmount: "",
    initialClaimFreeAmount: "",
    free: "",
    enabled: false,
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
    vestingPoolName: "",
    campaignTotalAmount: [],
    campaignCurrentAmount: [],
  };
}

export const Campaign = {
  encode(message: Campaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.campaignType !== 0) {
      writer.uint32(40).int32(message.campaignType);
    }
    if (message.removableClaimRecords === true) {
      writer.uint32(48).bool(message.removableClaimRecords);
    }
    if (message.feegrantAmount !== "") {
      writer.uint32(58).string(message.feegrantAmount);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(66).string(message.initialClaimFreeAmount);
    }
    if (message.free !== "") {
      writer.uint32(74).string(message.free);
    }
    if (message.enabled === true) {
      writer.uint32(80).bool(message.enabled);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(90).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(98).fork()).ldelim();
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(106).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(114).fork()).ldelim();
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(122).string(message.vestingPoolName);
    }
    for (const v of message.campaignTotalAmount) {
      Coin.encode(v!, writer.uint32(130).fork()).ldelim();
    }
    for (const v of message.campaignCurrentAmount) {
      Coin.encode(v!, writer.uint32(138).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Campaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.campaignType = reader.int32() as any;
          break;
        case 6:
          message.removableClaimRecords = reader.bool();
          break;
        case 7:
          message.feegrantAmount = reader.string();
          break;
        case 8:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 9:
          message.free = reader.string();
          break;
        case 10:
          message.enabled = reader.bool();
          break;
        case 11:
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 12:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 13:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 14:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 15:
          message.vestingPoolName = reader.string();
          break;
        case 16:
          message.campaignTotalAmount.push(Coin.decode(reader, reader.uint32()));
          break;
        case 17:
          message.campaignCurrentAmount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Campaign {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      campaignType: isSet(object.campaignType) ? campaignTypeFromJSON(object.campaignType) : 0,
      removableClaimRecords: isSet(object.removableClaimRecords) ? Boolean(object.removableClaimRecords) : false,
      feegrantAmount: isSet(object.feegrantAmount) ? String(object.feegrantAmount) : "",
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      free: isSet(object.free) ? String(object.free) : "",
      enabled: isSet(object.enabled) ? Boolean(object.enabled) : false,
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      campaignTotalAmount: Array.isArray(object?.campaignTotalAmount)
        ? object.campaignTotalAmount.map((e: any) => Coin.fromJSON(e))
        : [],
      campaignCurrentAmount: Array.isArray(object?.campaignCurrentAmount)
        ? object.campaignCurrentAmount.map((e: any) => Coin.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Campaign): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.campaignType !== undefined && (obj.campaignType = campaignTypeToJSON(message.campaignType));
    message.removableClaimRecords !== undefined && (obj.removableClaimRecords = message.removableClaimRecords);
    message.feegrantAmount !== undefined && (obj.feegrantAmount = message.feegrantAmount);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.free !== undefined && (obj.free = message.free);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    if (message.campaignTotalAmount) {
      obj.campaignTotalAmount = message.campaignTotalAmount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.campaignTotalAmount = [];
    }
    if (message.campaignCurrentAmount) {
      obj.campaignCurrentAmount = message.campaignCurrentAmount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.campaignCurrentAmount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Campaign>, I>>(object: I): Campaign {
    const message = createBaseCampaign();
    message.id = object.id ?? 0;
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.campaignType = object.campaignType ?? 0;
    message.removableClaimRecords = object.removableClaimRecords ?? false;
    message.feegrantAmount = object.feegrantAmount ?? "";
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.free = object.free ?? "";
    message.enabled = object.enabled ?? false;
    message.startTime = object.startTime ?? undefined;
    message.endTime = object.endTime ?? undefined;
    message.lockupPeriod = (object.lockupPeriod !== undefined && object.lockupPeriod !== null)
      ? Duration.fromPartial(object.lockupPeriod)
      : undefined;
    message.vestingPeriod = (object.vestingPeriod !== undefined && object.vestingPeriod !== null)
      ? Duration.fromPartial(object.vestingPeriod)
      : undefined;
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.campaignTotalAmount = object.campaignTotalAmount?.map((e) => Coin.fromPartial(e)) || [];
    message.campaignCurrentAmount = object.campaignCurrentAmount?.map((e) => Coin.fromPartial(e)) || [];
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
