import type { FC } from "react";
import { Navigate, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import ProjectPage from "../pages/ProjectPage";
import TaskPage from "../pages/TaskPage";
import MemberPage from "../pages/MemberPage";
import InboxPage from "../pages/Inbox";
import ChatPage from "../pages/ChatPage";
import ProfilePage from "../pages/ProfilePage";
import FormPage from "../pages/FormPage";
import FormTempateDetail from "../pages/FormTempateDetail";
import FormDetail from "../pages/FormDetail";
import FormPublicPage from "../pages/FormPublicPage";
import ContactPage from "../pages/ContactPage";
import SettingPage from "../pages/SettingPage";
import GeminiAgentPage from "../pages/GeminiAgentPage";
import GeminiAgentDetail from "../pages/GeminiAgentDetail";
import TaskAttributePage from "../pages/TaskAttributePage";
import TaskAttributeDetail from "../pages/TaskAttributeDetail";
import ConnectionPage from "../pages/ConnectionPage";
import ConnectionDetail from "../pages/ConnectionDetail";
import WhatsappPage from "../pages/WhatsappPage";
import BroadcastPage from "../pages/BroadcastPage";
import BroadcastDetail from "../pages/BroadcastDetail";
import TagPage from "../pages/TagPage";
import ProductPage from "../pages/ProductPage";
import TemplatePage from "../pages/TemplatePage";
import TemplateDetail from "../pages/TemplateDetail";
import TelegramPage from "../pages/TelegramPage";
import PrivacyPolicy from "../pages/PrivacyPolicy";
import FacebookCallback from "../pages/FacebookCallback";
import InstagramPage from "../pages/InstagramPage";

interface PrivateRouteProps {}

const PrivateRoute: FC<PrivateRouteProps> = ({}) => {
  return (
    <Routes>
      <Route path="/public/form/:formCode" element={<FormPublicPage />} />
      <Route path="/" element={<Home />} />
      <Route path="/project" element={<ProjectPage />} />
      <Route path="/task" element={<TaskPage />} />
      <Route path="/member" element={<MemberPage />} />
      <Route path="/inbox" element={<InboxPage />} />
      <Route path="/chat" element={<ChatPage />} />
      <Route path="/profile" element={<ProfilePage />} />
      <Route path="/form" element={<FormPage />} />
      <Route path="/tag" element={<TagPage />} />
      <Route
        path="/form-template/:templateId"
        element={<FormTempateDetail />}
      />
      <Route path="/auth/facebook/callback" element={<FacebookCallback />} />
      <Route path="/privacy" element={<PrivacyPolicy />} />
      <Route path="/form/:formId" element={<FormDetail />} />
      <Route path="/chat/:channelId" element={<ChatPage />} />
      <Route path="/project/:projectId" element={<ProjectDetail />} />
      <Route path="/contact" element={<ContactPage />} />
      <Route path="/connection" element={<ConnectionPage />} />
      <Route path="/connection/:connectionId" element={<ConnectionDetail />} />
      <Route path="/setting" element={<SettingPage />} />
      <Route path="/gemini-agent" element={<GeminiAgentPage />} />
      <Route path="/task-attribute" element={<TaskAttributePage />} />
      <Route path="/product" element={<ProductPage />} />
      <Route path="/whatsapp" element={<WhatsappPage />} />
      <Route path="/telegram" element={<TelegramPage />} />
      <Route path="/telegram/:sessionId" element={<TelegramPage />} />
      <Route path="/instagram" element={<InstagramPage />} />
      <Route path="/instagram/:sessionId" element={<InstagramPage />} />
      <Route path="/broadcast" element={<BroadcastPage />} />
      <Route path="/template" element={<TemplatePage />} />
      <Route path="/template/:templateId" element={<TemplateDetail />} />
      <Route path="/broadcast/:broadcastId" element={<BroadcastDetail/>} />
      <Route path="/whatsapp/:sessionId" element={<WhatsappPage />} />
      <Route
        path="/task-attribute/:attributeId"
        element={<TaskAttributeDetail />}
      />
      <Route path="/gemini-agent/:agentId" element={<GeminiAgentDetail />} />
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
};
export default PrivateRoute;
