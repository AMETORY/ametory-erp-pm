import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getGeminiAgents = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/gemini/agent?${queryParams}`, {
    method: "GET",
  });
};

export const createGeminiAgent = async (data: any) => {
  return await customFetch(`api/v1/gemini/agent`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateGeminiAgent = async (agentId: string, data: any) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteGeminiAgent = async (agentId: string) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}`, {
    method: "DELETE",
  });
};

export const getGeminiAgentDetail = async (agentId: string) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}`, {
    method: "GET",
  });
};
export const getGeminiAgentHistories = async (agentId: string) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}/histories`, {
    method: "GET",
  });
};
export const deleteGeminiAgentHistory = async (agentId: string, historyId: string) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}/history/${historyId}`, {
    method: "DELETE",
  });
};
export const toggleGeminiAgentHistoryModel = async (agentId: string, historyId: string) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}/history/${historyId}/toggle-model`, {
    method: "PUT",
  });
};
export const updateGeminiAgentHistory = async (agentId: string, historyId: string, data: any) => {
  return await customFetch(`api/v1/gemini/agent/${agentId}/history/${historyId}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};


export const generateContent = async (content: string, agentId: string, skipHistory: boolean, skipSave: boolean) => {
  return await customFetch(`api/v1/gemini/generate?agent_id=${agentId}&skip_history=${skipHistory}&skip_save=${skipSave}`, {
    method: "POST",
    body: JSON.stringify({ content }),
  });
};
