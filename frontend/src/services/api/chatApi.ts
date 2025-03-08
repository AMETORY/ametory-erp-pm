import { customFetch } from "./baseApi";

export const getChannels = async () => {
  return customFetch(`api/v1/chat/channels`, {
    method: "GET",
  });
};

export const getChannelMessages = async (channelID: string) => {
  return customFetch(`api/v1/chat/channel/${channelID}/messages`, {
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

export const getMessageDetail = async (channelID: string, messageID: string) => {
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
