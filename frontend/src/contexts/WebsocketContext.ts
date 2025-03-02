import { createContext, useEffect, useState } from "react";

export const WebsocketContext = createContext<{
  setWsMsg: (message: any) => void;
  wsMsg: any;
  isWsConnected: boolean;
  setWsConnected: (loading: boolean) => void;
}>({
  setWsMsg: () => {},
  wsMsg: null,
  isWsConnected: false,
  setWsConnected: () => {},
});

