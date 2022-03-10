/* eslint-disable */
import { Params } from "../cfevesting/params";
import { VestingTypes } from "../cfevesting/vesting_types";
import { AccountVestingsList } from "../cfevesting/account_vesting";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

/** GenesisState defines the cfevesting module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  vestingTypes: VestingTypes | undefined;
  accountVestingsList: AccountVestingsList | undefined;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.vestingTypes !== undefined) {
      VestingTypes.encode(
        message.vestingTypes,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.accountVestingsList !== undefined) {
      AccountVestingsList.encode(
        message.accountVestingsList,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.vestingTypes = VestingTypes.decode(reader, reader.uint32());
          break;
        case 3:
          message.accountVestingsList = AccountVestingsList.decode(
            reader,
            reader.uint32()
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
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      message.vestingTypes = VestingTypes.fromJSON(object.vestingTypes);
    } else {
      message.vestingTypes = undefined;
    }
    if (
      object.accountVestingsList !== undefined &&
      object.accountVestingsList !== null
    ) {
      message.accountVestingsList = AccountVestingsList.fromJSON(
        object.accountVestingsList
      );
    } else {
      message.accountVestingsList = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.vestingTypes !== undefined &&
      (obj.vestingTypes = message.vestingTypes
        ? VestingTypes.toJSON(message.vestingTypes)
        : undefined);
    message.accountVestingsList !== undefined &&
      (obj.accountVestingsList = message.accountVestingsList
        ? AccountVestingsList.toJSON(message.accountVestingsList)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      message.vestingTypes = VestingTypes.fromPartial(object.vestingTypes);
    } else {
      message.vestingTypes = undefined;
    }
    if (
      object.accountVestingsList !== undefined &&
      object.accountVestingsList !== null
    ) {
      message.accountVestingsList = AccountVestingsList.fromPartial(
        object.accountVestingsList
      );
    } else {
      message.accountVestingsList = undefined;
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
