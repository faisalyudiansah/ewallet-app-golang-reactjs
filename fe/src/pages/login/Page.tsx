import React from 'react';
import style from './login.module.css';
import { NavAuth } from 'src/components/layout/navbar/nav-auth';
import { HeroLeft } from 'src/components/layout/login/Hero-left';
import { HeroCenter } from 'src/components/layout/login/Hero-center';
import { HeroRight } from 'src/components/layout/login/Hero-right';

export const Login: React.FC = () => {
  return (
    <>
      <NavAuth />
      <section className={style['section-auth']}>
        <HeroLeft />
        <HeroCenter />
        <HeroRight />
      </section>
    </>
  );
};
