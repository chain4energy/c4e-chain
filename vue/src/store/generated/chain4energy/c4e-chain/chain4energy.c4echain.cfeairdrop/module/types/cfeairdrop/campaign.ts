/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Duration } from "../google/protobuf/duration";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export enum CampaignType {
  CAMPAIGN_TYPE_UNSPECIFIED = 0,
  DYNAMIC = 1,
  DEFAULT = 2,
  VESTING_POOL = 3,
  UNRECOGNIZED = -1,
}

export function campaignTypeFromJSON(object: any): CampaignType {
  switch (object) {
    case 0:
    case "CAMPAIGN_TYPE_UNSPECIFIED":
      return CampaignType.CAMPAIGN_TYPE_UNSPECIFIED;
    case 1:
    case "DYNAMIC":
      return CampaignType.DYNAMIC;
    case 2:
    case "DEFAULT":
      return CampaignType.DEFAULT;
    case 3:
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
    case CampaignType.DYNAMIC:
      return "DYNAMIC";
    case CampaignType.DEFAULT:
      return "DEFAULT";
    case CampaignType.VESTING_POOL:
      return "VESTING_POOL";
    default:
      return "UNKNOWN";
  }
}

/** Campaign close action */
export enum CloseAction {
  /** CLOSE_ACTION_UNSPECIFIED - Campaign close action */
  CLOSE_ACTION_UNSPECIFIED = 0,
  SEND_TO_COMMUNITY_POOL = 1,
  BURN = 2,
  SEND_TO_OWNER = 3,
  UNRECOGNIZED = -1,
}

export function CloseActionFromJSON(object: any): CloseAction {
  switch (object) {
    case 0:
    case "CLOSE_ACTION_UNSPECIFIED":
      return CloseAction.CLOSE_ACTION_UNSPECIFIED;
    case 1:
    case "SEND_TO_COMMUNITY_POOL":
      return CloseAction.SEND_TO_COMMUNITY_POOL;
    case 2:
    case "BURN":
      return CloseAction.BURN;
    case 3:
    case "SEND_TO_OWNER":
      return CloseAction.SEND_TO_OWNER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CloseAction.UNRECOGNIZED;
  }
}

export function CloseActionToJSON(object: CloseAction): string {
  switch (object) {
    case CloseAction.CLOSE_ACTION_UNSPECIFIED:
      return "CLOSE_ACTION_UNSPECIFIED";
    case CloseAction.SEND_TO_COMMUNITY_POOL:
      return "SEND_TO_COMMUNITY_POOL";
    case CloseAction.BURN:
      return "BURN";
    case CloseAction.SEND_TO_OWNER:
      return "SEND_TO_OWNER";
    default:
      return "UNKNOWN";
  }
}

export interface Campaign {
  id: number;
  owner: string;
  name: string;
  description: string;
  campaignType: CampaignType;
  feegrant_amount: string;
  initial_claim_free_amount: string;
  enabled: boolean;
  start_time: Date | undefined;
  end_time: Date | undefined;
  /** period of locked coins from claim */
  lockup_period: Duration | undefined;
  /** period of vesting coins after lockup period */
  vesting_period: Duration | undefined;
}

export interface CampaignTotalAmount {
  campaign_id: number;
  amount: Coin[];
}

export interface CampaignCurrentAmount {
  campaign_id: number;
  amount: Coin[];
}

const baseCampaign: object = {
  id: 0,
  owner: "",
  name: "",
  description: "",
  campaignType: 0,
  feegrant_amount: "",
  initial_claim_free_amount: "",
  enabled: false,
};

