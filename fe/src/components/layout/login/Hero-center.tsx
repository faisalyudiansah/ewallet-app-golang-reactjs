import React from 'react';
import style from './login.module.css';
import loginPageImg from 'src/assets/login/login-image.png';

export const HeroCenter: React.FC = () => {
  return (
    <div>
      <img className={style['image']} src={loginPageImg} alt="login-img" />
    </div>
  );
};
