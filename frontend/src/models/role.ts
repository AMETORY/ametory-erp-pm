import { PermissionModel } from "./permission";

export interface RoleModel {
  id?: string;
  name: string;
  permissions: PermissionModel[];
  permission_names: string[];
  company_id?: string;
  is_admin: boolean;
  is_merchant: boolean;
  is_super_admin: boolean;
  is_owner: boolean;
}
