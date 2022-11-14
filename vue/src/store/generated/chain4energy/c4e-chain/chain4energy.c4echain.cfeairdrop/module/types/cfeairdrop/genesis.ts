/* eslint-disable */
import { Params } from "../cfeairdrop/params";
import { ClaimRecord, InitialClaim } from "../cfeairdrop/airdrop";
import { Mission } from "../cfeairdrop/mission";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** GenesisState defines the cfeairdrop module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  claimRecords: ClaimRecord[];
  initialClaims: InitialClaim[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  missionList: Mission[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.claimRecords) {
      ClaimRecord.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.initialClaims) {
      InitialClaim.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.missionList) {
      Mission.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.claimRecords = [];
    message.initialClaims = [];
    message.missionList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.claimRecords.push(
            ClaimRecord.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.initialClaims.push(
            InitialClaim.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.missionList.push(Mission.decode(reader, reader.uint32()));
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
    message.claimRecords = [];
    message.initialClaims = [];
    message.missionList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.claimRecords !== undefined && object.claimRecords !== null) {
      for (const e of object.claimRecords) {
        message.claimRecords.push(ClaimRecord.fromJSON(e));
      }
    }
    if (object.initialClaims !== undefined && object.initialClaims !== null) {
      for (const e of object.initialClaims) {
        message.initialClaims.push(InitialClaim.fromJSON(e));
      }
    }
    if (object.missionList !== undefined && object.missionList !== null) {
      for (const e of object.missionList) {
        message.missionList.push(Mission.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.claimRecords) {
      obj.claimRecords = message.claimRecords.map((e) =>
        e ? ClaimRecord.toJSON(e) : undefined
      );
    } else {
      obj.claimRecords = [];
    }
    if (message.initialClaims) {
      obj.initialClaims = message.initialClaims.map((e) =>
        e ? InitialClaim.toJSON(e) : undefined
      );
    } else {
      obj.initialClaims = [];
    }
    if (message.missionList) {
      obj.missionList = message.missionList.map((e) =>
        e ? Mission.toJSON(e) : undefined
      );
    } else {
      obj.missionList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.claimRecords = [];
    message.initialClaims = [];
    message.missionList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.claimRecords !== undefined && object.claimRecords !== null) {
      for (const e of object.claimRecords) {
        message.claimRecords.push(ClaimRecord.fromPartial(e));
      }
    }
    if (object.initialClaims !== undefined && object.initialClaims !== null) {
      for (const e of object.initialClaims) {
        message.initialClaims.push(InitialClaim.fromPartial(e));
      }
    }
    if (object.missionList !== undefined && object.missionList !== null) {
      for (const e of object.missionList) {
        message.missionList.push(Mission.fromPartial(e));
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
