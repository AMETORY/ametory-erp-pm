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

export const generateContent = async (content: string, agentId: string) => {
  return await customFetch(`api/v1/gemini/generate?agent_id=${agentId}`, {
    method: "POST",
    body: JSON.stringify({ content }),
  });
};
