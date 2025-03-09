import type { FC } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import ProjectPage from "../pages/ProjectPage";
import TaskPage from "../pages/TaskPage";
import MemberPage from "../pages/MemberPage";
import InboxPage from "../pages/Inbox";
import ChatPage from "../pages/ChatPage";

interface PrivateRouteProps {}

const PrivateRoute: FC<PrivateRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/project" element={<ProjectPage />} />
      <Route path="/task" element={<TaskPage />} />
      <Route path="/member" element={<MemberPage />} />
      <Route path="/inbox" element={<InboxPage />} />
      <Route path="/chat" element={<ChatPage />} />
      <Route path="/chat/:channelId" element={<ChatPage />} />
      <Route path="/project/:projectId" element={<ProjectDetail />} />
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
};
export default PrivateRoute;