export const Campaign = {
  encode(message: Campaign, writer: Writer = Writer.create()): Writer {
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
    if (message.feegrant_amount !== "") {
      writer.uint32(50).string(message.feegrant_amount);
    }
    if (message.initial_claim_free_amount !== "") {
      writer.uint32(58).string(message.initial_claim_free_amount);
    }
    if (message.enabled === true) {
      writer.uint32(64).bool(message.enabled);
    }
    if (message.start_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.start_time),
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.end_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.end_time),
        writer.uint32(82).fork()
      ).ldelim();
    }
    if (message.lockup_period !== undefined) {
      Duration.encode(message.lockup_period, writer.uint32(90).fork()).ldelim();
    }
    if (message.vesting_period !== undefined) {
      Duration.encode(
        message.vesting_period,
        writer.uint32(98).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Campaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaign } as Campaign;
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
          message.feegrant_amount = reader.string();
          break;
        case 7:
          message.initial_claim_free_amount = reader.string();
          break;
        case 8:
          message.enabled = reader.bool();
          break;
        case 9:
          message.start_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.end_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 11:
          message.lockup_period = Duration.decode(reader, reader.uint32());
          break;
        case 12:
          message.vesting_period = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Campaign {
    const message = { ...baseCampaign } as Campaign;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.campaignType !== undefined && object.campaignType !== null) {
      message.campaignType = campaignTypeFromJSON(object.campaignType);
    } else {
      message.campaignType = 0;
    }
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = String(object.feegrant_amount);
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = String(
        object.initial_claim_free_amount
      );
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = Boolean(object.enabled);
    } else {
      message.enabled = false;
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = fromJsonTimestamp(object.start_time);
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = fromJsonTimestamp(object.end_time);
    } else {
      message.end_time = undefined;
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromJSON(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromJSON(object.vesting_period);
    } else {
      message.vesting_period = undefined;
    }
    return message;
  },

  toJSON(message: Campaign): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.campaignType !== undefined &&
      (obj.campaignType = campaignTypeToJSON(message.campaignType));
    message.feegrant_amount !== undefined &&
      (obj.feegrant_amount = message.feegrant_amount);
    message.initial_claim_free_amount !== undefined &&
      (obj.initial_claim_free_amount = message.initial_claim_free_amount);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.start_time !== undefined &&
      (obj.start_time =
        message.start_time !== undefined
          ? message.start_time.toISOString()
          : null);
    message.end_time !== undefined &&
      (obj.end_time =
        message.end_time !== undefined ? message.end_time.toISOString() : null);
    message.lockup_period !== undefined &&
      (obj.lockup_period = message.lockup_period
        ? Duration.toJSON(message.lockup_period)
        : undefined);
    message.vesting_period !== undefined &&
      (obj.vesting_period = message.vesting_period
        ? Duration.toJSON(message.vesting_period)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Campaign>): Campaign {
    const message = { ...baseCampaign } as Campaign;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.campaignType !== undefined && object.campaignType !== null) {
      message.campaignType = object.campaignType;
    } else {
      message.campaignType = 0;
    }
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = object.feegrant_amount;
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = object.initial_claim_free_amount;
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = object.enabled;
    } else {
      message.enabled = false;
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = undefined;
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromPartial(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromPartial(object.vesting_period);
    } else {
      message.vesting_period = undefined;
    }
    return message;
  },
};

const baseCampaignTotalAmount: object = { campaign_id: 0 };

export const CampaignTotalAmount = {
  encode(
    message: CampaignTotalAmount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CampaignTotalAmount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaignTotalAmount } as CampaignTotalAmount;
    message.amount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CampaignTotalAmount {
    const message = { ...baseCampaignTotalAmount } as CampaignTotalAmount;
    message.amount = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: CampaignTotalAmount): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<CampaignTotalAmount>): CampaignTotalAmount {
    const message = { ...baseCampaignTotalAmount } as CampaignTotalAmount;
    message.amount = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseCampaignCurrentAmount: object = { campaign_id: 0 };

export const CampaignCurrentAmount = {
  encode(
    message: CampaignCurrentAmount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CampaignCurrentAmount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaignCurrentAmount } as CampaignCurrentAmount;
    message.amount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CampaignCurrentAmount {
    const message = { ...baseCampaignCurrentAmount } as CampaignCurrentAmount;
    message.amount = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: CampaignCurrentAmount): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<CampaignCurrentAmount>): CampaignCurrentAmount {
    const message = { ...baseCampaignCurrentAmount } as CampaignCurrentAmount;
    message.amount = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
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

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
