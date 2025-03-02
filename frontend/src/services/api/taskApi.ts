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

export const getTask = async (id: string) => {
  return await customFetch(`api/v1/task/${id}`, {
    method: "GET",
  });
};

export const updateTask = async (id: string, task: any) => {
  return await customFetch(`api/v1/task/${id}`, {
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
