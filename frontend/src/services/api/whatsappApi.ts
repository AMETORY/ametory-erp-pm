import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getWhatsappSessions = async (session_id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/whatsapp/sessions?${queryParams}`, {
    method: "GET",
  });
};


export const getWhatsappSessionDetail = async (session_id: string) => {
  return await customFetch(`api/v1/whatsapp/sessions/${session_id}`, {
    method: "GET",
  });
};


export const getWhatsappMessages = async (session_id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/whatsapp/messages?${queryParams}`, {
    method: "GET",
  });
};
