import { asyncStorage } from "../../utils/async_storage";
import {
  LOCAL_STORAGE_COMPANIES,
  LOCAL_STORAGE_COMPANY_ID,
  LOCAL_STORAGE_REMEMBER_TOKEN,
  LOCAL_STORAGE_TOKEN,
} from "../../utils/constants";

export const customFetch = async <T>(
  url: string,
  options: RequestInit & { isBlob?: boolean; isMultipart?: boolean } = {}
): Promise<T> => {
  let token = await asyncStorage.getItem(LOCAL_STORAGE_TOKEN);
  if (token) {
    const payload = JSON.parse(atob(token.split(".")[1]));
    const expired = payload.exp * 1000 < Date.now();
    if (expired) {
      const refreshToken = await asyncStorage.getItem(
        LOCAL_STORAGE_REMEMBER_TOKEN
      );

      // const response = await fetch(`${process.env.REACT_APP_BASE_URL}/auth/refresh-token`, {
      //   method: "POST",
      //   headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify({ refresh_token: refreshToken }),
      // });
      // const { token: accessToken, refresh_token } = await response.json();
      // await asyncStorage.setItem(LOCAL_STORAGE_TOKEN, accessToken);
      // await asyncStorage.setItem(LOCAL_STORAGE_REMEMBER_TOKEN, refresh_token);

      await asyncStorage.removeItem(LOCAL_STORAGE_TOKEN);
      await asyncStorage.removeItem(LOCAL_STORAGE_COMPANIES);
      await asyncStorage.removeItem(LOCAL_STORAGE_COMPANY_ID);
      window.location.reload();
    }
  }
  const response = await fetch(`${process.env.REACT_APP_BASE_URL}/${url}`, {
    ...options,
    headers: {
      // 'Content-Type': 'application/json',
      authorization: `Bearer ${token}`,
      "ID-Company": await asyncStorage.getItem(LOCAL_STORAGE_COMPANY_ID),
      ...options.headers,
    },
  });

  let data;
  if (options.isBlob) {
    data = await response.blob();
  } else {
    data = await response.json();
  }

  if (!response.ok) {
    const error = new Error(
      data.message || data.error || "Something went wrong"
    );
    // @ts-ignore
    error.response = response;
    throw error;
  }

  return data;
};
