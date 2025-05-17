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
import Forgot from "../pages/ForgotPage";
import PrivacyPolicy from "../pages/PrivacyPolicy";
import FacebookCallback from "../pages/FacebookCallback";

interface PublicRouteProps {}

const PublicRoute: FC<PublicRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/privacy" element={<PrivacyPolicy />} />
      <Route path="/public/form/:formCode" element={<FormPublicPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/forgot" element={<Forgot />} />
      <Route path="/register" element={<Registration />} />
      <Route path="/invitation/verify/:token" element={<AcceptInvitation />} />
      <Route path="/verify/:token" element={<Verify />} />
      <Route path="/member/register/:code" element={<MemberRegisterPage />} />
      <Route path="/auth/facebook/callback" element={<FacebookCallback />} />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
};
export default PublicRoute;
