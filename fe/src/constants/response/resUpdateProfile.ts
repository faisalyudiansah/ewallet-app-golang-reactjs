export type DataUpdatedProfile = {
  email: string;
  full_name: string;
  profile_image: string;
};

export type ResUpdateProfile = {
  message: string;
  data: DataUpdatedProfile;
};
