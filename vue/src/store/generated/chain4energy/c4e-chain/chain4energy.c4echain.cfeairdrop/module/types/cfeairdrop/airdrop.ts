/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Duration } from "../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export enum MissionType {
  MISSION_TYPE_UNSPECIFIED = 0,
  INITIAL_CLAIM = 1,
  DELEGATION = 2,
  VOTE = 3,
  CLAIM = 4,
  UNRECOGNIZED = -1,
}

export function missionTypeFromJSON(object: any): MissionType {
  switch (object) {
    case 0:
    case "MISSION_TYPE_UNSPECIFIED":
      return MissionType.MISSION_TYPE_UNSPECIFIED;
    case 1:
    case "INITIAL_CLAIM":
      return MissionType.INITIAL_CLAIM;
    case 2:
    case "DELEGATION":
      return MissionType.DELEGATION;
    case 3:
    case "VOTE":
      return MissionType.VOTE;
    case 4:
    case "CLAIM":
      return MissionType.CLAIM;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MissionType.UNRECOGNIZED;
  }
}

export function missionTypeToJSON(object: MissionType): string {
  switch (object) {
    case MissionType.MISSION_TYPE_UNSPECIFIED:
      return "MISSION_TYPE_UNSPECIFIED";
    case MissionType.INITIAL_CLAIM:
      return "INITIAL_CLAIM";
    case MissionType.DELEGATION:
      return "DELEGATION";
    case MissionType.VOTE:
      return "VOTE";
    case MissionType.CLAIM:
      return "CLAIM";
    default:
      return "UNKNOWN";
  }
}

/** Campaign close action */
export enum CampaignCloseAction {
  /** CLOSE_ACTION_UNSPECIFIED - Campaign close action */
  CLOSE_ACTION_UNSPECIFIED = 0,
  SEND_TO_COMMUNITY_POOL = 1,
  BURN = 2,
  SEND_TO_OWNER = 3,
  UNRECOGNIZED = -1,
}

export function campaignCloseActionFromJSON(object: any): CampaignCloseAction {
  switch (object) {
    case 0:
    case "CLOSE_ACTION_UNSPECIFIED":
      return CampaignCloseAction.CLOSE_ACTION_UNSPECIFIED;
    case 1:
    case "SEND_TO_COMMUNITY_POOL":
      return CampaignCloseAction.SEND_TO_COMMUNITY_POOL;
    case 2:
    case "BURN":
      return CampaignCloseAction.BURN;
    case 3:
    case "SEND_TO_OWNER":
      return CampaignCloseAction.SEND_TO_OWNER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CampaignCloseAction.UNRECOGNIZED;
  }
}

export function campaignCloseActionToJSON(object: CampaignCloseAction): string {
  switch (object) {
    case CampaignCloseAction.CLOSE_ACTION_UNSPECIFIED:
      return "CLOSE_ACTION_UNSPECIFIED";
    case CampaignCloseAction.SEND_TO_COMMUNITY_POOL:
      return "SEND_TO_COMMUNITY_POOL";
    case CampaignCloseAction.BURN:
      return "BURN";
    case CampaignCloseAction.SEND_TO_OWNER:
      return "SEND_TO_OWNER";
    default:
      return "UNKNOWN";
  }
}

export enum CampaignType {
  CAMPAIGN_TYPE_UNSPECIFIED = 0,
  TEAMDROP = 1,
  DEFAULT = 2,
  SALE = 3,
  UNRECOGNIZED = -1,
}

export function campaignTypeFromJSON(object: any): CampaignType {
  switch (object) {
    case 0:
    case "CAMPAIGN_TYPE_UNSPECIFIED":
      return CampaignType.CAMPAIGN_TYPE_UNSPECIFIED;
    case 1:
    case "TEAMDROP":
      return CampaignType.TEAMDROP;
    case 2:
    case "DEFAULT":
      return CampaignType.DEFAULT;
    case 3:
    case "SALE":
      return CampaignType.SALE;
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
    case CampaignType.TEAMDROP:
      return "TEAMDROP";
    case CampaignType.DEFAULT:
      return "DEFAULT";
    case CampaignType.SALE:
      return "SALE";
    default:
      return "UNKNOWN";
  }
}

export interface UserEntry {
  address: string;
  claim_address: string;
  claim_records: ClaimRecord[];
}

export interface ClaimRecord {
  campaign_id: number;
  address: string;
  amount: Coin[];
  completedMissions: number[];
  claimedMissions: number[];
}

export interface CampaignTotalAmount {
  campaign_id: number;
  amount: Coin[];
}

