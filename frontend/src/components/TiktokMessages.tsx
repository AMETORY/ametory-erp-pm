import { useEffect, useState, type FC } from "react";
import { ConnectionModel } from "../models/connection";
import { PaginationResponse } from "../objects/pagination";
import { connect } from "http2";
import { getTiktokSessionMessages } from "../services/api/tiktokApi";
import { TiktokMessage, TiktokMessageSession } from "../models/tiktok";

interface TiktokMessagesProps {
  sessionId: string;
  session: TiktokMessageSession;
  connection: ConnectionModel;
}

const TiktokMessages: FC<TiktokMessagesProps> = ({ sessionId, connection }) => {
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();

    const [messages, setMessages] = useState<TiktokMessage[]>([]);

  useEffect(() => {
    if (sessionId && connection) {
      getTiktokSessionMessages(sessionId, {
        page: page,
        size: size,
        search: search,
        connection_session: connection.id,
      }).then((resp: any) => {
        setMessages(resp.data.messages);
        // setPagination(getPagination(resp.data));
      });
    }

    return () => {};
  }, [sessionId, connection]);
  return <h1>{}</h1>;
};
export default TiktokMessages;
