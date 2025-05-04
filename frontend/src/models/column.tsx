import { TaskModel } from "./task";

export interface ColumnModel {
  id?: string | null;
  project_id?: string;
  name?: string;
  icon?: string;
  order?: number;
  color?: string;
  tasks?: TaskModel[];
  count_tasks?: number;
  actions?: ColumnActionModel[];
}


export interface ColumnActionModel {
  id: string;
  created_at: string;
  name: string;
  column_id: string;
  action: string;
  action_trigger: string;
  action_data: any;
  status: string;
}
