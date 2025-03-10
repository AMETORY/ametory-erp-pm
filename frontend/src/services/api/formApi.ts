import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getFormTemplates = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/form-template/list?${queryParams}`, {
    method: "GET",
  });
};

export const getFormTemplate = async (id: string) => {
  return await customFetch(`api/v1/form-template/${id}`, {
    method: "GET",
  });
};

export const createFormTemplate = async (data: any) => {
  return await customFetch("api/v1/form-template/create", {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateFormTemplate = async (
  id: string,
  data: any
) => {
  return await customFetch(`api/v1/form-template/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteFormTemplate = async (id: string) => {
  return await customFetch(`api/v1/form-template/${id}`, {
    method: "DELETE",
  });
};
