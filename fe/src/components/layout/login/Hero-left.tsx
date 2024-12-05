import React from 'react';
import { Link } from 'react-router-dom';
import style from './login.module.css';

export const HeroLeft: React.FC = () => {
  return (
    <div className={style['container-title']}>
      <div className={style['flex-column']}>
        <span className={style['title-large']}>Sign in to </span>
        <span className={style['title-medium']}>Sea Wallet</span>
      </div>
      <div className={style['flex-column']}>
        <span>If you don’t have an account register</span>
        <span>
          You can{' '}
          <Link to={'/register'} className={style['sub-title']}>
            Register Here !
          </Link>
        </span>
      </div>
    </div>
  );
};
