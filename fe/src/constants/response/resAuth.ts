export type ResAuth = {
  message: string;
  data: {
    token?: string;
    username?: string;
    email?: string;
    full_name?: string;
    wallet_id?: string;
    wallet_number?: string;
  };
};

export type ErrorAuth = {
  message: string;
  details: [
    {
      field: string;
      message: string;
    },
  ];
};
