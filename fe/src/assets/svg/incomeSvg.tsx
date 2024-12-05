import React from 'react';

export const IncomeSvg: React.FC<{ customClass?: string }> = ({
  customClass,
}) => {
  return (
    <svg
      className={customClass}
      width="74"
      height="55"
      viewBox="0 0 74 55"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M69.4236 24.3786C70.1528 16.4802 70.1534 12.142 69.4236 4.58849C62.4016 3.80363 58.3695 3.80405 51.027 4.58849"
        stroke="black"
        strokeWidth="7.81003"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M69.3987 4.68872L39.9619 36.3559L28.8967 24.4524L4 51.2352"
        stroke="black"
        strokeWidth="6.84858"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
};
