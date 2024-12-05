import { InputLogin, InputRegister } from 'src/constants/types/typeUser';
import { validateEmail } from './validateEmail';

export const validateFormRegister = (
  inputRegister: InputRegister,
  setInputRegister: React.Dispatch<React.SetStateAction<InputRegister>>,
): boolean => {
  const errors: InputRegister['errors'] = {};
  if (!inputRegister.email) {
    errors.email = 'Email is required';
  } else if (!validateEmail(inputRegister.email)) {
    errors.email = 'Email is not valid';
  }
  if (!inputRegister.full_name) {
    errors.full_name = 'Full name is required';
  }
  if (!inputRegister.username) {
    errors.username = 'Username is required';
  }
  if (!inputRegister.password) {
    errors.password = 'Password is required';
  }
  if (!inputRegister.confirmPassword) {
    errors.confirmPassword = 'Confirm password is required';
  } else if (inputRegister.password !== inputRegister.confirmPassword) {
    if (!errors.password) {
      errors.password = 'Passwords do not match';
    }
    errors.confirmPassword = 'Passwords do not match';
  }
  setInputRegister((prevState) => ({ ...prevState, errors }));
  return !Object.keys(errors).length;
};

export const validateLogin = (
  inputLogin: InputLogin,
  setInputLogin: React.Dispatch<React.SetStateAction<InputLogin>>,
): boolean => {
  const errors: InputLogin['errors'] = {};
  if (!inputLogin.email) {
    errors.email = 'Email is required';
  } else if (!validateEmail(inputLogin.email)) {
    errors.email = 'Email is not valid';
  }
  if (!inputLogin.password) {
    errors.password = 'Password is required';
  }
  setInputLogin((prevState) => ({
    ...prevState,
    errors,
  }));
  return !Object.keys(errors).length;
};
