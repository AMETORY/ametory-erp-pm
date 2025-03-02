import { useEffect, useState } from "react";
import AppRoutes from "./routes/AppRoutes";
import { asyncStorage } from "./utils/async_storage";
import { LOCAL_STORAGE_TOKEN } from "./utils/constants";

function App() {
  const [token, setToken] = useState<string | null>(null);
  const [mounted, setMounted] = useState(false);
 
  return <AppRoutes token={localStorage.getItem(LOCAL_STORAGE_TOKEN)} />;
}

export default App;
