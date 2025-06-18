import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getConnections = async (req: PaginationRequest) => {
    const queryParams = new URLSearchParams();
    queryParams.set("page", String(req.page));
    queryParams.set("size", String(req.size));
    if (req.search) queryParams.set("search", req.search);
    return await customFetch(`api/v1/connection/list?${queryParams}`, {
        method: "GET",
    });
};

export const getConnection = async (id: string) => {
    return await customFetch(`api/v1/connection/${id}`);
};

export const createConnection = async (connection: any) => {
    return await customFetch("api/v1/connection/create", {
        method: "POST",
        body: JSON.stringify(connection),
    });
};

export const updateConnection = async (id: string, connection: any) => {
    return await customFetch(`api/v1/connection/${id}`, {
        method: "PUT",
        body: JSON.stringify(connection),
    });
};
export const getShopeeAuthURL = async () => {
    return await customFetch(`api/v1/connection/auth-url/shopee`, {
        method: "GET",
    });
};
export const getLazadaAuthURL = async () => {
    return await customFetch(`api/v1/connection/auth-url/lazada`, {
        method: "GET",
    });
};
export const authorizeConnection = async (id: string, data: any) => {
    console.log(`api/v1/connection/${id}/authorize`)
    return await customFetch(`api/v1/connection/${id}/authorize`, {
        method: "PUT",
        body: JSON.stringify(data),
    });
};
export const syncContactConnection = async (id: string) => {
    return await customFetch(`api/v1/connection/${id}/sync-contact`, {
        method: "PUT",
    });
};
export const connectDevice = async (id: string) => {
    return await customFetch(`api/v1/connection/${id}/connect`, {
        method: "PUT",
    });
};
export const getQr = async (id: string, sessionName: string) => {
    return await customFetch(`api/v1/connection/${id}/get-qr/${sessionName}`, {
        method: "PUT",
    });
};
 
export const deleteConnection = async (id: string) => {
    await customFetch(`api/v1/connection/${id}`, {
        method: "DELETE",
    });
};
