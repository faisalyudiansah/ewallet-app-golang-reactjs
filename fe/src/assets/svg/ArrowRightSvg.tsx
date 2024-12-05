import React from 'react';

export const ArrowRight: React.FC<{ color?: string }> = ({
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
      <path
        d="M9 10.4399V6.43994H0.0800028L0.0500031 4.42994H9V0.439941L14 5.43994L9 10.4399Z"
        fill={color}
      />
    </svg>
  );
};
