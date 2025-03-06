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