import React from 'react';
import style from './register.module.css';
import registerPageImg from 'src/assets/register/register.png';

export const HeroCenter: React.FC = () => {
  return (
    <div>
      <img
        className={style['image']}
        src={registerPageImg}
        alt="register-img"
      />
    </div>
  );
};
