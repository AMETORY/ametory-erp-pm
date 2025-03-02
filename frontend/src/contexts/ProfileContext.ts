import { createContext } from "react";
import { UserModel } from "../models/user";

export const ProfileContext = createContext<{
  profile: UserModel | null;
  setProfile: (profile: UserModel | null) => void;
}>({
  profile: null,
  setProfile: () => {},
});
