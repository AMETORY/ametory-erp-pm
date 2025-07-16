export interface TiktokMessageSession {
    can_send_message: boolean;
    create_time: number;
    id: string;
    latest_message: TiktokMessage;
    participant_count: number;
    participants: Array<{
        avatar: string;
        im_user_id: string;
        nickname: string;
        role: 'SYSTEM' | 'SHOP' | 'BUYER';
        user_id?: string;
        buyer_platform?: 'TIKTOK_SHOP';
    }>;
    unread_count: number;
}


export interface TiktokMessage {
    content: any
    create_time: number;
    id: string;
    index: string;
    is_visible: boolean;
    sender: {
        avatar: string;
        im_user_id: string;
        nickname: string;
        role: 'SYSTEM' | 'SHOP' | 'BUYER';
    };
    type: 'NOTIFICATION';
}
