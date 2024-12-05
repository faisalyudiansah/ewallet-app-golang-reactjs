import React from 'react';

export const ArrowLeft: React.FC<{ color?: string }> = ({
  color = '#282828',
}) => {
  return (
    <svg
      width="14"
      height="11"
      viewBox="0 0 14 11"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M5 10.5V6.5H13.92L13.95 4.49H5V0.5L0 5.5L5 10.5Z" fill={color} />
    </svg>
  );
};