export interface CampaignAmountLeft {
  campaign_id: number;
  amount: Coin[];
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

export interface Mission {
  id: number;
  campaign_id: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
  claim_start_date: Date | undefined;
}

const baseUserEntry: object = { address: "", claim_address: "" };

export const UserEntry = {
  encode(message: UserEntry, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.claim_address !== "") {
      writer.uint32(18).string(message.claim_address);
    }
    for (const v of message.claim_records) {
      ClaimRecord.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UserEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserEntry } as UserEntry;
    message.claim_records = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.claim_address = reader.string();
          break;
        case 3:
          message.claim_records.push(
            ClaimRecord.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserEntry {
    const message = { ...baseUserEntry } as UserEntry;
    message.claim_records = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.claim_address !== undefined && object.claim_address !== null) {
      message.claim_address = String(object.claim_address);
    } else {
      message.claim_address = "";
    }
    if (object.claim_records !== undefined && object.claim_records !== null) {
      for (const e of object.claim_records) {
        message.claim_records.push(ClaimRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: UserEntry): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.claim_address !== undefined &&
      (obj.claim_address = message.claim_address);
    if (message.claim_records) {
      obj.claim_records = message.claim_records.map((e) =>
        e ? ClaimRecord.toJSON(e) : undefined
      );
    } else {
      obj.claim_records = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<UserEntry>): UserEntry {
    const message = { ...baseUserEntry } as UserEntry;
    message.claim_records = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.claim_address !== undefined && object.claim_address !== null) {
      message.claim_address = object.claim_address;
    } else {
      message.claim_address = "";
    }
    if (object.claim_records !== undefined && object.claim_records !== null) {
      for (const e of object.claim_records) {
        message.claim_records.push(ClaimRecord.fromPartial(e));
      }
    }
    return message;
  },
};

const baseClaimRecord: object = {
  campaign_id: 0,
  address: "",
  completedMissions: 0,
  claimedMissions: 0,
};

export const ClaimRecord = {
  encode(message: ClaimRecord, writer: Writer = Writer.create()): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    writer.uint32(34).fork();
    for (const v of message.completedMissions) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(42).fork();
    for (const v of message.claimedMissions) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ClaimRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.amount = [];
    message.completedMissions = [];
    message.claimedMissions = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.completedMissions.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.completedMissions.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.claimedMissions.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.claimedMissions.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimRecord {
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.amount = [];
    message.completedMissions = [];
    message.claimedMissions = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    if (
      object.completedMissions !== undefined &&
      object.completedMissions !== null
    ) {
      for (const e of object.completedMissions) {
        message.completedMissions.push(Number(e));
      }
    }
    if (
      object.claimedMissions !== undefined &&
      object.claimedMissions !== null
    ) {
      for (const e of object.claimedMissions) {
        message.claimedMissions.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: ClaimRecord): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.address !== undefined && (obj.address = message.address);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    if (message.completedMissions) {
      obj.completedMissions = message.completedMissions.map((e) => e);
    } else {
      obj.completedMissions = [];
    }
    if (message.claimedMissions) {
      obj.claimedMissions = message.claimedMissions.map((e) => e);
    } else {
      obj.claimedMissions = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ClaimRecord>): ClaimRecord {
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.amount = [];
    message.completedMissions = [];
    message.claimedMissions = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    if (
      object.completedMissions !== undefined &&
      object.completedMissions !== null
    ) {
      for (const e of object.completedMissions) {
        message.completedMissions.push(e);
      }
    }
    if (
      object.claimedMissions !== undefined &&
      object.claimedMissions !== null
    ) {
      for (const e of object.claimedMissions) {
        message.claimedMissions.push(e);
      }
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

const baseCampaignAmountLeft: object = { campaign_id: 0 };

export const CampaignAmountLeft = {
  encode(
    message: CampaignAmountLeft,
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

  decode(input: Reader | Uint8Array, length?: number): CampaignAmountLeft {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaignAmountLeft } as CampaignAmountLeft;
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

  fromJSON(object: any): CampaignAmountLeft {
    const message = { ...baseCampaignAmountLeft } as CampaignAmountLeft;
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

  toJSON(message: CampaignAmountLeft): unknown {
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

  fromPartial(object: DeepPartial<CampaignAmountLeft>): CampaignAmountLeft {
    const message = { ...baseCampaignAmountLeft } as CampaignAmountLeft;
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

const baseMission: object = {
  id: 0,
  campaign_id: 0,
  name: "",
  description: "",
  missionType: 0,
  weight: "",
};

export const Mission = {
  encode(message: Mission, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.missionType !== 0) {
      writer.uint32(40).int32(message.missionType);
    }
    if (message.weight !== "") {
      writer.uint32(50).string(message.weight);
    }
    if (message.claim_start_date !== undefined) {
      Timestamp.encode(
        toTimestamp(message.claim_start_date),
        writer.uint32(58).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Mission {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMission } as Mission;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.missionType = reader.int32() as any;
          break;
        case 6:
          message.weight = reader.string();
          break;
        case 7:
          message.claim_start_date = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Mission {
    const message = { ...baseMission } as Mission;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
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
    if (object.missionType !== undefined && object.missionType !== null) {
      message.missionType = missionTypeFromJSON(object.missionType);
    } else {
      message.missionType = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = String(object.weight);
    } else {
      message.weight = "";
    }
    if (
      object.claim_start_date !== undefined &&
      object.claim_start_date !== null
    ) {
      message.claim_start_date = fromJsonTimestamp(object.claim_start_date);
    } else {
      message.claim_start_date = undefined;
    }
    return message;
  },

  toJSON(message: Mission): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.missionType !== undefined &&
      (obj.missionType = missionTypeToJSON(message.missionType));
    message.weight !== undefined && (obj.weight = message.weight);
    message.claim_start_date !== undefined &&
      (obj.claim_start_date =
        message.claim_start_date !== undefined
          ? message.claim_start_date.toISOString()
          : null);
    return obj;
  },

  fromPartial(object: DeepPartial<Mission>): Mission {
    const message = { ...baseMission } as Mission;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
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
    if (object.missionType !== undefined && object.missionType !== null) {
      message.missionType = object.missionType;
    } else {
      message.missionType = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = object.weight;
    } else {
      message.weight = "";
    }
    if (
      object.claim_start_date !== undefined &&
      object.claim_start_date !== null
    ) {
      message.claim_start_date = object.claim_start_date;
    } else {
      message.claim_start_date = undefined;
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
