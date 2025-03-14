import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const createTask = async (projectId: string, task: any) => {
  return await customFetch(`api/v1/project/${projectId}/task/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(task),
  });
};

export const getTasks = async (projectId: string, columnId: string) => {
  return await customFetch(
    `api/v1/project/${projectId}/task/list?column_id=${columnId}`,
    {
      method: "GET",
    }
  );
};
export const getMyTasks = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/task/my?${queryParams}`, {
    method: "GET",
  });
};
export const getMyWatchedTasks = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/task/watched?${queryParams}`, {
    method: "GET",
  });
};

export const getTask = async (projectId: string, id: string) => {
  return await customFetch(`api/v1/project/${projectId}/task/${id}/detail`, {
    method: "GET",
  });
};
export const moveTask = async (projectId: string, id: string, data: any) => {
  return await customFetch(`api/v1/project/${projectId}/task/${id}/move`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};
export const rearrangeTask = async (projectId: string, data: any) => {
  return await customFetch(`api/v1/project/${projectId}/task/rearrange`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

export const addComment = async (projectId: string, id: string, data: any) => {
  return await customFetch(`api/v1/project/${projectId}/task/${id}/comment`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const updateTask = async (projectId: string, id: string, task: any) => {
  return await customFetch(`api/v1/project/${projectId}/task/${id}/update`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(task),
  });
};

export const deleteTask = async (id: string) => {
  return await customFetch(`api/v1/task/${id}`, {
    method: "DELETE",
  });
};

export const addTaskPlugin = async (
  projectId: string,
  taskId: string,
  data: any
) => {
  return await customFetch(
    `api/v1/project/${projectId}/task/${taskId}/add-plugin`,
    {
      method: "PUT",
      body: JSON.stringify(data),
    }
  );
};
export const getTaskPlugins = async (projectId: string, taskId: string) => {
  return await customFetch(
    `api/v1/project/${projectId}/task/${taskId}/get-plugins`,
    {}
  );
};
export const getTaskPluginData = async (
  projectId: string,
  taskId: string,
  pluginId: string
) => {
  return await customFetch(
    `api/v1/project/${projectId}/task/${taskId}/plugin/${pluginId}`,
    {}
  );
};


// /api/v1/project/:id/task/:taskId/plugin/:pluginId
//  /api/v1/project/5f9bdc52-c556-4d69-8f1f-e3b52b626094/task/78b71d13-a06a-471d-8252-9e9294ede71b/plugin/05e9ff1e-d3ef-4bae-b49a-a10870ccf5d7