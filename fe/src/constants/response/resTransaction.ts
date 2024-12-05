export type ExpenseSumByMonth = {
  month: string;
  sum: number;
};

export type ResExpenseSumByMonth = {
  message: string;
  data: ExpenseSumByMonth;
};

export type TransactionType = {
  type_id: number;
  type_name: string;
};

export type ResTransactionType = {
  message: string;
  data: TransactionType[];
};

export type ResListTransaction = {
  transaction_id: number;
  transaction_ref_id: number | null;
  wallet_id: number;
  transaction_type_id: number;
  transaction_additional_detail_id: number;
  amount: string;
  description: string;
  transaction_date: string;
};

export type ResListTransactionWithTotalRow = {
  data: ResListTransaction[];
  total_row: number;
};

export type ErrorTransactions = {
  message: string;
  details: [
    {
      field: string;
      message: string;
    },
  ];
};
