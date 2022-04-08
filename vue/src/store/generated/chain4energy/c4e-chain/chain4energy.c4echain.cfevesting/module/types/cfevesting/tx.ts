/* eslint-disable */
import {
  VoteOption,
  voteOptionFromJSON,
  voteOptionToJSON,
} from "../cosmos/gov/v1beta1/gov";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface MsgVest {
  creator: string;
  /** uint64 amount = 2; */
  amount: string;
  vesting_type: string;
}

export interface MsgVestResponse {}

export interface MsgWithdrawAllAvailable {
  creator: string;
}

export interface MsgWithdrawAllAvailableResponse {}

export interface MsgDelegate {
  delegator_address: string;
  validator_address: string;
  amount: Coin | undefined;
}

export interface MsgDelegateResponse {}

export interface MsgUndelegate {
  delegator_address: string;
  validator_address: string;
  amount: Coin | undefined;
}

export interface MsgUndelegateResponse {
  completion_time: Date | undefined;
}

export interface MsgBeginRedelegate {
  delegator_address: string;
  validator_src_address: string;
  validator_dst_address: string;
  amount: Coin | undefined;
}

export interface MsgBeginRedelegateResponse {
  completion_time: Date | undefined;
}

export interface MsgWithdrawDelegatorReward {
  delegator_address: string;
  validator_address: string;
}

export interface MsgWithdrawDelegatorRewardResponse {}

export interface MsgSendVesting {
  from_address: string;
  to_address: string;
  vesting_id: number;
  amount: string;
  restart_vesting: boolean;
}

export interface MsgSendVestingResponse {}

export interface MsgVote {
  proposal_id: number;
  voter: string;
  option: VoteOption;
}

export interface MsgVoteResponse {}

export interface MsgVoteWeighted {
  voter: string;
  proposalId: string;
  options: string;
}

export interface MsgVoteWeightedResponse {}

const baseMsgVest: object = { creator: "", amount: "", vesting_type: "" };

export const MsgVest = {
  encode(message: MsgVest, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.vesting_type !== "") {
      writer.uint32(26).string(message.vesting_type);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVest } as MsgVest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.vesting_type = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVest {
    const message = { ...baseMsgVest } as MsgVest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = String(object.vesting_type);
    } else {
      message.vesting_type = "";
    }
    return message;
  },

  toJSON(message: MsgVest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.vesting_type !== undefined &&
      (obj.vesting_type = message.vesting_type);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVest>): MsgVest {
    const message = { ...baseMsgVest } as MsgVest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = object.vesting_type;
    } else {
      message.vesting_type = "";
    }
    return message;
  },
};

const baseMsgVestResponse: object = {};

export const MsgVestResponse = {
  encode(_: MsgVestResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVestResponse } as MsgVestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgVestResponse {
    const message = { ...baseMsgVestResponse } as MsgVestResponse;
    return message;
  },

  toJSON(_: MsgVestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgVestResponse>): MsgVestResponse {
    const message = { ...baseMsgVestResponse } as MsgVestResponse;
    return message;
  },
};

const baseMsgWithdrawAllAvailable: object = { creator: "" };

