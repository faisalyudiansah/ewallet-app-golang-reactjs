import React, { useEffect, useState } from 'react';
import { AppDispatch, RootState } from 'src/store';
import { getMe } from 'src/store/profileSlice';
import { useDispatch, useSelector } from 'react-redux';
import { NavProfile } from 'src/components/layout/navbar/nav-profile';
import { ModalTransfer } from 'src/components/ui/modal/ModalTransfer';
import style from './transfer.module.css';
import { Button } from 'src/components/ui/button/Button';

export const Transfer: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [isOpenModalTransfer, setIsOpenModalTransfer] = useState(true);

  const { dataUserProfile } = useSelector(
    (state: RootState) => state.profileReducer,
  );

  function getDataUserProfile() {
    dispatch(getMe());
  }

  function openModalTransfer() {
    setIsOpenModalTransfer(true);
  }

  function closeModalTransfer() {
    setIsOpenModalTransfer(false);
  }

  useEffect(() => {
    getDataUserProfile();
  }, []);

  return (
    <>
      <NavProfile navTitle={'Transfer'} />
      <div className={style['section-transfer']}>
        <Button variant="primary" onClick={openModalTransfer}>
          Transfer
        </Button>
        {isOpenModalTransfer && (
          <ModalTransfer
            closeModal={closeModalTransfer}
            dataUserProfile={dataUserProfile}
          />
        )}
      </div>
    </>
  );
};
