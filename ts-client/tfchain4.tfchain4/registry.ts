import { GeneratedType } from "@cosmjs/proto-signing";
import { Params } from "./types/tfchain4/tfchain4/params";
import { MsgUpdateParams } from "./types/tfchain4/tfchain4/tx";
import { QueryParamsRequest } from "./types/tfchain4/tfchain4/query";
import { MsgUpdateParamsResponse } from "./types/tfchain4/tfchain4/tx";
import { GenesisState } from "./types/tfchain4/tfchain4/genesis";
import { QueryParamsResponse } from "./types/tfchain4/tfchain4/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/tfchain4.tfchain4.Params", Params],
    ["/tfchain4.tfchain4.MsgUpdateParams", MsgUpdateParams],
    ["/tfchain4.tfchain4.QueryParamsRequest", QueryParamsRequest],
    ["/tfchain4.tfchain4.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/tfchain4.tfchain4.GenesisState", GenesisState],
    ["/tfchain4.tfchain4.QueryParamsResponse", QueryParamsResponse],
    
];

export { msgTypes }