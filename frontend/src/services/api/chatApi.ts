import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getChannels = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return customFetch(`api/v1/chat/channels?${queryParams}`, {
    method: "GET",
  });
};

export const getChannelMessages = async (
  channelID: string,
  req: PaginationRequest
) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  return customFetch(
    `api/v1/chat/channel/${channelID}/messages?${queryParams}`,
    {
      method: "GET",
    }
  );
};

export const getChannelDetail = async (channelID: string) => {
  return customFetch(`api/v1/chat/channel/${channelID}`, {
    method: "GET",
  });
};

export const createChannel = async (channel: any) => {
  return customFetch(`api/v1/chat/channel`, {
    method: "POST",
    body: JSON.stringify(channel),
  });
};

export const createMessage = async (channelID: string, message: any) => {
  return customFetch(`api/v1/chat/channel/${channelID}/message`, {
    method: "POST",
    body: JSON.stringify(message),
  });
};

export const getMessageDetail = async (
  channelID: string,
  messageID: string
) => {
  return customFetch(`api/v1/chat/channel/${channelID}/message/${messageID}`, {
    method: "GET",
  });
};

export const updateChannel = async (channelID: string, channel: any) => {
  return customFetch(`api/v1/chat/channel/${channelID}`, {
    method: "PUT",
    body: JSON.stringify(channel),
  });
};

export const deleteChannel = async (channelID: string) => {
  return customFetch(`api/v1/chat/channel/${channelID}`, {
    method: "DELETE",
  });
};

export const deleteMessage = async (channelID: string, messageID: string) => {
  return customFetch(`api/v1/chat/channel/${channelID}/message/${messageID}`, {
    method: "DELETE",
  });
};
