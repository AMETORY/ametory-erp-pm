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
    preference?: ProjectPreference;
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


export interface ProjectPreference {
    project_id: string;
    rapid_api_enabled: boolean;
    contact_enabled: boolean;
    custom_attribute_enabled: boolean;
    gemini_enabled: boolean;
    form_enabled: boolean;
}
