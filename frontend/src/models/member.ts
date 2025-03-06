import { CompanyModel } from "./company";
import { RoleModel } from "./role";
import { TeamModel } from "./team";
import { UserModel } from "./user";

export interface MemberModel {
    id?: string;
    company_id?: string;
    company?: CompanyModel;
    user_id: string;
    user: UserModel;
    role_id?: string;
    role?: RoleModel;
    team_id?: string;
    team?: TeamModel;
}



export interface MemberInvitationModel {
    id: string;
    company_id: string;
    user_id: string;
    user: UserModel;
    full_name: string;
    role_id: string;
    role: RoleModel;
    inviter_id: string;
    inviter: UserModel;
    expired_at: string;
    email: string;
}

