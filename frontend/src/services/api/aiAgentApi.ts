import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getAiAgents = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/ai-agent/agent?${queryParams}`, {
    method: "GET",
  });
};

export const createAiAgent = async (data: any) => {
  return await customFetch(`api/v1/ai-agent/agent`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateAiAgent = async (agentId: string, data: any) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const deleteAiAgent = async (agentId: string) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}`, {
    method: "DELETE",
  });
};

export const getAiAgentDetail = async (agentId: string) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}`, {
    method: "GET",
  });
};
export const getAiAgentHistories = async (agentId: string) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}/histories`, {
    method: "GET",
  });
};
export const deleteAiAgentHistory = async (agentId: string, historyId: string) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}/history/${historyId}`, {
    method: "DELETE",
  });
};
export const toggleAiAgentHistoryModel = async (agentId: string, historyId: string) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}/history/${historyId}/toggle-model`, {
    method: "PUT",
  });
};
export const updateAiAgentHistory = async (agentId: string, historyId: string, data: any) => {
  return await customFetch(`api/v1/ai-agent/agent/${agentId}/history/${historyId}`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};


export const generateContent = async (content: string, agentId: string, skipHistory: boolean, skipSave: boolean) => {
  return await customFetch(`api/v1/ai-agent/generate?agent_id=${agentId}&skip_history=${skipHistory}&skip_save=${skipSave}`, {
    method: "POST",
    body: JSON.stringify({ content }),
  });
};
