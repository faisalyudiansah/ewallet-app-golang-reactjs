import { configureStore } from '@reduxjs/toolkit';
import authReducer from './authSlice';
import profileReducer from './profileSlice';
import transactionReducer from './transactionSlice';

const store = configureStore({
  reducer: {
    authReducer,
    profileReducer,
    transactionReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
