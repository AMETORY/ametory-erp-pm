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
export const deleteWhatsappSession = async (session_id: string) => {
  return await customFetch(`api/v1/whatsapp/sessions/${session_id}`, {
    method: "DELETE",
  });
};

export const updateWhatsappSession = async (session_id: string, data: any) => {
  return await customFetch(`api/v1/whatsapp/sessions/${session_id}`, {
    method: "PUT",
    body: JSON.stringify(data),
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


export const markAsRead = async ( message_id: string) => {
  return customFetch(`api/v1/whatsapp/messages/${message_id}/read`, {
    method: "PUT",
  });
};

export const createWAMessage = async (sessionId: string, message: any) => {
  return customFetch(`api/v1/whatsapp/${sessionId}/message`, {
    method: "POST",
    body: JSON.stringify(message),
  });
};