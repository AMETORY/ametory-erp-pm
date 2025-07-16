import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getTiktokSessions = async (
  req: PaginationRequest
) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  if (req.tag_ids) queryParams.set("tag_ids", req.tag_ids);
  if (req.connection_session)
    queryParams.set("connection_id", req.connection_session);
  if (req.is_unread) queryParams.set("is_unread", "1");
  if (req.is_unreplied) queryParams.set("is_unreplied", "1");
  return customFetch(`api/v1/tiktok/sessions?${queryParams}`, {
    method: "GET",
  });
};

export const getTiktokSessionMessages = async (
  session_id: string,
  req: PaginationRequest
) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.connection_session)
    queryParams.set("connection_id", req.connection_session);
  if (req.search) queryParams.set("search", req.search);
  return customFetch(`api/v1/tiktok/sessions/${session_id}/messages?${queryParams}`, {
    method: "GET",
  });
};
