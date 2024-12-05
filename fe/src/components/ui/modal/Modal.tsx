import React from 'react';
import style from './modal.module.css';
import { Button } from 'src/components/ui/button/Button';

export const Modal: React.FC<{
  closeModal: () => void;
}> = ({ closeModal }) => {
  return (
    <>
      <div
        role="presentation"
        className={style['modal-overlay']}
        onClick={closeModal}
      ></div>
      <div role="presentation" className={style['modal-content']}>
        <Button variant="edit-profile" type="button">
          Hallo
        </Button>
      </div>
    </>
  );
};
