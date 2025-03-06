import type { FC } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import ProjectPage from "../pages/ProjectPage";
import TaskPage from "../pages/TaskPage";

interface PrivateRouteProps {}

const PrivateRoute: FC<PrivateRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/project" element={<ProjectPage />} />
      <Route path="/task" element={<TaskPage />} />
      <Route path="/project/:projectId" element={<ProjectDetail />} />
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
};
export default PrivateRoute;
