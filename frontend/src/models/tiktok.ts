export interface TiktokMessageSession {
  can_send_message: boolean;
  create_time: number;
  id: string;
  latest_message: TiktokMessage;
  participant_count: number;
  participants: Array<TiktokParticipant>;
  unread_count: number;
}

export interface TiktokParticipant {
    avatar: string;
    buyer_platform: string;
    im_user_id: string;
    nickname: string;
    role: string;
    user_id: string;
}
export interface TiktokSessionDetail {
  id: string;
  created_at: string;
  session: string;
  session_name: string;
  create_time: string;
  last_message: {
    content: string;
  };
  last_msg_time: string;
  company_id: string;
  participant: TiktokParticipant;
  ref_id: string;
  ref_type: "connection";
  is_human_agent: boolean;
  count_unread: number;
}
export interface TiktokMessage {
  content: any;
  create_time: number;
  id: string;
  index: string;
  is_visible: boolean;
  sender: TiktokParticipant;
  type: string;
}
