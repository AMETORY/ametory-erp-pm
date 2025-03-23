import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Button } from "flowbite-react";
import {
  WhatsappMessageModel,
  WhatsappMessageSessionModel,
} from "../models/whatsapp_message";
import {
  LOCAL_STORAGE_DEFAULT_CHANNEL,
  LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
} from "../utils/constants";
import { useNavigate, useParams } from "react-router-dom";
import { asyncStorage } from "../utils/async_storage";
import { PaginationResponse } from "../objects/pagination";
import { getPagination } from "../utils/helper";
import { getWhatsappSessions } from "../services/api/whatsappApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import WhatsappMessages from "../components/WhatsappMessages";

interface WhatsappPageProps {}

const WhatsappPage: FC<WhatsappPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);

  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);
  const [sessions, setSessions] = useState<WhatsappMessageSessionModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const { sessionId } = useParams();
  const nav = useNavigate();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllSessions();
    }
  }, [mounted, sessionId]);

  const getAllSessions = async () => {
    try {
      setLoading(true);
      const resp: any = await getWhatsappSessions(sessionId!, {
        page,
        size,
        search,
      });
      setSessions(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Whatsapp</h1>
          <div className="flex gap-2">
            <Button
              gradientDuoTone="purpleToBlue"
              pill
              onClick={() => {
                // setOpenChannelForm(true);
              }}
            >
              + Chat
            </Button>
          </div>
        </div>
        <div className="flex flex-row w-full h-full flex-1 gap-2">
          <div className="w-[300px] h-full">
            <ul className="space-y-2">
              {sessions.map((e) => (
                <li
                  className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                  key={e.id}
                  onClick={() => {
                    nav(`/whatsapp/${e.id}`);
                    asyncStorage.setItem(
                      LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
                      e.id
                    );
                  }}
                  style={{ background: sessionId == e.id ? "#e5e7eb" : "" }}
                >
                  <div className="flex gap-2 items-center">
                    {e.contact?.avatar && (
                      <img
                        src={e.contact?.avatar.url}
                        className=" aspect-square rounded-full object-cover w-8 h-8"
                      />
                    )}
                    <div className="flex flex-col">
                      <span className="font-semibold">{e.contact?.name}</span>
                      <small className="">{e.last_message}</small>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          </div>
          <div className="w-full border-l relative">
            {sessionId && <WhatsappMessages sessionId={sessionId} />}
          </div>
        </div>
      </div>
    </AdminLayout>
  );
};
export default WhatsappPage;
