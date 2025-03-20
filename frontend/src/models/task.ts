import { ColumnModel } from "./column";
import { FileModel } from "./file";
import { FormSection } from "./form";
import { MemberModel } from "./member";
import { ProjectActivityModel, ProjectModel } from "./project";
import { TaskAttributeModel } from "./task_attribute";

export interface TaskModel {
  id?: string;
  name?: string;
  description?: string;
  project_id?: string;
  project?: ProjectModel;
  column_id?: string;
  task_attribute_id?: string | null;
  task_attribute?: TaskAttributeModel;
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
  comments?: TaskCommentModel[];
  activities?: ProjectActivityModel[];
  comment_count?: number
  percentage?: number;
  priority?: string
  severity?: string
  form_response?: {
    sections: FormSection[]
  }
}

export interface TaskCommentModel {
  id?: string;
  task_id?: string;
  member_id?: string;
  member?: MemberModel;
  comment?: string;
  status?: string;
  published_at?: Date;
}

