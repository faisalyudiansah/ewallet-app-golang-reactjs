import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Provider } from 'react-redux';
import store from './store';
import './style/main.css';
import { Layout } from './Layout';
import { Login } from './pages/login/Page';
import { Dashboard } from './pages/dashboard/Page';
import { Register } from './pages/register/Page';
import { Transactions } from './pages/transactions/Page';
import { Transfer } from './pages/transfer/Page';
import { TopUp } from './pages/topup/Page';
import { authHome, authLogin } from './helpers/loaderRouter';

const router = createBrowserRouter([
  {
    path: '/login',
    element: <Login />,
    loader: authLogin,
  },
  {
    path: '/register',
    element: <Register />,
    loader: authLogin,
  },
  {
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <Dashboard />,
        loader: authHome,
      },
      {
        path: '/transactions',
        element: <Transactions />,
        loader: authHome,
      },
      {
        path: '/transfer',
        element: <Transfer />,
        loader: authHome,
      },
      {
        path: '/top-up',
        element: <TopUp />,
        loader: authHome,
      },
    ],
  },
]);

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
