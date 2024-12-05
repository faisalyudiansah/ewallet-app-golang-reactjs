export type InputLogin = {
  email: string;
  password: string;
  errors?: {
    email?: string;
    password?: string;
  };
};

export type TouchedFieldsLogin = {
  email?: boolean;
  password?: boolean;
  error?: {
    email?: boolean;
    password?: boolean;
  };
};

export type InputRegister = {
  email: string;
  username: string;
  password: string;
  confirmPassword: string;
  full_name: string;
  errors?: {
    email?: string;
    username?: string;
    password?: string;
    confirmPassword?: string;
    full_name?: string;
  };
};

export type TouchedFieldsRegister = {
  email?: boolean;
  full_name?: boolean;
  password?: boolean;
  confirmPassword?: boolean;
  username?: boolean;
  errors?: {
    email?: boolean;
    full_name?: boolean;
    password?: boolean;
    confirmPassword?: boolean;
    username?: boolean;
  };
};
