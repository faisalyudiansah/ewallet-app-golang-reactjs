import React from 'react';
import style from './nav-auth.module.css';

export const NavAuth: React.FC = () => {
  return (
    <nav className={style['container-nav-auth']}>
      <span className={style['title-nav-auth']}>Sea Wallet</span>
    </nav>
  );
};
