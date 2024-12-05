import React from 'react';
import style from './modal.module.css';
import { Button } from 'src/components/ui/button/Button';
import { useNavigate } from 'react-router-dom';

export const ModalConfirmLogout: React.FC<{
  closeModal: () => void;
}> = ({ closeModal }) => {
  const navigate = useNavigate();

  function logout() {
    localStorage.removeItem('access_token');
    navigate('/login');
  }

  return (
    <>
      <div
        role="presentation"
        className={style['modal-overlay']}
        onClick={closeModal}
      ></div>
      <div role="presentation" className={style['modal-content']}>
        <Button onClick={logout} variant="primary" type="button">
          Logout
        </Button>
        <Button onClick={closeModal} variant="secondary" type="button">
          Cancel
        </Button>
      </div>
    </>
  );
};
