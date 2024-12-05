import { redirect } from 'react-router-dom';

export const authHome = (): null => {
  const access_token = localStorage.access_token;
  if (!access_token) {
    throw redirect('/login');
  }
  return null;
};

export const authLogin = (): null => {
  const access_token = localStorage.access_token;
  if (access_token) {
    throw redirect('/');
  }
  return null;
};
