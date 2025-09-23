import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getWhatsappSessions = async (session_id: string, req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  if  (session_id != "") queryParams.set("session_id", session_id);
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  if (req.tag_ids) queryParams.set("tag_ids", req.tag_ids);
  if (req.connection_session) queryParams.set("connection_session", req.connection_session);
  if (req.is_unread) queryParams.set("is_unread", "1");
  if (req.is_unreplied) queryParams.set("is_unreplied", "1");
  if (req.type) queryParams.set("type", req.type);
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
export const clearWhatsappSession = async (session_id: string) => {
  return await customFetch(`api/v1/whatsapp/sessions/${session_id}/clear`, {
    method: "PUT",
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


export const markAsRead = async ( message_id: string, session_id: string) => {
  return customFetch(`api/v1/whatsapp/messages/${message_id}/read?session_id=${session_id}`, {
    method: "PUT",
  });
};

export const createWAMessage = async (sessionId: string, message: any) => {
  return customFetch(`api/v1/whatsapp/${sessionId}/message`, {
    method: "POST",
    body: JSON.stringify(message),
  });
};
export const markAllMsgAsRead = async (sessionId: string) => {
  return customFetch(`api/v1/whatsapp/${sessionId}/read-all`, {
    method: "POST",
  });
};
export const exportXls = async (data: any) => {

  return  await customFetch(`api/v1/whatsapp/export`, {
    method: "POST",
    body: JSON.stringify(data),
    isBlob: true
  });
  
};


export const getAllMessageTemplates = async (businessID: string) => {
  return await customFetch(`api/v1/whatsapp-api/message-templates/${businessID}`, {
    method: "GET",
  });
};
export const getMessageTemplateByName = async (businessID: string, templateID: string) => {
  return await customFetch(`api/v1/whatsapp-api/message-templates/${businessID}/${templateID}`, {
    method: "GET",
  });
};
