import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getTemplates = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/template/list?${queryParams}`, {
    method: "GET",
  });
};


export const getTemplateDetail = async (id: string) => {
  return await customFetch(`api/v1/template/${id}`, {
    method: "GET",
  });
};

export const createTemplate = async (data: any) => {
  return await customFetch(`api/v1/template/create`, {
    method: "POST",

    body: JSON.stringify(data),
  });
};

export const updateTemplate = async (id: string, data: any) => {
  return await customFetch(`api/v1/template/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteTemplate = async (id: string) => {
  return await customFetch(`api/v1/template/${id}`, {
    method: "DELETE",
  });
};


export const addProductTemplate = async (id: string, msgId: string, data: any) => {
  return await customFetch(`api/v1/template/${id}/add-image/${msgId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};
export const addMessageTemplate = async (id: string, data: any) => {
  return await customFetch(`api/v1/template/${id}/add-message`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};
export const deleteMessageTemplate = async (id: string, messageId: string) => {
  return await customFetch(`api/v1/template/${id}/delete-message/${messageId}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
};