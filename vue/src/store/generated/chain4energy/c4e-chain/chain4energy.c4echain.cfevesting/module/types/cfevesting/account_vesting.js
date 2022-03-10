/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "chain4energy.c4echain.cfevesting";
const baseAccountVestingsList = {};
export const AccountVestingsList = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vestings) {
            AccountVestings.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAccountVestingsList };
        message.vestings = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vestings.push(AccountVestings.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseAccountVestingsList };
        message.vestings = [];
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(AccountVestings.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.vestings) {
            obj.vestings = message.vestings.map((e) => e ? AccountVestings.toJSON(e) : undefined);
        }
        else {
            obj.vestings = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAccountVestingsList };
        message.vestings = [];
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(AccountVestings.fromPartial(e));
            }
        }
        return message;
    },
};
const baseAccountVestings = { address: "", delegableAddress: "" };
export const AccountVestings = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        if (message.delegableAddress !== "") {
            writer.uint32(18).string(message.delegableAddress);
        }
        for (const v of message.vestings) {
            Vesting.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseAccountVestings };
        message.vestings = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.delegableAddress = reader.string();
                    break;
                case 3:
                    message.vestings.push(Vesting.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseAccountVestings };
        message.vestings = [];
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.delegableAddress !== undefined &&
            object.delegableAddress !== null) {
            message.delegableAddress = String(object.delegableAddress);
        }
        else {
            message.delegableAddress = "";
        }
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(Vesting.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.delegableAddress !== undefined &&
            (obj.delegableAddress = message.delegableAddress);
        if (message.vestings) {
            obj.vestings = message.vestings.map((e) => e ? Vesting.toJSON(e) : undefined);
        }
        else {
            obj.vestings = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseAccountVestings };
        message.vestings = [];
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.delegableAddress !== undefined &&
            object.delegableAddress !== null) {
            message.delegableAddress = object.delegableAddress;
        }
        else {
            message.delegableAddress = "";
        }
        if (object.vestings !== undefined && object.vestings !== null) {
            for (const e of object.vestings) {
                message.vestings.push(Vesting.fromPartial(e));
            }
        }
        return message;
    },
};
const baseVesting = {
    vestingType: "",
    vestingStartBlock: 0,
    lockEndBlock: 0,
    vestingEndBlock: 0,
    vested: 0,
    claimable: 0,
    lastFreeingBlock: 0,
    freeCoinsBlockPeriod: 0,
    freeCoinsPerPeriod: 0,
    delegationAllowed: false,
    withdrawn: 0,
};
export const Vesting = {
    encode(message, writer = Writer.create()) {
        if (message.vestingType !== "") {
            writer.uint32(10).string(message.vestingType);
        }
        if (message.vestingStartBlock !== 0) {
            writer.uint32(16).int64(message.vestingStartBlock);
        }
        if (message.lockEndBlock !== 0) {
            writer.uint32(24).int64(message.lockEndBlock);
        }
        if (message.vestingEndBlock !== 0) {
            writer.uint32(32).int64(message.vestingEndBlock);
        }
        if (message.vested !== 0) {
            writer.uint32(40).uint64(message.vested);
        }
        if (message.claimable !== 0) {
            writer.uint32(48).uint64(message.claimable);
        }
        if (message.lastFreeingBlock !== 0) {
            writer.uint32(56).int64(message.lastFreeingBlock);
        }
        if (message.freeCoinsBlockPeriod !== 0) {
            writer.uint32(64).int64(message.freeCoinsBlockPeriod);
        }
        if (message.freeCoinsPerPeriod !== 0) {
            writer.uint32(72).uint64(message.freeCoinsPerPeriod);
        }
        if (message.delegationAllowed === true) {
            writer.uint32(80).bool(message.delegationAllowed);
        }
        if (message.withdrawn !== 0) {
            writer.uint32(88).uint64(message.withdrawn);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVesting };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vestingType = reader.string();
                    break;
                case 2:
                    message.vestingStartBlock = longToNumber(reader.int64());
                    break;
                case 3:
                    message.lockEndBlock = longToNumber(reader.int64());
                    break;
                case 4:
                    message.vestingEndBlock = longToNumber(reader.int64());
                    break;
                case 5:
                    message.vested = longToNumber(reader.uint64());
                    break;
                case 6:
                    message.claimable = longToNumber(reader.uint64());
                    break;
                case 7:
                    message.lastFreeingBlock = longToNumber(reader.int64());
                    break;
                case 8:
                    message.freeCoinsBlockPeriod = longToNumber(reader.int64());
                    break;
                case 9:
                    message.freeCoinsPerPeriod = longToNumber(reader.uint64());
                    break;
                case 10:
                    message.delegationAllowed = reader.bool();
                    break;
                case 11:
                    message.withdrawn = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVesting };
        if (object.vestingType !== undefined && object.vestingType !== null) {
            message.vestingType = String(object.vestingType);
        }
        else {
            message.vestingType = "";
        }
        if (object.vestingStartBlock !== undefined &&
            object.vestingStartBlock !== null) {
            message.vestingStartBlock = Number(object.vestingStartBlock);
        }
        else {
            message.vestingStartBlock = 0;
        }
        if (object.lockEndBlock !== undefined && object.lockEndBlock !== null) {
            message.lockEndBlock = Number(object.lockEndBlock);
        }
        else {
            message.lockEndBlock = 0;
        }
        if (object.vestingEndBlock !== undefined &&
            object.vestingEndBlock !== null) {
            message.vestingEndBlock = Number(object.vestingEndBlock);
        }
        else {
            message.vestingEndBlock = 0;
        }
        if (object.vested !== undefined && object.vested !== null) {
            message.vested = Number(object.vested);
        }
        else {
            message.vested = 0;
        }
        if (object.claimable !== undefined && object.claimable !== null) {
            message.claimable = Number(object.claimable);
        }
        else {
            message.claimable = 0;
        }
        if (object.lastFreeingBlock !== undefined &&
            object.lastFreeingBlock !== null) {
            message.lastFreeingBlock = Number(object.lastFreeingBlock);
        }
        else {
            message.lastFreeingBlock = 0;
        }
        if (object.freeCoinsBlockPeriod !== undefined &&
            object.freeCoinsBlockPeriod !== null) {
            message.freeCoinsBlockPeriod = Number(object.freeCoinsBlockPeriod);
        }
        else {
            message.freeCoinsBlockPeriod = 0;
        }
        if (object.freeCoinsPerPeriod !== undefined &&
            object.freeCoinsPerPeriod !== null) {
            message.freeCoinsPerPeriod = Number(object.freeCoinsPerPeriod);
        }
        else {
            message.freeCoinsPerPeriod = 0;
        }
        if (object.delegationAllowed !== undefined &&
            object.delegationAllowed !== null) {
            message.delegationAllowed = Boolean(object.delegationAllowed);
        }
        else {
            message.delegationAllowed = false;
        }
        if (object.withdrawn !== undefined && object.withdrawn !== null) {
            message.withdrawn = Number(object.withdrawn);
        }
        else {
            message.withdrawn = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vestingType !== undefined &&
            (obj.vestingType = message.vestingType);
        message.vestingStartBlock !== undefined &&
            (obj.vestingStartBlock = message.vestingStartBlock);
        message.lockEndBlock !== undefined &&
            (obj.lockEndBlock = message.lockEndBlock);
        message.vestingEndBlock !== undefined &&
            (obj.vestingEndBlock = message.vestingEndBlock);
        message.vested !== undefined && (obj.vested = message.vested);
        message.claimable !== undefined && (obj.claimable = message.claimable);
        message.lastFreeingBlock !== undefined &&
            (obj.lastFreeingBlock = message.lastFreeingBlock);
        message.freeCoinsBlockPeriod !== undefined &&
            (obj.freeCoinsBlockPeriod = message.freeCoinsBlockPeriod);
        message.freeCoinsPerPeriod !== undefined &&
            (obj.freeCoinsPerPeriod = message.freeCoinsPerPeriod);
        message.delegationAllowed !== undefined &&
            (obj.delegationAllowed = message.delegationAllowed);
        message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVesting };
        if (object.vestingType !== undefined && object.vestingType !== null) {
            message.vestingType = object.vestingType;
        }
        else {
            message.vestingType = "";
        }
        if (object.vestingStartBlock !== undefined &&
            object.vestingStartBlock !== null) {
            message.vestingStartBlock = object.vestingStartBlock;
        }
        else {
            message.vestingStartBlock = 0;
        }
        if (object.lockEndBlock !== undefined && object.lockEndBlock !== null) {
            message.lockEndBlock = object.lockEndBlock;
        }
        else {
            message.lockEndBlock = 0;
        }
        if (object.vestingEndBlock !== undefined &&
            object.vestingEndBlock !== null) {
            message.vestingEndBlock = object.vestingEndBlock;
        }
        else {
            message.vestingEndBlock = 0;
        }
        if (object.vested !== undefined && object.vested !== null) {
            message.vested = object.vested;
        }
        else {
            message.vested = 0;
        }
        if (object.claimable !== undefined && object.claimable !== null) {
            message.claimable = object.claimable;
        }
        else {
            message.claimable = 0;
        }
        if (object.lastFreeingBlock !== undefined &&
            object.lastFreeingBlock !== null) {
            message.lastFreeingBlock = object.lastFreeingBlock;
        }
        else {
            message.lastFreeingBlock = 0;
        }
        if (object.freeCoinsBlockPeriod !== undefined &&
            object.freeCoinsBlockPeriod !== null) {
            message.freeCoinsBlockPeriod = object.freeCoinsBlockPeriod;
        }
        else {
            message.freeCoinsBlockPeriod = 0;
        }
        if (object.freeCoinsPerPeriod !== undefined &&
            object.freeCoinsPerPeriod !== null) {
            message.freeCoinsPerPeriod = object.freeCoinsPerPeriod;
        }
        else {
            message.freeCoinsPerPeriod = 0;
        }
        if (object.delegationAllowed !== undefined &&
            object.delegationAllowed !== null) {
            message.delegationAllowed = object.delegationAllowed;
        }
        else {
            message.delegationAllowed = false;
        }
        if (object.withdrawn !== undefined && object.withdrawn !== null) {
            message.withdrawn = object.withdrawn;
        }
        else {
            message.withdrawn = 0;
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
