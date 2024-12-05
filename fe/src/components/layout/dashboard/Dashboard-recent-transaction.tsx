import React, { useEffect, useState } from 'react';
import { CalculatorGreenSvg } from 'src/assets/svg/calculatorGreenSvg';
import { CalculatorRedSvg } from 'src/assets/svg/calculatorRedSvg';
import style from './dashboard.module.css';
import emptyRecentTransaction from 'src/assets/dashboard/empty-recent-transaction.png';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from 'src/store';
import {
  getTransactionList,
  resetErrorListTransactions,
} from 'src/store/transactionSlice';
import { formatterDate, getThisWeek } from 'src/helpers/date';
import { CurrencyFormatterIDR } from 'src/helpers/currencyFormatter';
import { AlertError } from 'src/components/ui/alert/AlertError';

export const DashboardRecentTransaction: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [showAlert, setShowAlert] = useState(false);
  const { start_date, end_date } = getThisWeek();
  const configTransaction = {
    sortBy: 'created_at',
    sortDir: 'asc',
    transactionType: 1,
    limit: 3,
    page: 1,
    start_date,
    end_date,
  };
  const {
    listTransactions,
    isErrorListTransactions,
    errorMsgListTransactions,
  } = useSelector((state: RootState) => state.transactionReducer);

  async function fetchTransactionsList() {
    try {
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
      setShowAlert(true);
    }
  }

  useEffect(() => {
    fetchTransactionsList();
  }, []);

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetErrorListTransactions());
      setShowAlert(false);
    }, 4000);
  }, [isErrorListTransactions]);

  return (
    <>
      {showAlert && isErrorListTransactions && (
        <AlertError
          message={errorMsgListTransactions || "Something's wrong. Try again!"}
        />
      )}
      <div className={style['container-recent-transactions']}>
        <div className={style['container-header-recent-transactions']}>
          <span className={style['header-text']}>Recent Transactions</span>
          <span className={style['header-secondary-text']}>This Week</span>
        </div>
        <div className={style['container-card-recent-transaction']}>
          {listTransactions?.data.length === 0 ? (
            <div className={style['container-empty-recent-transaction']}>
              <img
                src={emptyRecentTransaction}
                alt="empty-Recent-Transaction"
              />
              <div className={style['container-text-empty-recent-transaction']}>
                <span
                  className={style['text-primary-empty-recent-transaction']}
                >
                  No recent transactions
                </span>
                <span
                  className={style['text-secondary-empty-recent-transaction']}
                >
                  Go to transactions page to add some!
                </span>
              </div>
            </div>
          ) : (
            <>
              {listTransactions?.data.map((data, i) => {
                return (
                  <div key={i + 1} className={style['card-recent-transaction']}>
                    <span
                      className={`${
                        style['text-money-card-recent-transaction-top']
                      } ${
                        configTransaction.transactionType === 2
                          ? style['text-red']
                          : style['text-green']
                      }`}
                    >
                      {data.amount}
                    </span>
                    {configTransaction.transactionType === 1 ? (
                      <CalculatorGreenSvg
                        customClass={style['svg-recent-transaction']}
                        width={39}
                        height={42}
                      />
                    ) : (
                      <CalculatorRedSvg
                        customClass={style['svg-recent-transaction']}
                        width={39}
                        height={42}
                      />
                    )}
                    <span className={style['text-card-recent-transaction']}>
                      {data.description}
                    </span>
                    <span className={style['text-card-recent-transaction']}>
                      {formatterDate(data.transaction_date)}
                    </span>
                    <span
                      className={`${
                        style['text-money-card-recent-transaction-bottom']
                      } ${
                        configTransaction.transactionType === 2
                          ? style['text-red']
                          : style['text-green']
                      }`}
                    >
                      IDR {CurrencyFormatterIDR(Number(data.amount))}
                    </span>
                  </div>
                );
              })}
            </>
          )}
        </div>
      </div>
    </>
  );
};
