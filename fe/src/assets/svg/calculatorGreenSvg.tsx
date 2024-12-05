import { PropsCalculatorSvg } from '@/constants/types/props';
import React from 'react';

export const CalculatorGreenSvg: React.FC<PropsCalculatorSvg> = ({
  width,
  height,
  customClass,
}) => {
  return (
    <svg
      className={customClass}
      width={width}
      height={height}
      viewBox="0 0 28 30"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M0 0V1.875V5.625H1.74296V1.875H26.1444V11.25H27.8873V0H0ZM0 7.5V9.375V30H1.74296H10.4577V28.125H1.74296V9.375H15.6866V11.25H17.4296V7.5H0ZM3.48592 13.125V15H10.4577V13.125H3.48592ZM3.48592 18.75V20.625H6.97183V18.75H3.48592ZM3.48592 22.5V24.375H10.4577V22.5H3.48592Z"
        fill="#4D4D4D"
      />
      <path
        d="M12.2007 13.125V15V30H13.9437H26.1444H27.8873V15V13.125H12.2007ZM13.9437 15H26.1444V28.125H13.9437V15ZM15.6866 16.875V20.625H24.4014V16.875H15.6866ZM19.1725 22.5V24.375H20.9155V22.5H19.1725ZM22.6584 22.5V24.375H24.4014V22.5H22.6584Z"
        fill="#27AE60"
      />
    </svg>
  );
};
