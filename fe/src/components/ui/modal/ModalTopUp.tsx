import React, { useEffect, useState } from 'react';
import style from './modalTopUp.module.css';
import { Button } from 'src/components/ui/button/Button';
import {
  CurrencyFormatterIDR,
  CurrencyFormatterIDRInput,
} from 'src/helpers/currencyFormatter';
import { GreenChecklist } from 'src/assets/svg/GreenChecklist';
import { FormatDate } from 'src/helpers/dataFormatter';
import { DropdownDown } from 'src/assets/svg/DropdownDown';
import { DropdownUp } from 'src/assets/svg/DropdownUp';
import { AppDispatch, RootState } from 'src/store';
import { useDispatch, useSelector } from 'react-redux';
import { resetErrorTopUp, topUpMoney } from 'src/store/transactionSlice';
import { AlertError } from '../alert/AlertError';

export const ModalTopUp: React.FC<{
  closeModal: () => void;
}> = ({ closeModal }) => {
  const dispatch = useDispatch<AppDispatch>();
  const [dataTopUp, setDataTopUp] = useState({
    source_of_funds_id: 0,
    amount: '',
  });
  const [successTopUp, setSuccessTopUp] = useState(false);
  const [dropdownSelectVisible, setDropdownSelectVisible] = useState(false);
  const dropdownSourceOfFundOptions = ['Credit Card', 'Cash', 'Rewards'];

  const { isErrorTopUp, errorMsgTopUp } = useSelector(
    (state: RootState) => state.transactionReducer,
  );

  const handleInputChange = (
    e:
      | React.ChangeEvent<HTMLInputElement>
      | React.ChangeEvent<HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    if (name === 'amount') {
      const numericValue = value.replace(/[^0-9]/g, '');
      setDataTopUp((prevState) => ({
        ...prevState,
        [name]: numericValue,
      }));
    } else {
      setDataTopUp((prevState) => ({
        ...prevState,
        [name]: value,
      }));
    }
  };

  const toggleDropdownSelect = () => {
    setDropdownSelectVisible(!dropdownSelectVisible);
  };

  const handleOptionClick = (option: number) => {
    setDataTopUp((prevState) => ({
      ...prevState,
      source_of_funds_id: option,
    }));
    setDropdownSelectVisible(false);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      dispatch(resetErrorTopUp());
      await dispatch(
        topUpMoney({
          source_of_funds_id: dataTopUp.source_of_funds_id,
          amount: Number(dataTopUp.amount),
        }),
      );
      setSuccessTopUp(true);
    } catch (error) {
      setSuccessTopUp(false);
    }
  };

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetErrorTopUp());
    }, 4000);
  }, [isErrorTopUp]);

  return (
    <>
      {isErrorTopUp && <AlertError message={errorMsgTopUp} />}
      <div
        role="presentation"
        className={style['modal-overlay']}
        onClick={closeModal}
      ></div>
      <div role="presentation" className={style['modal-content']}>
        {!successTopUp ? (
          <>
            <span className={style['modal-header-title']}>Top Up</span>
            <form onSubmit={handleSubmit} className={style['modal-form']}>
              <div className={style['modal-container-dropdown-select']}>
                <div
                  role="presentation"
                  onClick={toggleDropdownSelect}
                  className={`${style['modal-dropdown-select']} ${
                    dataTopUp.source_of_funds_id !== 0
                      ? style['text-black']
                      : ''
                  }`}
                >
                  {dropdownSourceOfFundOptions[
                    dataTopUp.source_of_funds_id - 1
                  ] || 'Choose source of funds'}
                  {dropdownSelectVisible ? <DropdownUp /> : <DropdownDown />}
                </div>
                {dropdownSelectVisible && (
                  <>
                    <div>
                      <div
                        role="presentation"
                        className={style['modal-dropdown-select-overlay']}
                        onClick={() => setDropdownSelectVisible(false)}
                      ></div>
                      <div className={style['modal-option-select']}>
                        {dropdownSourceOfFundOptions.map((option, i) => (
                          <span
                            role="presentation"
                            key={i + 1}
                            onClick={() => handleOptionClick(i + 1)}
                            className={style['modal-option-item']}
                          >
                            {option}
                          </span>
                        ))}
                      </div>
                    </div>
                  </>
                )}
              </div>
              <div className={style['input-container']}>
                <span className={style['currency-prefix']}>IDR</span>
                <input
                  type="text"
                  name="amount"
                  value={
                    dataTopUp.amount
                      ? CurrencyFormatterIDRInput(Number(dataTopUp.amount))
                      : ''
                  }
                  onChange={handleInputChange}
                  placeholder="Enter amount here"
                  className={`${style['modal-input']} ${style['input-amount']}`}
                />
              </div>
              <span className={style['modal-balance-info']}>
                Top Up Value must be between IDR 50,000 - IDR 10,000,000
              </span>
              <div className={style['modal-container-btn-submit']}>
                <Button
                  variant="submit-modal"
                  type="submit"
                  disabled={!dataTopUp.source_of_funds_id || !dataTopUp.amount}
                >
                  Submit
                </Button>
              </div>
            </form>
          </>
        ) : (
          <>
            <div className={style['modal-success-container-info']}>
              <div className={style['modal-success-svg']}>
                <GreenChecklist />
              </div>
              <span className={style['modal-success-title']}>
                Top Up Success
              </span>
              <span className={style['modal-success-amount']}>
                IDR {CurrencyFormatterIDR(Number(dataTopUp.amount))}
              </span>
              <span className={style['modal-success-date']}>
                {FormatDate(new Date())}
              </span>
              <Button onClick={closeModal} variant="close-modal" type="button">
                Close
              </Button>
            </div>
          </>
        )}
      </div>
    </>
  );
};
