import { createContext } from "react";
import { UserModel } from "../models/user";
import { MemberModel } from "../models/member";

export const ProfileContext = createContext<{
  profile: UserModel | null;
  setProfile: (profile: UserModel | null) => void;
}>({
  profile: null,
  setProfile: () => {},
});


export const MemberContext = createContext<{
  member: MemberModel | null;
  setMember: (member: MemberModel | null) => void;
}>({
  member: null,
  setMember: () => {},
});
