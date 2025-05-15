import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const setUpWebhook = async (connectionId: string) => {
  return customFetch(`api/v1/telegram/webhook/connection/${connectionId}`, {
    method: "POST",
  });
};

export const getTelegramSessions = async (
  session_id: string,
  req: PaginationRequest
) => {
  const queryParams = new URLSearchParams();
  if (session_id != "") queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  if (req.tag_ids) queryParams.set("tag_ids", req.tag_ids);
  if (req.is_unread) queryParams.set("is_unread", "1");
  if (req.is_unreplied) queryParams.set("is_unreplied", "1");
  return customFetch(`api/v1/telegram/sessions?${queryParams}`, {
    method: "GET",
  });
};


export const getTelegramSessionDetail = async (session_id: string) => {
  return customFetch(`api/v1/telegram/sessions/${session_id}`, {
    method: "GET",
  });
};


export const getTelegramMessages = async (session_id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return customFetch(`api/v1/telegram/messages?${queryParams}`, {
    method: "GET",
  });
};


export const createTelegramMessage = async (sessionId: string, message: any) => {
  return customFetch(`api/v1/telegram/${sessionId}/message`, {
    method: "POST",
    body: JSON.stringify(message),
  });
};
