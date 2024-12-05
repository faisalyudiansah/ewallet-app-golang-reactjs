import React, { useEffect, useState } from 'react';
import style from './transactions.module.css';
import { useDispatch, useSelector } from 'react-redux';
import { NavProfile } from 'src/components/layout/navbar/nav-profile';
import { TransactionsHeaderTitle } from 'src/components/layout/transactions/Transactions-header-title';
import { TransactionsTable } from 'src/components/layout/transactions/Transactions-table';
import { TransactionsBtnTfTopup } from 'src/components/layout/transactions/Transactions-btn-tf-topup';
import { AppDispatch, RootState } from 'src/store';
import { getMe } from 'src/store/profileSlice';
import { AlertError } from 'src/components/ui/alert/AlertError';
import {
  getTransactionList,
  resetErrorListTransactions,
} from 'src/store/transactionSlice';
import { getThisWeek } from 'src/helpers/date';
import { PropsConfigTransactions } from 'src/constants/types/props';

export const Transactions: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [isError, setIsError] = useState(false);
  const { start_date, end_date } = getThisWeek();
  const [configTransaction, setConfigTransaction] =
    useState<PropsConfigTransactions>({
      sortBy: 'amount',
      sortDir: 'asc',
      transactionType: 2,
      limit: 8,
      page: 1,
      start_date,
      end_date,
    });

  const { dataUserProfile } = useSelector(
    (state: RootState) => state.profileReducer,
  );

  const {
    listTransactions,
    isListTransactionsLoading,
    isErrorListTransactions,
    errorMsgListTransactions,
  } = useSelector((state: RootState) => state.transactionReducer);

  async function fetchData() {
    try {
      await dispatch(resetErrorListTransactions());
      await dispatch(getMe());
      await dispatch(
        getTransactionList(
          configTransaction.start_date,
          configTransaction.end_date,
          configTransaction.sortBy,
          configTransaction.sortDir,
          configTransaction.transactionType,
          configTransaction.limit,
          configTransaction.page,
        ),
      );
    } catch (error) {
      setIsError(true);
    }
  }

  useEffect(() => {
    fetchData();
  }, [configTransaction]);

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetErrorListTransactions());
    }, 4000);
  }, [isErrorListTransactions]);

  return (
    <>
      {isError && isErrorListTransactions && (
        <AlertError
          message={errorMsgListTransactions || "Something's wrong. Try again!"}
        />
      )}
      <div>
        {isListTransactionsLoading ? (
          <div className="loader"></div>
        ) : (
          <>
            <NavProfile navTitle={'Transactions'} />
            <section className={style['tx-section']}>
              <TransactionsHeaderTitle dataUserProfile={dataUserProfile} />
              <TransactionsTable
                listTransactions={listTransactions}
                configTransaction={configTransaction}
                setConfigTransaction={setConfigTransaction}
              />
              <TransactionsBtnTfTopup dataUserProfile={dataUserProfile} />
            </section>
          </>
        )}
      </div>
    </>
  );
};
