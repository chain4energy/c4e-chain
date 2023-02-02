/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

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
