// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateOrder } from "./types/market/tx";
import { MsgCreatePool } from "./types/market/tx";
import { MsgCreateDrop } from "./types/market/tx";
import { MsgRedeemDrop } from "./types/market/tx";


const types = [
  ["/onomyprotocol.market.market.MsgCreateOrder", MsgCreateOrder],
  ["/onomyprotocol.market.market.MsgCreatePool", MsgCreatePool],
  ["/onomyprotocol.market.market.MsgCreateDrop", MsgCreateDrop],
  ["/onomyprotocol.market.market.MsgRedeemDrop", MsgRedeemDrop],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateOrder: (data: MsgCreateOrder): EncodeObject => ({ typeUrl: "/onomyprotocol.market.market.MsgCreateOrder", value: MsgCreateOrder.fromPartial( data ) }),
    msgCreatePool: (data: MsgCreatePool): EncodeObject => ({ typeUrl: "/onomyprotocol.market.market.MsgCreatePool", value: MsgCreatePool.fromPartial( data ) }),
    msgCreateDrop: (data: MsgCreateDrop): EncodeObject => ({ typeUrl: "/onomyprotocol.market.market.MsgCreateDrop", value: MsgCreateDrop.fromPartial( data ) }),
    msgRedeemDrop: (data: MsgRedeemDrop): EncodeObject => ({ typeUrl: "/onomyprotocol.market.market.MsgRedeemDrop", value: MsgRedeemDrop.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
