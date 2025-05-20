import { createContext } from "react";


export const ScrollContext = createContext<{
  scrollPositions: {[key: string]: number};
  setScrollPositions: (val: {[key: string]: number}) => void;
}>({
  scrollPositions: {},
  setScrollPositions: (val) => {},
});
