import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "chain4energy.c4echain.cfevesting";
export interface AccountVestingsList {
    vestings: AccountVestings[];
}
export interface AccountVestings {
    address: string;
    delegableAddress: string;
    /** uint64 delegated = 12; */
    vestings: Vesting[];
}
export interface Vesting {
    id: number;
    vestingType: string;
    vestingStartBlock: number;
    lockEndBlock: number;
    vestingEndBlock: number;
    vested: string;
    /**
     * uint64 claimable = 6;
     * int64 last_freeing_block = 7;
     */
    freeCoinsBlockPeriod: number;
    /** uint64 free_coins_per_period = 9; */
    delegationAllowed: boolean;
    withdrawn: string;
}
export declare const AccountVestingsList: {
    encode(message: AccountVestingsList, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): AccountVestingsList;
    fromJSON(object: any): AccountVestingsList;
    toJSON(message: AccountVestingsList): unknown;
    fromPartial(object: DeepPartial<AccountVestingsList>): AccountVestingsList;
};
export declare const AccountVestings: {
    encode(message: AccountVestings, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): AccountVestings;
    fromJSON(object: any): AccountVestings;
    toJSON(message: AccountVestings): unknown;
    fromPartial(object: DeepPartial<AccountVestings>): AccountVestings;
};
export declare const Vesting: {
    encode(message: Vesting, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Vesting;
    fromJSON(object: any): Vesting;
    toJSON(message: Vesting): unknown;
    fromPartial(object: DeepPartial<Vesting>): Vesting;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
