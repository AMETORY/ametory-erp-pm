import { FileModel } from "./file";
import { MemberModel } from "./member";

export interface ChatChannelModel {
    id: string;
    name: string;
    description: string;
    icon: string;
    color: string;
    created_by_member_id: string;
    avatar?: FileModel;
    participant_members?: MemberModel[];
}


export interface ChatMessageModel {
    id: string;
    chat_channel_id: string;
    sender_member_id: string;
    sender_member: MemberModel;
    message: string;
    type: string;
    date: string;
    files: FileModel[];
}
