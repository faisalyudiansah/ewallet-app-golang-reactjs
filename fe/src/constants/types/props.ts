export type PropsSocmedAuth = {
  style: {
    [key: string]: string;
  };
};

export type PropsComponentButton = {
  type?: 'button' | 'submit' | 'reset';
  variant?:
    | 'primary'
    | 'secondary'
    | 'login'
    | 'register'
    | 'edit-profile'
    | 'submit-modal'
    | 'close-modal';
  customClass?: string | string[];
  onClick?: () => void;
  children: React.ReactNode;
  disabled?: boolean;
};

export type PropsColorMenuSvg = {
  color: string;
};

export type PropsCalculatorSvg = {
  width: number;
  height: number;
  customClass?: string;
};

export type PropsConfigTransactions = {
  sortBy: string;
  sortDir: string;
  transactionType: number;
  limit: number;
  page: number;
  start_date: string;
  end_date: string;
};
