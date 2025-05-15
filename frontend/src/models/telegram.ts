import { CompanyModel } from "./company";
import { ContactModel } from "./contact";
import { MemberModel } from "./member";
import { UserModel } from "./user";



export interface TelegramMessageSessionModel {
    id: string;
    session: string;
    session_name: string;
    last_online_at?: Date;
    last_message: string;
    company_id: string | null;
    company: CompanyModel | null;
    contact_id: string | null;
    contact: ContactModel | null;
    ref_id: string | null;
    ref_type: string | null;
    ref: any;
    is_human_agent: boolean;
    count_unread: number;
}


export interface TelegramMessage {
    id: string;
    message: string;
    media_url: string | null;
    mime_type: string | null;
    session: string;
    contact_id: string | null;
    contact: ContactModel | null;
    company_id: string | null;
    company: CompanyModel | null;
    is_from_me: boolean;
    is_group: boolean;
    is_replied: boolean;
    sent_at: Date | null;
    is_read: boolean;
    message_id: string | null;
    response_time: number | null;
    member_id: string | null;
    member: MemberModel | null;
    user_id: string | null;
    user: UserModel | null;
    is_new: boolean;
    ref_id: string | null;
    is_auto_pilot: boolean;
    telegram_message_session_id: string | null;
    telegram_message_session: TelegramMessageSessionModel | null;
}
