import React, { useState, useEffect } from 'react';
import style from './button.module.css';
import { PropsComponentButton } from '@/constants/types/props';

export const Button: React.FC<PropsComponentButton> = ({
  type = 'button',
  variant = 'primary',
  customClass = '',
  onClick,
  children,
  disabled,
}) => {
  const [buttonClass, setButtonClass] = useState('');

  useEffect(() => {
    let className = '';
    switch (variant) {
      case 'primary':
        className = style['button-primary'];
        break;
      case 'secondary':
        className = style['button-secondary'];
        break;
      case 'login':
        className = style['button-login'];
        break;
      case 'register':
        className = style['button-register'];
        break;
      case 'edit-profile':
        className = style['button-editProfile'];
        break;
      case 'submit-modal':
        className = style['button-submitModal'];
        break;
      case 'close-modal':
        className = style['button-closeModal'];
        break;
      default:
        className = '';
    }
    const customClasses = Array.isArray(customClass)
      ? customClass.join(' ')
      : customClass;
    setButtonClass(`${style.button} ${className} ${customClasses}`);
  }, [variant, customClass]);

  return (
    <button
      type={type}
      disabled={disabled}
      className={buttonClass}
      onClick={onClick}
    >
      {children}
    </button>
  );
};
