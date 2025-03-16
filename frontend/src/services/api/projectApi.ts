import { customFetch } from "./baseApi";
import { ProjectModel, ProjectPreference } from "../../models/project";
import { PaginationRequest } from "../../objects/pagination";

export const createProject = async (project: ProjectModel) => {
  return await customFetch("api/v1/project/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(project),
  });
};

export const getProjects = async (req: PaginationRequest) => {
  const queryParams = new URLSearchParams();
  queryParams.set("page", String(req.page));
  queryParams.set("size", String(req.size));
  if (req.search) queryParams.set("search", req.search);
  return await customFetch(`api/v1/project/list?${queryParams}`, {
    method: "GET",
  });
};

export const getProject = async (id: string) => {
  return await customFetch(`api/v1/project/${id}`, {
    method: "GET",
  });
};
export const getProjectMembers = async (id: string) => {
  return await customFetch(`api/v1/project/${id}/members`, {
    method: "GET",
  });
};
export const getProjectAddMember = async (id: string, data: any) => {
  return await customFetch(`api/v1/project/${id}/add-member`, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const addNewColumn = async (id: string, data: any) => {
  return await customFetch(`api/v1/project/${id}/add-column`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};
export const rearrangeColumns = async (id: string, data: any) => {
  return await customFetch(`api/v1/project/${id}/rearrange-columns`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};


export const updateColumn = async (id: string, data: any) => {
  return await customFetch(`api/v1/project/${id}/update-column`, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};
export const getProjectTemplates = async () => {
  return await customFetch(`api/v1/project/templates`, {
    method: "GET",
  });
};

export const updateProject = async (id: string, project: ProjectModel) => {
  return await customFetch(`api/v1/project/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(project),
  });
};
export const updateProjectPreference = async (id: string, data: ProjectPreference) => {
  return await customFetch(`api/v1/project/${id}/preference`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
};

export const deleteProject = async (id: string) => {
  return await customFetch(`api/v1/project/${id}`, {
    method: "DELETE",
  });
};



export const countCompletedTasks = async (id: string, days = 7) => {
  return await customFetch(`api/v1/project/${id}/count-completed?days=${days}`, {
    method: "GET",
  });
};

export const countUpdatedTasks = async (id: string, days = 7) => {
  return await customFetch(`api/v1/project/${id}/count-updated?days=${days}`, {
    method: "GET",
  });
};

export const countCreatedTasks = async (id: string, days = 7) => {
  return await customFetch(`api/v1/project/${id}/count-created?days=${days}`, {
    method: "GET",
  });
};

export const countDueTasks = async (id: string, days = 7) => {
  return await customFetch(`api/v1/project/${id}/count-due-tasks?days=${days}`, {
    method: "GET",
  });
};

export const countNextDueTasks = async (id: string, days = 7) => {
  return await customFetch(`api/v1/project/${id}/count-next-due-tasks?days=${days}`, {
    method: "GET",
  });
};

export const getRecentActivities = async (id: string) => {
  return await customFetch(`api/v1/project/${id}/recent-activities`, {
    method: "GET",
  });
};

export const countTasksByPriority = async (id: string) => {
  return await customFetch(`api/v1/project/${id}/count-tasks-by-priority`, {
    method: "GET",
  });
};

export const countTasksBySeverity = async (id: string) => {
  return await customFetch(`api/v1/project/${id}/count-tasks-by-severity`, {
    method: "GET",
  });
};
export const countTasksByColumn = async (id: string) => {
  return await customFetch(`api/v1/project/${id}/count-tasks-by-column`, {
    method: "GET",
  });
};






