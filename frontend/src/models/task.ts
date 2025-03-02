import { ColumnModel } from "./column";
import { FileModel } from "./file";
import { MemberModel } from "./member";
import { ProjectModel } from "./project";

export interface TaskModel {
  id?: string;
  name?: string;
  description?: string;
  project_id?: string;
  project?: ProjectModel;
  column_id?: string;
  column?: ColumnModel;
  created_by_id?: string;
  created_by?: MemberModel;
  assignee_id?: string;
  assignee?: MemberModel;
  parent_id?: string;
  parent?: TaskModel;
  children?: TaskModel[];
  order?: number;
  status?: string;
  completed?: boolean;
  completed_date?: Date;
  start_date?: Date;
  end_date?: Date;
  files?: FileModel[];
  cover?: FileModel;
  watchers?: MemberModel[];
}
