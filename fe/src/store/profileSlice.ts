import { createSlice } from '@reduxjs/toolkit';
import {
  DataUserProfile,
  ResErrorUpdateProfile,
  ResUserProfile,
} from 'src/constants/response/resProfile';
import { Dispatch } from 'redux';
import { UpdateProfile } from 'src/constants/types/typeProfile';

interface AppState {
  dataUserProfile: DataUserProfile | null;
  isUpdateProfileError: boolean;
  isUpdateProfileSuccess: boolean;
  errorUpdateProfileMsg: string | null;
}

const initialState: AppState = {
  dataUserProfile: null,
  isUpdateProfileError: false,
  isUpdateProfileSuccess: false,
  errorUpdateProfileMsg: null,
};

export const appSliceProfile = createSlice({
  name: 'sea_wallet_profile',
  initialState: initialState,
  reducers: {
    changeDataUserProfile: (state, action) => {
      state.dataUserProfile = action.payload;
    },
    changeIsUpdateProfileError: (state, action) => {
      state.isUpdateProfileError = action.payload;
    },
    changeIsUpdateProfileSuccess: (state, action) => {
      state.isUpdateProfileSuccess = action.payload;
    },
    changeErrorUpdateProfileMsg: (state, action) => {
      state.errorUpdateProfileMsg = action.payload;
    },
  },
});

export const {
  changeDataUserProfile,
  changeIsUpdateProfileSuccess,
  changeErrorUpdateProfileMsg,
  changeIsUpdateProfileError,
} = appSliceProfile.actions;

export const getMe = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link = process.env.REACT_APP_BASE_URL + `/users/me`;
      const response = await fetch(link, {
        method: 'GET',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
      });
      const data: ResUserProfile = await response.json();
      dispatch(changeDataUserProfile(data.data));
    } catch (error) {
      console.log(error);
      throw error;
    }
  };
};

export const updateProfile = (input: UpdateProfile) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      const link = process.env.REACT_APP_BASE_URL + `/users/me`;
      const response = await fetch(link, {
        method: 'POST',
        headers: {
          Authorization: 'Bearer ' + localStorage.access_token,
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ResErrorUpdateProfile = await response.json();
        if (!errorMsg.details) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.details[0].message);
      }
      dispatch(changeIsUpdateProfileSuccess(true));
    } catch (error) {
      dispatch(changeIsUpdateProfileError(true));
      if (error instanceof Error) {
        dispatch(changeErrorUpdateProfileMsg(error.message));
      } else {
        dispatch(
          changeErrorUpdateProfileMsg('Something wrong with the server'),
        );
      }
      throw error;
    }
  };
};

export const resetUpdateProfileError = () => {
  return async (dispatch: Dispatch): Promise<void> => {
    dispatch(changeIsUpdateProfileError(false));
    dispatch(changeErrorUpdateProfileMsg(null));
  };
};

export default appSliceProfile.reducer;
