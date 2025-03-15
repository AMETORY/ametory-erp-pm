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
}
