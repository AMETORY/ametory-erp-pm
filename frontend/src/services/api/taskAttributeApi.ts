import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getTaskAttributes = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/task-attribute/list?${queryParams}`, {
    method: "GET",
  });
};


export const getTaskAttributeDetail = async (id: string) => {
  return await customFetch(`api/v1/task-attribute/${id}`, {
    method: "GET",
  });
};

export const createTaskAttribute = async (data: any) => {
  return await customFetch(`api/v1/task-attribute/create`, {
    method: "POST",

    body: JSON.stringify(data),
  });
};

export const updateTaskAttribute = async (id: string, data: any) => {
  return await customFetch(`api/v1/task-attribute/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteTaskAttribute = async (id: string) => {
  return await customFetch(`api/v1/task-attribute/${id}`, {
    method: "DELETE",
  });
};
