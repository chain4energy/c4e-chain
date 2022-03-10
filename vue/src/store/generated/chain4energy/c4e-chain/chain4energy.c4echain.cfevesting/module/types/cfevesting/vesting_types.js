/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "chain4energy.c4echain.cfevesting";
const baseVestingTypes = {};
export const VestingTypes = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vestingTypes) {
            VestingType.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVestingTypes };
        message.vestingTypes = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vestingTypes.push(VestingType.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVestingTypes };
        message.vestingTypes = [];
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            for (const e of object.vestingTypes) {
                message.vestingTypes.push(VestingType.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.vestingTypes) {
            obj.vestingTypes = message.vestingTypes.map((e) => e ? VestingType.toJSON(e) : undefined);
        }
        else {
            obj.vestingTypes = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVestingTypes };
        message.vestingTypes = [];
        if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
            for (const e of object.vestingTypes) {
                message.vestingTypes.push(VestingType.fromPartial(e));
            }
        }
        return message;
    },
};
const baseVestingType = {
    name: "",
    lockupPeriod: 0,
    vestingPeriod: 0,
    tokenReleasingPeriod: 0,
    delegationsAllowed: false,
};
export const VestingType = {
    encode(message, writer = Writer.create()) {
        if (message.name !== "") {
            writer.uint32(10).string(message.name);
        }
        if (message.lockupPeriod !== 0) {
            writer.uint32(16).int64(message.lockupPeriod);
        }
        if (message.vestingPeriod !== 0) {
            writer.uint32(24).int64(message.vestingPeriod);
        }
        if (message.tokenReleasingPeriod !== 0) {
            writer.uint32(32).int64(message.tokenReleasingPeriod);
        }
        if (message.delegationsAllowed === true) {
            writer.uint32(40).bool(message.delegationsAllowed);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVestingType };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.lockupPeriod = longToNumber(reader.int64());
                    break;
                case 3:
                    message.vestingPeriod = longToNumber(reader.int64());
                    break;
                case 4:
                    message.tokenReleasingPeriod = longToNumber(reader.int64());
                    break;
                case 5:
                    message.delegationsAllowed = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVestingType };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = "";
        }
        if (object.lockupPeriod !== undefined && object.lockupPeriod !== null) {
            message.lockupPeriod = Number(object.lockupPeriod);
        }
        else {
            message.lockupPeriod = 0;
        }
        if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
            message.vestingPeriod = Number(object.vestingPeriod);
        }
        else {
            message.vestingPeriod = 0;
        }
        if (object.tokenReleasingPeriod !== undefined &&
            object.tokenReleasingPeriod !== null) {
            message.tokenReleasingPeriod = Number(object.tokenReleasingPeriod);
        }
        else {
            message.tokenReleasingPeriod = 0;
        }
        if (object.delegationsAllowed !== undefined &&
            object.delegationsAllowed !== null) {
            message.delegationsAllowed = Boolean(object.delegationsAllowed);
        }
        else {
            message.delegationsAllowed = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.lockupPeriod !== undefined &&
            (obj.lockupPeriod = message.lockupPeriod);
        message.vestingPeriod !== undefined &&
            (obj.vestingPeriod = message.vestingPeriod);
        message.tokenReleasingPeriod !== undefined &&
            (obj.tokenReleasingPeriod = message.tokenReleasingPeriod);
        message.delegationsAllowed !== undefined &&
            (obj.delegationsAllowed = message.delegationsAllowed);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVestingType };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = "";
        }
        if (object.lockupPeriod !== undefined && object.lockupPeriod !== null) {
            message.lockupPeriod = object.lockupPeriod;
        }
        else {
            message.lockupPeriod = 0;
        }
        if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
            message.vestingPeriod = object.vestingPeriod;
        }
        else {
            message.vestingPeriod = 0;
        }
        if (object.tokenReleasingPeriod !== undefined &&
            object.tokenReleasingPeriod !== null) {
            message.tokenReleasingPeriod = object.tokenReleasingPeriod;
        }
        else {
            message.tokenReleasingPeriod = 0;
        }
        if (object.delegationsAllowed !== undefined &&
            object.delegationsAllowed !== null) {
            message.delegationsAllowed = object.delegationsAllowed;
        }
        else {
            message.delegationsAllowed = false;
        }
        return message;
    },
};
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
