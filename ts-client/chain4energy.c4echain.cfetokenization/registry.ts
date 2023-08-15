import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAuthorizeCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgAcceptDevice } from "./types/c4echain/cfetokenization/tx";
import { MsgBuyCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgBurnCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgAddCertificateToMarketplace } from "./types/c4echain/cfetokenization/tx";
import { MsgAssignDeviceToUser } from "./types/c4echain/cfetokenization/tx";
import { MsgAddMeasurement } from "./types/c4echain/cfetokenization/tx";
import { MsgCreateUserCertificates } from "./types/c4echain/cfetokenization/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfetokenization.MsgAuthorizeCertificate", MsgAuthorizeCertificate],
    ["/chain4energy.c4echain.cfetokenization.MsgAcceptDevice", MsgAcceptDevice],
    ["/chain4energy.c4echain.cfetokenization.MsgBuyCertificate", MsgBuyCertificate],
    ["/chain4energy.c4echain.cfetokenization.MsgBurnCertificate", MsgBurnCertificate],
    ["/chain4energy.c4echain.cfetokenization.MsgAddCertificateToMarketplace", MsgAddCertificateToMarketplace],
    ["/chain4energy.c4echain.cfetokenization.MsgAssignDeviceToUser", MsgAssignDeviceToUser],
    ["/chain4energy.c4echain.cfetokenization.MsgAddMeasurement", MsgAddMeasurement],
    ["/chain4energy.c4echain.cfetokenization.MsgCreateUserCertificates", MsgCreateUserCertificates],
    
];

export { msgTypes }