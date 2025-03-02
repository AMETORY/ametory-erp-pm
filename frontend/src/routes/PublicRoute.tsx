import type { FC } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import Login from "../pages/Login";
import Registration from "../pages/Register";
import Verify from "../pages/Verify";

interface PublicRouteProps {}

const PublicRoute: FC<PublicRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Registration />} />
      <Route path="/verify/:token" element={<Verify />} />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
};
export default PublicRoute;
