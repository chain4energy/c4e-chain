import { Reader, Writer } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "chain4energy.c4echain.cfevesting";
export interface MsgVest {
    creator: string;
    /** uint64 amount = 2; */
    amount: string;
    vestingType: string;
}
export interface MsgVestResponse {
}
export interface MsgWithdrawAllAvailable {
    creator: string;
}
export interface MsgWithdrawAllAvailableResponse {
}
export interface MsgDelegate {
    delegatorAddress: string;
    validatorAddress: string;
    amount: Coin | undefined;
}
export interface MsgDelegateResponse {
}
export interface MsgUndelegate {
    delegatorAddress: string;
    validatorAddress: string;
    amount: Coin | undefined;
}
export interface MsgUndelegateResponse {
    completionTime: Date | undefined;
}
export interface MsgBeginRedelegate {
    delegatorAddress: string;
    validatorSrcAddress: string;
    validatorDstAddress: string;
    amount: Coin | undefined;
}
export interface MsgBeginRedelegateResponse {
    completionTime: Date | undefined;
}
export interface MsgWithdrawDelegatorReward {
    delegatorAddress: string;
    validatorAddress: string;
}
export interface MsgWithdrawDelegatorRewardResponse {
}
export interface MsgSendVesting {
    fromAddress: string;
    toAddress: string;
    vestingId: string;
    amount: string;
    restartVesting: string;
}
export interface MsgSendVestingResponse {
}
export declare const MsgVest: {
    encode(message: MsgVest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVest;
    fromJSON(object: any): MsgVest;
    toJSON(message: MsgVest): unknown;
    fromPartial(object: DeepPartial<MsgVest>): MsgVest;
};
export declare const MsgVestResponse: {
    encode(_: MsgVestResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgVestResponse;
    fromJSON(_: any): MsgVestResponse;
    toJSON(_: MsgVestResponse): unknown;
    fromPartial(_: DeepPartial<MsgVestResponse>): MsgVestResponse;
};
export declare const MsgWithdrawAllAvailable: {
    encode(message: MsgWithdrawAllAvailable, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgWithdrawAllAvailable;
    fromJSON(object: any): MsgWithdrawAllAvailable;
    toJSON(message: MsgWithdrawAllAvailable): unknown;
    fromPartial(object: DeepPartial<MsgWithdrawAllAvailable>): MsgWithdrawAllAvailable;
};
export declare const MsgWithdrawAllAvailableResponse: {
    encode(_: MsgWithdrawAllAvailableResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgWithdrawAllAvailableResponse;
    fromJSON(_: any): MsgWithdrawAllAvailableResponse;
    toJSON(_: MsgWithdrawAllAvailableResponse): unknown;
    fromPartial(_: DeepPartial<MsgWithdrawAllAvailableResponse>): MsgWithdrawAllAvailableResponse;
};
export declare const MsgDelegate: {
    encode(message: MsgDelegate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDelegate;
    fromJSON(object: any): MsgDelegate;
    toJSON(message: MsgDelegate): unknown;
    fromPartial(object: DeepPartial<MsgDelegate>): MsgDelegate;
};
export declare const MsgDelegateResponse: {
    encode(_: MsgDelegateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDelegateResponse;
    fromJSON(_: any): MsgDelegateResponse;
    toJSON(_: MsgDelegateResponse): unknown;
    fromPartial(_: DeepPartial<MsgDelegateResponse>): MsgDelegateResponse;
};
export declare const MsgUndelegate: {
    encode(message: MsgUndelegate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUndelegate;
    fromJSON(object: any): MsgUndelegate;
    toJSON(message: MsgUndelegate): unknown;
    fromPartial(object: DeepPartial<MsgUndelegate>): MsgUndelegate;
};
export declare const MsgUndelegateResponse: {
    encode(message: MsgUndelegateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUndelegateResponse;
    fromJSON(object: any): MsgUndelegateResponse;
    toJSON(message: MsgUndelegateResponse): unknown;
    fromPartial(object: DeepPartial<MsgUndelegateResponse>): MsgUndelegateResponse;
};
export declare const MsgBeginRedelegate: {
    encode(message: MsgBeginRedelegate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBeginRedelegate;
    fromJSON(object: any): MsgBeginRedelegate;
    toJSON(message: MsgBeginRedelegate): unknown;
    fromPartial(object: DeepPartial<MsgBeginRedelegate>): MsgBeginRedelegate;
};
export declare const MsgBeginRedelegateResponse: {
    encode(message: MsgBeginRedelegateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgBeginRedelegateResponse;
    fromJSON(object: any): MsgBeginRedelegateResponse;
    toJSON(message: MsgBeginRedelegateResponse): unknown;
    fromPartial(object: DeepPartial<MsgBeginRedelegateResponse>): MsgBeginRedelegateResponse;
};
export declare const MsgWithdrawDelegatorReward: {
    encode(message: MsgWithdrawDelegatorReward, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgWithdrawDelegatorReward;
    fromJSON(object: any): MsgWithdrawDelegatorReward;
    toJSON(message: MsgWithdrawDelegatorReward): unknown;
    fromPartial(object: DeepPartial<MsgWithdrawDelegatorReward>): MsgWithdrawDelegatorReward;
};
export declare const MsgWithdrawDelegatorRewardResponse: {
    encode(_: MsgWithdrawDelegatorRewardResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgWithdrawDelegatorRewardResponse;
    fromJSON(_: any): MsgWithdrawDelegatorRewardResponse;
    toJSON(_: MsgWithdrawDelegatorRewardResponse): unknown;
    fromPartial(_: DeepPartial<MsgWithdrawDelegatorRewardResponse>): MsgWithdrawDelegatorRewardResponse;
};
export declare const MsgSendVesting: {
    encode(message: MsgSendVesting, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendVesting;
    fromJSON(object: any): MsgSendVesting;
    toJSON(message: MsgSendVesting): unknown;
    fromPartial(object: DeepPartial<MsgSendVesting>): MsgSendVesting;
};
export declare const MsgSendVestingResponse: {
    encode(_: MsgSendVestingResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSendVestingResponse;
    fromJSON(_: any): MsgSendVestingResponse;
    toJSON(_: MsgSendVestingResponse): unknown;
    fromPartial(_: DeepPartial<MsgSendVestingResponse>): MsgSendVestingResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    Vest(request: MsgVest): Promise<MsgVestResponse>;
    WithdrawAllAvailable(request: MsgWithdrawAllAvailable): Promise<MsgWithdrawAllAvailableResponse>;
    Delegate(request: MsgDelegate): Promise<MsgDelegateResponse>;
    Undelegate(request: MsgUndelegate): Promise<MsgUndelegateResponse>;
    BeginRedelegate(request: MsgBeginRedelegate): Promise<MsgBeginRedelegateResponse>;
    WithdrawDelegatorReward(request: MsgWithdrawDelegatorReward): Promise<MsgWithdrawDelegatorRewardResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SendVesting(request: MsgSendVesting): Promise<MsgSendVestingResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Vest(request: MsgVest): Promise<MsgVestResponse>;
    WithdrawAllAvailable(request: MsgWithdrawAllAvailable): Promise<MsgWithdrawAllAvailableResponse>;
    Delegate(request: MsgDelegate): Promise<MsgDelegateResponse>;
    Undelegate(request: MsgUndelegate): Promise<MsgUndelegateResponse>;
    BeginRedelegate(request: MsgBeginRedelegate): Promise<MsgBeginRedelegateResponse>;
    WithdrawDelegatorReward(request: MsgWithdrawDelegatorReward): Promise<MsgWithdrawDelegatorRewardResponse>;
    SendVesting(request: MsgSendVesting): Promise<MsgSendVestingResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
