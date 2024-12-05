import React, { ChangeEvent, useEffect, useState } from 'react';
import style from './login.module.css';
import { Button } from 'src/components/ui/button/Button';
import { InputLogin, TouchedFieldsLogin } from '@/constants/types/typeUser';
import { SocmedAuth } from 'src/components/ui/socmedAuth/SocmedAuth';
import { InputPasswordHiddenSvg } from 'src/assets/svg/inputPasswordHiddenSvg';
import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { login, resetLoginError } from 'src/store/authSlice';
import { AppDispatch, RootState } from 'src/store';
import { validateLogin } from 'src/helpers/validateFormAuth';
import { AlertError } from 'src/components/ui/alert/AlertError';

export const HeroRight: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const navigate = useNavigate();
  const [inputLogin, setInputLogin] = useState<InputLogin>({
    email: '',
    password: '',
    errors: {},
  });
  const [showPassword, setShowPassword] = useState(false);
  const [touchedFields, setTouchedFields] = useState<TouchedFieldsLogin>({});
  const [showAlert, setShowAlert] = useState(false);

  const { isLoginError, errorLoginMsg } = useSelector(
    (state: RootState) => state.authReducer,
  );

  const handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setInputLogin({
      ...inputLogin,
      [name]: value,
    });
  };

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  const handleBlur = (e: ChangeEvent<HTMLInputElement>) => {
    const { name } = e.target;
    setTouchedFields({
      ...touchedFields,
      [name]: true,
    });
    validateLogin(inputLogin, setInputLogin);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    dispatch(resetLoginError());
    const allTouchedFields: TouchedFieldsLogin = {
      email: true,
      password: true,
    };
    setTouchedFields(allTouchedFields);
    if (validateLogin(inputLogin, setInputLogin)) {
      try {
        await dispatch(login(inputLogin));
        navigate('/');
      } catch (error) {
        setShowAlert(true);
      }
    }
  };

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetLoginError());
      setShowAlert(false);
    }, 4000);
  }, [isLoginError]);

  return (
    <>
      {isLoginError && showAlert && <AlertError message={errorLoginMsg} />}
      <div className={style['container-auth']}>
        <span className={style['title-input']}>Sign in</span>
        <form onSubmit={handleSubmit} className={style['form-input-auth']}>
          <div className={style['container-input']}>
            <div className={style['input-and-error']}>
              <input
                onChange={handleInput}
                onBlur={handleBlur}
                className={`${style['input-login']} ${
                  inputLogin.errors?.email && touchedFields.email
                    ? style['input-error']
                    : ''
                }`}
                type="email"
                name="email"
                placeholder="Enter Email"
              />
              {inputLogin.errors?.email && touchedFields.email && (
                <span className={style['error-text']}>
                  {inputLogin.errors?.email}
                </span>
              )}
            </div>
            <div className={style['input-and-error']}>
              <div className={style['container-input']}>
                <input
                  onChange={handleInput}
                  onBlur={handleBlur}
                  className={`${style['input-login']} ${
                    style['input-password']
                  } ${
                    inputLogin.errors?.password && touchedFields.password
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
              {inputLogin.errors?.password && touchedFields.password && (
                <span className={style['error-text']}>
                  {inputLogin.errors?.password}
                </span>
              )}
            </div>
          </div>
          <span
            className={`${style['text-secondary']} ${style['text-forgot-password']}`}
          >
            {' '}
            Forgot Password ?
          </span>
          <Button type="submit" variant="login">
            Login
          </Button>
        </form>
        <div className={style['flex-column-secondary']}>
          <span>If you don’t have an account register</span>
          <span>
            You can{' '}
            <Link to={'/register'} className={style['sub-title']}>
              Register Here !
            </Link>
          </span>
        </div>
        <SocmedAuth style={style} />
      </div>
    </>
  );
};
