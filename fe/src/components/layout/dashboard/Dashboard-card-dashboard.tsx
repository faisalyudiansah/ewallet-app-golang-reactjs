import React, { useState } from 'react';
import style from './dashboard.module.css';
import { IncomeSvg } from 'src/assets/svg/incomeSvg';
import { ExpenseSvg } from 'src/assets/svg/expenseSvg';
import { CalculatorGreenSvg } from 'src/assets/svg/calculatorGreenSvg';
import { CalculatorRedSvg } from 'src/assets/svg/calculatorRedSvg';
import { InputPasswordHiddenSvg } from 'src/assets/svg/inputPasswordHiddenSvg';
import { CurrencyFormatterIDR } from 'src/helpers/currencyFormatter';
import { DataUserProfile } from 'src/constants/response/resProfile';
import { ExpenseSumByMonth } from 'src/constants/response/resTransaction';
import { getThisMonth } from 'src/helpers/date';

export const DashboardCardDashboard: React.FC<{
  dataUserProfile: DataUserProfile | null;
  expenseSumByMonth: ExpenseSumByMonth | null;
}> = ({ dataUserProfile, expenseSumByMonth }) => {
  const [showBalance, setShowBalance] = useState(false);

  function toggleShowBalance() {
    setShowBalance(!showBalance);
  }

  return (
    <div className={style['container-card-dashboard']}>
      <div className={style['card-dashboard-1']}>
        <div className={style['container-balance']}>
          <span className={style['text-balance']}>Balance</span>
          <InputPasswordHiddenSvg
            className={style['icon-password']}
            onClick={toggleShowBalance}
          />
        </div>
        <div className={style['text-money-info']}>
          <span>IDR</span>
          <span>
            {showBalance
              ? CurrencyFormatterIDR(Number(dataUserProfile?.amount))
              : '**********'}
          </span>
        </div>
        <span className={style['text-wallet-number']}>
          {dataUserProfile?.wallet_number}
        </span>
      </div>
      <div className={style['card-dashboard-2']}>
        <span className={style['text-balance']}>
          <CalculatorGreenSvg width={28} height={30} />
        </span>
        <div className={style['container-text-money-logo']}>
          <div className={style['text-money-info']}>
            <span>IDR</span>
            <span>750.000</span>
          </div>
          <IncomeSvg customClass={style['svg-card']} />
        </div>
        <span className={style['text-income-date']}>
          Income - {getThisMonth()} 2024
        </span>
      </div>
      <div className={style['card-dashboard-3']}>
        <span className={style['text-balance']}>
          <CalculatorRedSvg width={28} height={30} />
        </span>
        <div className={style['container-text-money-logo']}>
          <div className={style['text-money-info']}>
            <span>IDR</span>
            <span>{CurrencyFormatterIDR(Number(expenseSumByMonth?.sum))}</span>
          </div>
          <ExpenseSvg customClass={style['svg-card']} />
        </div>
        <span className={style['text-expense-date']}>
          Expense - {getThisMonth()} 2024
        </span>
      </div>
    </div>
  );
};
