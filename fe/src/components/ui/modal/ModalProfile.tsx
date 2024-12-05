import React, { useState, useRef, useEffect } from 'react';
import style from './modalProfile.module.css';
import photoProfile from 'src/assets/frieren.jpg';
import { Button } from 'src/components/ui/button/Button';
import { PencilSvg } from 'src/assets/svg/PencilSvg';
import { useDispatch, useSelector } from 'react-redux';
import {
  changeIsUpdateProfileSuccess,
  resetUpdateProfileError,
  updateProfile,
} from 'src/store/profileSlice';
import { RootState, AppDispatch } from 'src/store';
import { AlertError } from '../alert/AlertError';
import { AlertSuccess } from '../alert/alertSuccess';

export const ModalProfile: React.FC<{
  getDataUserProfile: (() => void) | undefined;
  closeModal: () => void;
}> = ({ closeModal, getDataUserProfile }) => {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const dispatch = useDispatch<AppDispatch>();
  const {
    isUpdateProfileError,
    isUpdateProfileSuccess,
    errorUpdateProfileMsg,
    dataUserProfile,
  } = useSelector((state: RootState) => state.profileReducer);
  const [showAlert, setShowAlert] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [dataUser, setDataUser] = useState({
    full_name: '',
    email: '',
  });

  const handleProfileImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      alert(`Uploaded file: ${file.name}`);
    }
  };

  const handleEditClick = () => {
    setIsEditing(true);
    setDataUser({
      full_name: '',
      email: '',
    });
  };

  const handleSaveClick = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await dispatch(updateProfile(dataUser));
      setDataUser({
        full_name: '',
        email: '',
      });
      if (getDataUserProfile) {
        getDataUserProfile();
      }
    } catch (error) {
      setShowAlert(true);
    }
  };

  const handleCancelClick = () => {
    setIsEditing(false);
    setDataUser({
      full_name: '',
      email: '',
    });
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setDataUser((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  useEffect(() => {
    setTimeout(() => {
      dispatch(resetUpdateProfileError());
      dispatch(changeIsUpdateProfileSuccess(false));
      setShowAlert(false);
    }, 3000);
  }, [isUpdateProfileError, isUpdateProfileSuccess]);

  return (
    <>
      {isUpdateProfileError && !isUpdateProfileSuccess && showAlert && (
        <AlertError message={errorUpdateProfileMsg} />
      )}
      {!isUpdateProfileError && isUpdateProfileSuccess && !showAlert && (
        <AlertSuccess message={'Profile Updated'} />
      )}
      <div
        role="presentation"
        className={style['modal-overlay']}
        onClick={closeModal}
      ></div>
      <div role="presentation" className={style['modal-content']}>
        <div className={style['modal-container-image']}>
          <img
            className={style['modal-profile-image']}
            role="presentation"
            onClick={handleProfileImageClick}
            src={photoProfile}
            alt="img-profile"
          />
          <input
            type="file"
            accept="image/*"
            ref={fileInputRef}
            onChange={handleFileChange}
            style={{ display: 'none' }}
          />
          <div
            role="presentation"
            className={style['modal-svg-container']}
            onClick={() => fileInputRef.current?.click()}
          >
            <PencilSvg customClass={style['modal-svg-pencil']} />
          </div>
        </div>
        {isEditing ? (
          <form
            onSubmit={handleSaveClick}
            className={style['modal-profile-edit-container']}
          >
            <div>
              <label
                className={style['modal-label-input-edit-profile']}
                htmlFor="email"
              >
                Email
              </label>
              <input
                id="email"
                type="email"
                name="email"
                value={!dataUser.email ? '' : dataUser.email}
                placeholder={dataUserProfile?.email}
                onChange={handleInputChange}
                className={style['modal-profile-input']}
              />
            </div>
            <div>
              <label
                className={style['modal-label-input-edit-profile']}
                htmlFor="fullname"
              >
                Full Name
              </label>
              <input
                id="fullname"
                type="text"
                name="full_name"
                value={!dataUser.full_name ? '' : dataUser.full_name}
                placeholder={dataUserProfile?.name}
                onChange={handleInputChange}
                className={style['modal-profile-input']}
              />
            </div>
            <div className={style['modal-profile-buttons']}>
              <Button
                variant="primary"
                type="submit"
                disabled={dataUser.email === '' && dataUser.full_name === ''}
              >
                Save
              </Button>
              <Button
                variant="secondary"
                type="button"
                onClick={handleCancelClick}
              >
                Cancel
              </Button>
            </div>
          </form>
        ) : (
          <>
            <div className={style['modal-profile-text-container']}>
              <span className={style['modal-profile-name']}>
                {dataUserProfile?.name}
              </span>
              <span className={style['modal-profile-email']}>
                {dataUserProfile?.email}
              </span>
            </div>
            <Button
              variant="edit-profile"
              type="button"
              onClick={handleEditClick}
            >
              Edit Profile
            </Button>
          </>
        )}
      </div>
    </>
  );
};
