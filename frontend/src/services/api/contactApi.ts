import { PaginationRequest } from "../../objects/pagination";
import { customFetch } from "./baseApi";

export const getContacts = async (req: PaginationRequest) => {
    const queryParams = new URLSearchParams();
    queryParams.set("page", String(req.page));
    queryParams.set("size", String(req.size));
    if (req.search) queryParams.set("search", req.search);
	return await customFetch(`api/v1/contact/list?${queryParams}`, {
		method: "GET",
	});
};

export const getContact = async (id: string) => {
	return await customFetch(`api/v1/contact/${id}`);
};

export const createContact = async (contact: any) => {
	return await customFetch("api/v1/contact/create", {
        method: "POST",
        body: JSON.stringify(contact),
    });
};

export const updateContact = async (id: string, contact: any) => {
	return await customFetch(`api/v1/contact/${id}`, {
        method: "PUT",
        body: JSON.stringify(contact),
    });
};
 
export const deleteContact = async (id: string) => {
	await customFetch(`api/v1/contact/${id}`);
};
