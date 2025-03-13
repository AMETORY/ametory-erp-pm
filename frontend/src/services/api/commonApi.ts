import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getMembers = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/members?${queryParams}`, {
    method: "GET",
  });
};
export const getRoles = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/roles?${queryParams}`, {
    method: "GET",
  });
};
export const inviteMember = async (req: any) => {
  return await customFetch(`api/v1/invite-member`, {
    method: "POST",
    body: JSON.stringify(req),
  });
};
export const getInvitedMembers = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/invited?${queryParams}`, {
    method: "GET",
  });
};
export const deleteInvitation = async (id: string) => {
  return await customFetch(`api/v1/invited/${id}`, {
    method: "DELETE",
  });
};

export const acceptInvitation = async (token: string) => {
  return await customFetch(`api/v1/accept-invitation/${token}`, {
    method: "GET",
  });
};

export const uploadFile = async (
  file: File,
  reqs: Record<string, any>,
  onProgress: (progress: number) => void
) => {
  const formData = new FormData();
  formData.append("file", file);
  Object.keys(reqs).forEach((key) => {
    formData.append(key, reqs[key]);
  });

  const options = {
    onUploadProgress: (progressEvent: ProgressEvent) => {
      const progress = Math.round(
        (progressEvent.loaded * 100) / progressEvent.total
      );
      onProgress(progress);
    },
  };

  return await customFetch("api/v1/file/upload", {
    method: "POST",
    body: formData,
    ...options,
    isMultipart: true,
  });
};




export const getSetting = async () => {
  return await customFetch(`api/v1/setting`, {
    method: "GET",
  });
};
export const updateSetting = async (data: any) => {
  return await customFetch(`api/v1/setting`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};
export const getRapidAPIPlugins = async () => {
  return await customFetch(`api/v1/rapid-api-plugins`, {
  });
};