import { ColumnModel } from "./column";
import { MemberModel } from "./member";
import { TaskModel } from "./task";

export interface ProjectModel {
    id?: string;
    name: string;
    description?: string;
    deadline?: Date;
    status?: string; // e.g., "ongoing", "completed"
    columns?: ColumnModel[];
    members?: MemberModel[];
    tasks?: TaskModel[];
}
export interface IndustryColumn {
    name: string;
    icon: string;
    color: string;
}

export interface Industry {
    industry: string;
    columns: IndustryColumn[];
}



export interface ProjectActivityModel {
    id?: string;
    project_id?: string;
    project?: ProjectModel;
    member_id?: string;
    member?: MemberModel;
    activity_type?: string;
    notes?: string;
    column_id?: string;
    column?: ColumnModel;
    task_id?: string;
    task?: TaskModel;
    activity_date?: Date;
}
