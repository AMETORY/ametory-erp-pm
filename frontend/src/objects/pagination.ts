export interface PaginationRequest {
    page: number;
    size: number;
    search?: string;
    status?: string;
    type?: string;
    project_id?: string;
    team_id?: string;
    order?: string;
    tag_ids?: string
  }
  
  

export interface PaginationResponse {
    page: number;
    size: number;
    max_page?: number;
    total_pages: number;
    total: number;
    last?: boolean;
    first?: boolean;
    visible?: number;
    
  }