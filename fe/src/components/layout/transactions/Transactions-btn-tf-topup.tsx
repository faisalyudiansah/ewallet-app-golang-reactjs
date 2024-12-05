import React, { useState } from 'react';
import style from './transactions.module.css';
import { NewTransferSvg } from 'src/assets/svg/NewTransferSvg';
import { NewTopUpSvg } from 'src/assets/svg/NewTopUpSvg';
import { ModalTransfer } from 'src/components/ui/modal/ModalTransfer';
import { ModalTopUp } from 'src/components/ui/modal/ModalTopUp';
import { useNavigate } from 'react-router-dom';
import { DataUserProfile } from 'src/constants/response/resProfile';

export const TransactionsBtnTfTopup: React.FC<{
  dataUserProfile: DataUserProfile | null;
}> = ({ dataUserProfile }) => {
  const navigate = useNavigate();
  const [isOpenModalTransfer, setIsOpenModalTransfer] = useState(false);
  const [isOpenModalTopUp, setIsOpenModalTopUp] = useState(false);

  function openModalTransfer() {
    setIsOpenModalTransfer(true);
  }

  function closeModalTransfer() {
    setIsOpenModalTransfer(false);
    navigate(0);
  }

  function openModalTopUp() {
    setIsOpenModalTopUp(true);
  }

  function closeModalTopUp() {
    setIsOpenModalTopUp(false);
    navigate(0);
  }

  return (
    <>
      <div className={style['tx-container-btn-newTf-newTopUp']}>
        <button
          onClick={openModalTransfer}
          className={style['tx-btn-newTf-newTopUp']}
        >
          <div className={style['tx-btn-newTf-newTopUp-svg-text']}>
            <NewTransferSvg />
            <span>Transfer +</span>
          </div>
        </button>
        <button
          onClick={openModalTopUp}
          className={style['tx-btn-newTf-newTopUp']}
        >
          <div className={style['tx-btn-newTf-newTopUp-svg-text']}>
            <NewTopUpSvg />
            <span>Top Up +</span>
          </div>
        </button>
      </div>
      {isOpenModalTransfer && (
        <ModalTransfer
          closeModal={closeModalTransfer}
          dataUserProfile={dataUserProfile}
        />
      )}
      {isOpenModalTopUp && <ModalTopUp closeModal={closeModalTopUp} />}
    </>
  );
};
