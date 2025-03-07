import { FileModel } from "./file";
import { MemberModel } from "./member";
import { UserModel } from "./user";

export interface InboxModel {
  id: string;
  user_id: string;
  member_id: string;
  name: string;
  is_trash: boolean;
  is_default: boolean;
}

export interface InboxMessageModel {
  id?: string;
  inbox_id?: string;
  inbox?: InboxModel;
  sender_user_id?: string;
  sender_user?: UserModel;
  sender_member_id?: string;
  sender_member?: MemberModel;
  recipient_user_id?: string;
  recipient_user?: UserModel;
  recipient_member_id?: string;
  recipient_member?: MemberModel;
  cc_users?: UserModel[];
  cc_members?: MemberModel[];
  subject: string;
  message: string;
  read: boolean;
  parent_id?: string;
  parent?: InboxMessageModel;
  attachments?: FileModel[];
  replies?: InboxMessageModel[];
  favorited_by_users?: UserModel[];
  favorited_by_members?: MemberModel[];
  date?: Date;
}
