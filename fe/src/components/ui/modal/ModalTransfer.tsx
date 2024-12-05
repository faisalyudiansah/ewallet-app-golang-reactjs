import React, { useEffect, useState } from 'react';
import style from './modalTransfer.module.css';
import { Button } from 'src/components/ui/button/Button';
import { GreenChecklist } from 'src/assets/svg/GreenChecklist';
import {
  CurrencyFormatterIDR,
  CurrencyFormatterIDRInput,
} from 'src/helpers/currencyFormatter';
import { FormatDate } from 'src/helpers/dataFormatter';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from 'src/store';
import { resetErrorTransfer, transferMoney } from 'src/store/transactionSlice';
import { AlertError } from '../alert/AlertError';
import { DataUserProfile } from 'src/constants/response/resProfile';

export const ModalTransfer: React.FC<{
  closeModal?: () => void;
  dataUserProfile: DataUserProfile | null;
}> = ({ closeModal, dataUserProfile }) => {
  const dispatch = useDispatch<AppDispatch>();
  const [dataTransfer, setDataTransfer] = useState({
    amount: '',
    wallet_to: '',
    description: '',
  });
  const [successTransfer, setSuccessTransfer] = useState(false);

  const { isErrorTransfer, errorMsgTransfer } = useSelector(
    (state: RootState) => state.transactionReducer,
  );

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    if (name === 'amount') {
      const numericValue = value.replace(/[^0-9]/g, '');
      setDataTransfer((prevState) => ({
        ...prevState,
        [name]: numericValue,
      }));
    } else {
      setDataTransfer((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      dispatch(resetErrorTransfer());
      await dispatch(
        transferMoney({
          amount: Number(dataTransfer.amount),
          wallet_to: dataTransfer.wallet_to,
          description: dataTransfer.description,
        }),
      );
      setSuccessTransfer(true);
    } catch (error) {
      setSuccessTransfer(false);
    }
  };

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetErrorTransfer());
    }, 4000);
  }, [isErrorTransfer]);

  return (
    <>
      {isErrorTransfer && <AlertError message={errorMsgTransfer} />}
      <div
        role="presentation"
        className={style['modal-overlay']}
        onClick={closeModal}
      ></div>
      <div role="presentation" className={style['modal-content']}>
        {successTransfer && !isErrorTransfer ? (
          <>
            <div className={style['modal-success-container-info']}>
              <div className={style['modal-success-svg']}>
                <GreenChecklist />
              </div>
              <span className={style['modal-success-title']}>
                Transfer Success
              </span>
              <span className={style['modal-success-amount']}>
                IDR {CurrencyFormatterIDR(Number(dataTransfer.amount))}
              </span>
              <span className={style['modal-success-date']}>
                {FormatDate(new Date())}
              </span>
              <Button onClick={closeModal} variant="close-modal" type="button">
                Close
              </Button>
            </div>
          </>
        ) : (
          <>
            <span className={style['modal-header-title']}>Transfer</span>
            <form onSubmit={handleSubmit} className={style['modal-form']}>
              <input
                type="text"
                name="wallet_to"
                onChange={handleInputChange}
                placeholder="Enter destination account number"
                className={style['modal-input']}
              />
              <div className={style['input-container']}>
                <span className={style['currency-prefix']}>IDR</span>
                <input
                  type="text"
                  name="amount"
                  value={
                    dataTransfer.amount
                      ? CurrencyFormatterIDRInput(Number(dataTransfer.amount))
                      : ''
                  }
                  onChange={handleInputChange}
                  placeholder="Enter amount here"
                  className={`${style['modal-input']} ${style['input-amount']}`}
                />
              </div>
              <span className={style['modal-balance-info']}>
                Remaining Balance: IDR{' '}
                {CurrencyFormatterIDRInput(Number(dataUserProfile?.amount))}
              </span>
              <input
                type="text"
                name="description"
                onChange={handleInputChange}
                placeholder="Enter description"
                className={style['modal-input']}
              />
              <div className={style['modal-container-btn-submit']}>
                <Button
                  variant="submit-modal"
                  type="submit"
                  disabled={
                    dataTransfer.amount === '' ||
                    dataTransfer.wallet_to === '' ||
                    dataTransfer.description === ''
                  }
                >
                  Submit
                </Button>
              </div>
            </form>
          </>
        )}
      </div>
    </>
  );
};
