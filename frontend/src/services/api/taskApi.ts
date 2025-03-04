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
  return await customFetch(`api/v1/project/${projectId}/task/list?column_id=${columnId}`, {
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
