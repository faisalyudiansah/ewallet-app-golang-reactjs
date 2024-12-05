import React from 'react';
import style from './register.module.css';
import { NavAuth } from 'src/components/layout/navbar/nav-auth';
import { HeroLeft } from 'src/components/layout/register/Hero-left';
import { HeroCenter } from 'src/components/layout/register/Hero-center';
import { HeroRight } from 'src/components/layout/register/Hero-right';

export const Register: React.FC = () => {
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
