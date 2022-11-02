// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgGrantAllowance } from "./types/cosmos/feegrant/v1beta1/tx";
import { MsgRevokeAllowance } from "./types/cosmos/feegrant/v1beta1/tx";


export { MsgGrantAllowance, MsgRevokeAllowance };

type sendMsgGrantAllowanceParams = {
  value: MsgGrantAllowance,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeAllowanceParams = {
  value: MsgRevokeAllowance,
  fee?: StdFee,
  memo?: string
};


type msgGrantAllowanceParams = {
  value: MsgGrantAllowance,
};

type msgRevokeAllowanceParams = {
  value: MsgRevokeAllowance,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "https://rpc.astranetwork.zone", prefix: "astra" }) => {

  return {
		
		async sendMsgGrantAllowance({ value, fee, memo }: sendMsgGrantAllowanceParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgGrantAllowance: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgGrantAllowance({ value: MsgGrantAllowance.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgGrantAllowance: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeAllowance({ value, fee, memo }: sendMsgRevokeAllowanceParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeAllowance: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeAllowance({ value: MsgRevokeAllowance.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeAllowance: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgGrantAllowance({ value }: msgGrantAllowanceParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.feegrant.v1beta1.MsgGrantAllowance", value: MsgGrantAllowance.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgGrantAllowance: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeAllowance({ value }: msgRevokeAllowanceParams): EncodeObject {
			try {
				return { typeUrl: "/cosmos.feegrant.v1beta1.MsgRevokeAllowance", value: MsgRevokeAllowance.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeAllowance: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "https://astranetwork.zone" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "astra" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			CosmosFeegrantV1Beta1: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;