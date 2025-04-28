import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getProductCategories = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(
    `api/v1/product-category/list?${queryParams}`,
    {
      method: "GET",
    }
  );
};

export const createProductCategory = async (data: any) => {
  return await customFetch("api/v1/product-category/create", {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateProductCategory = async (id: string, data: any) => {
  return await customFetch(`api/v1/product-category/${id}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteProductCategory = async (id: string) => {
  return await customFetch(`api/v1/product-category/${id}`, {
    method: "DELETE",
  });
};
