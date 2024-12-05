import React, { useState } from 'react';
import style from './nav-profile.module.css';
import photoProfile from 'src/assets/frieren.jpg';
import { PersonSvg } from 'src/assets/svg/PersonSvg';
import { GoToSvg } from 'src/assets/svg/GoToSvg';
import { ModalProfile } from 'src/components/ui/modal/ModalProfile';
import { ModalConfirmLogout } from 'src/components/ui/modal/ModalConfirmLogout';

export const NavProfile: React.FC<{
  navTitle: string;
  getDataUserProfile?: () => void;
}> = ({ navTitle, getDataUserProfile }) => {
  const [modalConfirm, setModalConfirm] = useState(false);
  const [dropdownVisible, setDropdownVisible] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);

  const toggleDropdown = () => {
    setDropdownVisible(!dropdownVisible);
  };

  const openModal = () => {
    setModalVisible(true);
    setDropdownVisible(false);
  };

  const closeModal = () => {
    setDropdownVisible(false);
    setModalVisible(false);
  };

  return (
    <>
      {modalConfirm && (
        <ModalConfirmLogout closeModal={() => setModalConfirm(false)} />
      )}
      <nav className={style['nav-section']}>
        <span className={style['nav-title']}>{navTitle}</span>
        <img
          className={style['nav-profile-img']}
          src={photoProfile}
          alt="pic-profile"
          role="presentation"
          onClick={toggleDropdown}
        />
        {dropdownVisible && (
          <>
            <div
              role="presentation"
              className={style['dropdown-overlay']}
              onClick={() => setDropdownVisible(false)}
            ></div>
            <div className={style['dropdown-menu']}>
              <div
                role="presentation"
                className={style['dropdown-item']}
                onClick={openModal}
              >
                <GoToSvg customClass={style['svg']} />
                <span className={style['list-text-item']}>Profile</span>
              </div>
              <div
                role="presentation"
                className={style['dropdown-item']}
                onClick={() => {
                  setDropdownVisible(false);
                  setModalConfirm(!modalConfirm);
                }}
              >
                <PersonSvg customClass={style['svg']} />
                <span role="presentation" className={style['list-text-item']}>
                  Logout
                </span>
              </div>
            </div>
          </>
        )}
      </nav>
      {modalVisible && (
        <ModalProfile
          getDataUserProfile={getDataUserProfile}
          closeModal={closeModal}
        />
      )}
    </>
  );
};
