import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const sendMessage = async (data: any) => {
  return customFetch("api/v1/inbox/send", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};

export const getInboxMessagesCount = async () => {
  return customFetch("api/v1/inbox/count", {
    method: "GET",
  });
};
export const getSentMessagesCount = async () => {
  return customFetch("api/v1/inbox/count-sent", {
    method: "GET",
  });
};


export const getInboxes = async () => {
  return customFetch("api/v1/inbox/inboxes", {
    method: "GET",
  });
};

export const getInboxMessages = async (inboxID: string, req: PaginationRequest) => {
    const queryParams = new URLSearchParams();
    queryParams.set("inbox_id", String(inboxID));
    queryParams.set("page", String(req.page));
    queryParams.set("size", String(req.size));
    if (req.search) queryParams.set("search", req.search);

  return customFetch(`api/v1/inbox/messages?${queryParams}`, {
    method: "GET",
  });
};
export const getSentMessages = async ( req: PaginationRequest) => {
    const queryParams = new URLSearchParams();
    queryParams.set("page", String(req.page));
    queryParams.set("size", String(req.size));
    if (req.search) queryParams.set("search", req.search);

  return customFetch(`api/v1/inbox/sent?${queryParams}`, {
    method: "GET",
  });
};

export const getInboxMessageDetail = async (messageID: string) => {
  return customFetch(`api/v1/inbox/message/${messageID}`, {
    method: "GET",
  });
};


export const deleteMessage = async (messageID: string) => {
  return customFetch(`api/v1/inbox/message/${messageID}`, {
    method: "DELETE",
  });
};
