import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";



export const getInstagramSessions = async (
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
  return customFetch(`api/v1/facebook/instagram/sessions?${queryParams}`, {
    method: "GET",
  });
};


export const getInstagramSessionDetail = async (session_id: string) => {
  return customFetch(`api/v1/facebook/instagram/sessions/${session_id}`, {
    method: "GET",
  });
};

export const deleteInstagramSession = async (session_id: string) => {
  return customFetch(`api/v1/facebook/instagram/sessions/${session_id}`, {
    method: "DELETE",
  });
};


export const getInstagramMessages = async (session_id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return customFetch(`api/v1/facebook/instagram/messages?${queryParams}`, {
    method: "GET",
  });
};


export const createInstagramMessage = async (sessionId: string, message: any) => {
  return customFetch(`api/v1/facebook/instagram/${sessionId}/message`, {
    method: "POST",
    body: JSON.stringify(message),
  });
};



