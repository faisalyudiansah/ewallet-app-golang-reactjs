import React, { ChangeEvent, useEffect, useState } from 'react';
import style from './register.module.css';
import { Button } from 'src/components/ui/button/Button';
import {
  InputRegister,
  TouchedFieldsRegister,
} from 'src/constants/types/typeUser';
import { SocmedAuth } from 'src/components/ui/socmedAuth/SocmedAuth';
import { InputPasswordHiddenSvg } from 'src/assets/svg/inputPasswordHiddenSvg';
import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { register, resetRegisterError } from 'src/store/authSlice';
import { AppDispatch, RootState } from 'src/store';
import { validateFormRegister } from 'src/helpers/validateFormAuth';
import { AlertError } from 'src/components/ui/alert/AlertError';

export const HeroRight: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();
  const [inputRegister, setInputRegister] = useState<InputRegister>({
    email: '',
    full_name: '',
    password: '',
    confirmPassword: '',
    username: '',
    errors: {},
  });
  const [showPassword, setShowPassword] = useState(false);
  const [showPasswordConfirm, setShowPasswordConfirm] = useState(false);
  const [touchedFields, setTouchedFields] = useState<TouchedFieldsRegister>({});
  const [showAlert, setShowAlert] = useState(false);

  const { isRegisterError, errorRegisterMsg } = useSelector(
    (state: RootState) => state.authReducer,
  );

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  const togglePasswordConfirmVisibility = () => {
    setShowPasswordConfirm(!showPasswordConfirm);
  };

  const handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setInputRegister({
      ...inputRegister,
      [name]: value,
    });
  };

  const handleBlur = (e: ChangeEvent<HTMLInputElement>) => {
    const { name } = e.target;
    setTouchedFields({
      ...touchedFields,
      [name]: true,
    });
    validateFormRegister(inputRegister, setInputRegister);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    dispatch(resetRegisterError());
    const allTouchedFields: TouchedFieldsRegister = {
      email: true,
      full_name: true,
      password: true,
      confirmPassword: true,
      username: true,
    };
    setTouchedFields(allTouchedFields);
    if (validateFormRegister(inputRegister, setInputRegister)) {
      try {
        await dispatch(register(inputRegister));
        navigate('/');
      } catch (error) {
        setShowAlert(true);
      }
    }
  };

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetRegisterError());
      setShowAlert(false);
    }, 4000);
  }, [isRegisterError]);

  return (
    <>
      {isRegisterError && showAlert && (
        <AlertError message={errorRegisterMsg} />
      )}
      <div className={style['container-auth']}>
        <span className={style['title-input']}>Register</span>
        <form onSubmit={handleSubmit} className={style['form-input-auth']}>
          <div className={style['container-input']}>
            <div className={style['input-and-error']}>
              <input
                onChange={handleInput}
                onBlur={handleBlur}
                className={`${style['input-register']} ${
                  inputRegister.errors?.email && touchedFields.email
                    ? style['input-error']
                    : ''
                }`}
                type="text"
                name="email"
                placeholder="Enter Email"
              />
              {inputRegister.errors?.email && touchedFields.email && (
                <span className={style['error-text']}>
                  {inputRegister.errors?.email}
                </span>
              )}
            </div>

            <div className={style['input-and-error']}>
              <input
                onChange={handleInput}
                onBlur={handleBlur}
                className={`${style['input-register']} ${
                  inputRegister.errors?.full_name && touchedFields.full_name
                    ? style['input-error']
                    : ''
                }`}
                type="text"
                name="full_name"
                placeholder="Enter full name"
              />
              {inputRegister.errors?.full_name && touchedFields.full_name && (
                <span className={style['error-text']}>
                  {inputRegister.errors?.full_name}
                </span>
              )}
            </div>

            <div className={style['input-and-error']}>
              <input
                onChange={handleInput}
                onBlur={handleBlur}
                className={`${style['input-register']} ${
                  inputRegister.errors?.username && touchedFields.username
                    ? style['input-error']
                    : ''
                }`}
                type="text"
                name="username"
                placeholder="Enter username"
              />
              {inputRegister.errors?.username && touchedFields.username && (
                <span className={style['error-text']}>
                  {inputRegister.errors?.username}
                </span>
              )}
            </div>

            <div className={style['input-and-error']}>
              <div className={style['container-input']}>
                <input
                  onChange={handleInput}
                  onBlur={handleBlur}
                  className={`${style['input-register']} ${
                    style['input-password']
                  } ${
                    inputRegister.errors?.password && touchedFields.password
                      ? style['input-error']
                      : ''
                  }`}
                  type={showPassword ? 'text' : 'password'}
                  name="password"
                  placeholder="Password"
                />
                <InputPasswordHiddenSvg
                  className={style['icon-password']}
                  onClick={togglePasswordVisibility}
                />
              </div>
              {inputRegister.errors?.password && touchedFields.password && (
                <span className={style['error-text']}>
                  {inputRegister.errors?.password}
                </span>
              )}
            </div>

            <div className={style['input-and-error']}>
              <div className={style['container-input']}>
                <input
                  onChange={handleInput}
                  onBlur={handleBlur}
                  className={`${style['input-register']} ${
                    style['input-password']
                  } ${
                    inputRegister.errors?.confirmPassword &&
                    touchedFields.confirmPassword
                      ? style['input-error']
                      : ''
                  }`}
                  type={showPasswordConfirm ? 'text' : 'password'}
                  name="confirmPassword"
                  placeholder="Confirm password"
                />
                <InputPasswordHiddenSvg
                  className={style['icon-password']}
                  onClick={togglePasswordConfirmVisibility}
                />
              </div>
              {inputRegister.errors?.confirmPassword &&
                touchedFields.confirmPassword && (
                  <span className={style['error-text']}>
                    {inputRegister.errors?.confirmPassword}
                  </span>
                )}
            </div>
          </div>
          <span className={style['divider']}></span>
          <Button type="submit" variant="register">
            Register
          </Button>
        </form>
        <div className={style['flex-column-secondary']}>
          <span>Already have an account?</span>
          <span>
            You can{' '}
            <Link to={'/login'} className={style['sub-title']}>
              Login Here !
            </Link>
          </span>
        </div>
        <SocmedAuth style={style} />
      </div>
    </>
  );
};
