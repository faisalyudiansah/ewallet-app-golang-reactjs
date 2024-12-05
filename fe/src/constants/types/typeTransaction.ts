export type ReqTransferMoney = {
  amount: number;
  wallet_to: string;
  description: string;
};

export type ReqTopUp = {
  source_of_funds_id: number;
  amount: number;
};
