import { customFetch } from "./baseApi";

export const processRegister = async (data: any) => {
  return await customFetch("api/v1/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};

export const processLogin = async (data: any) => {
  return await customFetch("api/v1/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};


export const verifyEmail = async (token: string) => {
  return await customFetch(`api/v1/auth/verification/${token}`, {
    method: "GET",
  });
};
export const getProfile = async () => {
  return await customFetch(`api/v1/auth/profile`, {
    method: "GET",
  });
};
