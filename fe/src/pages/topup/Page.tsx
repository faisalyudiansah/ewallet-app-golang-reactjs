import React, { useEffect, useState } from 'react';
import { AppDispatch } from 'src/store';
import { getMe } from 'src/store/profileSlice';
import { useDispatch } from 'react-redux';
import { NavProfile } from 'src/components/layout/navbar/nav-profile';
import { ModalTopUp } from 'src/components/ui/modal/ModalTopUp';
import style from './topup.module.css';
import { Button } from 'src/components/ui/button/Button';

export const TopUp: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [isOpenModalTopup, setIsOpenModalTopup] = useState(true);

  function getDataUserProfile() {
    dispatch(getMe());
  }

  function openModalTopup() {
    setIsOpenModalTopup(true);
  }

  function closeModalTopUp() {
    setIsOpenModalTopup(false);
  }

  useEffect(() => {
    getDataUserProfile();
  }, []);

  return (
    <>
      <NavProfile navTitle={'Top Up'} />
      <div className={style['section-topup']}>
        <Button variant="primary" onClick={openModalTopup}>
          TopUp
        </Button>
        {isOpenModalTopup && <ModalTopUp closeModal={closeModalTopUp} />}
      </div>
    </>
  );
};
