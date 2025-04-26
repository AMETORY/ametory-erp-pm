import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getBroadcasts = async (req: PaginationRequest) => {
    const queryParams = new URLSearchParams();
    queryParams.set("page", String(req.page));
    queryParams.set("size", String(req.size));
    if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/broadcast/list?${queryParams}`, {
    method: "GET",
  });
};

export const getBroadcast = async (id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/broadcast/${id}?${queryParams}`, {
    method: "GET",
  });
};

export const createBroadcast = async (data: any) => {
  return await customFetch(`api/v1/broadcast/create`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateBroadcast = async (id: string, data: any) => {
  return await customFetch(`api/v1/broadcast/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};
export const sendBroadcast = async (id: string) => {
  return await customFetch(`api/v1/broadcast/${id}/send`, {
    method: "PUT",
  });
};
export const addContactBroadcast = async (id: string, data: any) => {
  return await customFetch(`api/v1/broadcast/${id}/add-contact`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};
export const deleteContactBroadcast = async (id: string, data: any) => {
  return await customFetch(`api/v1/broadcast/${id}/delete-contact`, {
    method: "DELETE",
    body: JSON.stringify(data),
  });
};

export const deleteBroadcast = async (id: string) => {
  return await customFetch(`api/v1/broadcast/${id}`, {
    method: "DELETE",
  });
};
