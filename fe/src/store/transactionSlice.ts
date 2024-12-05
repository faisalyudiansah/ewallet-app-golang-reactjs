import { createSlice } from '@reduxjs/toolkit';
import {
  ReqTopUp,
  ReqTransferMoney,
} from 'src/constants/types/typeTransaction';
import { Dispatch } from 'redux';
import {
  ErrorTransactions,
  ExpenseSumByMonth,
  ResExpenseSumByMonth,
  ResListTransaction,
  ResListTransactionWithTotalRow,
  ResTransactionType,
  TransactionType,
} from 'src/constants/response/resTransaction';
import { dataGetThisWeek, getThisMonth } from 'src/helpers/date';

interface AppState {
  transactionType: TransactionType[];
  isTransactionTypeLoading: boolean;
  expenseSumByMonth: ExpenseSumByMonth | null;
  isErrorTransfer: boolean;
  errorMsgTransfer: string | null;
  isErrorTopUp: boolean;
  errorMsgTopUp: string | null;
  listTransactions: ResListTransactionWithTotalRow | null;
  isListTransactionsLoading: boolean;
  isErrorListTransactions: boolean;
  errorMsgListTransactions: string | null;
}

const initialState: AppState = {
  transactionType: [],
  isTransactionTypeLoading: true,
  expenseSumByMonth: null,
  isErrorTransfer: false,
  errorMsgTransfer: null,
  isErrorTopUp: false,
  errorMsgTopUp: null,
  listTransactions: null,
  isListTransactionsLoading: true,
  isErrorListTransactions: false,
  errorMsgListTransactions: null,
};

export const appSliceTransaction = createSlice({
  name: 'sea_wallet_transaction',
  initialState: initialState,
  reducers: {
    changeTransactionType: (state, action) => {
      state.transactionType = action.payload;
    },
    changeIsTransactionTypeLoading: (state, action) => {
      state.isTransactionTypeLoading = action.payload;
    },
    changeExpenseSumByMonth: (state, action) => {
      state.expenseSumByMonth = action.payload;
    },
    changeIsErrorTransfer: (state, action) => {
      state.isErrorTransfer = action.payload;
    },
    changeErrorMsgTransfer: (state, action) => {
      state.errorMsgTransfer = action.payload;
    },
    changeIsErrorTopUp: (state, action) => {
      state.isErrorTopUp = action.payload;
    },
    changeErrorMsgTopUp: (state, action) => {
      state.errorMsgTopUp = action.payload;
    },
    changeListTransactions: (state, action) => {
      state.listTransactions = action.payload;
    },
    changeIsListTransactionsLoading: (state, action) => {
      state.isListTransactionsLoading = action.payload;
    },
    changeIsErrorListTransactions: (state, action) => {
      state.isErrorListTransactions = action.payload;
    },
    changeErrorMsgListTransactions: (state, action) => {
      state.errorMsgListTransactions = action.payload;
    },
  },
});

export const {
  changeTransactionType,
  changeIsTransactionTypeLoading,
  changeExpenseSumByMonth,
  changeIsErrorTransfer,
  changeErrorMsgTransfer,
  changeIsErrorTopUp,
  changeErrorMsgTopUp,
  changeListTransactions,
  changeIsListTransactionsLoading,
  changeIsErrorListTransactions,
  changeErrorMsgListTransactions,
} = appSliceTransaction.actions;

export const getTransactionList = (
  start_date = dataGetThisWeek.start_date,
  end_date = dataGetThisWeek.end_date,
  sortBy = 'amount',
  sortDir = 'asc',
  transactionType = 2,
  limit = 3,
  page = 1,
) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      dispatch(changeIsListTransactionsLoading(true));
      const link =
        process.env.REACT_APP_BASE_URL +
        `/transactions?start_date=${start_date}&end_date${end_date}=&sort_by[]=${sortBy}&sort_dir[]=${sortDir}&transaction_type=${transactionType}&limit=${limit}&page=${page}`;
      const response = await fetch(link, {
        method: 'GET',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
      });
      if (!response.ok) {
        const errorMsg: ErrorTransactions = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
      const { data } = await response.json();
      const listTransactions: ResListTransaction[] = data.entries;
      const totalRow = data.page_info.total_row;
      dispatch(
        changeListTransactions({
          data: listTransactions,
          total_row: totalRow,
        }),
      );
    } catch (error) {
      dispatch(changeIsErrorListTransactions(true));
      if (error instanceof Error) {
        dispatch(changeErrorMsgListTransactions(error.message));
      } else {
        dispatch(
          changeErrorMsgListTransactions('Something wrong with the server'),
        );
      }
      throw error;
    } finally {
      dispatch(changeIsListTransactionsLoading(false));
    }
  };
};

export const getTransactionType = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      dispatch(changeIsTransactionTypeLoading(true));
      const link = process.env.REACT_APP_BASE_URL + `/transactions/types`;
      const response = await fetch(link, {
        method: 'GET',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
      });
      const data: ResTransactionType = await response.json();
      dispatch(changeTransactionType(data.data));
    } catch (error) {
      console.log('from redux: ' + error);
      throw error;
    } finally {
      dispatch(changeIsTransactionTypeLoading(false));
    }
  };
};

export const getExpense = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link =
        process.env.REACT_APP_BASE_URL +
        `/transactions/expense/${getThisMonth()}`;
      const response = await fetch(link, {
        method: 'GET',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
      });
      const data: ResExpenseSumByMonth = await response.json();
      dispatch(changeExpenseSumByMonth(data.data));
    } catch (error) {
      console.log('from redux: ' + error);
      throw error;
    }
  };
};

export const transferMoney = (input: ReqTransferMoney) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link = process.env.REACT_APP_BASE_URL + `/transactions/transfers`;
      const response = await fetch(link, {
        method: 'POST',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ErrorTransactions = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
    } catch (error) {
      dispatch(changeIsErrorTransfer(true));
      if (error instanceof Error) {
        dispatch(changeErrorMsgTransfer(error.message));
      } else {
        dispatch(changeErrorMsgTransfer('Something wrong with the server'));
      }
      throw error;
    }
  };
};

export const topUpMoney = (input: ReqTopUp) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link = process.env.REACT_APP_BASE_URL + `/transactions/top-ups`;
      const response = await fetch(link, {
        method: 'POST',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ErrorTransactions = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
    } catch (error) {
      dispatch(changeIsErrorTopUp(true));
      if (error instanceof Error) {
        dispatch(changeErrorMsgTopUp(error.message));
      } else {
        dispatch(changeErrorMsgTopUp('Something wrong with the server'));
      }
      throw error;
    }
  };
};

export const resetErrorTransfer = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(changeIsErrorTransfer(false));
    dispatch(changeErrorMsgTransfer(null));
  };
};

export const resetErrorTopUp = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(changeIsErrorTopUp(false));
    dispatch(changeErrorMsgTopUp(null));
  };
};

export const resetErrorListTransactions = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(changeIsErrorListTransactions(false));
    dispatch(changeErrorMsgListTransactions(null));
  };
};

export default appSliceTransaction.reducer;
