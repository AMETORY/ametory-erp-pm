import { createContext } from "react";

export const SearchContext = createContext<{
  search: string;
  setSearch: (search: string) => void;
}>({
  search: "",
  setSearch: () => {},
});
