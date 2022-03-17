import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDelegate } from "./types/cfevesting/tx";
import { MsgWithdrawDelegatorReward } from "./types/cfevesting/tx";
import { MsgVest } from "./types/cfevesting/tx";
import { MsgWithdrawAllAvailable } from "./types/cfevesting/tx";
import { MsgBeginRedelegate } from "./types/cfevesting/tx";
import { MsgUndelegate } from "./types/cfevesting/tx";
import { MsgSendVesting } from "./types/cfevesting/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgDelegate: (data: MsgDelegate) => EncodeObject;
    msgWithdrawDelegatorReward: (data: MsgWithdrawDelegatorReward) => EncodeObject;
    msgVest: (data: MsgVest) => EncodeObject;
    msgWithdrawAllAvailable: (data: MsgWithdrawAllAvailable) => EncodeObject;
    msgBeginRedelegate: (data: MsgBeginRedelegate) => EncodeObject;
    msgUndelegate: (data: MsgUndelegate) => EncodeObject;
    msgSendVesting: (data: MsgSendVesting) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
