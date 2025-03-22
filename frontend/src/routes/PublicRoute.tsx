import type { FC } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import Login from "../pages/Login";
import Registration from "../pages/Register";
import Verify from "../pages/Verify";
import AcceptInvitation from "../pages/AcceptInvitation";
import FormPublicPage from "../pages/FormPublicPage";
import MemberRegisterPage from "../pages/MemberRegisterPage";

interface PublicRouteProps {}

const PublicRoute: FC<PublicRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/public/form/:formCode" element={<FormPublicPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Registration />} />
      <Route path="/invitation/verify/:token" element={<AcceptInvitation />} />
      <Route path="/verify/:token" element={<Verify />} />
      <Route path="/member/register/:code" element={<MemberRegisterPage />} />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
};
export default PublicRoute;
