import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const createProduct = async (data: any) => {
  return await customFetch("api/v1/product/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};

export const getProducts = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/product/list?${queryParams}`, {
    method: "GET",
  });
};

export const getProduct = async (id: string) => {
  return await customFetch(`api/v1/product/${id}`, {
    method: "GET",
  });
};

export const updateProduct = async (id: string, data: any) => {
  return await customFetch(`api/v1/product/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};

export const deleteProduct = async (id: string) => {
  return await customFetch(`api/v1/product/${id}`, {
    method: "DELETE",
  });
};
