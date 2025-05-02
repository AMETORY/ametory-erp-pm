import { customFetch } from "./baseApi";


export const getCustomerInteraction= async (startDate: Date, endDate: Date, memberIDs: string[]) => {
    const queryParams = new URLSearchParams();
  return await customFetch(`api/v1/analytic/customer-interaction?${queryParams}`, {
      method: "POST",
      body: JSON.stringify({
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        member_ids: memberIDs,
      }),
  });
};
export const getATRCustomerInteraction= async (startDate: Date, endDate: Date, memberIDs: string[]) => {
    const queryParams = new URLSearchParams();
  return await customFetch(`api/v1/analytic/average-time-reply?${queryParams}`, {
      method: "POST",
      body: JSON.stringify({
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        member_ids: memberIDs,
      }),
  });
};

export const getHourlyCustomerInteraction = async (startDate: Date, endDate: Date, memberIDs: string[]) => {
    const queryParams = new URLSearchParams();
  return await customFetch(`api/v1/analytic/hourly-customer-interaction?${queryParams}`, {
      method: "POST",
      body: JSON.stringify({
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        member_ids: memberIDs,
      }),
  });
};
export const getHourlyATR = async (startDate: Date, endDate: Date, memberIDs: string[]) => {
    const queryParams = new URLSearchParams();
  return await customFetch(`api/v1/analytic/hourly-average-time-reply?${queryParams}`, {
      method: "POST",
      body: JSON.stringify({
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        member_ids: memberIDs,
      }),
  });
};