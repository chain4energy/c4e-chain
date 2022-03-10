import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "chain4energy.c4echain.cfevesting";
export interface VestingTypes {
    vestingTypes: VestingType[];
}
export interface VestingType {
    name: string;
    lockupPeriod: number;
    vestingPeriod: number;
    tokenReleasingPeriod: number;
    delegationsAllowed: boolean;
}
export declare const VestingTypes: {
    encode(message: VestingTypes, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VestingTypes;
    fromJSON(object: any): VestingTypes;
    toJSON(message: VestingTypes): unknown;
    fromPartial(object: DeepPartial<VestingTypes>): VestingTypes;
};
export declare const VestingType: {
    encode(message: VestingType, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VestingType;
    fromJSON(object: any): VestingType;
    toJSON(message: VestingType): unknown;
    fromPartial(object: DeepPartial<VestingType>): VestingType;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
