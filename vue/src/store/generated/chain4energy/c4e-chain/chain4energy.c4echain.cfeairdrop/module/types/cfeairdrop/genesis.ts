/* eslint-disable */
import { Params } from "../cfeairdrop/params";
import { ClaimRecord } from "../cfeairdrop/airdrop";
import { InitialClaim } from "../cfeairdrop/initial_claim";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

/** GenesisState defines the cfeairdrop module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  claimRecordList: ClaimRecord[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  initialClaimList: InitialClaim[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.claimRecordList) {
      ClaimRecord.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.initialClaimList) {
      InitialClaim.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.claimRecordList = [];
    message.initialClaimList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.claimRecordList.push(
            ClaimRecord.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.initialClaimList.push(
            InitialClaim.decode(reader, reader.uint32())
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
    message.claimRecordList = [];
    message.initialClaimList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.claimRecordList !== undefined &&
      object.claimRecordList !== null
    ) {
      for (const e of object.claimRecordList) {
        message.claimRecordList.push(ClaimRecord.fromJSON(e));
      }
    }
    if (
      object.initialClaimList !== undefined &&
      object.initialClaimList !== null
    ) {
      for (const e of object.initialClaimList) {
        message.initialClaimList.push(InitialClaim.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.claimRecordList) {
      obj.claimRecordList = message.claimRecordList.map((e) =>
        e ? ClaimRecord.toJSON(e) : undefined
      );
    } else {
      obj.claimRecordList = [];
    }
    if (message.initialClaimList) {
      obj.initialClaimList = message.initialClaimList.map((e) =>
        e ? InitialClaim.toJSON(e) : undefined
      );
    } else {
      obj.initialClaimList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.claimRecordList = [];
    message.initialClaimList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.claimRecordList !== undefined &&
      object.claimRecordList !== null
    ) {
      for (const e of object.claimRecordList) {
        message.claimRecordList.push(ClaimRecord.fromPartial(e));
      }
    }
    if (
      object.initialClaimList !== undefined &&
      object.initialClaimList !== null
    ) {
      for (const e of object.initialClaimList) {
        message.initialClaimList.push(InitialClaim.fromPartial(e));
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
