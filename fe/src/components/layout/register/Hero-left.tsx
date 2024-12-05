import React from 'react';
import { Link } from 'react-router-dom';
import style from './register.module.css';

export const HeroLeft: React.FC = () => {
  return (
    <div className={style['container-title']}>
      <div className={style['flex-column']}>
        <span className={style['title-large']}>Join Us!</span>
        <span className={style['title-medium']}>Sea Wallet</span>
      </div>
      <div className={style['flex-column']}>
        <span>Already have an account?</span>
        <span>
          You can{' '}
          <Link to={'/login'} className={style['sub-title']}>
            Login Here !
          </Link>
        </span>
      </div>
    </div>
  );
};
