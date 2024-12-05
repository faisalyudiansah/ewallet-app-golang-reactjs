import { createSlice } from '@reduxjs/toolkit';
import { InputLogin, InputRegister } from 'src/constants/types/typeUser';
import { ErrorAuth, ResAuth } from 'src/constants/response/resAuth';
import { Dispatch } from 'redux';

interface AppState {
  isRegisterError: boolean;
  isRegisterSuccess: boolean;
  errorRegisterMsg: string | null;
  isLoginError: boolean;
  errorLoginMsg: string | null;
}

const initialState: AppState = {
  isRegisterError: false,
  isRegisterSuccess: false,
  errorRegisterMsg: null,
  isLoginError: false,
  errorLoginMsg: null,
};

export const appSlice = createSlice({
  name: 'sea_wallet_auth',
  initialState: initialState,
  reducers: {
    setIsRegisterError: (state, action) => {
      state.isRegisterError = action.payload;
    },
    setErrorRegisterMsg: (state, action) => {
      state.errorRegisterMsg = action.payload;
    },
    setRegisterSuccess: (state, action) => {
      state.isRegisterSuccess = action.payload;
    },
    setIsLoginError: (state, action) => {
      state.isLoginError = action.payload;
    },
    setErrorLoginMsg: (state, action) => {
      state.errorLoginMsg = action.payload;
    },
  },
});

export const {
  setIsRegisterError,
  setErrorRegisterMsg,
  setRegisterSuccess,
  setIsLoginError,
  setErrorLoginMsg,
} = appSlice.actions;

export const login = (input: InputLogin) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link = process.env.REACT_APP_BASE_URL + `/auth/login`;
      const response = await fetch(link, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ErrorAuth = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
      const data: ResAuth = await response.json();
      localStorage.access_token = data.data.token;
    } catch (error) {
      dispatch(setIsLoginError(true));
      if (error instanceof Error) {
        dispatch(setErrorLoginMsg(error.message));
      } else {
        dispatch(setErrorLoginMsg('Something wrong with the server'));
      }
      throw error;
    }
  };
};

export const register = (input: InputRegister) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      let response = await fetch(
        process.env.REACT_APP_BASE_URL + `/auth/register`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(input),
        },
      );
      if (!response.ok) {
        const errorMsg: ErrorAuth = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }

      response = await fetch(process.env.REACT_APP_BASE_URL + `/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: input.email,
          password: input.password,
        }),
      });
      if (!response.ok) {
        const errorMsg: ErrorAuth = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
      const data: ResAuth = await response.json();
      localStorage.access_token = data.data.token;
      dispatch(setIsRegisterError(false));
      dispatch(setErrorRegisterMsg(null));
      dispatch(setRegisterSuccess(true));
    } catch (error) {
      dispatch(setIsRegisterError(true));
      if (error instanceof Error) {
        dispatch(setErrorRegisterMsg(error.message));
      } else {
        dispatch(setErrorRegisterMsg('Something wrong with the server'));
      }
      throw error;
    }
  };
};

export const resetRegisterError = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(setIsRegisterError(false));
    dispatch(setErrorRegisterMsg(null));
  };
};

export const resetLoginError = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(setIsLoginError(false));
    dispatch(setErrorLoginMsg(null));
  };
};

export default appSlice.reducer;