export const MsgWithdrawAllAvailable = {
  encode(
    message: MsgWithdrawAllAvailable,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawAllAvailable {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawAllAvailable {
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawAllAvailable): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgWithdrawAllAvailable>
  ): MsgWithdrawAllAvailable {
    const message = {
      ...baseMsgWithdrawAllAvailable,
    } as MsgWithdrawAllAvailable;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgWithdrawAllAvailableResponse: object = {};

export const MsgWithdrawAllAvailableResponse = {
  encode(
    _: MsgWithdrawAllAvailableResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawAllAvailableResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawAllAvailableResponse {
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    return message;
  },

  toJSON(_: MsgWithdrawAllAvailableResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawAllAvailableResponse>
  ): MsgWithdrawAllAvailableResponse {
    const message = {
      ...baseMsgWithdrawAllAvailableResponse,
    } as MsgWithdrawAllAvailableResponse;
    return message;
  },
};

const baseMsgDelegate: object = {
  delegator_address: "",
  validator_address: "",
};

export const MsgDelegate = {
  encode(message: MsgDelegate, writer: Writer = Writer.create()): Writer {
    if (message.delegator_address !== "") {
      writer.uint32(10).string(message.delegator_address);
    }
    if (message.validator_address !== "") {
      writer.uint32(18).string(message.validator_address);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDelegate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDelegate } as MsgDelegate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegator_address = reader.string();
          break;
        case 2:
          message.validator_address = reader.string();
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDelegate {
    const message = { ...baseMsgDelegate } as MsgDelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = String(object.delegator_address);
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },

  toJSON(message: MsgDelegate): unknown {
    const obj: any = {};
    message.delegator_address !== undefined &&
      (obj.delegator_address = message.delegator_address);
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDelegate>): MsgDelegate {
    const message = { ...baseMsgDelegate } as MsgDelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = object.delegator_address;
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },
};

const baseMsgDelegateResponse: object = {};

export const MsgDelegateResponse = {
  encode(_: MsgDelegateResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDelegateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDelegateResponse } as MsgDelegateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDelegateResponse {
    const message = { ...baseMsgDelegateResponse } as MsgDelegateResponse;
    return message;
  },

  toJSON(_: MsgDelegateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDelegateResponse>): MsgDelegateResponse {
    const message = { ...baseMsgDelegateResponse } as MsgDelegateResponse;
    return message;
  },
};

const baseMsgUndelegate: object = {
  delegator_address: "",
  validator_address: "",
};

export const MsgUndelegate = {
  encode(message: MsgUndelegate, writer: Writer = Writer.create()): Writer {
    if (message.delegator_address !== "") {
      writer.uint32(10).string(message.delegator_address);
    }
    if (message.validator_address !== "") {
      writer.uint32(18).string(message.validator_address);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUndelegate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUndelegate } as MsgUndelegate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegator_address = reader.string();
          break;
        case 2:
          message.validator_address = reader.string();
          break;
        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUndelegate {
    const message = { ...baseMsgUndelegate } as MsgUndelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = String(object.delegator_address);
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },

  toJSON(message: MsgUndelegate): unknown {
    const obj: any = {};
    message.delegator_address !== undefined &&
      (obj.delegator_address = message.delegator_address);
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUndelegate>): MsgUndelegate {
    const message = { ...baseMsgUndelegate } as MsgUndelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = object.delegator_address;
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },
};

const baseMsgUndelegateResponse: object = {};

export const MsgUndelegateResponse = {
  encode(
    message: MsgUndelegateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.completion_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.completion_time),
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUndelegateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUndelegateResponse } as MsgUndelegateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.completion_time = fromTimestamp(
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

  fromJSON(object: any): MsgUndelegateResponse {
    const message = { ...baseMsgUndelegateResponse } as MsgUndelegateResponse;
    if (
      object.completion_time !== undefined &&
      object.completion_time !== null
    ) {
      message.completion_time = fromJsonTimestamp(object.completion_time);
    } else {
      message.completion_time = undefined;
    }
    return message;
  },

  toJSON(message: MsgUndelegateResponse): unknown {
    const obj: any = {};
    message.completion_time !== undefined &&
      (obj.completion_time =
        message.completion_time !== undefined
          ? message.completion_time.toISOString()
          : null);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUndelegateResponse>
  ): MsgUndelegateResponse {
    const message = { ...baseMsgUndelegateResponse } as MsgUndelegateResponse;
    if (
      object.completion_time !== undefined &&
      object.completion_time !== null
    ) {
      message.completion_time = object.completion_time;
    } else {
      message.completion_time = undefined;
    }
    return message;
  },
};

const baseMsgBeginRedelegate: object = {
  delegator_address: "",
  validator_src_address: "",
  validator_dst_address: "",
};

export const MsgBeginRedelegate = {
  encode(
    message: MsgBeginRedelegate,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegator_address !== "") {
      writer.uint32(10).string(message.delegator_address);
    }
    if (message.validator_src_address !== "") {
      writer.uint32(18).string(message.validator_src_address);
    }
    if (message.validator_dst_address !== "") {
      writer.uint32(26).string(message.validator_dst_address);
    }
    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgBeginRedelegate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgBeginRedelegate } as MsgBeginRedelegate;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegator_address = reader.string();
          break;
        case 2:
          message.validator_src_address = reader.string();
          break;
        case 3:
          message.validator_dst_address = reader.string();
          break;
        case 4:
          message.amount = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBeginRedelegate {
    const message = { ...baseMsgBeginRedelegate } as MsgBeginRedelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = String(object.delegator_address);
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_src_address !== undefined &&
      object.validator_src_address !== null
    ) {
      message.validator_src_address = String(object.validator_src_address);
    } else {
      message.validator_src_address = "";
    }
    if (
      object.validator_dst_address !== undefined &&
      object.validator_dst_address !== null
    ) {
      message.validator_dst_address = String(object.validator_dst_address);
    } else {
      message.validator_dst_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromJSON(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },

  toJSON(message: MsgBeginRedelegate): unknown {
    const obj: any = {};
    message.delegator_address !== undefined &&
      (obj.delegator_address = message.delegator_address);
    message.validator_src_address !== undefined &&
      (obj.validator_src_address = message.validator_src_address);
    message.validator_dst_address !== undefined &&
      (obj.validator_dst_address = message.validator_dst_address);
    message.amount !== undefined &&
      (obj.amount = message.amount ? Coin.toJSON(message.amount) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgBeginRedelegate>): MsgBeginRedelegate {
    const message = { ...baseMsgBeginRedelegate } as MsgBeginRedelegate;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = object.delegator_address;
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_src_address !== undefined &&
      object.validator_src_address !== null
    ) {
      message.validator_src_address = object.validator_src_address;
    } else {
      message.validator_src_address = "";
    }
    if (
      object.validator_dst_address !== undefined &&
      object.validator_dst_address !== null
    ) {
      message.validator_dst_address = object.validator_dst_address;
    } else {
      message.validator_dst_address = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Coin.fromPartial(object.amount);
    } else {
      message.amount = undefined;
    }
    return message;
  },
};

const baseMsgBeginRedelegateResponse: object = {};

export const MsgBeginRedelegateResponse = {
  encode(
    message: MsgBeginRedelegateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.completion_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.completion_time),
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgBeginRedelegateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgBeginRedelegateResponse,
    } as MsgBeginRedelegateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.completion_time = fromTimestamp(
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

  fromJSON(object: any): MsgBeginRedelegateResponse {
    const message = {
      ...baseMsgBeginRedelegateResponse,
    } as MsgBeginRedelegateResponse;
    if (
      object.completion_time !== undefined &&
      object.completion_time !== null
    ) {
      message.completion_time = fromJsonTimestamp(object.completion_time);
    } else {
      message.completion_time = undefined;
    }
    return message;
  },

  toJSON(message: MsgBeginRedelegateResponse): unknown {
    const obj: any = {};
    message.completion_time !== undefined &&
      (obj.completion_time =
        message.completion_time !== undefined
          ? message.completion_time.toISOString()
          : null);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgBeginRedelegateResponse>
  ): MsgBeginRedelegateResponse {
    const message = {
      ...baseMsgBeginRedelegateResponse,
    } as MsgBeginRedelegateResponse;
    if (
      object.completion_time !== undefined &&
      object.completion_time !== null
    ) {
      message.completion_time = object.completion_time;
    } else {
      message.completion_time = undefined;
    }
    return message;
  },
};

const baseMsgWithdrawDelegatorReward: object = {
  delegator_address: "",
  validator_address: "",
};

export const MsgWithdrawDelegatorReward = {
  encode(
    message: MsgWithdrawDelegatorReward,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.delegator_address !== "") {
      writer.uint32(10).string(message.delegator_address);
    }
    if (message.validator_address !== "") {
      writer.uint32(18).string(message.validator_address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawDelegatorReward {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawDelegatorReward,
    } as MsgWithdrawDelegatorReward;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.delegator_address = reader.string();
          break;
        case 2:
          message.validator_address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawDelegatorReward {
    const message = {
      ...baseMsgWithdrawDelegatorReward,
    } as MsgWithdrawDelegatorReward;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = String(object.delegator_address);
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = String(object.validator_address);
    } else {
      message.validator_address = "";
    }
    return message;
  },

  toJSON(message: MsgWithdrawDelegatorReward): unknown {
    const obj: any = {};
    message.delegator_address !== undefined &&
      (obj.delegator_address = message.delegator_address);
    message.validator_address !== undefined &&
      (obj.validator_address = message.validator_address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgWithdrawDelegatorReward>
  ): MsgWithdrawDelegatorReward {
    const message = {
      ...baseMsgWithdrawDelegatorReward,
    } as MsgWithdrawDelegatorReward;
    if (
      object.delegator_address !== undefined &&
      object.delegator_address !== null
    ) {
      message.delegator_address = object.delegator_address;
    } else {
      message.delegator_address = "";
    }
    if (
      object.validator_address !== undefined &&
      object.validator_address !== null
    ) {
      message.validator_address = object.validator_address;
    } else {
      message.validator_address = "";
    }
    return message;
  },
};

const baseMsgWithdrawDelegatorRewardResponse: object = {};

export const MsgWithdrawDelegatorRewardResponse = {
  encode(
    _: MsgWithdrawDelegatorRewardResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawDelegatorRewardResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawDelegatorRewardResponse,
    } as MsgWithdrawDelegatorRewardResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawDelegatorRewardResponse {
    const message = {
      ...baseMsgWithdrawDelegatorRewardResponse,
    } as MsgWithdrawDelegatorRewardResponse;
    return message;
  },

  toJSON(_: MsgWithdrawDelegatorRewardResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawDelegatorRewardResponse>
  ): MsgWithdrawDelegatorRewardResponse {
    const message = {
      ...baseMsgWithdrawDelegatorRewardResponse,
    } as MsgWithdrawDelegatorRewardResponse;
    return message;
  },
};

const baseMsgSendVesting: object = {
  from_address: "",
  to_address: "",
  vesting_id: 0,
  amount: "",
  restart_vesting: false,
};

export const MsgSendVesting = {
  encode(message: MsgSendVesting, writer: Writer = Writer.create()): Writer {
    if (message.from_address !== "") {
      writer.uint32(10).string(message.from_address);
    }
    if (message.to_address !== "") {
      writer.uint32(18).string(message.to_address);
    }
    if (message.vesting_id !== 0) {
      writer.uint32(24).int32(message.vesting_id);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restart_vesting === true) {
      writer.uint32(40).bool(message.restart_vesting);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendVesting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendVesting } as MsgSendVesting;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.from_address = reader.string();
          break;
        case 2:
          message.to_address = reader.string();
          break;
        case 3:
          message.vesting_id = reader.int32();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.restart_vesting = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendVesting {
    const message = { ...baseMsgSendVesting } as MsgSendVesting;
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = String(object.from_address);
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = String(object.to_address);
    } else {
      message.to_address = "";
    }
    if (object.vesting_id !== undefined && object.vesting_id !== null) {
      message.vesting_id = Number(object.vesting_id);
    } else {
      message.vesting_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (
      object.restart_vesting !== undefined &&
      object.restart_vesting !== null
    ) {
      message.restart_vesting = Boolean(object.restart_vesting);
    } else {
      message.restart_vesting = false;
    }
    return message;
  },

  toJSON(message: MsgSendVesting): unknown {
    const obj: any = {};
    message.from_address !== undefined &&
      (obj.from_address = message.from_address);
    message.to_address !== undefined && (obj.to_address = message.to_address);
    message.vesting_id !== undefined && (obj.vesting_id = message.vesting_id);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restart_vesting !== undefined &&
      (obj.restart_vesting = message.restart_vesting);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendVesting>): MsgSendVesting {
    const message = { ...baseMsgSendVesting } as MsgSendVesting;
    if (object.from_address !== undefined && object.from_address !== null) {
      message.from_address = object.from_address;
    } else {
      message.from_address = "";
    }
    if (object.to_address !== undefined && object.to_address !== null) {
      message.to_address = object.to_address;
    } else {
      message.to_address = "";
    }
    if (object.vesting_id !== undefined && object.vesting_id !== null) {
      message.vesting_id = object.vesting_id;
    } else {
      message.vesting_id = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (
      object.restart_vesting !== undefined &&
      object.restart_vesting !== null
    ) {
      message.restart_vesting = object.restart_vesting;
    } else {
      message.restart_vesting = false;
    }
    return message;
  },
};

const baseMsgSendVestingResponse: object = {};

export const MsgSendVestingResponse = {
  encode(_: MsgSendVestingResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendVestingResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendVestingResponse } as MsgSendVestingResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSendVestingResponse {
    const message = { ...baseMsgSendVestingResponse } as MsgSendVestingResponse;
    return message;
  },

  toJSON(_: MsgSendVestingResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSendVestingResponse>): MsgSendVestingResponse {
    const message = { ...baseMsgSendVestingResponse } as MsgSendVestingResponse;
    return message;
  },
};

const baseMsgVote: object = { proposal_id: 0, voter: "", option: 0 };

export const MsgVote = {
  encode(message: MsgVote, writer: Writer = Writer.create()): Writer {
    if (message.proposal_id !== 0) {
      writer.uint32(8).uint64(message.proposal_id);
    }
    if (message.voter !== "") {
      writer.uint32(18).string(message.voter);
    }
    if (message.option !== 0) {
      writer.uint32(24).int32(message.option);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVote {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVote } as MsgVote;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposal_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.voter = reader.string();
          break;
        case 3:
          message.option = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVote {
    const message = { ...baseMsgVote } as MsgVote;
    if (object.proposal_id !== undefined && object.proposal_id !== null) {
      message.proposal_id = Number(object.proposal_id);
    } else {
      message.proposal_id = 0;
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    if (object.option !== undefined && object.option !== null) {
      message.option = voteOptionFromJSON(object.option);
    } else {
      message.option = 0;
    }
    return message;
  },

  toJSON(message: MsgVote): unknown {
    const obj: any = {};
    message.proposal_id !== undefined &&
      (obj.proposal_id = message.proposal_id);
    message.voter !== undefined && (obj.voter = message.voter);
    message.option !== undefined &&
      (obj.option = voteOptionToJSON(message.option));
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVote>): MsgVote {
    const message = { ...baseMsgVote } as MsgVote;
    if (object.proposal_id !== undefined && object.proposal_id !== null) {
      message.proposal_id = object.proposal_id;
    } else {
      message.proposal_id = 0;
    }
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    if (object.option !== undefined && object.option !== null) {
      message.option = object.option;
    } else {
      message.option = 0;
    }
    return message;
  },
};

const baseMsgVoteResponse: object = {};

export const MsgVoteResponse = {
  encode(_: MsgVoteResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgVoteResponse {
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    return message;
  },

  toJSON(_: MsgVoteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgVoteResponse>): MsgVoteResponse {
    const message = { ...baseMsgVoteResponse } as MsgVoteResponse;
    return message;
  },
};

const baseMsgVoteWeighted: object = { voter: "", proposalId: "", options: "" };

export const MsgVoteWeighted = {
  encode(message: MsgVoteWeighted, writer: Writer = Writer.create()): Writer {
    if (message.voter !== "") {
      writer.uint32(10).string(message.voter);
    }
    if (message.proposalId !== "") {
      writer.uint32(18).string(message.proposalId);
    }
    if (message.options !== "") {
      writer.uint32(26).string(message.options);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteWeighted {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoteWeighted } as MsgVoteWeighted;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.voter = reader.string();
          break;
        case 2:
          message.proposalId = reader.string();
          break;
        case 3:
          message.options = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgVoteWeighted {
    const message = { ...baseMsgVoteWeighted } as MsgVoteWeighted;
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = String(object.voter);
    } else {
      message.voter = "";
    }
    if (object.proposalId !== undefined && object.proposalId !== null) {
      message.proposalId = String(object.proposalId);
    } else {
      message.proposalId = "";
    }
    if (object.options !== undefined && object.options !== null) {
      message.options = String(object.options);
    } else {
      message.options = "";
    }
    return message;
  },

  toJSON(message: MsgVoteWeighted): unknown {
    const obj: any = {};
    message.voter !== undefined && (obj.voter = message.voter);
    message.proposalId !== undefined && (obj.proposalId = message.proposalId);
    message.options !== undefined && (obj.options = message.options);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVoteWeighted>): MsgVoteWeighted {
    const message = { ...baseMsgVoteWeighted } as MsgVoteWeighted;
    if (object.voter !== undefined && object.voter !== null) {
      message.voter = object.voter;
    } else {
      message.voter = "";
    }
    if (object.proposalId !== undefined && object.proposalId !== null) {
      message.proposalId = object.proposalId;
    } else {
      message.proposalId = "";
    }
    if (object.options !== undefined && object.options !== null) {
      message.options = object.options;
    } else {
      message.options = "";
    }
    return message;
  },
};

const baseMsgVoteWeightedResponse: object = {};

export const MsgVoteWeightedResponse = {
  encode(_: MsgVoteWeightedResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoteWeightedResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgVoteWeightedResponse,
    } as MsgVoteWeightedResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgVoteWeightedResponse {
    const message = {
      ...baseMsgVoteWeightedResponse,
    } as MsgVoteWeightedResponse;
    return message;
  },

  toJSON(_: MsgVoteWeightedResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgVoteWeightedResponse>
  ): MsgVoteWeightedResponse {
    const message = {
      ...baseMsgVoteWeightedResponse,
    } as MsgVoteWeightedResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Vest(request: MsgVest): Promise<MsgVestResponse>;
  WithdrawAllAvailable(
    request: MsgWithdrawAllAvailable
  ): Promise<MsgWithdrawAllAvailableResponse>;
  Delegate(request: MsgDelegate): Promise<MsgDelegateResponse>;
  Undelegate(request: MsgUndelegate): Promise<MsgUndelegateResponse>;
  BeginRedelegate(
    request: MsgBeginRedelegate
  ): Promise<MsgBeginRedelegateResponse>;
  WithdrawDelegatorReward(
    request: MsgWithdrawDelegatorReward
  ): Promise<MsgWithdrawDelegatorRewardResponse>;
  SendVesting(request: MsgSendVesting): Promise<MsgSendVestingResponse>;
  Vote(request: MsgVote): Promise<MsgVoteResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  VoteWeighted(request: MsgVoteWeighted): Promise<MsgVoteWeightedResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Vest(request: MsgVest): Promise<MsgVestResponse> {
    const data = MsgVest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "Vest",
      data
    );
    return promise.then((data) => MsgVestResponse.decode(new Reader(data)));
  }

  WithdrawAllAvailable(
    request: MsgWithdrawAllAvailable
  ): Promise<MsgWithdrawAllAvailableResponse> {
    const data = MsgWithdrawAllAvailable.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "WithdrawAllAvailable",
      data
    );
    return promise.then((data) =>
      MsgWithdrawAllAvailableResponse.decode(new Reader(data))
    );
  }

  Delegate(request: MsgDelegate): Promise<MsgDelegateResponse> {
    const data = MsgDelegate.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "Delegate",
      data
    );
    return promise.then((data) => MsgDelegateResponse.decode(new Reader(data)));
  }

  Undelegate(request: MsgUndelegate): Promise<MsgUndelegateResponse> {
    const data = MsgUndelegate.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "Undelegate",
      data
    );
    return promise.then((data) =>
      MsgUndelegateResponse.decode(new Reader(data))
    );
  }

  BeginRedelegate(
    request: MsgBeginRedelegate
  ): Promise<MsgBeginRedelegateResponse> {
    const data = MsgBeginRedelegate.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "BeginRedelegate",
      data
    );
    return promise.then((data) =>
      MsgBeginRedelegateResponse.decode(new Reader(data))
    );
  }

  WithdrawDelegatorReward(
    request: MsgWithdrawDelegatorReward
  ): Promise<MsgWithdrawDelegatorRewardResponse> {
    const data = MsgWithdrawDelegatorReward.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "WithdrawDelegatorReward",
      data
    );
    return promise.then((data) =>
      MsgWithdrawDelegatorRewardResponse.decode(new Reader(data))
    );
  }

  SendVesting(request: MsgSendVesting): Promise<MsgSendVestingResponse> {
    const data = MsgSendVesting.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "SendVesting",
      data
    );
    return promise.then((data) =>
      MsgSendVestingResponse.decode(new Reader(data))
    );
  }

  Vote(request: MsgVote): Promise<MsgVoteResponse> {
    const data = MsgVote.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "Vote",
      data
    );
    return promise.then((data) => MsgVoteResponse.decode(new Reader(data)));
  }

  VoteWeighted(request: MsgVoteWeighted): Promise<MsgVoteWeightedResponse> {
    const data = MsgVoteWeighted.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfevesting.Msg",
      "VoteWeighted",
      data
    );
    return promise.then((data) =>
      MsgVoteWeightedResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
