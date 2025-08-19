
export interface WhatsappInteractiveModel {
    id?: string;
    title?: string;
    description?: string;
    type?: string;
    ref_id?: string;
    ref_type?: string;
    data?: any;
}


export interface WhatsappInteractiveListHeader {
    type: string;
    text: string;
}

export interface WhatsappInteractiveListBody {
    text: string;
}

export interface WhatsappInteractiveListFooter {
    text: string;
}

export interface WhatsappInteractiveListRow {
    id: string;
    title: string;
    description: string;
}

export interface WhatsappInteractiveListSection {
    title: string;
    rows: WhatsappInteractiveListRow[];
}

export interface WhatsappInteractiveListAction {
    button: string;
    sections: WhatsappInteractiveListSection[];
}

export interface WhatsappInteractiveList {
    type: string;
    header?: WhatsappInteractiveListHeader;
    body: WhatsappInteractiveListBody;
    footer?: WhatsappInteractiveListFooter;
    action: WhatsappInteractiveListAction;
}


export interface WhatsappInteractiveCTAUrl {
    type: string;
    header?: WhatsappInteractiveListHeader;
    body: WhatsappInteractiveListBody;
    footer?: WhatsappInteractiveListFooter;
    action: WhatsappInteractiveCTAUrlAction;
}

export interface WhatsappInteractiveCTAUrlAction {
    name: string;
    parameters: {
        display_text: string;
        url: string;
    };
}
