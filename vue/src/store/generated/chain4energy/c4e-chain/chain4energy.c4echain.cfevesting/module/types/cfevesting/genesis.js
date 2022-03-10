/* eslint-disable */
import { Params } from "../cfevesting/params";
import { VestingTypes } from "../cfevesting/vesting_types";
import { AccountVestingsList } from "../cfevesting/account_vesting";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "chain4energy.c4echain.cfevesting";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        if (message.vestingTypes !== undefined) {
            VestingTypes.encode(message.vestingTypes, writer.uint32(18).fork()).ldelim();
        }
        if (message.accountVestingsList !== undefined) {
            AccountVestingsList.encode(message.accountVestingsList, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
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
                    message.accountVestingsList = AccountVestingsList.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            message.vestingTypes = VestingTypes.fromJSON(object.vestingTypes);
        }
        else {
            message.vestingTypes = undefined;
        }
        if (object.accountVestingsList !== undefined &&
            object.accountVestingsList !== null) {
            message.accountVestingsList = AccountVestingsList.fromJSON(object.accountVestingsList);
        }
        else {
            message.accountVestingsList = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = { ...baseGenesisState };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            message.vestingTypes = VestingTypes.fromPartial(object.vestingTypes);
        }
        else {
            message.vestingTypes = undefined;
        }
        if (object.accountVestingsList !== undefined &&
            object.accountVestingsList !== null) {
            message.accountVestingsList = AccountVestingsList.fromPartial(object.accountVestingsList);
        }
        else {
            message.accountVestingsList = undefined;
        }
        return message;
    },
};
