import React, { useEffect, useState } from 'react';
import style from './dashboard.module.css';
import { NavProfile } from 'src/components/layout/navbar/nav-profile';
import { DashboardHeadertitle } from 'src/components/layout/dashboard/Dashboard-Header-title';
import { DashboardCardDashboard } from 'src/components/layout/dashboard/Dashboard-card-dashboard';
import { DashboardRecentTransaction } from 'src/components/layout/dashboard/Dashboard-recent-transaction';
import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from 'src/store';
import { AlertSuccessRegister } from 'src/components/ui/alert/AlertSuccessRegister';
import { getMe } from 'src/store/profileSlice';
import { getExpense } from 'src/store/transactionSlice';
import { setRegisterSuccess } from 'src/store/authSlice';
import { AlertError } from 'src/components/ui/alert/AlertError';

export const Dashboard: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { isRegisterSuccess } = useSelector(
    (state: RootState) => state.authReducer,
  );
  const { dataUserProfile } = useSelector(
    (state: RootState) => state.profileReducer,
  );
  const { expenseSumByMonth } = useSelector(
    (state: RootState) => state.transactionReducer,
  );
  const [showAlertRegister, setShowAlertSuccessRegister] = useState(false);
  const [isError, setIsError] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  function getDataUserProfile() {
    try {
      setIsLoading(true);
      dispatch(getMe());
      dispatch(getExpense());
    } catch (error) {
      console.log(error);
      setIsError(true);
    } finally {
      setIsLoading(false);
    }
  }

  useEffect(() => {
    setTimeout(() => {
      getDataUserProfile();
    }, 500);

    if (isRegisterSuccess) {
      setShowAlertSuccessRegister(true);
      setTimeout(() => {
        setShowAlertSuccessRegister(false);
      }, 4000);
      dispatch(setRegisterSuccess(false));
    }
  }, []);

  return (
    <>
      {showAlertRegister && (
        <AlertSuccessRegister message="Registered successfully" />
      )}
      {isError && <AlertError message={"Something's wrong. Try again!"} />}
      <div>
        {isLoading ? (
          <div className="loader"></div>
        ) : (
          <>
            <NavProfile
              navTitle={'Dashboard'}
              getDataUserProfile={getDataUserProfile}
            />
            <section className={style['section-dashboard']}>
              <DashboardHeadertitle dataUserProfile={dataUserProfile} />
              <DashboardCardDashboard
                dataUserProfile={dataUserProfile}
                expenseSumByMonth={expenseSumByMonth}
              />
              <DashboardRecentTransaction />
            </section>
          </>
        )}
      </div>
    </>
  );
};
