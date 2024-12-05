export type DataUserProfile = {
  user_id: number;
  name: string;
  email: string;
  wallet_id: number;
  wallet_number: string;
  profile_image: string;
  amount: string;
};

export type ResUserProfile = {
  message: string;
  data: DataUserProfile;
};

export type ResErrorUpdateProfile = {
  message: string;
  details: [
    {
      field: string;
      message: string;
    },
  ];
};
