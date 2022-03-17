/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cfevesting/params";
import { VestingTypes } from "../cfevesting/vesting_types";
import { Coin } from "../cosmos/base/v1beta1/coin";
export const protobufPackage = "chain4energy.c4echain.cfevesting";
const baseQueryParamsRequest = {};
export const QueryParamsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsRequest };
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
    fromJSON(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
};
const baseQueryParamsResponse = {};
export const QueryParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
};
const baseQueryVestingTypeRequest = {};
export const QueryVestingTypeRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryVestingTypeRequest,
        };
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
    fromJSON(_) {
        const message = {
            ...baseQueryVestingTypeRequest,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseQueryVestingTypeRequest,
        };
        return message;
    },
};
const baseQueryVestingTypeResponse = {};
export const QueryVestingTypeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.vestingTypes !== undefined) {
            VestingTypes.encode(message.vestingTypes, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryVestingTypeResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vestingTypes = VestingTypes.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryVestingTypeResponse,
        };
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            message.vestingTypes = VestingTypes.fromJSON(object.vestingTypes);
        }
        else {
            message.vestingTypes = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vestingTypes !== undefined &&
            (obj.vestingTypes = message.vestingTypes
                ? VestingTypes.toJSON(message.vestingTypes)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryVestingTypeResponse,
        };
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            message.vestingTypes = VestingTypes.fromPartial(object.vestingTypes);
        }
        else {
            message.vestingTypes = undefined;
        }
        return message;
    },
};
const baseQueryVestingRequest = { address: "" };
export const QueryVestingRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryVestingRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryVestingRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryVestingRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryVestingResponse = { delegableAddress: "" };
export const QueryVestingResponse = {
    encode(message, writer = Writer.create()) {
        if (message.delegableAddress !== "") {
            writer.uint32(10).string(message.delegableAddress);
        }
        for (const v of message.vestings) {
            VestingInfo.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryVestingResponse };
        message.vestings = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.delegableAddress = reader.string();
                    break;
                case 2:
                    message.vestings.push(VestingInfo.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryVestingResponse };
        message.vestings = [];
        if (object.delegableAddress !== undefined &&
            object.delegableAddress !== null) {
            message.delegableAddress = String(object.delegableAddress);
        }
        else {
            message.delegableAddress = "";
        }
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(VestingInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.delegableAddress !== undefined &&
            (obj.delegableAddress = message.delegableAddress);
        if (message.vestings) {
            obj.vestings = message.vestings.map((e) => e ? VestingInfo.toJSON(e) : undefined);
        }
        else {
            obj.vestings = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryVestingResponse };
        message.vestings = [];
        if (object.delegableAddress !== undefined &&
            object.delegableAddress !== null) {
            message.delegableAddress = object.delegableAddress;
        }
        else {
            message.delegableAddress = "";
        }
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(VestingInfo.fromPartial(e));
            }
        }
        return message;
    },
};
const baseVestingInfo = {
    id: 0,
    vestingType: "",
    vestingStartHeight: 0,
    lockEndHeight: 0,
    vestingEndHeight: 0,
    withdrawable: "",
    delegationAllowed: false,
    currentVestedAmount: "",
};
export const VestingInfo = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).int32(message.id);
        }
        if (message.vestingType !== "") {
            writer.uint32(18).string(message.vestingType);
        }
        if (message.vestingStartHeight !== 0) {
            writer.uint32(24).int64(message.vestingStartHeight);
        }
        if (message.lockEndHeight !== 0) {
            writer.uint32(32).int64(message.lockEndHeight);
        }
        if (message.vestingEndHeight !== 0) {
            writer.uint32(40).int64(message.vestingEndHeight);
        }
        if (message.withdrawable !== "") {
            writer.uint32(50).string(message.withdrawable);
        }
        if (message.delegationAllowed === true) {
            writer.uint32(56).bool(message.delegationAllowed);
        }
        if (message.vested !== undefined) {
            Coin.encode(message.vested, writer.uint32(66).fork()).ldelim();
        }
        if (message.currentVestedAmount !== "") {
            writer.uint32(74).string(message.currentVestedAmount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVestingInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = reader.int32();
                    break;
                case 2:
                    message.vestingType = reader.string();
                    break;
                case 3:
                    message.vestingStartHeight = longToNumber(reader.int64());
                    break;
                case 4:
                    message.lockEndHeight = longToNumber(reader.int64());
                    break;
                case 5:
                    message.vestingEndHeight = longToNumber(reader.int64());
                    break;
                case 6:
                    message.withdrawable = reader.string();
                    break;
                case 7:
                    message.delegationAllowed = reader.bool();
                    break;
                case 8:
                    message.vested = Coin.decode(reader, reader.uint32());
                    break;
                case 9:
                    message.currentVestedAmount = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVestingInfo };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.vestingType !== undefined && object.vestingType !== null) {
            message.vestingType = String(object.vestingType);
        }
        else {
            message.vestingType = "";
        }
        if (object.vestingStartHeight !== undefined &&
            object.vestingStartHeight !== null) {
            message.vestingStartHeight = Number(object.vestingStartHeight);
        }
        else {
            message.vestingStartHeight = 0;
        }
        if (object.lockEndHeight !== undefined && object.lockEndHeight !== null) {
            message.lockEndHeight = Number(object.lockEndHeight);
        }
        else {
            message.lockEndHeight = 0;
        }
        if (object.vestingEndHeight !== undefined &&
            object.vestingEndHeight !== null) {
            message.vestingEndHeight = Number(object.vestingEndHeight);
        }
        else {
            message.vestingEndHeight = 0;
        }
        if (object.withdrawable !== undefined && object.withdrawable !== null) {
            message.withdrawable = String(object.withdrawable);
        }
        else {
            message.withdrawable = "";
        }
        if (object.delegationAllowed !== undefined &&
            object.delegationAllowed !== null) {
            message.delegationAllowed = Boolean(object.delegationAllowed);
        }
        else {
            message.delegationAllowed = false;
        }
        if (object.vested !== undefined && object.vested !== null) {
            message.vested = Coin.fromJSON(object.vested);
        }
        else {
            message.vested = undefined;
        }
        if (object.currentVestedAmount !== undefined &&
            object.currentVestedAmount !== null) {
            message.currentVestedAmount = String(object.currentVestedAmount);
        }
        else {
            message.currentVestedAmount = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        message.vestingType !== undefined &&
            (obj.vestingType = message.vestingType);
        message.vestingStartHeight !== undefined &&
            (obj.vestingStartHeight = message.vestingStartHeight);
        message.lockEndHeight !== undefined &&
            (obj.lockEndHeight = message.lockEndHeight);
        message.vestingEndHeight !== undefined &&
            (obj.vestingEndHeight = message.vestingEndHeight);
        message.withdrawable !== undefined &&
            (obj.withdrawable = message.withdrawable);
        message.delegationAllowed !== undefined &&
            (obj.delegationAllowed = message.delegationAllowed);
        message.vested !== undefined &&
            (obj.vested = message.vested ? Coin.toJSON(message.vested) : undefined);
        message.currentVestedAmount !== undefined &&
            (obj.currentVestedAmount = message.currentVestedAmount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVestingInfo };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.vestingType !== undefined && object.vestingType !== null) {
            message.vestingType = object.vestingType;
        }
        else {
            message.vestingType = "";
        }
        if (object.vestingStartHeight !== undefined &&
            object.vestingStartHeight !== null) {
            message.vestingStartHeight = object.vestingStartHeight;
        }
        else {
            message.vestingStartHeight = 0;
        }
        if (object.lockEndHeight !== undefined && object.lockEndHeight !== null) {
            message.lockEndHeight = object.lockEndHeight;
        }
        else {
            message.lockEndHeight = 0;
        }
        if (object.vestingEndHeight !== undefined &&
            object.vestingEndHeight !== null) {
            message.vestingEndHeight = object.vestingEndHeight;
        }
        else {
            message.vestingEndHeight = 0;
        }
        if (object.withdrawable !== undefined && object.withdrawable !== null) {
            message.withdrawable = object.withdrawable;
        }
        else {
            message.withdrawable = "";
        }
        if (object.delegationAllowed !== undefined &&
            object.delegationAllowed !== null) {
            message.delegationAllowed = object.delegationAllowed;
        }
        else {
            message.delegationAllowed = false;
        }
        if (object.vested !== undefined && object.vested !== null) {
            message.vested = Coin.fromPartial(object.vested);
        }
        else {
            message.vested = undefined;
        }
        if (object.currentVestedAmount !== undefined &&
            object.currentVestedAmount !== null) {
            message.currentVestedAmount = object.currentVestedAmount;
        }
        else {
            message.currentVestedAmount = "";
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Params(request) {
        const data = QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "Params", data);
        return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
    }
    VestingType(request) {
        const data = QueryVestingTypeRequest.encode(request).finish();
        const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "VestingType", data);
        return promise.then((data) => QueryVestingTypeResponse.decode(new Reader(data)));
    }
    Vesting(request) {
        const data = QueryVestingRequest.encode(request).finish();
        const promise = this.rpc.request("chain4energy.c4echain.cfevesting.Query", "Vesting", data);
        return promise.then((data) => QueryVestingResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
