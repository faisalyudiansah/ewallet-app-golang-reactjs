import React, { useEffect } from 'react';
import style from './transactions.module.css';
import { ArrowLeft } from 'src/assets/svg/ArrowLeftSvg';
import { ArrowRight } from 'src/assets/svg/ArrowRightSvg';
import { TransactionsFilter } from './Transactions-filter';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from 'src/store';
import { getTransactionType } from 'src/store/transactionSlice';
import { ResListTransactionWithTotalRow } from '@/constants/response/resTransaction';
import { PropsConfigTransactions } from 'src/constants/types/props';
import { formatterDate } from 'src/helpers/date';
import { CurrencyFormatterIDR } from 'src/helpers/currencyFormatter';

export const TransactionsTable: React.FC<{
  listTransactions: ResListTransactionWithTotalRow | null;
  configTransaction: PropsConfigTransactions;
  setConfigTransaction: React.Dispatch<
    React.SetStateAction<PropsConfigTransactions>
  >;
}> = ({ listTransactions, configTransaction, setConfigTransaction }) => {
  const dispatch = useDispatch<AppDispatch>();
  const { transactionType, isTransactionTypeLoading } = useSelector(
    (state: RootState) => state.transactionReducer,
  );
  const totalPages = listTransactions
    ? Math.ceil(listTransactions.total_row / configTransaction.limit)
    : 1;

  const handlePrevPage = () => {
    if (configTransaction.page > 1) {
      setConfigTransaction((prev) => ({
        ...prev,
        page: prev.page - 1,
      }));
    }
  };

  const handleNextPage = () => {
    if (
      listTransactions &&
      configTransaction.page <
        Math.ceil(listTransactions.total_row / configTransaction.limit)
    ) {
      setConfigTransaction((prev) => ({
        ...prev,
        page: prev.page + 1,
      }));
    }
  };

  useEffect(() => {
    dispatch(getTransactionType());
  }, []);

  return (
    <>
      {!isTransactionTypeLoading && (
        <TransactionsFilter
          transactionType={transactionType}
          configTransaction={configTransaction}
          setConfigTransaction={setConfigTransaction}
        />
      )}
      <div className={style['tx-container-table']}>
        <table className={style['tx-table']}>
          <thead>
            <tr className={style['tx-table-head-container']}>
              <th className={style['tx-table-head-th-1']}>DATE</th>
              <th className={style['tx-table-head-th-1']}>DESCRIPTION</th>
              <th className={style['tx-table-head-th-1']}>TO / FROM</th>
              <th className={style['tx-table-head-th-2']}>AMOUNT</th>
            </tr>
          </thead>
          <tbody>
            {listTransactions?.data.map((data, i) => {
              const color =
                configTransaction.transactionType === 1
                  ? style['text-green']
                  : style['text-red'];
              const sign = configTransaction.transactionType === 1 ? '+' : '-';
              const borderBottom =
                listTransactions?.data.length - 1 == i
                  ? ''
                  : style['border-bottom-table'];
              return (
                <tr key={i + 1}>
                  <td className={`${style['tx-table-td-1']} ${borderBottom}`}>
                    {formatterDate(data?.transaction_date)}
                  </td>
                  <td className={`${style['tx-table-td-2']} ${borderBottom}`}>
                    {!data?.description ? '-' : data?.description}
                  </td>
                  <td className={`${style['tx-table-td-2']} ${borderBottom}`}>
                    {!data?.transaction_additional_detail_id
                      ? '-'
                      : data?.transaction_additional_detail_id}
                  </td>
                  <td
                    className={`${style['tx-table-td-3']} ${borderBottom} ${color}`}
                  >
                    {`${sign} IDR ${CurrencyFormatterIDR(
                      Number(data?.amount),
                    )}`}
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
      <div className={style['tx-container-btn-pagination']}>
        <button
          onClick={handlePrevPage}
          disabled={configTransaction.page === 1}
          className={style['tx-btn-pagination']}
        >
          <ArrowLeft />
        </button>
        <button
          onClick={handleNextPage}
          disabled={
            listTransactions?.total_row === 0 ||
            configTransaction.page === totalPages
          }
          className={style['tx-btn-pagination']}
        >
          <ArrowRight />
        </button>
      </div>
    </>
  );
};
