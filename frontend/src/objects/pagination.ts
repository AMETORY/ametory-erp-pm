export interface PaginationRequest {
    page: number;
    size: number;
    search?: string;
    status?: string;
    type?: string;
    project_id?: string;
    team_id?: string;
  }
  
  

export interface PaginationResponse {
    page: number;
    size: number;
    max_page: number;
    total_pages: number;
    total: number;
    last: boolean;
    first: boolean;
    visible: number;
  }