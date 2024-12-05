import React, { useState } from 'react';
import style from './transactions.module.css';
import { DataUserProfile } from 'src/constants/response/resProfile';
import { CurrencyFormatterIDR } from 'src/helpers/currencyFormatter';

export const TransactionsHeaderTitle: React.FC<{
  dataUserProfile: DataUserProfile | null;
}> = ({ dataUserProfile }) => {
  const [show, setShow] = useState(false);
  function showBalance() {
    return setShow(!show);
  }
  return (
    <div className={style['tx-container-header']}>
      <span
        onClick={showBalance}
        role="presentation"
        className={style['tx-header-title']}
      >
        IDR{' '}
        {show
          ? CurrencyFormatterIDR(Number(dataUserProfile?.amount))
          : '*********'}
      </span>
      <span className={style['tx-header-desc-title']}>
        Total balance from account {dataUserProfile?.wallet_number}
      </span>
    </div>
  );
};
